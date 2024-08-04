package tests

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	getbytickets "go-scores/gen/get_by_tickets"
	"go-scores/scores/interfaces/controllers"
)

var _ = Describe("get scores by tickets", func() {
	Context("given a week period", func() {
		var client *getbytickets.Client
		var from, to string

		BeforeEach(func() {
			endpoint := getbytickets.NewGetAggregatedScoresByTicketEndpoint(controllers.NewGetByTickets(DB))
			client = getbytickets.NewClient(endpoint)

			from = "2019-08-01"
			to = "2019-08-07"
		})

		When("calculating scores", func() {
			var ticketScores *getbytickets.GetAggregatedScoresByTicketResult
			var err error
			BeforeEach(func(ctx context.Context) {
				payload := &getbytickets.GetAggregatedScoresByTicketPayload{
					From: from,
					To:   to,
				}

				ticketScores, err = client.GetAggregatedScoresByTicket(ctx, payload)
				Expect(err).To(BeNil())
			})

			It("should includes correct meta", func() {
				Expect(ticketScores.Meta.Categories).To(HaveLen(4))
				Expect(ticketScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(ticketScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(ticketScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(3)))))
				Expect(ticketScores.Meta.Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(4)))))
			})

			It("should includes scores for each tickets", func() {
				Expect(ticketScores.Data).To(HaveLen(183))
			})

			It("each tickets should have each categories", func() {
				Expect(ticketScores.Data[0].Categories).To(HaveLen(4))
				Expect(ticketScores.Data[0].Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(1)))))
				Expect(ticketScores.Data[0].Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(2)))))
				Expect(ticketScores.Data[0].Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(3)))))
				Expect(ticketScores.Data[0].Categories).To(ContainElement(HaveField("ID", HaveValue(Equal(4)))))
			})

			It("scores should be correct", func() {
				Expect(ticketScores.Data[0].Categories[0].Score).To(HaveValue(Equal(float32(20))))
				Expect(ticketScores.Data[0].Categories[1].Score).To(HaveValue(Equal(float32(79.75))))
				Expect(ticketScores.Data[1].Categories[0].Score).To(HaveValue(Equal(float32(20))))
				Expect(ticketScores.Data[1].Categories[1].Score).To(HaveValue(Equal(float32(63.80000305175781))))
			})
		})
	})
})
