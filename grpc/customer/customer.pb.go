// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

/*
Package customer is a generated protocol buffer package.

It is generated from these files:
	customer.proto

It has these top-level messages:
	Address
	Customers
	Customer
	Response
	Filter
*/
package customer

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

type Address struct {
	Street            string `proto:"bytes,1,opt,name=street" json:"street,omitempty"`
	City              string `proto:"bytes,2,opt,name=city" json:"city,omitempty"`
	State             string `proto:"bytes,3,opt,name=state" json:"state,omitempty"`
	Zip               string `proto:"bytes,4,opt,name=zip" json:"zip,omitempty"`
	IsShippingAddress bool   `proto:"varint,5,opt,name=isShippingAddress" json:"isShippingAddress,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
}

type Customers struct {
	Data []*Customer `proto:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *Customers) Reset()                    { *m = Customers{} }
func (m *Customers) String() string            { return proto.CompactTextString(m) }
func (*Customers) ProtoMessage()               {}
func (*Customers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Customers) GetData() []*Customer {
	if m != nil {
		return m.Data
	}
	return nil
}

// Request message for creating a new customer
type Customer struct {
	Id        int32      `proto:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string     `proto:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email     string     `proto:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone     string     `proto:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Addresses []*Address `proto:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *Customer) Reset()                    { *m = Customer{} }
func (m *Customer) String() string            { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Customer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Customer) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Customer) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Response struct {
	Id      int32 `proto:"varint,1,opt,name=id" json:"id,omitempty"`
	Success bool  `proto:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Filter struct {
	Keyword string `proto:"bytes,1,opt,name=keyword" json:"keyword,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Filter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*Address)(nil), "customer.Address")
	proto.RegisterType((*Customers)(nil), "customer.Customers")
	proto.RegisterType((*Customer)(nil), "customer.Customer")
	proto.RegisterType((*Response)(nil), "customer.Response")
	proto.RegisterType((*Filter)(nil), "customer.Filter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Grpcservice service

type GrpcserviceClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Customers, error)
	// Create a new Customer - A simple RPC
	CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Response, error)
}

type grpcserviceClient struct {
	cc *grpc.ClientConn
}

func NewGrpcserviceClient(cc *grpc.ClientConn) GrpcserviceClient {
	return &grpcserviceClient{cc}
}

func (c *grpcserviceClient) GetCustomers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Customers, error) {
	out := new(Customers)
	err := grpc.Invoke(ctx, "/customer.grpcservice/GetCustomers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcserviceClient) CreateCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/customer.grpcservice/CreateCustomer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Grpcservice service

type GrpcserviceServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	GetCustomers(context.Context, *Filter) (*Customers, error)
	// Create a new Customer - A simple RPC
	CreateCustomer(context.Context, *Customer) (*Response, error)
}

func RegisterGrpcserviceServer(s *grpc.Server, srv GrpcserviceServer) {
	s.RegisterService(&_Grpcservice_serviceDesc, srv)
}

func _Grpcservice_GetCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcserviceServer).GetCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.grpcservice/GetCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcserviceServer).GetCustomers(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpcservice_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcserviceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.grpcservice/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcserviceServer).CreateCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpcservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.grpcservice",
	HandlerType: (*GrpcserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomers",
			Handler:    _Grpcservice_GetCustomers_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _Grpcservice_CreateCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xbb, 0xdb, 0x7f, 0xdb, 0xa9, 0x94, 0x76, 0x14, 0x09, 0x3d, 0x95, 0x1c, 0xa4, 0x07,
	0xa9, 0xd0, 0x0a, 0x82, 0x37, 0x29, 0xe8, 0x7d, 0xfd, 0x04, 0x71, 0x77, 0x68, 0x83, 0xed, 0x6e,
	0x48, 0x52, 0xa5, 0x1e, 0xbd, 0xfa, 0xa5, 0x25, 0xd9, 0x4d, 0x57, 0xa8, 0xb7, 0x79, 0x6f, 0x86,
	0xe4, 0x37, 0x2f, 0x81, 0x51, 0x76, 0x30, 0xb6, 0xdc, 0x93, 0x5e, 0x28, 0x5d, 0xda, 0x12, 0x93,
	0xa0, 0xf9, 0x4f, 0x04, 0xfd, 0xa7, 0x3c, 0xd7, 0x64, 0x0c, 0x5e, 0x43, 0xcf, 0x58, 0x4d, 0x64,
	0x59, 0x34, 0x8b, 0xe6, 0x83, 0xb4, 0x56, 0x88, 0xd0, 0xc9, 0xa4, 0x3d, 0xb2, 0xd8, 0xbb, 0xbe,
	0xc6, 0x2b, 0xe8, 0x1a, 0x2b, 0x2c, 0xb1, 0xb6, 0x37, 0x2b, 0x81, 0x63, 0x68, 0x7f, 0x49, 0xc5,
	0x3a, 0xde, 0x73, 0x25, 0xde, 0xc2, 0x44, 0x9a, 0xd7, 0xad, 0x54, 0x4a, 0x16, 0x9b, 0xfa, 0x22,
	0xd6, 0x9d, 0x45, 0xf3, 0x24, 0x3d, 0x6f, 0xf0, 0x15, 0x0c, 0xd6, 0x35, 0x99, 0xc1, 0x1b, 0xe8,
	0xe4, 0xc2, 0x0a, 0x16, 0xcd, 0xda, 0xf3, 0xe1, 0x12, 0x17, 0xa7, 0x1d, 0xc2, 0x48, 0xea, 0xfb,
	0x6e, 0x85, 0x24, 0x58, 0x38, 0x82, 0x58, 0xe6, 0x9e, 0xbf, 0x9b, 0xc6, 0x32, 0x77, 0xec, 0x85,
	0xd8, 0x53, 0x60, 0x77, 0xb5, 0x63, 0xa7, 0xbd, 0x90, 0xbb, 0xc0, 0xee, 0x85, 0x73, 0xd5, 0xb6,
	0x2c, 0xa8, 0xa6, 0xaf, 0x04, 0xde, 0xc1, 0x40, 0x54, 0x70, 0xe4, 0xb8, 0x1d, 0xc9, 0xa4, 0x21,
	0xa9, 0xb9, 0xd3, 0x66, 0x86, 0xdf, 0x43, 0x92, 0x92, 0x51, 0x65, 0x61, 0xe8, 0x0c, 0x86, 0x41,
	0xdf, 0x1c, 0xb2, 0xcc, 0x45, 0x10, 0xfb, 0x08, 0x82, 0xe4, 0x1c, 0x7a, 0xcf, 0x72, 0x67, 0x49,
	0xbb, 0x99, 0x77, 0x3a, 0x7e, 0x96, 0x3a, 0xaf, 0x5f, 0x21, 0xc8, 0xe5, 0x77, 0x04, 0xc3, 0x8d,
	0x56, 0x99, 0x21, 0xfd, 0x21, 0x33, 0xc2, 0x07, 0xb8, 0x78, 0x21, 0xdb, 0xe4, 0x35, 0x6e, 0xb8,
	0xaa, 0xb3, 0xa6, 0x97, 0xe7, 0x99, 0x19, 0xde, 0xc2, 0x47, 0x18, 0xad, 0x35, 0x09, 0x4b, 0xa7,
	0xd4, 0xfe, 0x09, 0x77, 0xfa, 0xc7, 0x0b, 0x0b, 0xf1, 0xd6, 0x5b, 0xcf, 0x7f, 0xa0, 0xd5, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x58, 0xcb, 0x13, 0x52, 0x02, 0x00, 0x00,
}
