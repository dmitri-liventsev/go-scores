package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
)

var _ = Describe("ticket score calculation for single rate", func() {
	Context("categories and rate are valid", func() {
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
			{
				ID:     3,
				Name:   "GDPR",
				Weight: float32(1.2),
			},
			{
				ID:     4,
				Name:   "Randomness",
				Weight: float32(0),
			},
		}

		When("calculating the score", func() {
			var calculator services.ScoreCalculator
			BeforeEach(func() {
				calculator = services.NewScoreCalculator(categories)
			})

			DescribeTable("score are correct", func(rate int, categoryId int, expectedScore float32) {
				Expect(calculator.Calculate(float32(rate), vo.NewCategoryID(categoryId))).To(Equal(&expectedScore))
			},
				Entry("Spelling - weight 1, rate 1", 1, 1, float32(20)),
				Entry("Spelling - weight 1, rate 2", 2, 1, float32(40)),
				Entry("Spelling - weight 1, rate 3", 3, 1, float32(60.000003814697266)),
				Entry("Spelling - weight 1, rate 4", 4, 1, float32(80)),
				Entry("Spelling - weight 1, rate 5", 5, 1, float32(100)),
				Entry("Spelling - weight 0.7, rate 1", 1, 2, float32(15.950000762939453)),
				Entry("Spelling - weight 0.7, rate 2", 2, 2, float32(31.900001525878906)),
				Entry("Spelling - weight 0.7, rate 3", 3, 2, float32(47.85000228881836)),
				Entry("Spelling - weight 0.7, rate 4", 4, 2, float32(63.80000305175781)),
				Entry("Spelling - weight 0.7, rate 5", 5, 2, float32(79.75)),
				Entry("Spelling - weight 1.2, rate 1", 1, 3, float32(23.200000762939453)),
				Entry("Spelling - weight 1.2, rate 2", 2, 3, float32(46.400001525878906)),
				Entry("Spelling - weight 1.2, rate 3", 3, 3, float32(69.60000610351562)),
				Entry("Spelling - weight 1.2, rate 4", 4, 3, float32(92.80000305175781)),
				Entry("Spelling - weight 1.2, rate 5", 5, 3, float32(100)),
				Entry("Spelling - weight 0, rate 1", 1, 4, float32(10)),
				Entry("Spelling - weight 0, rate 2", 2, 4, float32(20)),
				Entry("Spelling - weight 0, rate 3", 3, 4, float32(30.000001907348633)),
				Entry("Spelling - weight 0, rate 4", 4, 4, float32(40)),
				Entry("Spelling - weight 0, rate 5", 5, 4, float32(50)),
			)
		})
	})
})

var _ = Describe("ticket score calculation for multiple rates", func() {
	Context("categories and rate are valid", func() {
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
			{
				ID:     3,
				Name:   "GDPR",
				Weight: float32(1.2),
			},
			{
				ID:     4,
				Name:   "Randomness",
				Weight: float32(0),
			},
		}

		When("calculating the score", func() {
			var calculator services.ScoreCalculator
			var score *float32
			BeforeEach(func() {
				calculator = services.NewScoreCalculator(categories)
				score = calculator.CalculateRatingsScore([]int{1, 2}, 1)
			})

			It("score are correcty", func() {
				Expect(*score).To(Equal(float32(30.000001907348633)))
			})
		})
	})
})
