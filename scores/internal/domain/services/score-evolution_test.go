package services_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
	pkgtime "go-scores/scores/pkg/time"
	"time"
)

var _ = Describe("calculation of score delta between periods", func() {
	Context("given ratings and categories list", func() {
		var ratings []entities.Rating
		var periods []entities.Period

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

		service := services.NewScoreEvolution(categories)

		When("calculating score delta between weeks periods", func() {
			var scoreDeltas []vo.ScoreDelta
			BeforeEach(func() {
				w1Start := time.Date(2024, time.August, 25, 0, 0, 0, 0, time.UTC)
				w1End := time.Date(2024, time.August, 25, 0, 0, 0, 0, time.UTC)
				w2Start := time.Date(2024, time.August, 26, 0, 0, 0, 0, time.UTC)
				w2Date := time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC)
				w2End := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
				w3Start := time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC)
				w3End := time.Date(2024, time.September, 8, 0, 0, 0, 0, time.UTC)

				periods = []entities.Period{
					entities.NewPeriod(1, w1Start, w1End, pkgtime.PeriodTypeWeek),
					entities.NewPeriod(2, w2Start, w2End, pkgtime.PeriodTypeWeek),
					entities.NewPeriod(3, w3Start, w3End, pkgtime.PeriodTypeWeek),
				}

				ratings = []entities.Rating{
					entities.NewRating(1, 3, 1, 1, w2Date),
					entities.NewRating(2, 4, 2, 1, w1End),
					entities.NewRating(3, 5, 1, 2, w3End),
				}

				scoreDeltas = service.Calculate(ratings, periods)
			})

			It("should includes all periods", func() {
				Expect(scoreDeltas).To(HaveLen(3))
			})

			It("scores should be correct", func() {
				Expect(scoreDeltas[0].Delta).To(Equal(float32(100)))
				Expect(scoreDeltas[1].Delta).To(Equal(float32(-5.956111431121826)))
				Expect(scoreDeltas[2].Delta).To(Equal(float32(66.66665649414062)))
			})
		})

		When("calculating score delta between months periods", func() {
			var scoreDeltas []vo.ScoreDelta
			BeforeEach(func() {
				m1Start := time.Date(2024, time.August, 25, 0, 0, 0, 0, time.UTC)
				m1End := time.Date(2024, time.August, 31, 0, 0, 0, 0, time.UTC)
				m2Start := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
				m2Date := time.Date(2024, time.September, 28, 0, 0, 0, 0, time.UTC)
				m2End := time.Date(2024, time.September, 30, 0, 0, 0, 0, time.UTC)
				m3Start := time.Date(2024, time.October, 1, 0, 0, 0, 0, time.UTC)
				m3End := time.Date(2024, time.October, 31, 0, 0, 0, 0, time.UTC)

				periods = []entities.Period{
					entities.NewPeriod(1, m1Start, m1End, pkgtime.PeriodTypeMonth),
					entities.NewPeriod(2, m2Start, m2End, pkgtime.PeriodTypeMonth),
					entities.NewPeriod(3, m3Start, m3End, pkgtime.PeriodTypeMonth),
				}

				ratings = []entities.Rating{
					entities.NewRating(1, 3, 1, 1, m1Start),
					entities.NewRating(2, 4, 2, 1, m2Date),
					entities.NewRating(3, 5, 1, 2, m3End),
				}

				scoreDeltas = service.Calculate(ratings, periods)
			})

			It("should includes all periods", func() {
				Expect(scoreDeltas).To(HaveLen(3))
			})

			It("scores should be correct", func() {
				Expect(scoreDeltas[0].Delta).To(Equal(float32(100)))
				Expect(scoreDeltas[1].Delta).To(Equal(float32(6.33333158493042)))
				Expect(scoreDeltas[2].Delta).To(Equal(float32(56.73980712890625)))
			})
		})

		When("calculating score delta between year periods", func() {
			var scoreDeltas []vo.ScoreDelta
			BeforeEach(func() {
				y1Start := time.Date(2019, time.August, 25, 0, 0, 0, 0, time.UTC)
				y1End := time.Date(2019, time.December, 31, 0, 0, 0, 0, time.UTC)
				y2Start := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
				y2Date := time.Date(2020, time.May, 28, 0, 0, 0, 0, time.UTC)
				y2End := time.Date(2020, time.December, 30, 0, 0, 0, 0, time.UTC)
				y3Start := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
				y3End := time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)

				periods = []entities.Period{
					entities.NewPeriod(1, y1Start, y1End, pkgtime.PeriodTypeYear),
					entities.NewPeriod(2, y2Start, y2End, pkgtime.PeriodTypeYear),
					entities.NewPeriod(3, y3Start, y3End, pkgtime.PeriodTypeYear),
				}

				ratings = []entities.Rating{
					entities.NewRating(1, 3, 1, 1, y1Start),
					entities.NewRating(2, 4, 2, 1, y2Date),
					entities.NewRating(3, 5, 1, 2, y3End),
				}

				scoreDeltas = service.Calculate(ratings, periods)
			})

			It("should includes all periods", func() {
				Expect(scoreDeltas).To(HaveLen(3))
			})

			It("scores should be correct", func() {
				Expect(scoreDeltas[0].Delta).To(Equal(float32(100)))
				Expect(scoreDeltas[1].Delta).To(Equal(float32(6.33333158493042)))
				Expect(scoreDeltas[2].Delta).To(Equal(float32(56.73980712890625)))
			})
		})
	})
})
