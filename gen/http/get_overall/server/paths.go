// Code generated by goa v3.18.2, DO NOT EDIT.
//
// HTTP request path constructors for the get overall service.
//
// Command:
// $ goa gen go-scores/design

package server

import (
	"fmt"
)

// GetOverallScoreGetOverallPath returns the URL path to the get overall service getOverallScore HTTP endpoint.
func GetOverallScoreGetOverallPath(from string, to string) string {
	return fmt.Sprintf("/overall/%v/%v", from, to)
}
