// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: app/protos/app.proto

package __

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
	Book_CreateBook_FullMethodName = "/app.Book/CreateBook"
	Book_GetBook_FullMethodName    = "/app.Book/GetBook"
	Book_UpdateBook_FullMethodName = "/app.Book/UpdateBook"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	CreateBook(ctx context.Context, in *CreateBookRequestBody, opts ...grpc.CallOption) (*CreateBookResponseBody, error)
	GetBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*GetBookResponseBody, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequestBody, opts ...grpc.CallOption) (*CreateBookResponseBody, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) CreateBook(ctx context.Context, in *CreateBookRequestBody, opts ...grpc.CallOption) (*CreateBookResponseBody, error) {
	out := new(CreateBookResponseBody)
	err := c.cc.Invoke(ctx, Book_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*GetBookResponseBody, error) {
	out := new(GetBookResponseBody)
	err := c.cc.Invoke(ctx, Book_GetBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) UpdateBook(ctx context.Context, in *UpdateBookRequestBody, opts ...grpc.CallOption) (*CreateBookResponseBody, error) {
	out := new(CreateBookResponseBody)
	err := c.cc.Invoke(ctx, Book_UpdateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility
type BookServer interface {
	CreateBook(context.Context, *CreateBookRequestBody) (*CreateBookResponseBody, error)
	GetBook(context.Context, *BookID) (*GetBookResponseBody, error)
	UpdateBook(context.Context, *UpdateBookRequestBody) (*CreateBookResponseBody, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (UnimplementedBookServer) CreateBook(context.Context, *CreateBookRequestBody) (*CreateBookResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServer) GetBook(context.Context, *BookID) (*GetBookResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServer) UpdateBook(context.Context, *UpdateBookRequestBody) (*CreateBookResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).CreateBook(ctx, req.(*CreateBookRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).UpdateBook(ctx, req.(*UpdateBookRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _Book_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Book_GetBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Book_UpdateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/protos/app.proto",
}

const (
	Auth_Login_FullMethodName           = "/app.Auth/Login"
	Auth_TokenValidation_FullMethodName = "/app.Auth/TokenValidation"
	Auth_SignUp_FullMethodName          = "/app.Auth/SignUp"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *LoginRequestBody, opts ...grpc.CallOption) (*LoginResponseBody, error)
	TokenValidation(ctx context.Context, in *Token, opts ...grpc.CallOption) (*GetTokenResponseBody, error)
	SignUp(ctx context.Context, in *SignUpRequestBody, opts ...grpc.CallOption) (*SignUpResponseBody, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequestBody, opts ...grpc.CallOption) (*LoginResponseBody, error) {
	out := new(LoginResponseBody)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) TokenValidation(ctx context.Context, in *Token, opts ...grpc.CallOption) (*GetTokenResponseBody, error) {
	out := new(GetTokenResponseBody)
	err := c.cc.Invoke(ctx, Auth_TokenValidation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignUp(ctx context.Context, in *SignUpRequestBody, opts ...grpc.CallOption) (*SignUpResponseBody, error) {
	out := new(SignUpResponseBody)
	err := c.cc.Invoke(ctx, Auth_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Login(context.Context, *LoginRequestBody) (*LoginResponseBody, error)
	TokenValidation(context.Context, *Token) (*GetTokenResponseBody, error)
	SignUp(context.Context, *SignUpRequestBody) (*SignUpResponseBody, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Login(context.Context, *LoginRequestBody) (*LoginResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) TokenValidation(context.Context, *Token) (*GetTokenResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenValidation not implemented")
}
func (UnimplementedAuthServer) SignUp(context.Context, *SignUpRequestBody) (*SignUpResponseBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_TokenValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).TokenValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_TokenValidation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).TokenValidation(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUp(ctx, req.(*SignUpRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "TokenValidation",
			Handler:    _Auth_TokenValidation_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Auth_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/protos/app.proto",
}
