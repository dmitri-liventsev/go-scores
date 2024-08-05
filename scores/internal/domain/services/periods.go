package services

import (
	"go-scores/scores/internal/domain/entities"
	pkgtime "go-scores/scores/pkg/time"
	"time"
)

// DividePeriods function to get periods between two dates. Period type will be determined automatically daily or weekly.
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

// DivideByType function to get periods between two dates by specific type.
func DivideByType(start, end time.Time, periodType string) []entities.Period {
	switch periodType {
	case pkgtime.PeriodTypeDay:
		return divideByDays(start, end)
	case pkgtime.PeriodTypeWeek:
		return divideByWeeks(start, end)
	case pkgtime.PeriodTypeMonth:
		return divideByMonth(start, end)
	case pkgtime.PeriodTypeYear:
		return divideByYear(start, end)
	}

	return nil
}

func divideByYear(start, end time.Time) []entities.Period {
	var periods []entities.Period
	var endOfTheYear time.Time

	startOfTheYear := start
	daysToEndOfTheYear := pkgtime.DaysUntilEndOfYear(start)

	id := 1
	for {
		endOfTheYear = startOfTheYear.AddDate(0, 0, daysToEndOfTheYear)
		if endOfTheYear.After(end) {
			endOfTheYear = end
		}
		periods = append(periods, entities.NewPeriod(id, startOfTheYear, endOfTheYear, pkgtime.PeriodTypeYear))
		startOfTheYear = endOfTheYear.AddDate(0, 0, 1)

		if endOfTheYear.Equal(end) {
			break
		}

		daysToEndOfTheYear = pkgtime.DaysUntilEndOfYear(startOfTheYear)
		id++
	}

	return periods
}

func divideByMonth(start, end time.Time) []entities.Period {
	var periods []entities.Period
	var endOfTheMonth time.Time

	startOfTheMonth := start
	daysToEndOfTheMonth := pkgtime.DaysUntilEndOfMonth(start)

	id := 1
	for {
		endOfTheMonth = startOfTheMonth.AddDate(0, 0, daysToEndOfTheMonth)
		if endOfTheMonth.After(end) {
			endOfTheMonth = end
		}
		periods = append(periods, entities.NewPeriod(id, startOfTheMonth, endOfTheMonth, pkgtime.PeriodTypeMonth))
		startOfTheMonth = endOfTheMonth.AddDate(0, 0, 1)

		if endOfTheMonth.Equal(end) {
			break
		}

		daysToEndOfTheMonth = pkgtime.DaysUntilEndOfMonth(startOfTheMonth)
		id++
	}

	return periods
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
