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