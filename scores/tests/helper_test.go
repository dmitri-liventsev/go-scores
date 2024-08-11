package tests

import "time"

func getNumberOfWeeks() int {
	today := time.Now()
	pastDate := time.Date(2019, 2, 17, 0, 0, 0, 0, time.UTC)
	duration := today.Sub(pastDate)
	weeks := int(duration.Hours() / (24 * 7))

	return weeks
}

func getNumberOfMonths() int {
	today := time.Now()
	pastDate := time.Date(2019, 2, 17, 0, 0, 0, 0, time.UTC)
	months := int((today.Year()-pastDate.Year())*12 + int(today.Month()) - int(pastDate.Month()))

	return months + 1
}
