// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/commentpost.proto

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

type CommentPostMessage struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Images               [][]byte `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	PostId               int32    `protobuf:"varint,4,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentPostMessage) Reset()         { *m = CommentPostMessage{} }
func (m *CommentPostMessage) String() string { return proto.CompactTextString(m) }
func (*CommentPostMessage) ProtoMessage()    {}
func (*CommentPostMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_576dce2973bf60db, []int{0}
}

func (m *CommentPostMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentPostMessage.Unmarshal(m, b)
}
func (m *CommentPostMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentPostMessage.Marshal(b, m, deterministic)
}
func (m *CommentPostMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentPostMessage.Merge(m, src)
}
func (m *CommentPostMessage) XXX_Size() int {
	return xxx_messageInfo_CommentPostMessage.Size(m)
}
func (m *CommentPostMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentPostMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CommentPostMessage proto.InternalMessageInfo

func (m *CommentPostMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CommentPostMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *CommentPostMessage) GetImages() [][]byte {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *CommentPostMessage) GetPostId() int32 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type CommentPostResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentPostResponse) Reset()         { *m = CommentPostResponse{} }
func (m *CommentPostResponse) String() string { return proto.CompactTextString(m) }
func (*CommentPostResponse) ProtoMessage()    {}
func (*CommentPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_576dce2973bf60db, []int{1}
}

func (m *CommentPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommentPostResponse.Unmarshal(m, b)
}
func (m *CommentPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommentPostResponse.Marshal(b, m, deterministic)
}
func (m *CommentPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentPostResponse.Merge(m, src)
}
func (m *CommentPostResponse) XXX_Size() int {
	return xxx_messageInfo_CommentPostResponse.Size(m)
}
func (m *CommentPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentPostResponse proto.InternalMessageInfo

func (m *CommentPostResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*CommentPostMessage)(nil), "pupapi.CommentPostMessage")
	proto.RegisterType((*CommentPostResponse)(nil), "pupapi.CommentPostResponse")
}

func init() { proto.RegisterFile("proto/commentpost.proto", fileDescriptor_576dce2973bf60db) }

var fileDescriptor_576dce2973bf60db = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0xc5, 0x40,
	0x10, 0x85, 0x8d, 0xf7, 0x66, 0xa3, 0xa3, 0xd5, 0x28, 0x66, 0x89, 0x4d, 0x48, 0x95, 0x2a, 0x01,
	0x7d, 0x04, 0x1b, 0x2d, 0x04, 0xd9, 0xc6, 0x52, 0xf2, 0x33, 0x84, 0x05, 0x93, 0x59, 0x9c, 0x4d,
	0xe1, 0xdb, 0x4b, 0x36, 0x06, 0x22, 0xde, 0x6e, 0xbf, 0x73, 0x58, 0x3e, 0xce, 0x40, 0xea, 0xbe,
	0xd8, 0x73, 0xdd, 0xf1, 0x38, 0xd2, 0xe4, 0x1d, 0x8b, 0xaf, 0x42, 0x82, 0xca, 0xcd, 0xae, 0x71,
	0xb6, 0x60, 0xc0, 0xa7, 0xb5, 0x7c, 0x63, 0xf1, 0xaf, 0x24, 0xd2, 0x0c, 0x84, 0xb7, 0x10, 0x7b,
	0xeb, 0x3f, 0x49, 0x47, 0x79, 0x54, 0x5e, 0x9a, 0x15, 0x10, 0xe1, 0xd8, 0x72, 0xff, 0xad, 0xcf,
	0x43, 0x18, 0xde, 0x78, 0x07, 0xca, 0x8e, 0xcd, 0x40, 0xa2, 0x0f, 0xf9, 0xa1, 0xbc, 0x36, 0xbf,
	0x84, 0x29, 0x24, 0x8b, 0xed, 0xc3, 0xf6, 0xfa, 0x98, 0x47, 0x65, 0x6c, 0xd4, 0x82, 0x2f, 0x7d,
	0x51, 0xc3, 0xcd, 0x4e, 0x68, 0x48, 0x1c, 0x4f, 0x42, 0xa8, 0x21, 0x91, 0xb9, 0xeb, 0x48, 0x24,
	0x38, 0x2f, 0xcc, 0x86, 0x0f, 0xef, 0x70, 0xb5, 0xfb, 0x80, 0xcf, 0x7f, 0x31, 0xab, 0xd6, 0x21,
	0xd5, 0xff, 0x15, 0xd9, 0xfd, 0x89, 0x6e, 0x13, 0x16, 0x67, 0xad, 0x0a, 0x97, 0x78, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0xbf, 0xc7, 0xd2, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentPostClient is the client API for CommentPost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentPostClient interface {
	CommentPost(ctx context.Context, in *CommentPostMessage, opts ...grpc.CallOption) (*CommentPostResponse, error)
}

type commentPostClient struct {
	cc *grpc.ClientConn
}

func NewCommentPostClient(cc *grpc.ClientConn) CommentPostClient {
	return &commentPostClient{cc}
}

func (c *commentPostClient) CommentPost(ctx context.Context, in *CommentPostMessage, opts ...grpc.CallOption) (*CommentPostResponse, error) {
	out := new(CommentPostResponse)
	err := c.cc.Invoke(ctx, "/pupapi.CommentPost/CommentPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentPostServer is the server API for CommentPost service.
type CommentPostServer interface {
	CommentPost(context.Context, *CommentPostMessage) (*CommentPostResponse, error)
}

// UnimplementedCommentPostServer can be embedded to have forward compatible implementations.
type UnimplementedCommentPostServer struct {
}

func (*UnimplementedCommentPostServer) CommentPost(ctx context.Context, req *CommentPostMessage) (*CommentPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentPost not implemented")
}

func RegisterCommentPostServer(s *grpc.Server, srv CommentPostServer) {
	s.RegisterService(&_CommentPost_serviceDesc, srv)
}

func _CommentPost_CommentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentPostMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentPostServer).CommentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pupapi.CommentPost/CommentPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentPostServer).CommentPost(ctx, req.(*CommentPostMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentPost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pupapi.CommentPost",
	HandlerType: (*CommentPostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentPost",
			Handler:    _CommentPost_CommentPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/commentpost.proto",
}
