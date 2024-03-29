// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/newpost.proto

package pupapi

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

type NewPostMessage struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Images               [][]byte `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPostMessage) Reset()         { *m = NewPostMessage{} }
func (m *NewPostMessage) String() string { return proto.CompactTextString(m) }
func (*NewPostMessage) ProtoMessage()    {}
func (*NewPostMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f90978f11127b201, []int{0}
}

func (m *NewPostMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPostMessage.Unmarshal(m, b)
}
func (m *NewPostMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPostMessage.Marshal(b, m, deterministic)
}
func (m *NewPostMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPostMessage.Merge(m, src)
}
func (m *NewPostMessage) XXX_Size() int {
	return xxx_messageInfo_NewPostMessage.Size(m)
}
func (m *NewPostMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPostMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NewPostMessage proto.InternalMessageInfo

func (m *NewPostMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NewPostMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *NewPostMessage) GetImages() [][]byte {
	if m != nil {
		return m.Images
	}
	return nil
}

type NewPostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewPostResponse) Reset()         { *m = NewPostResponse{} }
func (m *NewPostResponse) String() string { return proto.CompactTextString(m) }
func (*NewPostResponse) ProtoMessage()    {}
func (*NewPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f90978f11127b201, []int{1}
}

func (m *NewPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewPostResponse.Unmarshal(m, b)
}
func (m *NewPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewPostResponse.Marshal(b, m, deterministic)
}
func (m *NewPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewPostResponse.Merge(m, src)
}
func (m *NewPostResponse) XXX_Size() int {
	return xxx_messageInfo_NewPostResponse.Size(m)
}
func (m *NewPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewPostResponse proto.InternalMessageInfo

func (m *NewPostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*NewPostMessage)(nil), "pupapi.NewPostMessage")
	proto.RegisterType((*NewPostResponse)(nil), "pupapi.NewPostResponse")
}

func init() { proto.RegisterFile("proto/newpost.proto", fileDescriptor_f90978f11127b201) }

var fileDescriptor_f90978f11127b201 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4b, 0x2d, 0x2f, 0xc8, 0x2f, 0x2e, 0xd1, 0x03, 0xf3, 0x84, 0xd8, 0x0a, 0x4a,
	0x0b, 0x12, 0x0b, 0x32, 0x95, 0x82, 0xb8, 0xf8, 0xfc, 0x52, 0xcb, 0x03, 0xf2, 0x8b, 0x4b, 0x7c,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a,
	0x09, 0x26, 0xb0, 0x20, 0x98, 0x2d, 0x24, 0xc6, 0xc5, 0x96, 0x99, 0x9b, 0x98, 0x9e, 0x5a, 0x2c,
	0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x13, 0x04, 0xe5, 0x29, 0x69, 0x73, 0xf1, 0x43, 0xcd, 0x0c, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d,
	0x2e, 0x06, 0x1b, 0xcb, 0x11, 0x04, 0xe3, 0x1a, 0xb9, 0x73, 0xb1, 0x43, 0x15, 0x0b, 0xd9, 0x20,
	0x98, 0x62, 0x7a, 0x10, 0xf7, 0xe9, 0xa1, 0x3a, 0x4e, 0x4a, 0x1c, 0x4d, 0x1c, 0x66, 0x81, 0x12,
	0x43, 0x12, 0x1b, 0xd8, 0x63, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x40, 0xee, 0x29,
	0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NewPostClient is the client API for NewPost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewPostClient interface {
	NewPost(ctx context.Context, in *NewPostMessage, opts ...grpc.CallOption) (*NewPostResponse, error)
}

type newPostClient struct {
	cc *grpc.ClientConn
}

func NewNewPostClient(cc *grpc.ClientConn) NewPostClient {
	return &newPostClient{cc}
}

func (c *newPostClient) NewPost(ctx context.Context, in *NewPostMessage, opts ...grpc.CallOption) (*NewPostResponse, error) {
	out := new(NewPostResponse)
	err := c.cc.Invoke(ctx, "/pupapi.NewPost/NewPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewPostServer is the server API for NewPost service.
type NewPostServer interface {
	NewPost(context.Context, *NewPostMessage) (*NewPostResponse, error)
}

// UnimplementedNewPostServer can be embedded to have forward compatible implementations.
type UnimplementedNewPostServer struct {
}

func (*UnimplementedNewPostServer) NewPost(ctx context.Context, req *NewPostMessage) (*NewPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPost not implemented")
}

func RegisterNewPostServer(s *grpc.Server, srv NewPostServer) {
	s.RegisterService(&_NewPost_serviceDesc, srv)
}

func _NewPost_NewPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPostMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewPostServer).NewPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pupapi.NewPost/NewPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewPostServer).NewPost(ctx, req.(*NewPostMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewPost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pupapi.NewPost",
	HandlerType: (*NewPostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPost",
			Handler:    _NewPost_NewPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/newpost.proto",
}
