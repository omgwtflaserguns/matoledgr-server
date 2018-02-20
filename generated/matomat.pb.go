// Code generated by protoc-gen-go. DO NOT EDIT.
// source: matomat.proto

/*
Package matomat is a generated protocol buffer package.

It is generated from these files:
	matomat.proto

It has these top-level messages:
	ProductRequest
	ProductList
	Product
	AccountRequest
	RegisterResponse
	LoginResponse
	User
*/
package matomat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterStatus int32

const (
	RegisterStatus_REGISTER_OK                         RegisterStatus = 0
	RegisterStatus_REGISTER_FAILED                     RegisterStatus = 1
	RegisterStatus_REGISTER_FAILED_NAME_ALREADY_IN_USE RegisterStatus = 2
	RegisterStatus_REGISTER_FAILED_PASSWORD_INVALID    RegisterStatus = 3
)

var RegisterStatus_name = map[int32]string{
	0: "REGISTER_OK",
	1: "REGISTER_FAILED",
	2: "REGISTER_FAILED_NAME_ALREADY_IN_USE",
	3: "REGISTER_FAILED_PASSWORD_INVALID",
}
var RegisterStatus_value = map[string]int32{
	"REGISTER_OK":                         0,
	"REGISTER_FAILED":                     1,
	"REGISTER_FAILED_NAME_ALREADY_IN_USE": 2,
	"REGISTER_FAILED_PASSWORD_INVALID":    3,
}

