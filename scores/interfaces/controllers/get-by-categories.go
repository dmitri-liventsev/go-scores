package controllers

import (
	"context"
	getbycategories "go-scores/gen/get_by_categories"
	"go-scores/scores"
	"go-scores/scores/internal/domain/entities"
	"go-scores/scores/internal/domain/repositories"
	vo "go-scores/scores/internal/domain/value-objects"
	"gorm.io/gorm"
	"time"
)

// getByCategories implementation. Separate interface layer from domain
type getByCategories struct {
	query scores.GetByCategories
}

// GetAggregatedScores Get aggregated scores for categories within a specified period.
func (s *getByCategories) GetAggregatedScores(ctx context.Context, p *getbycategories.GetAggregatedScoresPayload) (*getbycategories.GetAggregatedScoresResult, error) {
	layout := "2006-01-02"

	from, err := time.Parse(layout, p.From)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse(layout, p.To)
	if err != nil {
		return nil, err
	}

	catScores, periods, categories, err := s.query.Execute(from, to)
	if err != nil {
		return nil, err
	}

	return transform(catScores, periods, categories), nil
}

func transform(catScores []vo.CategoryScore, periods []entities.Period, categories []entities.Category) *getbycategories.GetAggregatedScoresResult {
	return &getbycategories.GetAggregatedScoresResult{
		Meta: transformMeta(periods, categories),
		Data: transformData(catScores),
	}
}

func transformMeta(periods []entities.Period, categories []entities.Category) *getbycategories.CategoryMeta {
	mPeriod := make([]*getbycategories.Period, len(periods))
	for i, p := range periods {
		start := p.Start.Unix()
		end := p.End.Unix()
		pID := p.ID.ToInt()
		mPeriod[i] = &getbycategories.Period{
			ID:    &pID,
			Start: &start,
			End:   &end,
		}
	}

	mCategory := make([]*getbycategories.Category, len(categories))
	for i, c := range categories {
		cID := c.ID.ToInt()
		mCategory[i] = &getbycategories.Category{
			ID:   &cID,
			Name: &c.Name,
		}
	}

	return &getbycategories.CategoryMeta{
		Periods:    mPeriod,
		Categories: mCategory,
	}
}

func transformData(catScores []vo.CategoryScore) []*getbycategories.CategoryScore {
	dCatScores := make([]*getbycategories.CategoryScore, len(catScores))

	for i, c := range catScores {
		periods := make([]*getbycategories.ScorePeriod, len(c.Periods))
		for j, p := range c.Periods {
			periodID := p.PeriodID.ToInt()
			periods[j] = &getbycategories.ScorePeriod{
				ID:    &periodID,
				Score: p.Score,
			}
		}
		categoryID := c.CategoryID.ToInt()
		dCatScores[i] = &getbycategories.CategoryScore{
			CategoryID:   &categoryID,
			NumOfReviews: &c.NumOfReviews,
			Periods:      periods,
			TotalScore:   c.TotalScore,
		}
	}

	return dCatScores
}

// NewGetByCategories returns the get by categories service implementation.
func NewGetByCategories(db *gorm.DB) getbycategories.Service {
	return &getByCategories{
		query: scores.NewGetByCategories(
			repositories.NewRatings(db),
			repositories.NewCategories(db),
		),
	}
}
