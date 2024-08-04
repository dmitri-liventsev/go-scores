// Code generated with goa v3.18.2, DO NOT EDIT.
//
// get changes by periods protocol buffer definition
//
// Command:
// $ goa gen go-scores/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: goagen_go-scores_get_changes_by_periods.proto

package get_changes_by_periodspb

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
	GetChangesByPeriods_GetChangesByPeriods_FullMethodName = "/get_changes_by_periods.GetChangesByPeriods/GetChangesByPeriods"
)

// GetChangesByPeriodsClient is the client API for GetChangesByPeriods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Period over period score change
type GetChangesByPeriodsClient interface {
	// Get the score change from a selected period over the previous period.
	GetChangesByPeriods(ctx context.Context, in *GetChangesByPeriodsRequest, opts ...grpc.CallOption) (*GetChangesByPeriodsResponse, error)
}

type getChangesByPeriodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGetChangesByPeriodsClient(cc grpc.ClientConnInterface) GetChangesByPeriodsClient {
	return &getChangesByPeriodsClient{cc}
}

func (c *getChangesByPeriodsClient) GetChangesByPeriods(ctx context.Context, in *GetChangesByPeriodsRequest, opts ...grpc.CallOption) (*GetChangesByPeriodsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChangesByPeriodsResponse)
	err := c.cc.Invoke(ctx, GetChangesByPeriods_GetChangesByPeriods_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetChangesByPeriodsServer is the server API for GetChangesByPeriods service.
// All implementations must embed UnimplementedGetChangesByPeriodsServer
// for forward compatibility.
//
// Period over period score change
type GetChangesByPeriodsServer interface {
	// Get the score change from a selected period over the previous period.
	GetChangesByPeriods(context.Context, *GetChangesByPeriodsRequest) (*GetChangesByPeriodsResponse, error)
	mustEmbedUnimplementedGetChangesByPeriodsServer()
}

// UnimplementedGetChangesByPeriodsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetChangesByPeriodsServer struct{}

func (UnimplementedGetChangesByPeriodsServer) GetChangesByPeriods(context.Context, *GetChangesByPeriodsRequest) (*GetChangesByPeriodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangesByPeriods not implemented")
}
func (UnimplementedGetChangesByPeriodsServer) mustEmbedUnimplementedGetChangesByPeriodsServer() {}
func (UnimplementedGetChangesByPeriodsServer) testEmbeddedByValue()                             {}

// UnsafeGetChangesByPeriodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetChangesByPeriodsServer will
// result in compilation errors.
type UnsafeGetChangesByPeriodsServer interface {
	mustEmbedUnimplementedGetChangesByPeriodsServer()
}

func RegisterGetChangesByPeriodsServer(s grpc.ServiceRegistrar, srv GetChangesByPeriodsServer) {
	// If the following call pancis, it indicates UnimplementedGetChangesByPeriodsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetChangesByPeriods_ServiceDesc, srv)
}

func _GetChangesByPeriods_GetChangesByPeriods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangesByPeriodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetChangesByPeriodsServer).GetChangesByPeriods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetChangesByPeriods_GetChangesByPeriods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetChangesByPeriodsServer).GetChangesByPeriods(ctx, req.(*GetChangesByPeriodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetChangesByPeriods_ServiceDesc is the grpc.ServiceDesc for GetChangesByPeriods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetChangesByPeriods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "get_changes_by_periods.GetChangesByPeriods",
	HandlerType: (*GetChangesByPeriodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChangesByPeriods",
			Handler:    _GetChangesByPeriods_GetChangesByPeriods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_go-scores_get_changes_by_periods.proto",
}