// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by periods client
//
// Command:
// $ goa gen go-scores/design

package getbyperiods

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "get by periods" service client.
type Client struct {
	GetByPeriodsEndpoint goa.Endpoint
}

// NewClient initializes a "get by periods" service client given the endpoints.
func NewClient(getByPeriods goa.Endpoint) *Client {
	return &Client{
		GetByPeriodsEndpoint: getByPeriods,
	}
}

// GetByPeriods calls the "getByPeriods" endpoint of the "get by periods"
// service.
func (c *Client) GetByPeriods(ctx context.Context, p *GetByPeriodsPayload) (res *GetByPeriodsResult, err error) {
	var ires any
	ires, err = c.GetByPeriodsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetByPeriodsResult), nil
}