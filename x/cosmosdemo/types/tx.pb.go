// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmosdemo/cosmosdemo/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cosmosdemo/cosmosdemo/tx.proto", fileDescriptor_664a17a35edc4952) }

var fileDescriptor_664a17a35edc4952 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x4e, 0x49, 0xcd, 0xcd, 0xd7, 0x47, 0x62, 0x96, 0x54, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x29, 0x66, 0xe4, 0x17, 0xa7, 0x26, 0xa6, 0xa4, 0x96, 0xe5, 0x17, 0x14, 0xeb, 0x21,
	0x14, 0x20, 0x31, 0x8d, 0x58, 0xb9, 0x98, 0x7d, 0x8b, 0xd3, 0x9d, 0xfc, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0xc9, 0x38, 0x64, 0xfb, 0x2a, 0x50, 0x2c, 0xaf, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0x3b, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x81, 0x5c, 0x92, 0xa2, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hoseadevops.cosmosdemo.cosmosdemo.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cosmosdemo/cosmosdemo/tx.proto",
}
