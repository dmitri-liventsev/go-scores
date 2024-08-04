package e2e

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	get_by_periodspb "go-scores/gen/grpc/get_by_periods/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var _ = Describe("get scores by categories", func() {
	var client get_by_periodspb.GetByPeriodsClient

	BeforeEach(func() {
		{
			conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic(err)
			}
			client = get_by_periodspb.NewGetByPeriodsClient(conn)

			DeferCleanup(func() {
				_ = conn.Close()
			})
		}
	})

	Context("given a daily periods and ratings list", func() {
		var req *get_by_periodspb.GetByPeriodsRequest

		BeforeEach(func() {
			req = &get_by_periodspb.GetByPeriodsRequest{
				Period: "month",
			}
		})

		When("fetching scores", func() {
			var err error
			var res *get_by_periodspb.GetByPeriodsResponse

			BeforeEach(func(ctx context.Context) {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				DeferCleanup(cancel)

				res, err = client.GetByPeriods(ctx, req)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return data", func() {
				Expect(res.Data).To(HaveLen(133))
			})
		})
	})
})
