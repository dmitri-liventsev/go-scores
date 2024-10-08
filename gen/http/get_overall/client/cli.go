// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get overall HTTP client CLI support package
//
// Command:
// $ goa gen go-scores/design

package client

import (
	getoverall "go-scores/gen/get_overall"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetOverallScorePayload builds the payload for the get overall
// getOverallScore endpoint from CLI flags.
func BuildGetOverallScorePayload(getOverallGetOverallScoreFrom string, getOverallGetOverallScoreTo string) (*getoverall.GetOverallScorePayload, error) {
	var err error
	var from string
	{
		from = getOverallGetOverallScoreFrom
		err = goa.MergeErrors(err, goa.ValidateFormat("from", from, goa.FormatDate))
		err = goa.MergeErrors(err, goa.ValidatePattern("from", from, "^\\d{4}-\\d{2}-\\d{2}$"))
		if err != nil {
			return nil, err
		}
	}
	var to string
	{
		to = getOverallGetOverallScoreTo
		err = goa.MergeErrors(err, goa.ValidateFormat("to", to, goa.FormatDate))
		err = goa.MergeErrors(err, goa.ValidatePattern("to", to, "^\\d{4}-\\d{2}-\\d{2}$"))
		if err != nil {
			return nil, err
		}
	}
	v := &getoverall.GetOverallScorePayload{}
	v.From = from
	v.To = to

	return v, nil
}
