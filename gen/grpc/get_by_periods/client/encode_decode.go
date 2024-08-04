// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by periods gRPC client encoders and decoders
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"context"
	getbyperiods "go-scores/gen/get_by_periods"
	get_by_periodspb "go-scores/gen/grpc/get_by_periods/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildGetByPeriodsFunc builds the remote method to invoke for "get by
// periods" service "getByPeriods" endpoint.
func BuildGetByPeriodsFunc(grpccli get_by_periodspb.GetByPeriodsClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetByPeriods(ctx, reqpb.(*get_by_periodspb.GetByPeriodsRequest), opts...)
		}
		return grpccli.GetByPeriods(ctx, &get_by_periodspb.GetByPeriodsRequest{}, opts...)
	}
}

// EncodeGetByPeriodsRequest encodes requests sent to get by periods
// getByPeriods endpoint.
func EncodeGetByPeriodsRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*getbyperiods.GetByPeriodsPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("get by periods", "getByPeriods", "*getbyperiods.GetByPeriodsPayload", v)
	}
	return NewProtoGetByPeriodsRequest(payload), nil
}

// DecodeGetByPeriodsResponse decodes responses from the get by periods
// getByPeriods endpoint.
func DecodeGetByPeriodsResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*get_by_periodspb.GetByPeriodsResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("get by periods", "getByPeriods", "*get_by_periodspb.GetByPeriodsResponse", v)
	}
	if err := ValidateGetByPeriodsResponse(message); err != nil {
		return nil, err
	}
	res := NewGetByPeriodsResult(message)
	return res, nil
}
