// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: url_shorter/v1/url_shorter.proto

package v1

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

const (
	UrlShortenerService_CreateShortUrl_FullMethodName = "/url_shorter.v1.UrlShortenerService/CreateShortUrl"
	UrlShortenerService_GetRedirectURL_FullMethodName = "/url_shorter.v1.UrlShortenerService/GetRedirectURL"
)

// UrlShortenerServiceClient is the client API for UrlShortenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerServiceClient interface {
	// CreateUrlShortener creates a new short URL based on the provided long URL, short key, and biz tag.
	CreateShortUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlReply, error)
	// GetRedirectURL retrieves the long URL associated with the given short URL and returns it to the client.
	GetRedirectURL(ctx context.Context, in *JumpRequest, opts ...grpc.CallOption) (*JumpReply, error)
}

type urlShortenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerServiceClient(cc grpc.ClientConnInterface) UrlShortenerServiceClient {
	return &urlShortenerServiceClient{cc}
}

func (c *urlShortenerServiceClient) CreateShortUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlReply, error) {
	out := new(CreateUrlReply)
	err := c.cc.Invoke(ctx, UrlShortenerService_CreateShortUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerServiceClient) GetRedirectURL(ctx context.Context, in *JumpRequest, opts ...grpc.CallOption) (*JumpReply, error) {
	out := new(JumpReply)
	err := c.cc.Invoke(ctx, UrlShortenerService_GetRedirectURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServiceServer is the server API for UrlShortenerService service.
// All implementations must embed UnimplementedUrlShortenerServiceServer
// for forward compatibility
type UrlShortenerServiceServer interface {
	// CreateUrlShortener creates a new short URL based on the provided long URL, short key, and biz tag.
	CreateShortUrl(context.Context, *CreateUrlRequest) (*CreateUrlReply, error)
	// GetRedirectURL retrieves the long URL associated with the given short URL and returns it to the client.
	GetRedirectURL(context.Context, *JumpRequest) (*JumpReply, error)
	mustEmbedUnimplementedUrlShortenerServiceServer()
}

// UnimplementedUrlShortenerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortenerServiceServer struct {
}

func (UnimplementedUrlShortenerServiceServer) CreateShortUrl(context.Context, *CreateUrlRequest) (*CreateUrlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortUrl not implemented")
}
func (UnimplementedUrlShortenerServiceServer) GetRedirectURL(context.Context, *JumpRequest) (*JumpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRedirectURL not implemented")
}
func (UnimplementedUrlShortenerServiceServer) mustEmbedUnimplementedUrlShortenerServiceServer() {}

// UnsafeUrlShortenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortenerServiceServer will
// result in compilation errors.
type UnsafeUrlShortenerServiceServer interface {
	mustEmbedUnimplementedUrlShortenerServiceServer()
}

func RegisterUrlShortenerServiceServer(s grpc.ServiceRegistrar, srv UrlShortenerServiceServer) {
	s.RegisterService(&UrlShortenerService_ServiceDesc, srv)
}

func _UrlShortenerService_CreateShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServiceServer).CreateShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortenerService_CreateShortUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServiceServer).CreateShortUrl(ctx, req.(*CreateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortenerService_GetRedirectURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServiceServer).GetRedirectURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortenerService_GetRedirectURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServiceServer).GetRedirectURL(ctx, req.(*JumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortenerService_ServiceDesc is the grpc.ServiceDesc for UrlShortenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "url_shorter.v1.UrlShortenerService",
	HandlerType: (*UrlShortenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortUrl",
			Handler:    _UrlShortenerService_CreateShortUrl_Handler,
		},
		{
			MethodName: "GetRedirectURL",
			Handler:    _UrlShortenerService_GetRedirectURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "url_shorter/v1/url_shorter.proto",
}
