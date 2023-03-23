// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: token-info.proto

package proto

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

// TokenInfoClient is the client API for TokenInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenInfoClient interface {
	VerifyToken(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type tokenInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenInfoClient(cc grpc.ClientConnInterface) TokenInfoClient {
	return &tokenInfoClient{cc}
}

func (c *tokenInfoClient) VerifyToken(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenInfo/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenInfoServer is the server API for TokenInfo service.
// All implementations must embed UnimplementedTokenInfoServer
// for forward compatibility
type TokenInfoServer interface {
	VerifyToken(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedTokenInfoServer()
}

// UnimplementedTokenInfoServer must be embedded to have forward compatible implementations.
type UnimplementedTokenInfoServer struct {
}

func (UnimplementedTokenInfoServer) VerifyToken(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedTokenInfoServer) mustEmbedUnimplementedTokenInfoServer() {}

// UnsafeTokenInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenInfoServer will
// result in compilation errors.
type UnsafeTokenInfoServer interface {
	mustEmbedUnimplementedTokenInfoServer()
}

func RegisterTokenInfoServer(s grpc.ServiceRegistrar, srv TokenInfoServer) {
	s.RegisterService(&TokenInfo_ServiceDesc, srv)
}

func _TokenInfo_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenInfoServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenInfo/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenInfoServer).VerifyToken(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenInfo_ServiceDesc is the grpc.ServiceDesc for TokenInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TokenInfo",
	HandlerType: (*TokenInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _TokenInfo_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token-info.proto",
}