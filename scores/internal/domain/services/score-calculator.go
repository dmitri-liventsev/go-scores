package services

import (
	"go-scores/scores/internal/domain/entities"
	vo "go-scores/scores/internal/domain/value-objects"
	"math"
)

// ScoreCalculator calculating the scores based on ratings and category weights.
type ScoreCalculator struct {
	categoryWeights map[vo.CategoryID]float32
}

const MaxScore = 5

// Calculate returns score for one rating based on category weight.
// Y=0.25x^2 + 0.25x + 0.5
func (s ScoreCalculator) Calculate(avg float32, categoryId vo.CategoryID) *float32 {
	var result float32

	result = avg / MaxScore * 100

	weight := s.categoryWeights[categoryId]
	weightCoef := float32(0.25*math.Pow(float64(weight), 2) + float64(0.25*weight) + 0.5)

	result *= weightCoef

	if result > 100 {
		result = 100
	}

	return &result
}

// CalculateRatingsScore returns average score for set of ratings based on category weight.
// NB! All ratings should be from the same category.
func (s ScoreCalculator) CalculateRatingsScore(ratings []int, categoryID vo.CategoryID) *float32 {
	if len(ratings) == 0 {
		return nil
	}

	sumOfRatings := 0
	for _, rating := range ratings {
		sumOfRatings += rating
	}

	avg := float32(sumOfRatings) / float32(len(ratings))

	return s.Calculate(avg, categoryID)
}

func (s ScoreCalculator) min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// NewScoreCalculator returns new new ScoreCalculator instance.
func NewScoreCalculator(categories []entities.Category) ScoreCalculator {
	categoryWeights := make(map[vo.CategoryID]float32)
	for _, category := range categories {
		categoryWeights[category.ID] = category.Weight
	}

	return ScoreCalculator{categoryWeights}
}
