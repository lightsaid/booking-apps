// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: movie_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	ListMovie(ctx context.Context, in *ListMovieRequest, opts ...grpc.CallOption) (MovieService_ListMovieClient, error)
	GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) ListMovie(ctx context.Context, in *ListMovieRequest, opts ...grpc.CallOption) (MovieService_ListMovieClient, error) {
	stream, err := c.cc.NewStream(ctx, &MovieService_ServiceDesc.Streams[0], "/MovieService/ListMovie", opts...)
	if err != nil {
		return nil, err
	}
	x := &movieServiceListMovieClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MovieService_ListMovieClient interface {
	Recv() (*ListMovieResponse, error)
	grpc.ClientStream
}

type movieServiceListMovieClient struct {
	grpc.ClientStream
}

func (x *movieServiceListMovieClient) Recv() (*ListMovieResponse, error) {
	m := new(ListMovieResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *movieServiceClient) GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/MovieService/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	ListMovie(*ListMovieRequest, MovieService_ListMovieServer) error
	GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error)
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) ListMovie(*ListMovieRequest, MovieService_ListMovieServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMovie not implemented")
}
func (UnimplementedMovieServiceServer) GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_ListMovie_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMovieRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieServiceServer).ListMovie(m, &movieServiceListMovieServer{stream})
}

type MovieService_ListMovieServer interface {
	Send(*ListMovieResponse) error
	grpc.ServerStream
}

type movieServiceListMovieServer struct {
	grpc.ServerStream
}

func (x *movieServiceListMovieServer) Send(m *ListMovieResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MovieService_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieService/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovie(ctx, req.(*GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieService_GetMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMovie",
			Handler:       _MovieService_ListMovie_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "movie_service.proto",
}
