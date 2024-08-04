package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
)

var _ = Describe("score calculating for tickets", func() {
	Context("given ratings list", func() {
		var ratings []entities.Rating

		categories := []entities.Category{
			{
				ID:     1,
				Name:   "Spelling",
				Weight: float32(1),
			},
			{
				ID:     2,
				Name:   "Grammar",
				Weight: float32(0.7),
			},
		}

		BeforeEach(func() {
			ratings = []entities.Rating{
				{
					TicketID:   1,
					CategoryID: 1,
					Value:      1,
				},
				{
					TicketID:   1,
					CategoryID: 2,
					Value:      2,
				},
				{
					TicketID:   1,
					CategoryID: 2,
					Value:      3,
				},
				{
					TicketID:   2,
					CategoryID: 1,
					Value:      5,
				},
			}
		})

		When("calculating scores", func() {
			var ticketScores []vo.TicketScore

			BeforeEach(func() {
				service := services.NewTicketScores(categories)
				ticketScores = service.GetScores(ratings)
			})

			It("should includes scores for each tickets", func() {
				Expect(ticketScores).To(HaveLen(2))
				Expect(ticketScores).To(ContainElement(HaveField("TicketID", vo.NewTicketID(1))))
				Expect(ticketScores).To(ContainElement(HaveField("TicketID", vo.NewTicketID(2))))
			})

			It("each tickets should have each categories", func() {
				Expect(ticketScores[0].Categories).To(HaveLen(2))
				Expect(ticketScores[0].Categories).To(ContainElement(HaveField("CategoryID", vo.NewCategoryID(1))))
				Expect(ticketScores[0].Categories).To(ContainElement(HaveField("CategoryID", vo.NewCategoryID(2))))
				Expect(ticketScores[1].Categories).To(HaveLen(2))
				Expect(ticketScores[1].Categories).To(ContainElement(HaveField("CategoryID", vo.NewCategoryID(1))))
				Expect(ticketScores[1].Categories).To(ContainElement(HaveField("CategoryID", vo.NewCategoryID(2))))
			})

			It("scores should be correct", func() {
				Expect(ticketScores[0].Categories[0].Score).To(HaveValue(Equal(float32(20))))
				Expect(ticketScores[0].Categories[1].Score).To(HaveValue(Equal(float32(39.875))))
				Expect(ticketScores[1].Categories[0].Score).To(HaveValue(Equal(float32(100))))
				Expect(ticketScores[1].Categories[1].Score).To(BeNil())
			})
		})
	})
})
