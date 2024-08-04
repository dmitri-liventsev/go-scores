package e2e

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	get_changes_by_periodspb "go-scores/gen/grpc/get_changes_by_periods/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var _ = Describe("get scores by categories", func() {
	var client get_changes_by_periodspb.GetChangesByPeriodsClient

	BeforeEach(func() {
		{
			conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic(err)
			}
			client = get_changes_by_periodspb.NewGetChangesByPeriodsClient(conn)

			DeferCleanup(func() {
				_ = conn.Close()
			})
		}
	})

	Context("given a daily periods and ratings list", func() {
		var req *get_changes_by_periodspb.GetChangesByPeriodsRequest

		BeforeEach(func() {
			req = &get_changes_by_periodspb.GetChangesByPeriodsRequest{
				Period: "month",
			}
		})

		When("fetching scores", func() {
			var err error
			var res *get_changes_by_periodspb.GetChangesByPeriodsResponse

			BeforeEach(func(ctx context.Context) {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				DeferCleanup(cancel)

				res, err = client.GetChangesByPeriods(ctx, req)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return data", func() {
				Expect(res.Data).To(HaveLen(133))
			})
		})
	})
})
