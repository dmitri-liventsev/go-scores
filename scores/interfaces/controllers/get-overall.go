package controllers

import (
	"context"
	getoverall "go-scores/gen/get_overall"
	"gorm.io/gorm"

	"goa.design/clue/log"
)

// get overall service example implementation.
// The example methods log the requests and return zero values.
type getOverallsrvc struct{}

// NewGetOverall returns the get overall service implementation.
func NewGetOverall(db *gorm.DB) getoverall.Service {
	return &getOverallsrvc{}
}

// Get the overall aggregate score for a specified period.
func (s *getOverallsrvc) GetOverallScore(ctx context.Context, p *getoverall.GetOverallScorePayload) (res float32, err error) {
	log.Printf(ctx, "getOverall.getOverallScore")
	return
}
