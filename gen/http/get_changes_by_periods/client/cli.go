// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get changes by periods HTTP client CLI support package
//
// Command:
// $ goa gen go-scores/design

package client

import (
	getchangesbyperiods "go-scores/gen/get_changes_by_periods"
)

// BuildGetChangesByPeriodsPayload builds the payload for the get changes by
// periods getChangesByPeriods endpoint from CLI flags.
func BuildGetChangesByPeriodsPayload(getChangesByPeriodsGetChangesByPeriodsPeriod string) (*getchangesbyperiods.GetChangesByPeriodsPayload, error) {
	var period string
	{
		period = getChangesByPeriodsGetChangesByPeriodsPeriod
	}
	v := &getchangesbyperiods.GetChangesByPeriodsPayload{}
	v.Period = period

	return v, nil
}
