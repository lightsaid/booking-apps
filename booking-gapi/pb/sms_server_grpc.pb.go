// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: sms_server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SMSServiceClient is the client API for SMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSServiceClient interface {
	Send(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.Int32Value, error)
}

type sMSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSServiceClient(cc grpc.ClientConnInterface) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) Send(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.Int32Value, error) {
	out := new(wrapperspb.Int32Value)
	err := c.cc.Invoke(ctx, "/SMSService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServiceServer is the server API for SMSService service.
// All implementations must embed UnimplementedSMSServiceServer
// for forward compatibility
type SMSServiceServer interface {
	Send(context.Context, *wrapperspb.StringValue) (*wrapperspb.Int32Value, error)
	mustEmbedUnimplementedSMSServiceServer()
}

// UnimplementedSMSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSMSServiceServer struct {
}

func (UnimplementedSMSServiceServer) Send(context.Context, *wrapperspb.StringValue) (*wrapperspb.Int32Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSMSServiceServer) mustEmbedUnimplementedSMSServiceServer() {}

// UnsafeSMSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSServiceServer will
// result in compilation errors.
type UnsafeSMSServiceServer interface {
	mustEmbedUnimplementedSMSServiceServer()
}

func RegisterSMSServiceServer(s grpc.ServiceRegistrar, srv SMSServiceServer) {
	s.RegisterService(&SMSService_ServiceDesc, srv)
}

func _SMSService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SMSService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).Send(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// SMSService_ServiceDesc is the grpc.ServiceDesc for SMSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SMSService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms_server.proto",
}
