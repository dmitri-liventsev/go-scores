// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets client
//
// Command:
// $ goa gen go-scores/design

package getbytickets

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "get by tickets" service client.
type Client struct {
	GetAggregatedScoresByTicketEndpoint goa.Endpoint
}

// NewClient initializes a "get by tickets" service client given the endpoints.
func NewClient(getAggregatedScoresByTicket goa.Endpoint) *Client {
	return &Client{
		GetAggregatedScoresByTicketEndpoint: getAggregatedScoresByTicket,
	}
}

// GetAggregatedScoresByTicket calls the "getAggregatedScoresByTicket" endpoint
// of the "get by tickets" service.
func (c *Client) GetAggregatedScoresByTicket(ctx context.Context, p *GetAggregatedScoresByTicketPayload) (res *GetAggregatedScoresByTicketResult, err error) {
	var ires any
	ires, err = c.GetAggregatedScoresByTicketEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetAggregatedScoresByTicketResult), nil
}
