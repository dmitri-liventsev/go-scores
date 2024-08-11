package time

import (
	"fmt"
	"time"
)

// DaysBetween returns number of days between two dates.
func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}

// GetPeriodInvariant Get period invariant.
func GetPeriodInvariant(date time.Time, periodType string) string {
	switch periodType {
	case PeriodTypeDay:
		return date.Format("2006-01-02")
	case PeriodTypeWeek:
		year, week := date.ISOWeek()
		return fmt.Sprintf("%d-%02d", year, week)
	case PeriodTypeMonth:
		return date.Format("2006-01")
	case PeriodTypeYear:
		return date.Format("2006")
	}

	return ""
}

// IsDateBetween check at date are between two dates.
func IsDateBetween(date, start, end time.Time) bool {
	return (date.After(start) || date.Equal(start)) && (date.Before(end) || date.Equal(end))
}

// DaysUntilEndOfMonth returns number of days until end of the month.
func DaysUntilEndOfMonth(t time.Time) int {
	year, month, _ := t.Date()
	firstDayNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())
	lastDayOfMonth := firstDayNextMonth.AddDate(0, 0, -1)

	return DaysBetween(t, lastDayOfMonth)
}

// DaysUntilEndOfYear returns number of days until end of the month.
func DaysUntilEndOfYear(t time.Time) int {
	year, _, _ := t.Date()
	firstDayNextYear := time.Date(year+1, 1, 1, 0, 0, 0, 0, t.Location())
	lastDayOfYear := firstDayNextYear.AddDate(0, 0, -1)

	return DaysBetween(t, lastDayOfYear)
}
