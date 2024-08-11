package tests

import "time"

func getNumberOfweeks() int {
	today := time.Now()
	pastDate := time.Date(2019, 2, 17, 0, 0, 0, 0, time.UTC)
	duration := today.Sub(pastDate)
	weeks := int(duration.Hours() / (24 * 7))

	return weeks
}
