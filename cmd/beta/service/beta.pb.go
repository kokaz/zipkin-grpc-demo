// Code generated by protoc-gen-go.
// source: service/beta.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service/beta.proto

It has these top-level messages:
	GetRequest
	GetResponse
	SetRequest
	SetResponse
*/
package service

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

type GetRequest struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetResponse struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SetRequest struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *SetRequest) Reset()                    { *m = SetRequest{} }
func (m *SetRequest) String() string            { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()               {}
func (*SetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SetResponse struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *SetResponse) Reset()                    { *m = SetResponse{} }
func (m *SetResponse) String() string            { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()               {}
func (*SetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*GetRequest)(nil), "beta.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "beta.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "beta.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "beta.SetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for BetaService service

type BetaServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
}

type betaServiceClient struct {
	cc *grpc.ClientConn
}

func NewBetaServiceClient(cc *grpc.ClientConn) BetaServiceClient {
	return &betaServiceClient{cc}
}

func (c *betaServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/beta.BetaService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betaServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := grpc.Invoke(ctx, "/beta.BetaService/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BetaService service

type BetaServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
}

func RegisterBetaServiceServer(s *grpc.Server, srv BetaServiceServer) {
	s.RegisterService(&_BetaService_serviceDesc, srv)
}

func _BetaService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetaServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beta.BetaService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetaServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetaService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetaServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beta.BetaService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetaServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BetaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beta.BetaService",
	HandlerType: (*BetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BetaService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _BetaService_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("service/beta.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x4a, 0x2d, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0xb1, 0x95, 0x14, 0xb8, 0xb8, 0xdc, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x25, 0x45, 0x2e, 0x6e, 0xb0, 0x8a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xac, 0x4a, 0x14,
	0xb8, 0xb8, 0x82, 0x09, 0x1a, 0x12, 0x8c, 0xdf, 0x10, 0xa3, 0x54, 0x2e, 0x6e, 0xa7, 0xd4, 0x92,
	0xc4, 0x60, 0x88, 0x4b, 0x85, 0xb4, 0xb8, 0x98, 0xdd, 0x53, 0x4b, 0x84, 0x04, 0xf4, 0xc0, 0x4e,
	0x46, 0xb8, 0x51, 0x4a, 0x10, 0x49, 0x04, 0x6a, 0x9c, 0x16, 0x17, 0x73, 0x30, 0x42, 0x6d, 0x30,
	0x86, 0x5a, 0x24, 0xab, 0x9d, 0x38, 0xa3, 0xd8, 0xa1, 0x81, 0x91, 0xc4, 0x06, 0x0e, 0x08, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x24, 0xed, 0xba, 0x1e, 0x01, 0x00, 0x00,
}
