package services

import (
	"go-scores/scores/internal/domain/entities"
	vo "go-scores/scores/internal/domain/value-objects"
	"sort"
)

// TicketScores aggregates scores by tickets.
type TicketScores struct {
	categories []entities.Category
	calculator ScoreCalculator
}

// GetScores aggregates scores by tickets.
func (t TicketScores) GetScores(ratings []entities.Rating) []vo.TicketScore {
	ticketCategoryMap := t.getTicketCategoryMap(ratings)

	ticketIDs := make([]vo.TicketID, 0, len(ticketCategoryMap))
	for k := range ticketCategoryMap {
		ticketIDs = append(ticketIDs, k)
	}
	sort.Slice(ticketIDs, func(i, j int) bool {
		return ticketIDs[i] < ticketIDs[j]
	})

	result := make([]vo.TicketScore, len(ticketCategoryMap))
	for i, ticketID := range ticketIDs {
		result[i].TicketID = ticketID
		result[i].Categories = t.calculateCategoryScores(ticketCategoryMap[ticketID])
	}

	return result
}

func (t TicketScores) getTicketCategoryMap(ratings []entities.Rating) map[vo.TicketID]map[vo.CategoryID][]int {
	result := make(map[vo.TicketID]map[vo.CategoryID][]int)

	for _, rating := range ratings {
		if result[rating.TicketID] == nil {
			result[rating.TicketID] = make(map[vo.CategoryID][]int)
			for _, c := range t.categories {
				result[rating.TicketID][c.ID] = []int{}
			}
		}

		result[rating.TicketID][rating.CategoryID] = append(result[rating.TicketID][rating.CategoryID], rating.Value)
	}

	return result
}

func (t TicketScores) calculateCategoryScores(catScoresMap map[vo.CategoryID][]int) []vo.CategoryScores {
	categoryIds := make([]vo.CategoryID, 0, len(catScoresMap))
	for k := range catScoresMap {
		categoryIds = append(categoryIds, k)
	}
	sort.Slice(categoryIds, func(i, j int) bool {
		return categoryIds[i] < categoryIds[j]
	})

	result := make([]vo.CategoryScores, len(categoryIds))
	for i, categoryID := range categoryIds {
		result[i] = vo.CategoryScores{
			CategoryID: categoryID,
			Score:      t.calculator.CalculateRatingsScore(catScoresMap[categoryID], categoryID),
		}
	}

	return result
}

// NewTicketScores returns new TicketScores service instance.
func NewTicketScores(categories []entities.Category) TicketScores {
	return TicketScores{
		categories: categories,
		calculator: NewScoreCalculator(categories),
	}
}
