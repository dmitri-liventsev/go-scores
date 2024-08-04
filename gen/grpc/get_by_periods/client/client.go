// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by periods gRPC client
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"context"
	get_by_periodspb "go-scores/gen/grpc/get_by_periods/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli get_by_periodspb.GetByPeriodsClient
	opts    []grpc.CallOption
} // NewClient instantiates gRPC client for all the get by periods service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: get_by_periodspb.NewGetByPeriodsClient(cc),
		opts:    opts,
	}
} // GetByPeriods calls the "GetByPeriods" function in
// get_by_periodspb.GetByPeriodsClient interface.
func (c *Client) GetByPeriods() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildGetByPeriodsFunc(c.grpccli, c.opts...),
			EncodeGetByPeriodsRequest,
			DecodeGetByPeriodsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}