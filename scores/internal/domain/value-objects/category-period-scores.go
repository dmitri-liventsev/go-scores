package vo

// CategoryPeriodScores aggregation of scores by category and periods.
type CategoryPeriodScores struct {
	CategoryID   CategoryID
	NumOfReviews int
	Periods      []ScorePeriod
	TotalScore   *float32
}
