// Code generated with goa v3.18.2, DO NOT EDIT.
//
// get by categories protocol buffer definition
//
// Command:
// $ goa gen go-scores/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: goagen_go-scores_get_by_categories.proto

package get_by_categoriespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GetByCategories_GetAggregatedScores_FullMethodName = "/get_by_categories.GetByCategories/GetAggregatedScores"
)

// GetByCategoriesClient is the client API for GetByCategories service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Aggregated category scores over a period of time
type GetByCategoriesClient interface {
	// Get aggregated scores for categories within a specified period.
	GetAggregatedScores(ctx context.Context, in *GetAggregatedScoresRequest, opts ...grpc.CallOption) (*GetAggregatedScoresResponse, error)
}

type getByCategoriesClient struct {
	cc grpc.ClientConnInterface
}

func NewGetByCategoriesClient(cc grpc.ClientConnInterface) GetByCategoriesClient {
	return &getByCategoriesClient{cc}
}

func (c *getByCategoriesClient) GetAggregatedScores(ctx context.Context, in *GetAggregatedScoresRequest, opts ...grpc.CallOption) (*GetAggregatedScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAggregatedScoresResponse)
	err := c.cc.Invoke(ctx, GetByCategories_GetAggregatedScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetByCategoriesServer is the server API for GetByCategories service.
// All implementations must embed UnimplementedGetByCategoriesServer
// for forward compatibility.
//
// Aggregated category scores over a period of time
type GetByCategoriesServer interface {
	// Get aggregated scores for categories within a specified period.
	GetAggregatedScores(context.Context, *GetAggregatedScoresRequest) (*GetAggregatedScoresResponse, error)
	mustEmbedUnimplementedGetByCategoriesServer()
}

// UnimplementedGetByCategoriesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetByCategoriesServer struct{}

func (UnimplementedGetByCategoriesServer) GetAggregatedScores(context.Context, *GetAggregatedScoresRequest) (*GetAggregatedScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAggregatedScores not implemented")
}
func (UnimplementedGetByCategoriesServer) mustEmbedUnimplementedGetByCategoriesServer() {}
func (UnimplementedGetByCategoriesServer) testEmbeddedByValue()                         {}

// UnsafeGetByCategoriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetByCategoriesServer will
// result in compilation errors.
type UnsafeGetByCategoriesServer interface {
	mustEmbedUnimplementedGetByCategoriesServer()
}

func RegisterGetByCategoriesServer(s grpc.ServiceRegistrar, srv GetByCategoriesServer) {
	// If the following call pancis, it indicates UnimplementedGetByCategoriesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetByCategories_ServiceDesc, srv)
}

func _GetByCategories_GetAggregatedScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAggregatedScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetByCategoriesServer).GetAggregatedScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetByCategories_GetAggregatedScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetByCategoriesServer).GetAggregatedScores(ctx, req.(*GetAggregatedScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetByCategories_ServiceDesc is the grpc.ServiceDesc for GetByCategories service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetByCategories_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "get_by_categories.GetByCategories",
	HandlerType: (*GetByCategoriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAggregatedScores",
			Handler:    _GetByCategories_GetAggregatedScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_go-scores_get_by_categories.proto",
}
