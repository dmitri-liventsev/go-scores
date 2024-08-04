package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/services"
)

var _ = Describe("score calculating for tickets", func() {
	Context("given ratings and categories  list", func() {
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
			var score float32

			BeforeEach(func() {
				service := services.NewOverallScore(categories)
				score = service.GetScore(ratings)
			})

			It("should be correct", func() {
				Expect(score).To(Equal(float32(49.9375)))
			})
		})
	})
})
