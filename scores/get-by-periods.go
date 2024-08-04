package scores

import (
	"go-scores/scores/internal/domain/repositories"
	"time"
)

type GetByPeriods struct {
	RatingRepo   repositories.Ratings
	CategoryRepo repositories.Categories
}

// Execute returns scores by period.
func (q *GetByPeriods) Execute(start time.Time, end time.Time) (float32, error) {

	return float32(0), nil
}

// NewGetByPeriods returns new GetByPeriods query instance.
func NewGetByPeriods(ratingRepo repositories.Ratings, categoryRepo repositories.Categories) GetByPeriods {
	return GetByPeriods{
		RatingRepo:   ratingRepo,
		CategoryRepo: categoryRepo,
	}
}
