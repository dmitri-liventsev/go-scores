package vo

type CategoryPeriodScores struct {
	CategoryID   CategoryID
	NumOfReviews int
	Periods      []ScorePeriod
	TotalScore   *float32
}
