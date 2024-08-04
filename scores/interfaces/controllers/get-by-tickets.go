package controllers

import (
	"context"
	getbytickets "go-scores/gen/get_by_tickets"
	"gorm.io/gorm"

	"goa.design/clue/log"
)

// get by tickets service example implementation.
// The example methods log the requests and return zero values.
type getByTicketssrvc struct{}

// NewGetByTickets returns the get by tickets service implementation.
func NewGetByTickets(db *gorm.DB) getbytickets.Service {
	return &getByTicketssrvc{}
}

// Get aggregate category scores by ticket for a specified period.
func (s *getByTicketssrvc) GetAggregatedScoresByTicket(ctx context.Context, p *getbytickets.GetAggregatedScoresByTicketPayload) (res *getbytickets.GetAggregatedScoresByTicketResult, err error) {
	res = &getbytickets.GetAggregatedScoresByTicketResult{}
	log.Printf(ctx, "getByTickets.getAggregatedScoresByTicket")
	return
}
