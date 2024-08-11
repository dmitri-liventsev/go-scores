package controllers

import (
	"context"
	"fmt"
	getoverall "go-scores/gen/get_overall"
	"go-scores/scores"
	"go-scores/scores/internal/domain/repositories"
	"gorm.io/gorm"
	"time"
)

// get overall service example implementation.
// The example methods log the requests and return zero values.
type getOverallsrvc struct {
	query scores.GetOverall
}

// GetOverallScore Get the overall aggregate score for a specified period.
func (s *getOverallsrvc) GetOverallScore(ctx context.Context, p *getoverall.GetOverallScorePayload) (res float32, err error) {
	layout := "2006-01-02"

	from, err := time.Parse(layout, p.From)
	if err != nil {
		return 0, err
	}

	to, err := time.Parse(layout, p.To)
	if err != nil {
		return 0, err
	}

	if from.After(to) {
		return 0, fmt.Errorf("start date cant be before end date")
	}

	score, err := s.query.Execute(from, to)
	if err != nil {
		return 0, err
	}

	return score, nil
}

// NewGetOverall returns the get overall service implementation.
func NewGetOverall(db *gorm.DB) getoverall.Service {
	return &getOverallsrvc{
		query: scores.NewGetOverall(repositories.NewRatings(db), repositories.NewCategories(db)),
	}
}
