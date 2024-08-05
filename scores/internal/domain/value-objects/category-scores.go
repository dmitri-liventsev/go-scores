package vo

// CategoryScores aggregation of scores by category.
type CategoryScores struct {
	CategoryID CategoryID
	Score      *float32
}
