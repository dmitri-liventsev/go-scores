package controllers

import (
	"context"
	getbyperiods "go-scores/gen/get_by_periods"
	"gorm.io/gorm"

	"goa.design/clue/log"
)

// get by periods service example implementation.
// The example methods log the requests and return zero values.
type getByPeriodssrvc struct{}

// NewGetByPeriods returns the get by periods service implementation.
func NewGetByPeriods(db *gorm.DB) getbyperiods.Service {
	return &getByPeriodssrvc{}
}

// Get the score change from a selected period over the previous period.
func (s *getByPeriodssrvc) GetByPeriods(ctx context.Context, p *getbyperiods.GetByPeriodsPayload) (res *getbyperiods.GetByPeriodsResult, err error) {
	res = &getbyperiods.GetByPeriodsResult{}
	log.Printf(ctx, "getByPeriods.getByPeriods")
	return
}
