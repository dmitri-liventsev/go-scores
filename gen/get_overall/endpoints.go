// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get overall endpoints
//
// Command:
// $ goa gen go-scores/design

package getoverall

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "get overall" service endpoints.
type Endpoints struct {
	GetOverallScore goa.Endpoint
}

// NewEndpoints wraps the methods of the "get overall" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetOverallScore: NewGetOverallScoreEndpoint(s),
	}
}

// Use applies the given middleware to all the "get overall" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetOverallScore = m(e.GetOverallScore)
}

// NewGetOverallScoreEndpoint returns an endpoint function that calls the
// method "getOverallScore" of service "get overall".
func NewGetOverallScoreEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetOverallScorePayload)
		return s.GetOverallScore(ctx, p)
	}
}
