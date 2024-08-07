// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets client HTTP transport
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the get by tickets service endpoint HTTP clients.
type Client struct {
	// GetAggregatedScoresByTicket Doer is the HTTP client used to make requests to
	// the getAggregatedScoresByTicket endpoint.
	GetAggregatedScoresByTicketDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the get by tickets service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetAggregatedScoresByTicketDoer: doer,
		RestoreResponseBody:             restoreBody,
		scheme:                          scheme,
		host:                            host,
		decoder:                         dec,
		encoder:                         enc,
	}
}

// GetAggregatedScoresByTicket returns an endpoint that makes HTTP requests to
// the get by tickets service getAggregatedScoresByTicket server.
func (c *Client) GetAggregatedScoresByTicket() goa.Endpoint {
	var (
		decodeResponse = DecodeGetAggregatedScoresByTicketResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetAggregatedScoresByTicketRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAggregatedScoresByTicketDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("get by tickets", "getAggregatedScoresByTicket", err)
		}
		return decodeResponse(resp)
	}
}
