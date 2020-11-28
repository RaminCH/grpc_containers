// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consigment/consigment.proto

/*
Package consigment is a generated protocol buffer package.

It is generated from these files:
	proto/consigment/consigment.proto

It has these top-level messages:
	Container
	Command
	Response
*/
package consigment

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

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Command struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	Vessel      string       `protobuf:"bytes,5,opt,name=vessel" json:"vessel,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Command) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Command) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Command) GetVessel() string {
	if m != nil {
		return m.Vessel
	}
	return ""
}

type Response struct {
	Created bool     `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Command *Command `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func init() {
	proto.RegisterType((*Container)(nil), "consigment.Container")
	proto.RegisterType((*Command)(nil), "consigment.Command")
	proto.RegisterType((*Response)(nil), "consigment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/consigment.ShippingService/CreateCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceServer interface {
	CreateCommand(context.Context, *Command) (*Response, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consigment.ShippingService/CreateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consigment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommand",
			Handler:    _ShippingService_CreateCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consigment/consigment.proto",
}

func init() { proto.RegisterFile("proto/consigment/consigment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0xfc, 0xfa, 0x9b, 0x76, 0xa3, 0x0f, 0x24, 0xf3, 0x67, 0x71, 0xa1, 0xe4, 0xd4, 0x0b, 0x45,
	0x2a, 0xe2, 0xc6, 0x2d, 0xa7, 0x9e, 0x90, 0x9c, 0x07, 0x40, 0xc1, 0x5e, 0xa5, 0x2b, 0x11, 0x3b,
	0xb2, 0xdd, 0xf2, 0x3e, 0x3c, 0x29, 0x8a, 0x9d, 0x40, 0x10, 0xdc, 0x3c, 0x33, 0xeb, 0x9d, 0x19,
	0x1b, 0x6e, 0x1b, 0x6b, 0xbc, 0xb9, 0x97, 0x46, 0x3b, 0xaa, 0x6a, 0xd4, 0x7e, 0x70, 0xdc, 0x04,
	0x8d, 0xc1, 0x37, 0x93, 0xd5, 0xb0, 0xcc, 0x8d, 0xf6, 0x25, 0x69, 0xb4, 0xec, 0x04, 0xc6, 0xa4,
	0xf8, 0x68, 0x35, 0x5a, 0x2f, 0xc5, 0x98, 0x14, 0xbb, 0x81, 0x54, 0x1e, 0x9c, 0x37, 0x35, 0xda,
	0x17, 0x52, 0x7c, 0x1c, 0x04, 0xe8, 0xa9, 0x9d, 0x62, 0x97, 0x30, 0x37, 0x96, 0x2a, 0xd2, 0x7c,
	0x12, 0xb4, 0x0e, 0xb1, 0x2b, 0x48, 0x0e, 0x2e, 0x5e, 0x9a, 0x46, 0xa1, 0x85, 0x3b, 0x95, 0x7d,
	0x8c, 0x20, 0xc9, 0x4d, 0x5d, 0x97, 0x5a, 0xfd, 0x72, 0x5b, 0x41, 0xaa, 0xd0, 0x49, 0x4b, 0x8d,
	0x27, 0xa3, 0x3b, 0xb7, 0x21, 0xd5, 0xda, 0xbd, 0x23, 0x55, 0x7b, 0x1f, 0xec, 0x66, 0xa2, 0x43,
	0xec, 0x11, 0xda, 0x4a, 0xb1, 0x84, 0xe3, 0xd3, 0xd5, 0x64, 0x9d, 0x6e, 0x2f, 0x36, 0x83, 0xde,
	0x5f, 0x15, 0xc5, 0x60, 0xb0, 0x5d, 0x77, 0x44, 0xe7, 0xf0, 0x8d, 0xcf, 0x62, 0xc8, 0x88, 0xb2,
	0x02, 0x16, 0x02, 0x5d, 0x63, 0xb4, 0x43, 0xc6, 0x21, 0x91, 0x16, 0x4b, 0x8f, 0x31, 0xe9, 0x42,
	0xf4, 0x90, 0xdd, 0x41, 0x22, 0x63, 0x93, 0x10, 0x35, 0xdd, 0x9e, 0xfd, 0x74, 0x0c, 0x92, 0xe8,
	0x67, 0xb6, 0xcf, 0x70, 0x5a, 0xec, 0xa9, 0x69, 0x48, 0x57, 0x05, 0xda, 0x23, 0x49, 0x64, 0x4f,
	0xf0, 0x3f, 0x0f, 0xcb, 0xfa, 0x17, 0xf9, 0x6b, 0xc3, 0xf5, 0xf9, 0x90, 0xec, 0x73, 0x65, 0xff,
	0x5e, 0xe7, 0xe1, 0x33, 0x1f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x05, 0x48, 0xa1, 0x5f, 0xf1,
	0x01, 0x00, 0x00,
}