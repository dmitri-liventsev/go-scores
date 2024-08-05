package services

import (
	"go-scores/scores/internal/domain/entities"
	vo "go-scores/scores/internal/domain/value-objects"
)

// OverallScore calculates overall score for ratings list.
type OverallScore struct {
	categories []entities.Category
	calculator ScoreCalculator
}

// GetScore calculates overall score for ratings list.
func (o OverallScore) GetScore(ratings []entities.Rating) float32 {
	categoryMap := o.getCategoryMap(ratings)
	if len(categoryMap) == 0 {
		return 0
	}

	sum := float32(0)

	for categoryID, ratings := range categoryMap {
		score := o.calculator.CalculateRatingsScore(ratings, categoryID)
		if score != nil {
			sum += *score
		}
	}

	return sum / float32(len(categoryMap))
}

func (o OverallScore) getCategoryMap(ratings []entities.Rating) map[vo.CategoryID][]int {
	result := make(map[vo.CategoryID][]int)

	for _, rating := range ratings {
		result[rating.CategoryID] = append(result[rating.CategoryID], rating.Value)
	}

	return result
}

// NewOverallScore returns new OverallScore instance.
func NewOverallScore(categories []entities.Category) OverallScore {
	return OverallScore{
		categories: categories,
		calculator: NewScoreCalculator(categories),
	}
}
