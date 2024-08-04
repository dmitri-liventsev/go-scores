package scores

import (
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	"go-scores/scores/internal/domain/services"
	vo "go-scores/scores/internal/domain/value-objects"
	"time"
)

type GetByTickets struct {
	RatingRepo   repositories.Ratings
	CategoryRepo repositories.Categories
}

// Execute returns scores by categories.
func (q *GetByTickets) Execute(start time.Time, end time.Time) ([]vo.TicketScore, []entities.Category, error) {
	ratings, err := q.RatingRepo.GetByPeriod(start, end)
	if err != nil {
		return nil, nil, err
	}
	categories, err := q.CategoryRepo.FindAll()
	if err != nil {
		return nil, nil, err
	}

	service := services.NewTicketScores(categories)
	ticketScores := service.GetScores(ratings)

	return ticketScores, categories, nil
}

// NewGetByTickets returns new GetByTickets query instance.
func NewGetByTickets(ratingRepo repositories.Ratings, categoryRepo repositories.Categories) GetByTickets {
	return GetByTickets{
		RatingRepo:   ratingRepo,
		CategoryRepo: categoryRepo,
	}
}
