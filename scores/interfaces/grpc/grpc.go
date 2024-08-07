package grpc

import (
	"context"
	"fmt"
	getbycategories "go-scores/gen/get_by_categories"
	getbytickets "go-scores/gen/get_by_tickets"
	getchangesbyperiods "go-scores/gen/get_changes_by_periods"
	getoverall "go-scores/gen/get_overall"
	get_by_categoriespb "go-scores/gen/grpc/get_by_categories/pb"
	getbycategoriessvr "go-scores/gen/grpc/get_by_categories/server"
	get_by_ticketspb "go-scores/gen/grpc/get_by_tickets/pb"
	getbyticketssvr "go-scores/gen/grpc/get_by_tickets/server"
	get_changes_by_periodspb "go-scores/gen/grpc/get_changes_by_periods/pb"
	getbyperiodssvr "go-scores/gen/grpc/get_changes_by_periods/server"
	get_overallpb "go-scores/gen/grpc/get_overall/pb"
	getoverallsvr "go-scores/gen/grpc/get_overall/server"
	"net"
	"net/url"
	"sync"

	"goa.design/clue/debug"
	"goa.design/clue/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func Handle(ctx context.Context, u *url.URL, getOverallEndpoints *getoverall.Endpoints, getByCategoriesEndpoints *getbycategories.Endpoints, getByTicketsEndpoints *getbytickets.Endpoints, getByPeriodsEndpoints *getchangesbyperiods.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		getOverallServer      *getoverallsvr.Server
		getByCategoriesServer *getbycategoriessvr.Server
		getByTicketsServer    *getbyticketssvr.Server
		getByPeriodsServer    *getbyperiodssvr.Server
	)
	{
		getOverallServer = getoverallsvr.New(getOverallEndpoints, nil)
		getByCategoriesServer = getbycategoriessvr.New(getByCategoriesEndpoints, nil)
		getByTicketsServer = getbyticketssvr.New(getByTicketsEndpoints, nil)
		getByPeriodsServer = getbyperiodssvr.New(getByPeriodsEndpoints, nil)
	}

	// Create interceptor which sets up the logger in each request context.
	chain := grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx))
	if dbg {
		// Log request and response content if debug logs are enabled.
		chain = grpc.ChainUnaryInterceptor(log.UnaryServerInterceptor(ctx), debug.UnaryServerInterceptor())
	}

	// Initialize gRPC server
	srv := grpc.NewServer(chain)

	// Register the servers.
	get_overallpb.RegisterGetOverallServer(srv, getOverallServer)
	get_by_categoriespb.RegisterGetByCategoriesServer(srv, getByCategoriesServer)
	get_by_ticketspb.RegisterGetByTicketsServer(srv, getByTicketsServer)
	get_changes_by_periodspb.RegisterGetChangesByPeriodsServer(srv, getByPeriodsServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			log.Printf(ctx, "serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			if lis == nil {
				errc <- fmt.Errorf("failed to listen on %q", u.Host)
			}
			log.Printf(ctx, "gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
