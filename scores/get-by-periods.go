package scores

import (
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
	"time"
)

// GetByPeriods returns score deltas by period
type GetByPeriods struct {
	RatingRepo   repositories.Ratings
	CategoryRepo repositories.Categories
}

// Execute returns score deltas by period.
func (q *GetByPeriods) Execute(periodType string) ([]vo.ScoreDelta, []entities.Period, error) {
	ratings, err := q.RatingRepo.FindAll()
	if err != nil {
		return nil, nil, err
	}

	if len(ratings) == 0 {
		return []vo.ScoreDelta{}, []entities.Period{}, nil
	}

	categories, err := q.CategoryRepo.FindAll()
	if err != nil {
		return nil, nil, err
	}

	start := ratings[0].CreatedAt
	end := time.Now()
	periods := services.DevideByType(start, end, periodType)

	service := services.NewScoreEvolution(categories)
	scoreDeltas := service.Calculate(ratings, periods)

	return scoreDeltas, periods, nil
}

// NewGetByPeriods returns new GetByPeriods query instance.
func NewGetByPeriods(ratingRepo repositories.Ratings, categoryRepo repositories.Categories) GetByPeriods {
	return GetByPeriods{
		RatingRepo:   ratingRepo,
		CategoryRepo: categoryRepo,
	}
}
