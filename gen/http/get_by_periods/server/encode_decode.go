// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by periods HTTP server encoders and decoders
//
// Command:
// $ goa gen go-scores/design

package server

import (
	"context"
	getbyperiods "go-scores/gen/get_by_periods"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeGetByPeriodsResponse returns an encoder for responses returned by the
// get by periods getByPeriods endpoint.
func EncodeGetByPeriodsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*getbyperiods.GetByPeriodsResult)
		enc := encoder(ctx, w)
		body := NewGetByPeriodsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetByPeriodsRequest returns a decoder for requests sent to the get by
// periods getByPeriods endpoint.
func DecodeGetByPeriodsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			period string

			params = mux.Vars(r)
		)
		period = params["period"]
		payload := NewGetByPeriodsPayload(period)

		return payload, nil
	}
}

// marshalGetbyperiodsPeriodsMetaToPeriodsMetaResponseBody builds a value of
// type *PeriodsMetaResponseBody from a value of type *getbyperiods.PeriodsMeta.
func marshalGetbyperiodsPeriodsMetaToPeriodsMetaResponseBody(v *getbyperiods.PeriodsMeta) *PeriodsMetaResponseBody {
	res := &PeriodsMetaResponseBody{}
	if v.Periods != nil {
		res.Periods = make([]*PeriodResponseBody, len(v.Periods))
		for i, val := range v.Periods {
			res.Periods[i] = marshalGetbyperiodsPeriodToPeriodResponseBody(val)
		}
	}

	return res
}

// marshalGetbyperiodsPeriodToPeriodResponseBody builds a value of type
// *PeriodResponseBody from a value of type *getbyperiods.Period.
func marshalGetbyperiodsPeriodToPeriodResponseBody(v *getbyperiods.Period) *PeriodResponseBody {
	if v == nil {
		return nil
	}
	res := &PeriodResponseBody{
		ID:    v.ID,
		Start: v.Start,
		End:   v.End,
	}

	return res
}

// marshalGetbyperiodsPeriodScoreChangeToPeriodScoreChangeResponseBody builds a
// value of type *PeriodScoreChangeResponseBody from a value of type
// *getbyperiods.PeriodScoreChange.
func marshalGetbyperiodsPeriodScoreChangeToPeriodScoreChangeResponseBody(v *getbyperiods.PeriodScoreChange) *PeriodScoreChangeResponseBody {
	res := &PeriodScoreChangeResponseBody{
		PeriodID:  v.PeriodID,
		ScoreDiff: v.ScoreDiff,
	}

	return res
}
