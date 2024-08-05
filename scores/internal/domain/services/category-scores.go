package services

import (
	"go-scores/scores/internal/domain/entities"
	vo "go-scores/scores/internal/domain/value-objects"
	"go-scores/scores/pkg/time"
	"sort"
)

// CategoryScores aggregates ratings by category.
type CategoryScores struct {
	categories []entities.Category
	calculator ScoreCalculator
}

// GetScores scores by category. NB all periods should have the same type!
func (s CategoryScores) GetScores(ratings []entities.Rating, periods []entities.Period) []vo.CategoryPeriodScores {
	result := make([]vo.CategoryPeriodScores, len(s.categories))
	catPeriodMap := s.makeCategoryPeriodMap(ratings, periods)

	categoryIDs := make([]vo.CategoryID, 0, len(catPeriodMap))
	for k := range catPeriodMap {
		categoryIDs = append(categoryIDs, k)
	}
	sort.Slice(categoryIDs, func(i, j int) bool {
		return categoryIDs[i] < categoryIDs[j]
	})

	for i, categoryID := range categoryIDs {
		periodRatingsMap := catPeriodMap[categoryID]
		p := s.calcScorePeriods(periodRatingsMap, categoryID)
		result[i] = vo.CategoryPeriodScores{
			CategoryID:   categoryID,
			NumOfReviews: s.calcNumOfReviews(periodRatingsMap),
			Periods:      p,
			TotalScore:   s.calcTotalScore(p),
		}
	}

	return result
}

func (s CategoryScores) makeCategoryPeriodMap(ratings []entities.Rating, periods []entities.Period) map[vo.CategoryID]map[vo.PeriodID][]int {
	result := make(map[vo.CategoryID]map[vo.PeriodID][]int)
	for _, category := range s.categories {
		result[category.ID] = make(map[vo.PeriodID][]int)
		for _, period := range periods {
			result[category.ID][period.ID] = []int{}
		}
	}

	categoryMap := make(map[vo.CategoryID]entities.Category)
	periodMap := make(map[string]entities.Period)
	periodType := time.PeriodTypeDay

	for _, category := range s.categories {
		categoryMap[category.ID] = category
	}

	for _, period := range periods {
		periodType = period.Type
		periodMap[period.GetInvariant()] = period
	}

	for _, r := range ratings {
		periodInvariant := time.GetPeriodInvariant(r.CreatedAt, periodType)
		category := categoryMap[r.CategoryID]
		period := periodMap[periodInvariant]

		result[category.ID][period.ID] = append(result[category.ID][period.ID], r.Value)
	}

	return result
}

func (s CategoryScores) calcNumOfReviews(periodRatingsMap map[vo.PeriodID][]int) int {
	totalRatings := 0
	for _, ratings := range periodRatingsMap {
		totalRatings += len(ratings)
	}

	return totalRatings
}

func (s CategoryScores) calcScorePeriods(periodRatingsMap map[vo.PeriodID][]int, categoryID vo.CategoryID) []vo.ScorePeriod {
	periodIDs := make([]vo.PeriodID, 0, len(periodRatingsMap))
	for k := range periodRatingsMap {
		periodIDs = append(periodIDs, k)
	}
	sort.Slice(periodIDs, func(i, j int) bool {
		return periodIDs[i] < periodIDs[j]
	})

	result := make([]vo.ScorePeriod, len(periodRatingsMap))

	for i, periodID := range periodIDs {
		score := s.calculator.CalculateRatingsScore(periodRatingsMap[periodID], categoryID)
		result[i] = vo.ScorePeriod{
			PeriodID: periodID,
			Score:    score,
		}
	}

	return result
}

func (s CategoryScores) calcTotalScore(scorePeriods []vo.ScorePeriod) *float32 {
	if len(scorePeriods) == 0 {
		return nil
	}

	totalSum := float32(0)
	for _, period := range scorePeriods {
		if period.Score != nil {
			totalSum += *period.Score
		}
	}

	totalSum = totalSum / float32(len(scorePeriods))

	return &totalSum
}

// NewCategoryScores returns new CategoryScores service instance.
func NewCategoryScores(categories []entities.Category) CategoryScores {
	return CategoryScores{
		categories: categories,
		calculator: NewScoreCalculator(categories),
	}
}