func (x RegisterStatus) String() string {
	return proto.EnumName(RegisterStatus_name, int32(x))
}
func (RegisterStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoginStatus int32

const (
	LoginStatus_LOGIN_OK     LoginStatus = 0
	LoginStatus_LOGIN_FAILED LoginStatus = 1
)

var LoginStatus_name = map[int32]string{
	0: "LOGIN_OK",
	1: "LOGIN_FAILED",
}
var LoginStatus_value = map[string]int32{
	"LOGIN_OK":     0,
	"LOGIN_FAILED": 1,
}

func (x LoginStatus) String() string {
	return proto.EnumName(LoginStatus_name, int32(x))
}
func (LoginStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ProductRequest struct {
}

func (m *ProductRequest) Reset()                    { *m = ProductRequest{} }
func (m *ProductRequest) String() string            { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()               {}
func (*ProductRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProductList struct {
	Products []*Product `protobuf:"bytes,1,rep,name=products" json:"products,omitempty"`
}

func (m *ProductList) Reset()                    { *m = ProductList{} }
func (m *ProductList) String() string            { return proto.CompactTextString(m) }
func (*ProductList) ProtoMessage()               {}
func (*ProductList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProductList) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type Product struct {
	Id    int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price" json:"price,omitempty"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type AccountRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AccountRequest) Reset()                    { *m = AccountRequest{} }
func (m *AccountRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()               {}
func (*AccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	Status RegisterStatus `protobuf:"varint,1,opt,name=status,enum=matomat.RegisterStatus" json:"status,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterResponse) GetStatus() RegisterStatus {
	if m != nil {
		return m.Status
	}
	return RegisterStatus_REGISTER_OK
}

type LoginResponse struct {
	Status LoginStatus `protobuf:"varint,1,opt,name=status,enum=matomat.LoginStatus" json:"status,omitempty"`
	User   *User       `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginResponse) GetStatus() LoginStatus {
	if m != nil {
		return m.Status
	}
	return LoginStatus_LOGIN_OK
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "matomat.ProductRequest")
	proto.RegisterType((*ProductList)(nil), "matomat.ProductList")
	proto.RegisterType((*Product)(nil), "matomat.Product")
	proto.RegisterType((*AccountRequest)(nil), "matomat.AccountRequest")
	proto.RegisterType((*RegisterResponse)(nil), "matomat.RegisterResponse")
	proto.RegisterType((*LoginResponse)(nil), "matomat.LoginResponse")
	proto.RegisterType((*User)(nil), "matomat.User")
	proto.RegisterEnum("matomat.RegisterStatus", RegisterStatus_name, RegisterStatus_value)
	proto.RegisterEnum("matomat.LoginStatus", LoginStatus_name, LoginStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Products service

type ProductsClient interface {
	ListProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductList, error)
}

type productsClient struct {
	cc *grpc.ClientConn
}

func NewProductsClient(cc *grpc.ClientConn) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) ListProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := grpc.Invoke(ctx, "/matomat.Products/ListProducts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Products service

type ProductsServer interface {
	ListProducts(context.Context, *ProductRequest) (*ProductList, error)
}

func RegisterProductsServer(s *grpc.Server, srv ProductsServer) {
	s.RegisterService(&_Products_serviceDesc, srv)
}

func _Products_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matomat.Products/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).ListProducts(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Products_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matomat.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _Products_ListProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matomat.proto",
}

// Client API for Account service

type AccountClient interface {
	Login(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Login(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/matomat.Account/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Register(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/matomat.Account/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountServer interface {
	Login(context.Context, *AccountRequest) (*LoginResponse, error)
	Register(context.Context, *AccountRequest) (*RegisterResponse, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matomat.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matomat.Account/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Register(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matomat.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Account_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matomat.proto",
}

func init() { proto.RegisterFile("matomat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0xed, 0xf4, 0x63, 0x8b, 0x97, 0x96, 0x25, 0xd7, 0xc6, 0xad, 0x7d, 0xc2, 0xd1, 0xc4, 0x66,
	0xb3, 0xae, 0x09, 0xbe, 0xe9, 0x83, 0x92, 0x2d, 0xae, 0x64, 0xb1, 0xdd, 0x0c, 0xae, 0xc6, 0x27,
	0xc4, 0x76, 0xb2, 0xe1, 0xa1, 0x05, 0x99, 0x21, 0xfe, 0x02, 0xe3, 0xdf, 0x36, 0x0c, 0x30, 0x6e,
	0xbb, 0xd1, 0x37, 0xce, 0xbd, 0xe7, 0x1c, 0xce, 0x30, 0x07, 0x18, 0x6f, 0x13, 0x99, 0x6d, 0x13,
	0x79, 0x9e, 0x17, 0x99, 0xcc, 0x70, 0xd8, 0x40, 0x6a, 0x83, 0x75, 0x5d, 0x64, 0x9b, 0x72, 0x2d,
	0x19, 0xff, 0x51, 0x72, 0x21, 0xe9, 0x1b, 0x30, 0x9b, 0x49, 0x98, 0x0a, 0x89, 0x67, 0x60, 0xe4,
	0x35, 0x14, 0x53, 0xe2, 0xf4, 0xe6, 0xa6, 0x6b, 0x9f, 0xb7, 0x5e, 0xad, 0x52, 0x33, 0xe8, 0x05,
	0x0c, 0x9b, 0x21, 0x5a, 0xd0, 0x4d, 0x37, 0x53, 0xe2, 0x90, 0xf9, 0x80, 0x75, 0xd3, 0x0d, 0x22,
	0xf4, 0x77, 0xc9, 0x96, 0x4f, 0xbb, 0x0e, 0x99, 0x3f, 0x60, 0xea, 0x19, 0x27, 0x30, 0xc8, 0x8b,
	0x74, 0xcd, 0xa7, 0x3d, 0x87, 0xcc, 0xbb, 0xac, 0x06, 0xf4, 0x03, 0x58, 0xde, 0x7a, 0x9d, 0x95,
	0xbb, 0x36, 0x13, 0xce, 0xc0, 0x28, 0x05, 0x2f, 0x94, 0x9e, 0x28, 0xbd, 0xc6, 0xd5, 0x2e, 0x4f,
	0x84, 0xf8, 0x99, 0x15, 0x9b, 0xc6, 0x5b, 0x63, 0x7a, 0x01, 0x36, 0xe3, 0xb7, 0xa9, 0x90, 0xbc,
	0x60, 0x5c, 0xe4, 0xd9, 0x4e, 0x70, 0x7c, 0x09, 0x47, 0x42, 0x26, 0xb2, 0x14, 0xca, 0xc9, 0x72,
	0x4f, 0xf4, 0x71, 0x5a, 0x6a, 0xa4, 0xd6, 0xac, 0xa1, 0xd1, 0x6f, 0x30, 0x0e, 0xb3, 0xdb, 0x74,
	0xa7, 0x1d, 0xce, 0x0e, 0x1c, 0x26, 0xda, 0x41, 0xf1, 0xf6, 0xe5, 0xf8, 0x04, 0xfa, 0x55, 0x56,
	0x95, 0xcd, 0x74, 0xc7, 0x9a, 0x7b, 0x23, 0x78, 0xc1, 0xd4, 0x8a, 0x52, 0xe8, 0x57, 0xe8, 0x7f,
	0xc7, 0x3c, 0xfd, 0x45, 0xc0, 0xda, 0x0f, 0x88, 0xc7, 0x60, 0x32, 0xff, 0x32, 0x88, 0x3e, 0xf9,
	0x2c, 0x5e, 0x5d, 0xd9, 0x1d, 0x7c, 0x08, 0xc7, 0x7a, 0xf0, 0xde, 0x0b, 0x42, 0x7f, 0x61, 0x13,
	0x7c, 0x0e, 0x4f, 0x0f, 0x86, 0xf1, 0xd2, 0xfb, 0xe8, 0xc7, 0x5e, 0xc8, 0x7c, 0x6f, 0xf1, 0x35,
	0x0e, 0x96, 0xf1, 0x4d, 0xe4, 0xdb, 0x5d, 0x7c, 0x06, 0xce, 0x21, 0xf1, 0xda, 0x8b, 0xa2, 0x2f,
	0x2b, 0xb6, 0x88, 0x83, 0xe5, 0x67, 0x2f, 0x0c, 0x16, 0x76, 0xef, 0xf4, 0x05, 0x98, 0x77, 0x4e,
	0x89, 0x23, 0x30, 0xc2, 0xd5, 0x65, 0xb0, 0xac, 0x03, 0xd8, 0x30, 0xaa, 0x51, 0xfb, 0x76, 0xf7,
	0x0a, 0x8c, 0xa6, 0x10, 0x02, 0xdf, 0xc2, 0xa8, 0xaa, 0x94, 0xc6, 0x27, 0xf7, 0x8a, 0x54, 0x5f,
	0xf7, 0x6c, 0x72, 0xb8, 0xa8, 0x64, 0xb4, 0xe3, 0xfe, 0x26, 0x30, 0x6c, 0x9a, 0x81, 0xaf, 0x61,
	0xa0, 0x72, 0xdc, 0x71, 0xd9, 0x2f, 0xcd, 0xec, 0xd1, 0xfe, 0xb5, 0xb4, 0xd7, 0x47, 0x3b, 0xf8,
	0x0e, 0x8c, 0xf6, 0x53, 0xfe, 0x5b, 0xfe, 0xf8, 0x5e, 0x2f, 0xfe, 0x3a, 0x7c, 0x3f, 0x52, 0xbf,
	0xd1, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x40, 0xf6, 0x45, 0x57, 0x03, 0x00, 0x00,
}
