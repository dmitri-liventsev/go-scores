package scores

import (
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
	"time"
)

// GetByCategories query to get scores grouped by category.
type GetByCategories struct {
	RatingRepo   repositories.Ratings
	CategoryRepo repositories.Categories
}

// Execute returns scores by categories.
func (q *GetByCategories) Execute(from time.Time, to time.Time) ([]vo.CategoryPeriodScores, []entities.Period, []entities.Category, error) {
	periods := services.DividePeriods(from, to)
	categories, err := q.CategoryRepo.FindAll()
	if err != nil {
		return nil, nil, nil, err
	}

	ratings, err := q.RatingRepo.FindByPeriod(from, to)
	if err != nil {
		return nil, nil, nil, err
	}

	catService := services.NewCategoryScores(categories)
	catScores := catService.GetScores(ratings, periods)

	return catScores, periods, categories, nil
}

// NewGetByCategories returns new GetByCategories query instance.
func NewGetByCategories(ratingRepo repositories.Ratings, categoryRepo repositories.Categories) GetByCategories {
	return GetByCategories{
		RatingRepo:   ratingRepo,
		CategoryRepo: categoryRepo,
	}
}
