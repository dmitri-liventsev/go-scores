// Code generated by goa v3.18.2, DO NOT EDIT.
//
// HTTP request path constructors for the get changes by periods service.
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"fmt"
)

// GetChangesByPeriodsGetChangesByPeriodsPath returns the URL path to the get changes by periods service getChangesByPeriods HTTP endpoint.
func GetChangesByPeriodsGetChangesByPeriodsPath(period string) string {
	return fmt.Sprintf("/periods/%v", period)
}