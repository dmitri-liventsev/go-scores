package scores

import (
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
	"time"
)

// GetChangesByPeriods returns score deltas by period
type GetChangesByPeriods struct {
	RatingRepo   repositories.Ratings
	CategoryRepo repositories.Categories
}

// Execute returns score deltas by period.
func (q *GetChangesByPeriods) Execute(periodType string) ([]vo.ScoreDelta, []entities.Period, error) {
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

// NewGetChangesByPeriods returns new GetChangesByPeriods query instance.
func NewGetChangesByPeriods(ratingRepo repositories.Ratings, categoryRepo repositories.Categories) GetChangesByPeriods {
	return GetChangesByPeriods{
		RatingRepo:   ratingRepo,
		CategoryRepo: categoryRepo,
	}
}
