// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package MovieProto

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

// GetMovieInfoClient is the client API for GetMovieInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetMovieInfoClient interface {
	GetMovie(ctx context.Context, in *SeachRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type getMovieInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewGetMovieInfoClient(cc grpc.ClientConnInterface) GetMovieInfoClient {
	return &getMovieInfoClient{cc}
}

func (c *getMovieInfoClient) GetMovie(ctx context.Context, in *SeachRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/MovieProto.GetMovieInfo/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetMovieInfoServer is the server API for GetMovieInfo service.
// All implementations must embed UnimplementedGetMovieInfoServer
// for forward compatibility
type GetMovieInfoServer interface {
	GetMovie(context.Context, *SeachRequest) (*SearchResponse, error)
}

// UnimplementedGetMovieInfoServer must be embedded to have forward compatible implementations.
type UnimplementedGetMovieInfoServer struct {
}

func (UnimplementedGetMovieInfoServer) GetMovie(context.Context, *SeachRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedGetMovieInfoServer) mustEmbedUnimplementedGetMovieInfoServer() {}

// UnsafeGetMovieInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetMovieInfoServer will
// result in compilation errors.
type UnsafeGetMovieInfoServer interface {
	mustEmbedUnimplementedGetMovieInfoServer()
}

func RegisterGetMovieInfoServer(s grpc.ServiceRegistrar, srv GetMovieInfoServer) {
	s.RegisterService(&GetMovieInfo_ServiceDesc, srv)
}

func _GetMovieInfo_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetMovieInfoServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieProto.GetMovieInfo/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetMovieInfoServer).GetMovie(ctx, req.(*SeachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetMovieInfo_ServiceDesc is the grpc.ServiceDesc for GetMovieInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetMovieInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MovieProto.GetMovieInfo",
	HandlerType: (*GetMovieInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _GetMovieInfo_GetMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/proto/MovieProto/Movie.proto",
}
