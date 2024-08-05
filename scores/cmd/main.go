package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/glebarez/sqlite"
	getbycategories "go-scores/gen/get_by_categories"
	getbytickets "go-scores/gen/get_by_tickets"
	getchangesbyperiods "go-scores/gen/get_changes_by_periods"
	getoverall "go-scores/gen/get_overall"
	"go-scores/scores/interfaces/controllers"
	"go-scores/scores/interfaces/grpc"
	"gorm.io/gorm"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"goa.design/clue/debug"
	"goa.design/clue/log"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		grpcPortF = flag.String("grpc-port", "", "gRPC port (overrides host gRPC port specified in service design)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	format := log.FormatJSON
	if log.IsTerminal() {
		format = log.FormatTerminal
	}
	ctx := log.Context(context.Background(), log.WithFormat(format))
	if *dbgF {
		ctx = log.Context(ctx, log.WithDebug())
		log.Debugf(ctx, "debug logs enabled")
	}
	log.Print(ctx, log.KV{K: "http-port", V: *httpPortF})

	//Initialize DB
	dsn := "../../database.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(ctx, err, "Failed to connect to the database")
	}

	// Initialize the services.
	var (
		getOverallSvc      getoverall.Service
		getByCategoriesSvc getbycategories.Service
		getByTicketsSvc    getbytickets.Service
		getByPeriodsSvc    getchangesbyperiods.Service
	)
	{
		getOverallSvc = controllers.NewGetOverall(db)
		getByCategoriesSvc = controllers.NewGetByCategories(db)
		getByTicketsSvc = controllers.NewGetByTickets(db)
		getByPeriodsSvc = controllers.NewGetByPeriods(db)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		getOverallEndpoints      *getoverall.Endpoints
		getByCategoriesEndpoints *getbycategories.Endpoints
		getByTicketsEndpoints    *getbytickets.Endpoints
		getByPeriodsEndpoints    *getchangesbyperiods.Endpoints
	)
	{
		getOverallEndpoints = getoverall.NewEndpoints(getOverallSvc)
		getOverallEndpoints.Use(debug.LogPayloads())
		getOverallEndpoints.Use(log.Endpoint)
		getByCategoriesEndpoints = getbycategories.NewEndpoints(getByCategoriesSvc)
		getByCategoriesEndpoints.Use(debug.LogPayloads())
		getByCategoriesEndpoints.Use(log.Endpoint)
		getByTicketsEndpoints = getbytickets.NewEndpoints(getByTicketsSvc)
		getByTicketsEndpoints.Use(debug.LogPayloads())
		getByTicketsEndpoints.Use(log.Endpoint)
		getByPeriodsEndpoints = getchangesbyperiods.NewEndpoints(getByPeriodsSvc)
		getByPeriodsEndpoints.Use(debug.LogPayloads())
		getByPeriodsEndpoints.Use(log.Endpoint)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "grpc://0.0.0.0:8080"
			u, err := url.Parse(addr)
			if err != nil {
				log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					log.Fatalf(ctx, err, "invalid URL %#v\n", u.Host)
				}
				u.Host = net.JoinHostPort(h, *grpcPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "8080")
			}
			grpc.Handle(ctx, u, getOverallEndpoints, getByCategoriesEndpoints, getByTicketsEndpoints, getByPeriodsEndpoints, &wg, errc, *dbgF)
		}

	default:
		log.Fatal(ctx, fmt.Errorf("invalid host argument: %q (valid hosts: localhost)", *hostF))
	}

	// Wait for signal.
	log.Printf(ctx, "exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	log.Printf(ctx, "exited")
}
