// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: expense-tracker/sanitary/v1/service.proto

package v1

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
	SanitaryService_Ping_FullMethodName = "/expense_tracker.sanitary.v1.SanitaryService/Ping"
)

// SanitaryServiceClient is the client API for SanitaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SanitaryServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type sanitaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSanitaryServiceClient(cc grpc.ClientConnInterface) SanitaryServiceClient {
	return &sanitaryServiceClient{cc}
}

func (c *sanitaryServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, SanitaryService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SanitaryServiceServer is the server API for SanitaryService service.
// All implementations must embed UnimplementedSanitaryServiceServer
// for forward compatibility.
type SanitaryServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedSanitaryServiceServer()
}

// UnimplementedSanitaryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSanitaryServiceServer struct{}

func (UnimplementedSanitaryServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSanitaryServiceServer) mustEmbedUnimplementedSanitaryServiceServer() {}
func (UnimplementedSanitaryServiceServer) testEmbeddedByValue()                         {}

// UnsafeSanitaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SanitaryServiceServer will
// result in compilation errors.
type UnsafeSanitaryServiceServer interface {
	mustEmbedUnimplementedSanitaryServiceServer()
}

func RegisterSanitaryServiceServer(s grpc.ServiceRegistrar, srv SanitaryServiceServer) {
	// If the following call pancis, it indicates UnimplementedSanitaryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SanitaryService_ServiceDesc, srv)
}

func _SanitaryService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SanitaryServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SanitaryService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SanitaryServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SanitaryService_ServiceDesc is the grpc.ServiceDesc for SanitaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SanitaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "expense_tracker.sanitary.v1.SanitaryService",
	HandlerType: (*SanitaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SanitaryService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expense-tracker/sanitary/v1/service.proto",
}
