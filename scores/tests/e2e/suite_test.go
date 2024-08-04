package e2e

import (
	"context"
	"github.com/glebarez/sqlite"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	getbycategories "go-scores/gen/get_by_categories"
	getbyperiods "go-scores/gen/get_by_periods"
	getbytickets "go-scores/gen/get_by_tickets"
	getoverall "go-scores/gen/get_overall"
	get_by_categoriespb "go-scores/gen/grpc/get_by_categories/pb"
	"go-scores/scores/interfaces/controllers"
	"go-scores/scores/interfaces/grpc"
	"goa.design/clue/log"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"net"
	"net/url"
	"sync"
	"testing"
)

var DB *gorm.DB
var client get_by_categoriespb.GetByCategoriesClient
var _ = BeforeSuite(func(ctx context.Context) {
	{
		var err error
		DB, err = gorm.Open(sqlite.Open("../../../database.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		DeferCleanup(func() {
			sqlDB, _ := DB.DB()
			_ = sqlDB.Close()
		})
	}

	{
		var (
			getOverallSvc      getoverall.Service
			getByCategoriesSvc getbycategories.Service
			getByTicketsSvc    getbytickets.Service
			getByPeriodsSvc    getbyperiods.Service
		)

		{
			getOverallSvc = controllers.NewGetOverall(DB)
			getByCategoriesSvc = controllers.NewGetByCategories(DB)
			getByTicketsSvc = controllers.NewGetByTickets(DB)
			getByPeriodsSvc = controllers.NewGetByPeriods(DB)
		}

		var (
			getOverallEndpoints      *getoverall.Endpoints
			getByCategoriesEndpoints *getbycategories.Endpoints
			getByTicketsEndpoints    *getbytickets.Endpoints
			getByPeriodsEndpoints    *getbyperiods.Endpoints
		)
		{
			getOverallEndpoints = getoverall.NewEndpoints(getOverallSvc)
			getByCategoriesEndpoints = getbycategories.NewEndpoints(getByCategoriesSvc)
			getByTicketsEndpoints = getbytickets.NewEndpoints(getByTicketsSvc)
			getByPeriodsEndpoints = getbyperiods.NewEndpoints(getByPeriodsSvc)
		}

		addr := "grpc://localhost"
		u, err := url.Parse(addr)
		if err != nil {
			panic(err)
		}
		u.Host = net.JoinHostPort(u.Host, "8080")

		format := log.FormatJSON
		if log.IsTerminal() {
			format = log.FormatTerminal
		}
		ctx := log.Context(context.Background(), log.WithFormat(format))

		ctx, cancel := context.WithCancel(ctx)
		errc := make(chan error)

		var wg sync.WaitGroup
		grpc.Handle(ctx, u, getOverallEndpoints, getByCategoriesEndpoints, getByTicketsEndpoints, getByPeriodsEndpoints, &wg, errc, false)

		DeferCleanup(func() {
			cancel()
			wg.Wait()
		})
	}

	{
		conn, err := ggrpc.NewClient("localhost:8080", ggrpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		client = get_by_categoriespb.NewGetByCategoriesClient(conn)

		DeferCleanup(func() {
			_ = conn.Close()
		})
	}
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "wallet/tests")
}