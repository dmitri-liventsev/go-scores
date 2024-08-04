// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by tickets HTTP client types
//
// Command:
// $ goa gen go-scores/design

package client

import (
	getbytickets "go-scores/gen/get_by_tickets"

	goa "goa.design/goa/v3/pkg"
)

// GetAggregatedScoresByTicketResponseBody is the type of the "get by tickets"
// service "getAggregatedScoresByTicket" endpoint HTTP response body.
type GetAggregatedScoresByTicketResponseBody struct {
	Meta *TicketMetaResponseBody    `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Data []*TicketScoreResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// TicketMetaResponseBody is used to define fields on response body types.
type TicketMetaResponseBody struct {
	Categories []*CategoryResponseBody `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
}

// CategoryResponseBody is used to define fields on response body types.
type CategoryResponseBody struct {
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// TicketScoreResponseBody is used to define fields on response body types.
type TicketScoreResponseBody struct {
	TicketID   *int                               `form:"ticket_id,omitempty" json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	Categories []*CategoryScoreDetailResponseBody `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
}

// CategoryScoreDetailResponseBody is used to define fields on response body
// types.
type CategoryScoreDetailResponseBody struct {
	ID    *int     `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Score *float32 `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

// NewGetAggregatedScoresByTicketResultOK builds a "get by tickets" service
// "getAggregatedScoresByTicket" endpoint result from a HTTP "OK" response.
func NewGetAggregatedScoresByTicketResultOK(body *GetAggregatedScoresByTicketResponseBody) *getbytickets.GetAggregatedScoresByTicketResult {
	v := &getbytickets.GetAggregatedScoresByTicketResult{}
	v.Meta = unmarshalTicketMetaResponseBodyToGetbyticketsTicketMeta(body.Meta)
	v.Data = make([]*getbytickets.TicketScore, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalTicketScoreResponseBodyToGetbyticketsTicketScore(val)
	}

	return v
}

// ValidateGetAggregatedScoresByTicketResponseBody runs the validations defined
// on GetAggregatedScoresByTicketResponseBody
func ValidateGetAggregatedScoresByTicketResponseBody(body *GetAggregatedScoresByTicketResponseBody) (err error) {
	if body.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	return
}
