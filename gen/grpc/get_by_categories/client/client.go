// Code generated by goa v3.18.2, DO NOT EDIT.
//
// get by categories gRPC client
//
// Command:
// $ goa gen go-scores/design

package client

import (
	"context"
	get_by_categoriespb "go-scores/gen/grpc/get_by_categories/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli get_by_categoriespb.GetByCategoriesClient
	opts    []grpc.CallOption
} // NewClient instantiates gRPC client for all the get by categories service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: get_by_categoriespb.NewGetByCategoriesClient(cc),
		opts:    opts,
	}
} // GetAggregatedScores calls the "GetAggregatedScores" function in
// get_by_categoriespb.GetByCategoriesClient interface.
func (c *Client) GetAggregatedScores() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildGetAggregatedScoresFunc(c.grpccli, c.opts...),
			EncodeGetAggregatedScoresRequest,
			DecodeGetAggregatedScoresResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
