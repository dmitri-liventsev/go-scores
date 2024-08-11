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

			It("should be divided by weeks", func() {
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

var _ = Describe("divide period by months", func() {
	Context("given a period with a length of more than 30 days", func() {
		When("dividing periods to smaller periods", func() {
			date1 := time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC)
			date2 := time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC)
			periods := services.DivideByType(date1, date2, pkgtime.PeriodTypeMonth)

			It("should includes all weeks", func() {
				Expect(periods).To(HaveLen(13))
			})

			It("should be divided by months", func() {
				for i := 0; i <= 12; i++ {
					Expect(periods[i].Type).To(Equal(pkgtime.PeriodTypeMonth), fmt.Sprintf("index %d", i))
				}
			})

			It("months in periods are correct", func() {
				expectedStartDate := time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC)
				expectedEndDate := time.Date(2023, time.August, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[0].Start).To(Equal(expectedStartDate))
				Expect(periods[0].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2023, time.September, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2023, time.September, 30, 0, 0, 0, 0, time.UTC)
				Expect(periods[1].Start).To(Equal(expectedStartDate))
				Expect(periods[1].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2023, time.October, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2023, time.October, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[2].Start).To(Equal(expectedStartDate))
				Expect(periods[2].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2023, time.November, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2023, time.November, 30, 0, 0, 0, 0, time.UTC)
				Expect(periods[3].Start).To(Equal(expectedStartDate))
				Expect(periods[3].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[4].Start).To(Equal(expectedStartDate))
				Expect(periods[4].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[5].Start).To(Equal(expectedStartDate))
				Expect(periods[5].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)
				Expect(periods[6].Start).To(Equal(expectedStartDate))
				Expect(periods[6].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.March, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[7].Start).To(Equal(expectedStartDate))
				Expect(periods[7].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.April, 30, 0, 0, 0, 0, time.UTC)
				Expect(periods[8].Start).To(Equal(expectedStartDate))
				Expect(periods[8].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.May, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[9].Start).To(Equal(expectedStartDate))
				Expect(periods[9].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.June, 30, 0, 0, 0, 0, time.UTC)
				Expect(periods[10].Start).To(Equal(expectedStartDate))
				Expect(periods[10].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.July, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[11].Start).To(Equal(expectedStartDate))
				Expect(periods[11].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC)
				Expect(periods[12].Start).To(Equal(expectedStartDate))
				Expect(periods[12].End).To(Equal(expectedEndDate))

			})
		})
	})
})

var _ = Describe("divide period by years", func() {
	Context("given a period with a length of more than one year", func() {
		When("dividing periods to smaller periods", func() {
			date1 := time.Date(2019, time.August, 15, 0, 0, 0, 0, time.UTC)
			date2 := time.Date(2022, time.August, 30, 0, 0, 0, 0, time.UTC)
			periods := services.DivideByType(date1, date2, pkgtime.PeriodTypeYear)

			It("should includes all weeks", func() {
				Expect(periods).To(HaveLen(4))
			})

			It("should be divided by months", func() {
				for i := 0; i <= 3; i++ {
					Expect(periods[i].Type).To(Equal(pkgtime.PeriodTypeYear), fmt.Sprintf("index %d", i))
				}
			})

			It("months in periods are correct", func() {
				expectedStartDate := time.Date(2019, time.August, 15, 0, 0, 0, 0, time.UTC)
				expectedEndDate := time.Date(2019, time.December, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[0].Start).To(Equal(expectedStartDate))
				Expect(periods[0].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[1].Start).To(Equal(expectedStartDate))
				Expect(periods[1].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)
				Expect(periods[2].Start).To(Equal(expectedStartDate))
				Expect(periods[2].End).To(Equal(expectedEndDate))

				expectedStartDate = time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
				expectedEndDate = time.Date(2022, time.August, 30, 0, 0, 0, 0, time.UTC)
				Expect(periods[3].Start).To(Equal(expectedStartDate))
				Expect(periods[3].End).To(Equal(expectedEndDate))
			})
		})
	})
})
