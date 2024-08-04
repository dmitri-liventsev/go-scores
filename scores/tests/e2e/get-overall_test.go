package e2e

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	get_overallpb "go-scores/gen/grpc/get_overall/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var _ = Describe("get scores by categories", func() {
	var client get_overallpb.GetOverallClient

	BeforeEach(func() {
		{
			conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic(err)
			}
			client = get_overallpb.NewGetOverallClient(conn)

			DeferCleanup(func() {
				_ = conn.Close()
			})
		}
	})

	Context("given a daily periods and ratings list", func() {
		var from, to string
		var req *get_overallpb.GetOverallScoreRequest

		BeforeEach(func() {
			from = "2019-08-01"
			to = "2019-08-07"

			req = &get_overallpb.GetOverallScoreRequest{
				From: from,
				To:   to,
			}
		})

		When("fetching scores", func() {
			var err error
			var res *get_overallpb.GetOverallScoreResponse

			BeforeEach(func(ctx context.Context) {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				DeferCleanup(cancel)

				res, err = client.GetOverallScore(ctx, req)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return data", func() {
				Expect(res.Field).To(Equal(float32(41.67595672607422)))
			})
		})
	})
})
