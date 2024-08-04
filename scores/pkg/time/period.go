package time

import (
	"fmt"
	"time"
)

// DaysBetween returns number of days between two dates
func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days + 1
}

func GetPeriodInvariant(date time.Time, periodType string) string {
	if periodType == PeriodTypeDay {

		return date.Format("2006-01-02")
	}

	year, week := date.ISOWeek()
	return fmt.Sprintf("%d-%02d", year, week)
}

func IsDateBetween(date, start, end time.Time) bool {
	return (date.After(start) || date.Equal(start)) && (date.Before(end) || date.Equal(end))
}
