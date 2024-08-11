package controllers

import (
	"context"
	"fmt"
	getbytickets "go-scores/gen/get_by_tickets"
	"go-scores/scores"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	vo "go-scores/scores/internal/domain/value-objects"
	"gorm.io/gorm"
	"time"
)

// get by tickets service example implementation.
type getByTicketssrvc struct {
	query scores.GetByTickets
}

// GetAggregatedScoresByTicket Get aggregate scores by ticket for a specified period.
func (s *getByTicketssrvc) GetAggregatedScoresByTicket(_ context.Context, p *getbytickets.GetAggregatedScoresByTicketPayload) (*getbytickets.GetAggregatedScoresByTicketResult, error) {
	layout := "2006-01-02"

	from, err := time.Parse(layout, p.From)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse(layout, p.To)
	if err != nil {
		return nil, err
	}

	if from.After(to) {
		return nil, fmt.Errorf("start date cant be before end date")
	}

	ticketScores, categories, err := s.query.Execute(from, to)
	if err != nil {
		return nil, err
	}

	return &getbytickets.GetAggregatedScoresByTicketResult{
		Meta: s.transformMeta(categories),
		Data: s.transformData(ticketScores),
	}, nil
}

func (s getByTicketssrvc) transformMeta(categories []entities.Category) *getbytickets.TicketMeta {
	mCategories := make([]*getbytickets.Category, len(categories))
	for i, c := range categories {
		catId := c.ID.ToInt()

		mCategories[i] = &getbytickets.Category{
			ID:   &catId,
			Name: &c.Name,
		}
	}

	return &getbytickets.TicketMeta{
		Categories: mCategories,
	}
}

func (s getByTicketssrvc) transformData(ticketScores []vo.TicketScore) []*getbytickets.TicketScore {
	result := make([]*getbytickets.TicketScore, len(ticketScores))

	for i, t := range ticketScores {
		ticketID := t.TicketID.ToInt()
		result[i] = &getbytickets.TicketScore{
			TicketID:   &ticketID,
			Categories: transformCategoryScores(t.Categories),
		}
	}
	return result
}

func transformCategoryScores(categoryScores []vo.CategoryScores) []*getbytickets.CategoryScoreDetail {
	result := make([]*getbytickets.CategoryScoreDetail, len(categoryScores))

	for i, c := range categoryScores {
		categoryID := c.CategoryID.ToInt()
		result[i] = &getbytickets.CategoryScoreDetail{
			ID:    &categoryID,
			Score: c.Score,
		}
	}

	return result
}

// NewGetByTickets returns the get by tickets service implementation.
func NewGetByTickets(db *gorm.DB) getbytickets.Service {
	return &getByTicketssrvc{
		query: scores.NewGetByTickets(
			repositories.NewRatings(db),
			repositories.NewCategories(db),
		),
	}
}
