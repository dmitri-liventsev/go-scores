package services_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-scores/scores/internal/domain/services"
	pkgtime "go-scores/scores/pkg/time"
	"time"
)

var _ = Describe("divide period by days", func() {
	Context("given a period with a length of less than 31 days", func() {
		When("dividing periods to smaller periods", func() {
			date1 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
			date2 := time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC)
			periods := services.DividePeriods(date1, date2)

			It("should includes all days", func() {
				Expect(periods).To(HaveLen(30))
			})

			It("should be divided by days", func() {
				for i := 0; i <= 29; i++ {
					Expect(periods[i].Type).To(Equal(pkgtime.PeriodTypeDay), fmt.Sprintf("index %d", i))
				}
			})

			It("dates in periods are correct", func() {
				for i := 0; i <= 29; i++ {
					expectedDate := time.Date(2024, time.August, i+1, 0, 0, 0, 0, time.UTC)
					Expect(periods[i].Start).To(Equal(expectedDate), fmt.Sprintf("index %d", i))
					Expect(periods[i].End).To(Equal(expectedDate), fmt.Sprintf("index %d", i))
				}
			})
		})
	})
})

var _ = Describe("divide period by weeks", func() {
	Context("given a period with a length of more than 30 days", func() {
		When("dividing periods to smaller periods", func() {
			date1 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
			date2 := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
			periods := services.DividePeriods(date1, date2)

			It("should includes all weeks", func() {
				Expect(periods).To(HaveLen(5))
			})

			It("should be divided by days", func() {
				for i := 0; i <= 4; i++ {
					Expect(periods[i].Type).To(Equal(pkgtime.PeriodTypeWeek), fmt.Sprintf("index %d", i))
				}
			})

			It("weeks in periods are correct", func() {
				expectedStartDate := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate := time.Date(2024, time.August, 4, 0, 0, 0, 0, time.UTC)
				Expect(periods[0].Start).To(Equal(expectedStartDate))
				Expect(periods[0].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.August, 5, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.August, 11, 0, 0, 0, 0, time.UTC)
				Expect(periods[1].Start).To(Equal(expectedStartDate))
				Expect(periods[1].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.August, 12, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.August, 18, 0, 0, 0, 0, time.UTC)
				Expect(periods[2].Start).To(Equal(expectedStartDate))
				Expect(periods[2].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.August, 19, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.August, 25, 0, 0, 0, 0, time.UTC)
				Expect(periods[3].Start).To(Equal(expectedStartDate))
				Expect(periods[3].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.August, 26, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
				Expect(periods[4].Start).To(Equal(expectedStartDate))
				Expect(periods[4].End).To(Equal(expectedEndDate))

			})
		})
	})
})
