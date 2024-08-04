// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets gRPC client CLI support package
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"encoding/json"
	"fmt"
	getbytickets "go-scores/gen/get_by_tickets"
	get_by_ticketspb "go-scores/gen/grpc/get_by_tickets/pb"
)

// BuildGetAggregatedScoresByTicketPayload builds the payload for the get by
// tickets getAggregatedScoresByTicket endpoint from CLI flags.
func BuildGetAggregatedScoresByTicketPayload(getByTicketsGetAggregatedScoresByTicketMessage string) (*getbytickets.GetAggregatedScoresByTicketPayload, error) {
	var err error
	var message get_by_ticketspb.GetAggregatedScoresByTicketRequest
	{
		if getByTicketsGetAggregatedScoresByTicketMessage != "" {
			err = json.Unmarshal([]byte(getByTicketsGetAggregatedScoresByTicketMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"from\": \"Similique cupiditate ut facilis.\",\n      \"to\": \"Quod vitae amet quos sint nulla.\"\n   }'")
			}
		}
	}
	v := &getbytickets.GetAggregatedScoresByTicketPayload{
		From: message.From,
		To:   message.To,
	}

	return v, nil
}
