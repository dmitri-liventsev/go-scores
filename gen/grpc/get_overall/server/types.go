// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get overall gRPC server types
//
// Command:
// $ goa gen go-scores/design

package server

import (
	getoverall "go-scores/gen/get_overall"
	get_overallpb "go-scores/gen/grpc/get_overall/pb"

	goa "goa.design/goa/v3/pkg"
)

// NewGetOverallScorePayload builds the payload of the "getOverallScore"
// endpoint of the "get overall" service from the gRPC request type.
func NewGetOverallScorePayload(message *get_overallpb.GetOverallScoreRequest) *getoverall.GetOverallScorePayload {
	v := &getoverall.GetOverallScorePayload{
		From: message.From,
		To:   message.To,
	}
	return v
}

// NewProtoGetOverallScoreResponse builds the gRPC response type from the
// result of the "getOverallScore" endpoint of the "get overall" service.
func NewProtoGetOverallScoreResponse(result float32) *get_overallpb.GetOverallScoreResponse {
	message := &get_overallpb.GetOverallScoreResponse{}
	message.Field = result
	return message
}

// ValidateGetOverallScoreRequest runs the validations defined on
// GetOverallScoreRequest.
func ValidateGetOverallScoreRequest(message *get_overallpb.GetOverallScoreRequest) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.from", message.From, goa.FormatDate))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.from", message.From, "^\\d{4}-\\d{2}-\\d{2}$"))
	err = goa.MergeErrors(err, goa.ValidateFormat("message.to", message.To, goa.FormatDate))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.to", message.To, "^\\d{4}-\\d{2}-\\d{2}$"))
	return
}
