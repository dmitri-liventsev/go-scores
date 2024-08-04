package services

import (
	"go-scores/scores/internal/domain/entities"
	pkgtime "go-scores/scores/pkg/time"
	"time"
)

// DividePeriods function to get periods between two dates.
func DividePeriods(start, end time.Time) []entities.Period {
	if end.Before(start) {
		return nil
	}
	daysBetween := pkgtime.DaysBetween(start, end)
	if daysBetween > 30 {
		return divideByWeeks(start, end)
	}

	return divideByDays(start, end)
}

// divideByDays function to get an array of days between two dates.
func divideByDays(start, end time.Time) []entities.Period {
	var periods []entities.Period
	id := 1
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		periods = append(periods, entities.NewPeriod(id, d, d, pkgtime.PeriodTypeDay))
		id++
	}

	return periods
}

// divideByWeeks function to get an array of weeks between two dates.
func divideByWeeks(start, end time.Time) []entities.Period {
	var periods []entities.Period
	var nextSunday time.Time

	startWeek := start
	daysUntilSunday := (7 - int(start.Weekday())) % 7

	id := 1
	for {
		nextSunday = startWeek.AddDate(0, 0, daysUntilSunday)
		if nextSunday.After(end) {
			nextSunday = end
		}
		periods = append(periods, entities.NewPeriod(id, startWeek, nextSunday, pkgtime.PeriodTypeWeek))
		startWeek = nextSunday.AddDate(0, 0, 1)

		if nextSunday.Equal(end) {
			break
		}

		daysUntilSunday = 6
		id++
	}

	return periods
}
