package tests

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	getoverall "go-scores/gen/get_overall"
	"go-scores/scores/interfaces/controllers"
)

var _ = Describe("get overall score by period", func() {
	Context("given a week period", func() {
		var client *getoverall.Client
		var from, to string

		BeforeEach(func() {
			endpoint := getoverall.NewGetOverallScoreEndpoint(controllers.NewGetOverall(DB))
			client = getoverall.NewClient(endpoint)

			from = "2019-08-01"
			to = "2019-08-07"
		})

		When("calculating scores", func() {
			var score float32
			var err error
			BeforeEach(func(ctx context.Context) {
				payload := &getoverall.GetOverallScorePayload{
					From: from,
					To:   to,
				}

				score, err = client.GetOverallScore(ctx, payload)
				Expect(err).To(BeNil())
			})

			It("should be correct", func() {
				Expect(score).To(Equal(float32(41.67595672607422)))
			})
		})
	})
})
