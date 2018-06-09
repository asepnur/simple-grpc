// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package user

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

type Hello struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_b4667de6bf20d4b7, []int{0}
}
func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (dst *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(dst, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*Hello)(nil), "user.Hello")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	Ping(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Ping(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/user.Greeter/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	Ping(context.Context, *Hello) (*Hello, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Greeter/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Ping(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Greeter_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/grpc.proto",
}

func init() { proto.RegisterFile("grpc/grpc.proto", fileDescriptor_grpc_b4667de6bf20d4b7) }

var fileDescriptor_grpc_b4667de6bf20d4b7 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x2a, 0x48,
	0xd6, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0xa5, 0xc5, 0xa9, 0x45, 0x4a,
	0xca, 0x5c, 0xac, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x42, 0x52, 0x5c, 0x1c, 0xe9, 0x45, 0xa9, 0xa9,
	0x25, 0x99, 0x79, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x91, 0x2e, 0x17,
	0xbb, 0x3b, 0x88, 0x9d, 0x5a, 0x24, 0xa4, 0xc4, 0xc5, 0x12, 0x90, 0x99, 0x97, 0x2e, 0xc4, 0xad,
	0x07, 0xd2, 0xae, 0x07, 0xd6, 0x2b, 0x85, 0xcc, 0x51, 0x62, 0x70, 0xd2, 0xe5, 0x92, 0xc9, 0xcc,
	0xd7, 0x03, 0x5b, 0x95, 0x5a, 0x91, 0x98, 0x5b, 0x90, 0x93, 0x5a, 0xac, 0x57, 0x9c, 0x09, 0xa2,
	0x75, 0x41, 0x82, 0x4e, 0xdc, 0x48, 0x9c, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0x7b, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x23, 0x7d, 0xae, 0xa2, 0x00, 0x00, 0x00,
}