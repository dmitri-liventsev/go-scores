package services

import (
	"go-scores/scores/internal/domain/entities"
	vo "go-scores/scores/internal/domain/value-objects"
	"go-scores/scores/pkg/time"
	"sort"
)

// ScoreEvolution calculates score changes between periods.
type ScoreEvolution struct {
	overallScoreCalculator OverallScore
}

// Calculate returns list of score deltas between periods.
func (s ScoreEvolution) Calculate(ratings []entities.Rating, periods []entities.Period) []vo.ScoreDelta {
	ratingsByPeriodMap := s.getRatingsByPeriodMap(ratings, periods)
	scoresByPeriodMap := s.getScoresByPeriodMap(ratingsByPeriodMap, periods)

	return s.calculateEvolution(scoresByPeriodMap)
}

func (s ScoreEvolution) getRatingsByPeriodMap(ratings []entities.Rating, periods []entities.Period) map[vo.PeriodID][]entities.Rating {
	result := make(map[vo.PeriodID][]entities.Rating)
	periodMap := make(map[string]entities.Period)
	periodType := time.PeriodTypeDay

	for _, period := range periods {
		periodType = period.Type
		periodMap[period.GetInvariant()] = period
	}

	for _, r := range ratings {
		periodInvariant := time.GetPeriodInvariant(r.CreatedAt, periodType)
		period := periodMap[periodInvariant]

		result[period.ID] = append(result[period.ID], r)
	}

	return result
}

func (s ScoreEvolution) getScoresByPeriodMap(ratingsByPeriodMap map[vo.PeriodID][]entities.Rating, periods []entities.Period) []vo.ScorePeriod {
	result := make([]vo.ScorePeriod, len(periods))

	for i, period := range periods {
		zero := float32(0)
		result[i] = vo.ScorePeriod{
			PeriodID: period.ID,
			Score:    &zero,
		}
	}

	periodIDs := make([]vo.PeriodID, 0, len(ratingsByPeriodMap))
	for k := range ratingsByPeriodMap {
		periodIDs = append(periodIDs, k)
	}
	sort.Slice(periodIDs, func(i, j int) bool {
		return periodIDs[i] < periodIDs[j]
	})

	for i, periodID := range periodIDs {
		ratings := ratingsByPeriodMap[periodID]
		score := s.overallScoreCalculator.GetScore(ratings)

		result[i] = vo.ScorePeriod{
			PeriodID: periodID,
			Score:    &score,
		}
	}

	return result
}

func (s ScoreEvolution) calculateEvolution(scoresPerPeriod []vo.ScorePeriod) []vo.ScoreDelta {
	if len(scoresPerPeriod) == 0 {
		return []vo.ScoreDelta{}
	}
	result := make([]vo.ScoreDelta, len(scoresPerPeriod))
	result[0] = vo.ScoreDelta{
		PeriodID: scoresPerPeriod[0].PeriodID,
		Delta:    100,
	}

	for i := 1; i < len(scoresPerPeriod); i++ {
		delta := float32(0)
		prev := i - 1
		if scoresPerPeriod[prev].Score != nil && *scoresPerPeriod[prev].Score != 0 {
			currentScore := *scoresPerPeriod[i].Score
			prevScore := *scoresPerPeriod[prev].Score
			delta = ((currentScore - prevScore) / prevScore) * 100
		}

		result[i] = vo.ScoreDelta{
			PeriodID: scoresPerPeriod[i].PeriodID,
			Delta:    delta,
		}
	}

	return result
}

// NewScoreEvolution returns new ScoreEvolution instance.
func NewScoreEvolution(categories []entities.Category) ScoreEvolution {
	return ScoreEvolution{
		overallScoreCalculator: NewOverallScore(categories),
	}
}
