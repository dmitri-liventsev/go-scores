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

var _ = Describe("score calculating", func() {
	Context("given a daily periods and ratings list", func() {
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

		date1 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2024, time.August, 2, 0, 0, 0, 0, time.UTC)
		date3 := time.Date(2024, time.August, 3, 0, 0, 0, 0, time.UTC)

		periods := []entities.Period{
			entities.NewPeriod(1, date1, date1, pkgtime.PeriodTypeDay),
			entities.NewPeriod(2, date2, date2, pkgtime.PeriodTypeDay),
			entities.NewPeriod(3, date3, date3, pkgtime.PeriodTypeDay),
		}

		ratings := []entities.Rating{
			entities.NewRating(1, 0, 1, 1, date1),
			entities.NewRating(2, 2, 2, 1, date1),
			entities.NewRating(3, 1, 1, 2, date2),
		}

		When("calculating scores", func() {
			var categoryScores []vo.CategoryScore
			categoryScoresService := services.NewCategoryScores(categories)

			BeforeEach(func() {
				categoryScores = categoryScoresService.Calculate(ratings, periods)
			})

			It("score should includes each categories", func() {
				Expect(categoryScores).To(HaveLen(2))
				Expect(categoryScores[0].CategoryID.Value()).To(Equal(1))
				Expect(categoryScores[1].CategoryID.Value()).To(Equal(2))
			})

			It("score should includes correct NumOfReviews", func() {
				Expect(categoryScores[0].NumOfReviews).To(Equal(2))
				Expect(categoryScores[1].NumOfReviews).To(Equal(1))
			})

			It("score should includes all periods for each categories", func() {
				Expect(categoryScores[0].Periods).To(HaveLen(3))
				Expect(categoryScores[1].Periods).To(HaveLen(3))
			})

			It("score should includes periods are correct and sorted", func() {
				Expect(categoryScores[0].Periods[0].PeriodID.ToInt()).To(Equal(1))
				Expect(*categoryScores[0].Periods[0].Score).To(Equal(float32(0)))
				Expect(categoryScores[0].Periods[1].PeriodID.ToInt()).To(Equal(2))
				Expect(*categoryScores[0].Periods[1].Score).To(Equal(float32(20)))
				Expect(categoryScores[0].Periods[2].PeriodID.ToInt()).To(Equal(3))
				Expect(categoryScores[0].Periods[2].Score).To(BeNil())

				Expect(categoryScores[1].Periods[0].PeriodID.ToInt()).To(Equal(1))
				Expect(*categoryScores[1].Periods[0].Score).To(Equal(float32(31.900001525878906)))
				Expect(categoryScores[1].Periods[1].PeriodID.ToInt()).To(Equal(2))
				Expect(categoryScores[1].Periods[1].Score).To(BeNil())
				Expect(categoryScores[1].Periods[2].PeriodID.ToInt()).To(Equal(3))
				Expect(categoryScores[1].Periods[2].Score).To(BeNil())
			})

			It("score should includes correct NumOfReviews", func() {
				Expect(*categoryScores[0].TotalScore).To(Equal(float32(6.666666507720947)))
				Expect(*categoryScores[1].TotalScore).To(Equal(float32(10.633334159851074)))
			})
		})
	})

	Context("given a weekly periods and ratings list", func() {
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

		w1Start := time.Date(2024, time.August, 25, 0, 0, 0, 0, time.UTC)
		w1End := time.Date(2024, time.August, 25, 0, 0, 0, 0, time.UTC)
		w2Start := time.Date(2024, time.August, 26, 0, 0, 0, 0, time.UTC)
		w2Date := time.Date(2024, time.August, 27, 0, 0, 0, 0, time.UTC)
		w2End := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
		w3Start := time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC)
		w3End := time.Date(2024, time.September, 8, 0, 0, 0, 0, time.UTC)

		periods := []entities.Period{
			entities.NewPeriod(1, w1Start, w1End, pkgtime.PeriodTypeWeek),
			entities.NewPeriod(2, w2Start, w2End, pkgtime.PeriodTypeWeek),
			entities.NewPeriod(3, w3Start, w3End, pkgtime.PeriodTypeWeek),
		}

		ratings := []entities.Rating{
			entities.NewRating(1, 0, 1, 1, w1Start),
			entities.NewRating(2, 2, 2, 1, w2Start),
			entities.NewRating(3, 1, 1, 2, w2Date),
			entities.NewRating(3, 1, 1, 2, w2End),
		}

		When("calculating scores", func() {
			var categoryScores []vo.CategoryScore
			categoryScoresService := services.NewCategoryScores(categories)

			BeforeEach(func() {
				categoryScores = categoryScoresService.Calculate(ratings, periods)
			})

			It("score should includes each categories", func() {
				Expect(categoryScores).To(HaveLen(2))
				Expect(categoryScores[0].CategoryID.Value()).To(Equal(1))
				Expect(categoryScores[1].CategoryID.Value()).To(Equal(2))
			})

			It("score should includes correct NumOfReviews", func() {
				Expect(categoryScores[0].NumOfReviews).To(Equal(3))
				Expect(categoryScores[1].NumOfReviews).To(Equal(1))
			})

			It("score should includes all periods for each categories", func() {
				Expect(categoryScores[0].Periods).To(HaveLen(3))
				Expect(categoryScores[1].Periods).To(HaveLen(3))
			})

			It("score should includes periods are correct and sorted", func() {
				Expect(categoryScores[0].Periods[0].PeriodID.ToInt()).To(Equal(1))
				Expect(*categoryScores[0].Periods[0].Score).To(Equal(float32(0)))
				Expect(categoryScores[0].Periods[1].PeriodID.ToInt()).To(Equal(2))
				Expect(*categoryScores[0].Periods[1].Score).To(Equal(float32(20)))
				Expect(categoryScores[0].Periods[2].PeriodID.ToInt()).To(Equal(3))
				Expect(categoryScores[0].Periods[2].Score).To(BeNil())

				Expect(categoryScores[1].Periods[0].PeriodID.ToInt()).To(Equal(1))
				Expect(categoryScores[1].Periods[0].Score).To(BeNil())
				Expect(categoryScores[1].Periods[1].PeriodID.ToInt()).To(Equal(2))
				Expect(*categoryScores[1].Periods[1].Score).To(Equal(float32(31.900001525878906)))
				Expect(categoryScores[1].Periods[2].PeriodID.ToInt()).To(Equal(3))
				Expect(categoryScores[1].Periods[2].Score).To(BeNil())
			})

			It("score should includes correct NumOfReviews", func() {
				Expect(*categoryScores[0].TotalScore).To(Equal(float32(6.666666507720947)))
				Expect(*categoryScores[1].TotalScore).To(Equal(float32(10.633334159851074)))
			})
		})
	})
})
