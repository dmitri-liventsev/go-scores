package e2e

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	get_by_categoriespb "go-scores/gen/grpc/get_by_categories/pb"
	"time"
)

var _ = Describe("get scores by categories", func() {
	Context("given a daily periods and ratings list", func() {
		var from, to string
		var req *get_by_categoriespb.GetAggregatedScoresRequest

		BeforeEach(func() {
			from = "2019-08-01"
			to = "2019-08-02"

			req = &get_by_categoriespb.GetAggregatedScoresRequest{
				From: from,
				To:   to,
			}
		})

		When("fetching scores", func() {
			var err error
			var res *get_by_categoriespb.GetAggregatedScoresResponse

			BeforeEach(func(ctx context.Context) {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				DeferCleanup(cancel)

				res, err = client.GetAggregatedScores(ctx, req)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return data", func() {
				Expect(res.Data).To(HaveLen(4))
			})
		})
	})
})
