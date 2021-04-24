// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route_guide.grpc

package routeguide

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
// is compatible with the grpc package it is being compiled against.
// A compilation error at this line likely means your copy of the
// grpc package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the grpc package

// 0.Ping request
type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

// 0.Ping reply
type PingReply struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{1}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "routeguide.PingRequest")
	proto.RegisterType((*PingReply)(nil), "routeguide.PingReply")
}

func init() { proto.RegisterFile("route_guide.grpc", fileDescriptor_b7d679f20da65b7b) }

var fileDescriptor_b7d679f20da65b7b = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xca, 0x2f, 0x2d,
	0x49, 0x8d, 0x4f, 0x2f, 0xcd, 0x4c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02,
	0x0b, 0x81, 0x45, 0x94, 0x78, 0xb9, 0xb8, 0x03, 0x32, 0xf3, 0xd2, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x94, 0x14, 0xb9, 0x38, 0x21, 0xdc, 0x82, 0x9c, 0x4a, 0x21, 0x11, 0x2e, 0xd6, 0x22,
	0x10, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x31, 0x72, 0xe3, 0xe2, 0x0a, 0x02,
	0xe9, 0x77, 0x07, 0xe9, 0x17, 0xb2, 0xe0, 0x62, 0x01, 0x69, 0x10, 0x12, 0xd7, 0x43, 0x18, 0xaa,
	0x87, 0x64, 0xa2, 0x94, 0x28, 0xa6, 0x44, 0x41, 0x4e, 0xa5, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x31,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0xa4, 0xe2, 0xbe, 0xa1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteGuideClient interface {
	// 检查服务可用性 正常情况下会返回字符串 pong
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteGuideServer is the server API for RouteGuide service.
type RouteGuideServer interface {
	// 检查服务可用性 正常情况下会返回字符串 pong
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

// UnimplementedRouteGuideServer can be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (*UnimplementedRouteGuideServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RouteGuide_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route_guide.grpc",
}
