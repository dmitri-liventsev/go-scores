package time_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	pkgtime "go-scores/scores/pkg/time"
	"time"
)

var _ = Describe("calculate days between two days", func() {
	When("two days are equals", func() {
		date1 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)

		It("should return correct length", func() {
			Expect(pkgtime.DaysBetween(date1, date2)).To(Equal(0))
		})
	})

	When("diff are one day", func() {
		date1 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2024, time.August, 2, 0, 0, 0, 0, time.UTC)

		It("should return correct length", func() {
			Expect(pkgtime.DaysBetween(date1, date2)).To(Equal(1))
		})
	})

	When("diff are ten day", func() {
		date1 := time.Date(2024, time.August, 1, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2024, time.August, 10, 0, 0, 0, 0, time.UTC)

		It("should return correct length", func() {
			Expect(pkgtime.DaysBetween(date1, date2)).To(Equal(9))
		})
	})

	When("dates in different months", func() {
		date1 := time.Date(2024, time.August, 31, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)

		It("should return correct length", func() {
			Expect(pkgtime.DaysBetween(date1, date2)).To(Equal(1))
		})
	})

	When("dates in different years", func() {
		date1 := time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

		It("should return correct length", func() {
			Expect(pkgtime.DaysBetween(date1, date2)).To(Equal(1))
		})
	})
})
