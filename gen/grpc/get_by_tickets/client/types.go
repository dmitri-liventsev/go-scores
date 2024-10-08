// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets gRPC client types
//
// Command:
// $ goa gen go-scores/design

package client

import (
	getbytickets "go-scores/gen/get_by_tickets"
	get_by_ticketspb "go-scores/gen/grpc/get_by_tickets/pb"

	goa "goa.design/goa/v3/pkg"
)

// NewProtoGetAggregatedScoresByTicketRequest builds the gRPC request type from
// the payload of the "getAggregatedScoresByTicket" endpoint of the "get by
// tickets" service.
func NewProtoGetAggregatedScoresByTicketRequest(payload *getbytickets.GetAggregatedScoresByTicketPayload) *get_by_ticketspb.GetAggregatedScoresByTicketRequest {
	message := &get_by_ticketspb.GetAggregatedScoresByTicketRequest{
		From: payload.From,
		To:   payload.To,
	}
	return message
}

// NewGetAggregatedScoresByTicketResult builds the result type of the
// "getAggregatedScoresByTicket" endpoint of the "get by tickets" service from
// the gRPC response type.
func NewGetAggregatedScoresByTicketResult(message *get_by_ticketspb.GetAggregatedScoresByTicketResponse) *getbytickets.GetAggregatedScoresByTicketResult {
	result := &getbytickets.GetAggregatedScoresByTicketResult{}
	if message.Meta != nil {
		result.Meta = protobufGetByTicketspbTicketMetaToGetbyticketsTicketMeta(message.Meta)
	}
	if message.Data != nil {
		result.Data = make([]*getbytickets.TicketScore, len(message.Data))
		for i, val := range message.Data {
			result.Data[i] = &getbytickets.TicketScore{}
			if val.TicketId != nil {
				ticketID := int(*val.TicketId)
				result.Data[i].TicketID = &ticketID
			}
			if val.Categories != nil {
				result.Data[i].Categories = make([]*getbytickets.CategoryScoreDetail, len(val.Categories))
				for j, val := range val.Categories {
					result.Data[i].Categories[j] = &getbytickets.CategoryScoreDetail{
						Score: val.Score,
					}
					if val.Id != nil {
						id := int(*val.Id)
						result.Data[i].Categories[j].ID = &id
					}
				}
			}
		}
	}
	return result
}

// ValidateGetAggregatedScoresByTicketResponse runs the validations defined on
// GetAggregatedScoresByTicketResponse.
func ValidateGetAggregatedScoresByTicketResponse(message *get_by_ticketspb.GetAggregatedScoresByTicketResponse) (err error) {
	if message.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "message"))
	}
	if message.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "message"))
	}
	return
}

// svcGetbyticketsTicketMetaToGetByTicketspbTicketMeta builds a value of type
// *get_by_ticketspb.TicketMeta from a value of type *getbytickets.TicketMeta.
func svcGetbyticketsTicketMetaToGetByTicketspbTicketMeta(v *getbytickets.TicketMeta) *get_by_ticketspb.TicketMeta {
	res := &get_by_ticketspb.TicketMeta{}
	if v.Categories != nil {
		res.Categories = make([]*get_by_ticketspb.Category, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = &get_by_ticketspb.Category{
				Name: val.Name,
			}
			if val.ID != nil {
				id := int32(*val.ID)
				res.Categories[i].Id = &id
			}
		}
	}

	return res
}

// protobufGetByTicketspbTicketMetaToGetbyticketsTicketMeta builds a value of
// type *getbytickets.TicketMeta from a value of type
// *get_by_ticketspb.TicketMeta.
func protobufGetByTicketspbTicketMetaToGetbyticketsTicketMeta(v *get_by_ticketspb.TicketMeta) *getbytickets.TicketMeta {
	res := &getbytickets.TicketMeta{}
	if v.Categories != nil {
		res.Categories = make([]*getbytickets.Category, len(v.Categories))
		for i, val := range v.Categories {
			res.Categories[i] = &getbytickets.Category{
				Name: val.Name,
			}
			if val.Id != nil {
				id := int(*val.Id)
				res.Categories[i].ID = &id
			}
		}
	}

	return res
}
