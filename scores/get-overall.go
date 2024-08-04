package scores

import (
	"go-scores/scores/internal/domain/repositories"
	"go-scores/scores/internal/domain/services"
	"time"
)

type GetOverall struct {
	RatingRepo   repositories.Ratings
	CategoryRepo repositories.Categories
}

// Execute returns overall scores by period.
func (q *GetOverall) Execute(from time.Time, to time.Time) (float32, error) {
	ratings, err := q.RatingRepo.FindByPeriod(from, to)
	if err != nil {
		return 0, err
	}
	categories, err := q.CategoryRepo.FindAll()
	if err != nil {
		return 0, err
	}

	service := services.NewOverallScore(categories)
	score := service.GetScore(ratings)

	return score, nil
}

// NewGetOverall returns new GetOverall query instance.
func NewGetOverall(ratingRepo repositories.Ratings, categoryRepo repositories.Categories) GetOverall {
	return GetOverall{
		RatingRepo:   ratingRepo,
		CategoryRepo: categoryRepo,
	}
}
