// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets endpoints
//
// Command:
// $ goa gen go-scores/design

package getbytickets

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "get by tickets" service endpoints.
type Endpoints struct {
	GetAggregatedScoresByTicket goa.Endpoint
}

// NewEndpoints wraps the methods of the "get by tickets" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetAggregatedScoresByTicket: NewGetAggregatedScoresByTicketEndpoint(s),
	}
}

// Use applies the given middleware to all the "get by tickets" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAggregatedScoresByTicket = m(e.GetAggregatedScoresByTicket)
}

// NewGetAggregatedScoresByTicketEndpoint returns an endpoint function that
// calls the method "getAggregatedScoresByTicket" of service "get by tickets".
func NewGetAggregatedScoresByTicketEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetAggregatedScoresByTicketPayload)
		return s.GetAggregatedScoresByTicket(ctx, p)
	}
}