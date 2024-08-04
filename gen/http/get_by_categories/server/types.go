// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by categories HTTP server types
//
// Command:
// $ goa gen go-scores/design

package server

import (
	getbycategories "go-scores/gen/get_by_categories"
)

// GetAggregatedScoresResponseBody is the type of the "get by categories"
// service "getAggregatedScores" endpoint HTTP response body.
type GetAggregatedScoresResponseBody struct {
	Meta *CategoryMetaResponseBody    `form:"meta" json:"meta" xml:"meta"`
	Data []*CategoryScoreResponseBody `form:"data" json:"data" xml:"data"`
}

// CategoryMetaResponseBody is used to define fields on response body types.
type CategoryMetaResponseBody struct {
	Periods    []*PeriodResponseBody   `form:"periods,omitempty" json:"periods,omitempty" xml:"periods,omitempty"`
	Categories []*CategoryResponseBody `form:"categories,omitempty" json:"categories,omitempty" xml:"categories,omitempty"`
}

// PeriodResponseBody is used to define fields on response body types.
type PeriodResponseBody struct {
	ID    *int   `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Start *int64 `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End   *int64 `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
}

// CategoryResponseBody is used to define fields on response body types.
type CategoryResponseBody struct {
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CategoryScoreResponseBody is used to define fields on response body types.
type CategoryScoreResponseBody struct {
	CategoryID   *int                       `form:"category_id,omitempty" json:"category_id,omitempty" xml:"category_id,omitempty"`
	NumOfReviews *int                       `form:"num_of_reviews,omitempty" json:"num_of_reviews,omitempty" xml:"num_of_reviews,omitempty"`
	Periods      []*ScorePeriodResponseBody `form:"periods,omitempty" json:"periods,omitempty" xml:"periods,omitempty"`
	TotalScore   *float32                   `form:"total_score,omitempty" json:"total_score,omitempty" xml:"total_score,omitempty"`
}

// ScorePeriodResponseBody is used to define fields on response body types.
type ScorePeriodResponseBody struct {
	ID    *int     `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Score *float32 `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
}

// NewGetAggregatedScoresResponseBody builds the HTTP response body from the
// result of the "getAggregatedScores" endpoint of the "get by categories"
// service.
func NewGetAggregatedScoresResponseBody(res *getbycategories.GetAggregatedScoresResult) *GetAggregatedScoresResponseBody {
	body := &GetAggregatedScoresResponseBody{}
	if res.Meta != nil {
		body.Meta = marshalGetbycategoriesCategoryMetaToCategoryMetaResponseBody(res.Meta)
	}
	if res.Data != nil {
		body.Data = make([]*CategoryScoreResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalGetbycategoriesCategoryScoreToCategoryScoreResponseBody(val)
		}
	} else {
		body.Data = []*CategoryScoreResponseBody{}
	}
	return body
}

// NewGetAggregatedScoresPayload builds a get by categories service
// getAggregatedScores endpoint payload.
func NewGetAggregatedScoresPayload(from string, to string) *getbycategories.GetAggregatedScoresPayload {
	v := &getbycategories.GetAggregatedScoresPayload{}
	v.From = from
	v.To = to

	return v
}