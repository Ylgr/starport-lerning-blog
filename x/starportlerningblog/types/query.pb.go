// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: starportlerningblog/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("starportlerningblog/query.proto", fileDescriptor_740b3da3db25b20b) }

var fileDescriptor_740b3da3db25b20b = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0xae, 0x82, 0x30,
	0x14, 0x86, 0x61, 0xb8, 0xf7, 0x26, 0x8c, 0x77, 0x24, 0xa6, 0xee, 0x46, 0x38, 0x41, 0x9f, 0x40,
	0xdf, 0xc0, 0x55, 0xa7, 0x96, 0x34, 0xb5, 0x49, 0xe9, 0xa9, 0xed, 0xc1, 0xc8, 0x5b, 0xf8, 0x58,
	0x8e, 0x8c, 0x8e, 0x06, 0x5e, 0xc4, 0x00, 0xba, 0xb1, 0x9e, 0x7c, 0x39, 0xff, 0xf7, 0x25, 0xcb,
	0x40, 0xdc, 0x3b, 0xf4, 0x64, 0xa4, 0xb7, 0xda, 0x2a, 0x61, 0x50, 0xc1, 0xa5, 0x96, 0xbe, 0xc9,
	0x9d, 0x47, 0xc2, 0xff, 0x75, 0x63, 0x94, 0xcf, 0x67, 0xa8, 0xb9, 0x5b, 0xba, 0x50, 0x88, 0xca,
	0x48, 0xe0, 0x4e, 0x03, 0xb7, 0x16, 0x89, 0x93, 0x46, 0x1b, 0xa6, 0x5f, 0xe9, 0xaa, 0xc4, 0x50,
	0x61, 0x00, 0xc1, 0x83, 0x9c, 0x46, 0xe0, 0x5a, 0x08, 0x49, 0xbc, 0x00, 0xc7, 0x95, 0xb6, 0x23,
	0x3c, 0xb1, 0x9b, 0xbf, 0xe4, 0xe7, 0x30, 0x10, 0xfb, 0xd3, 0xa3, 0x63, 0x71, 0xdb, 0xb1, 0xf8,
	0xd5, 0xb1, 0xf8, 0xde, 0xb3, 0xa8, 0xed, 0x59, 0xf4, 0xec, 0x59, 0x74, 0xdc, 0x29, 0x4d, 0xe7,
	0x5a, 0xe4, 0x25, 0x56, 0x30, 0x58, 0xc2, 0xd7, 0x28, 0xfb, 0x28, 0x65, 0x63, 0xcd, 0x0d, 0xe6,
	0x1a, 0xa9, 0x71, 0x32, 0x88, 0xdf, 0x71, 0x6c, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x11, 0x02,
	0x82, 0x0c, 0x07, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ylgr.starportlerningblog.starportlerningblog.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "starportlerningblog/query.proto",
}
