// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get changes by periods HTTP client types
//
// Command:
// $ goa gen go-scores/design

package client

import (
	getchangesbyperiods "go-scores/gen/get_changes_by_periods"

	goa "goa.design/goa/v3/pkg"
)

// GetChangesByPeriodsResponseBody is the type of the "get changes by periods"
// service "getChangesByPeriods" endpoint HTTP response body.
type GetChangesByPeriodsResponseBody struct {
	Meta *PeriodsMetaResponseBody        `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
	Data []*PeriodScoreDeltaResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// PeriodsMetaResponseBody is used to define fields on response body types.
type PeriodsMetaResponseBody struct {
	Periods []*PeriodResponseBody `form:"periods,omitempty" json:"periods,omitempty" xml:"periods,omitempty"`
}

// PeriodResponseBody is used to define fields on response body types.
type PeriodResponseBody struct {
	ID    *int   `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Start *int64 `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End   *int64 `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
}

// PeriodScoreDeltaResponseBody is used to define fields on response body types.
type PeriodScoreDeltaResponseBody struct {
	PeriodID   *int     `form:"period_id,omitempty" json:"period_id,omitempty" xml:"period_id,omitempty"`
	ScoreDelta *float32 `form:"score_delta,omitempty" json:"score_delta,omitempty" xml:"score_delta,omitempty"`
}

// NewGetChangesByPeriodsResultOK builds a "get changes by periods" service
// "getChangesByPeriods" endpoint result from a HTTP "OK" response.
func NewGetChangesByPeriodsResultOK(body *GetChangesByPeriodsResponseBody) *getchangesbyperiods.GetChangesByPeriodsResult {
	v := &getchangesbyperiods.GetChangesByPeriodsResult{}
	v.Meta = unmarshalPeriodsMetaResponseBodyToGetchangesbyperiodsPeriodsMeta(body.Meta)
	v.Data = make([]*getchangesbyperiods.PeriodScoreDelta, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalPeriodScoreDeltaResponseBodyToGetchangesbyperiodsPeriodScoreDelta(val)
	}

	return v
}

// ValidateGetChangesByPeriodsResponseBody runs the validations defined on
// GetChangesByPeriodsResponseBody
func ValidateGetChangesByPeriodsResponseBody(body *GetChangesByPeriodsResponseBody) (err error) {
	if body.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	return
}
