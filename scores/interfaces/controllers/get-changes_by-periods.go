package controllers

import (
	"context"
	getbyperiods "go-scores/gen/get_changes_by_periods"
	"go-scores/scores"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	vo "go-scores/scores/internal/domain/value-objects"
	"gorm.io/gorm"
)

// get by periods service example implementation.
type getChangesByPeriodssrvc struct {
	query scores.GetByPeriods
}

// GetChangesByPeriods Get the score change from a selected period over the previous period.
func (s *getChangesByPeriodssrvc) GetChangesByPeriods(_ context.Context, p *getbyperiods.GetChangesByPeriodsPayload) (*getbyperiods.GetChangesByPeriodsResult, error) {
	scoreDeltas, periods, err := s.query.Execute(p.Period)
	if err != nil {
		return nil, err
	}

	return &getbyperiods.GetChangesByPeriodsResult{
		Meta: s.transformMeta(periods),
		Data: s.transformData(scoreDeltas),
	}, nil
}

func (s *getChangesByPeriodssrvc) transformMeta(periods []entities.Period) *getbyperiods.PeriodsMeta {
	mPeriod := make([]*getbyperiods.Period, len(periods))
	for i, p := range periods {
		start := p.Start.Unix()
		end := p.End.Unix()
		pID := p.ID.ToInt()
		mPeriod[i] = &getbyperiods.Period{
			ID:    &pID,
			Start: &start,
			End:   &end,
		}
	}

	return &getbyperiods.PeriodsMeta{
		Periods: mPeriod,
	}
}

func (s *getChangesByPeriodssrvc) transformData(scoreDeltas []vo.ScoreDelta) []*getbyperiods.PeriodScoreDelta {
	scoreDiff := make([]*getbyperiods.PeriodScoreDelta, len(scoreDeltas))
	for i, p := range scoreDeltas {
		pID := p.PeriodID.ToInt()
		score := p.Delta

		scoreDiff[i] = &getbyperiods.PeriodScoreDelta{
			PeriodID:   &pID,
			ScoreDelta: &score,
		}
	}

	return scoreDiff
}

// NewGetByPeriods returns the get by periods service implementation.
func NewGetByPeriods(db *gorm.DB) getbyperiods.Service {
	return &getChangesByPeriodssrvc{
		query: scores.NewGetByPeriods(
			repositories.NewRatings(db),
			repositories.NewCategories(db),
		),
	}
}
