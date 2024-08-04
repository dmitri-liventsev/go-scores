package vo

type CategoryScore struct {
	CategoryID   CategoryID
	NumOfReviews int
	Periods      []ScorePeriod
	TotalScore   *float32
}
