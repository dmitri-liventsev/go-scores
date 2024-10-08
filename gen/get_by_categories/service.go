// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by categories service
//
// Command:
// $ goa gen go-scores/design

package getbycategories

import (
	"context"
)

// Aggregated category scores over a period of time
type Service interface {
	// Get aggregated scores for categories within a specified period.
	GetAggregatedScores(context.Context, *GetAggregatedScoresPayload) (res *GetAggregatedScoresResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "get overall"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "get by categories"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"getAggregatedScores"}

type Category struct {
	ID   *int
	Name *string
}

type CategoryMeta struct {
	Periods    []*Period
	Categories []*Category
}

type CategoryScore struct {
	CategoryID   *int
	NumOfReviews *int
	Periods      []*ScorePeriod
	TotalScore   *float32
}

// GetAggregatedScoresPayload is the payload type of the get by categories
// service getAggregatedScores method.
type GetAggregatedScoresPayload struct {
	// Start date (YYYY-MM-DD)
	From string
	// End date (YYYY-MM-DD)
	To string
}

// GetAggregatedScoresResult is the result type of the get by categories
// service getAggregatedScores method.
type GetAggregatedScoresResult struct {
	Meta *CategoryMeta
	Data []*CategoryScore
}

type Period struct {
	ID    *int
	Start *int64
	End   *int64
}

type ScorePeriod struct {
	ID    *int
	Score *float32
}
