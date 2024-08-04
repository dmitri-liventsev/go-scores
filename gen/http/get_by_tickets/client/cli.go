// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets HTTP client CLI support package
//
// Command:
// $ goa gen go-scores/design

package client

import (
	getbytickets "go-scores/gen/get_by_tickets"
)

// BuildGetAggregatedScoresByTicketPayload builds the payload for the get by
// tickets getAggregatedScoresByTicket endpoint from CLI flags.
func BuildGetAggregatedScoresByTicketPayload(getByTicketsGetAggregatedScoresByTicketFrom string, getByTicketsGetAggregatedScoresByTicketTo string) (*getbytickets.GetAggregatedScoresByTicketPayload, error) {
	var from string
	{
		from = getByTicketsGetAggregatedScoresByTicketFrom
	}
	var to string
	{
		to = getByTicketsGetAggregatedScoresByTicketTo
	}
	v := &getbytickets.GetAggregatedScoresByTicketPayload{}
	v.From = from
	v.To = to

	return v, nil
}
