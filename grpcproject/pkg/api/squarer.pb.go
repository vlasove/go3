// Code generated by protoc-gen-go. DO NOT EDIT.
// source: squarer.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SquareRequest struct {
	Arg                  int32    `protobuf:"varint,1,opt,name=arg,proto3" json:"arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRequest) Reset()         { *m = SquareRequest{} }
func (m *SquareRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRequest) ProtoMessage()    {}
func (*SquareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af8d901bd6407ab, []int{0}
}

func (m *SquareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRequest.Unmarshal(m, b)
}
func (m *SquareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRequest.Marshal(b, m, deterministic)
}
func (m *SquareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRequest.Merge(m, src)
}
func (m *SquareRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRequest.Size(m)
}
func (m *SquareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRequest proto.InternalMessageInfo

func (m *SquareRequest) GetArg() int32 {
	if m != nil {
		return m.Arg
	}
	return 0
}

type SquareResponse struct {
	Answer               int32    `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareResponse) Reset()         { *m = SquareResponse{} }
func (m *SquareResponse) String() string { return proto.CompactTextString(m) }
func (*SquareResponse) ProtoMessage()    {}
func (*SquareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0af8d901bd6407ab, []int{1}
}

func (m *SquareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareResponse.Unmarshal(m, b)
}
func (m *SquareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareResponse.Marshal(b, m, deterministic)
}
func (m *SquareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareResponse.Merge(m, src)
}
func (m *SquareResponse) XXX_Size() int {
	return xxx_messageInfo_SquareResponse.Size(m)
}
func (m *SquareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareResponse proto.InternalMessageInfo

func (m *SquareResponse) GetAnswer() int32 {
	if m != nil {
		return m.Answer
	}
	return 0
}

func init() {
	proto.RegisterType((*SquareRequest)(nil), "api.SquareRequest")
	proto.RegisterType((*SquareResponse)(nil), "api.SquareResponse")
}

func init() { proto.RegisterFile("squarer.proto", fileDescriptor_0af8d901bd6407ab) }

var fileDescriptor_0af8d901bd6407ab = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x2c, 0x4d,
	0x2c, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52,
	0xe4, 0xe2, 0x0d, 0x06, 0x8b, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x70, 0x31,
	0x27, 0x16, 0xa5, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98, 0x4a, 0x1a, 0x5c, 0x7c,
	0x30, 0x25, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0x79, 0xc5, 0xe5,
	0xa9, 0x45, 0x50, 0x65, 0x50, 0x9e, 0x91, 0x1d, 0x17, 0x3b, 0x44, 0x65, 0x91, 0x90, 0x31, 0x17,
	0x1b, 0x84, 0x29, 0x24, 0xa4, 0x97, 0x58, 0x90, 0xa9, 0x87, 0x62, 0x89, 0x94, 0x30, 0x8a, 0x18,
	0xc4, 0x54, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xc3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31,
	0xa6, 0x28, 0xf8, 0xa9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SquarerClient is the client API for Squarer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquarerClient interface {
	Square(ctx context.Context, in *SquareRequest, opts ...grpc.CallOption) (*SquareResponse, error)
}

type squarerClient struct {
	cc grpc.ClientConnInterface
}

func NewSquarerClient(cc grpc.ClientConnInterface) SquarerClient {
	return &squarerClient{cc}
}

func (c *squarerClient) Square(ctx context.Context, in *SquareRequest, opts ...grpc.CallOption) (*SquareResponse, error) {
	out := new(SquareResponse)
	err := c.cc.Invoke(ctx, "/api.Squarer/Square", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquarerServer is the server API for Squarer service.
type SquarerServer interface {
	Square(context.Context, *SquareRequest) (*SquareResponse, error)
}

// UnimplementedSquarerServer can be embedded to have forward compatible implementations.
type UnimplementedSquarerServer struct {
}

func (*UnimplementedSquarerServer) Square(ctx context.Context, req *SquareRequest) (*SquareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Square not implemented")
}

func RegisterSquarerServer(s *grpc.Server, srv SquarerServer) {
	s.RegisterService(&_Squarer_serviceDesc, srv)
}

func _Squarer_Square_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquarerServer).Square(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Squarer/Square",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquarerServer).Square(ctx, req.(*SquareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Squarer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Squarer",
	HandlerType: (*SquarerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Square",
			Handler:    _Squarer_Square_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squarer.proto",
}
