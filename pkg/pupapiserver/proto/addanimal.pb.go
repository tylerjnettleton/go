// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/addanimal.proto

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

type AddAnimalMessage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bio                  string   `protobuf:"bytes,2,opt,name=bio,proto3" json:"bio,omitempty"`
	ProfileImage         []byte   `protobuf:"bytes,3,opt,name=profile_image,json=profileImage,proto3" json:"profile_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAnimalMessage) Reset()         { *m = AddAnimalMessage{} }
func (m *AddAnimalMessage) String() string { return proto.CompactTextString(m) }
func (*AddAnimalMessage) ProtoMessage()    {}
func (*AddAnimalMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f618ab4a7e8ed1ca, []int{0}
}

func (m *AddAnimalMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAnimalMessage.Unmarshal(m, b)
}
func (m *AddAnimalMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAnimalMessage.Marshal(b, m, deterministic)
}
func (m *AddAnimalMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAnimalMessage.Merge(m, src)
}
func (m *AddAnimalMessage) XXX_Size() int {
	return xxx_messageInfo_AddAnimalMessage.Size(m)
}
func (m *AddAnimalMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAnimalMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddAnimalMessage proto.InternalMessageInfo

func (m *AddAnimalMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddAnimalMessage) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *AddAnimalMessage) GetProfileImage() []byte {
	if m != nil {
		return m.ProfileImage
	}
	return nil
}

type AddAnimalResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAnimalResponse) Reset()         { *m = AddAnimalResponse{} }
func (m *AddAnimalResponse) String() string { return proto.CompactTextString(m) }
func (*AddAnimalResponse) ProtoMessage()    {}
func (*AddAnimalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f618ab4a7e8ed1ca, []int{1}
}

func (m *AddAnimalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAnimalResponse.Unmarshal(m, b)
}
func (m *AddAnimalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAnimalResponse.Marshal(b, m, deterministic)
}
func (m *AddAnimalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAnimalResponse.Merge(m, src)
}
func (m *AddAnimalResponse) XXX_Size() int {
	return xxx_messageInfo_AddAnimalResponse.Size(m)
}
func (m *AddAnimalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAnimalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddAnimalResponse proto.InternalMessageInfo

func (m *AddAnimalResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*AddAnimalMessage)(nil), "pupapi.AddAnimalMessage")
	proto.RegisterType((*AddAnimalResponse)(nil), "pupapi.AddAnimalResponse")
}

func init() { proto.RegisterFile("proto/addanimal.proto", fileDescriptor_f618ab4a7e8ed1ca) }

var fileDescriptor_f618ab4a7e8ed1ca = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x49, 0x49, 0xcc, 0xcb, 0xcc, 0x4d, 0xcc, 0xd1, 0x03, 0xf3, 0x85, 0xd8,
	0x0a, 0x4a, 0x0b, 0x12, 0x0b, 0x32, 0x95, 0x62, 0xb9, 0x04, 0x1c, 0x53, 0x52, 0x1c, 0xc1, 0x52,
	0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x52, 0x66, 0xbe, 0x04,
	0x13, 0x58, 0x08, 0xc4, 0x14, 0x52, 0xe6, 0xe2, 0x2d, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x8d,
	0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0xe2, 0x81, 0x0a, 0x7a,
	0x82, 0xc4, 0x94, 0x74, 0xb9, 0x04, 0xe1, 0xc6, 0x07, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0x49, 0x70, 0xb1, 0x17, 0x97, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x83, 0xad, 0xe0, 0x08, 0x82,
	0x71, 0x8d, 0xfc, 0xb9, 0x38, 0xe1, 0xca, 0x85, 0x9c, 0x90, 0x39, 0x12, 0x7a, 0x10, 0x07, 0xeb,
	0xa1, 0xbb, 0x56, 0x4a, 0x12, 0x43, 0x06, 0x66, 0x91, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xb7, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xb2, 0x38, 0x68, 0x06, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddAnimalClient is the client API for AddAnimal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddAnimalClient interface {
	AddAnimal(ctx context.Context, in *AddAnimalMessage, opts ...grpc.CallOption) (*AddAnimalResponse, error)
}

type addAnimalClient struct {
	cc *grpc.ClientConn
}

func NewAddAnimalClient(cc *grpc.ClientConn) AddAnimalClient {
	return &addAnimalClient{cc}
}

func (c *addAnimalClient) AddAnimal(ctx context.Context, in *AddAnimalMessage, opts ...grpc.CallOption) (*AddAnimalResponse, error) {
	out := new(AddAnimalResponse)
	err := c.cc.Invoke(ctx, "/pupapi.AddAnimal/AddAnimal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddAnimalServer is the server API for AddAnimal service.
type AddAnimalServer interface {
	AddAnimal(context.Context, *AddAnimalMessage) (*AddAnimalResponse, error)
}

// UnimplementedAddAnimalServer can be embedded to have forward compatible implementations.
type UnimplementedAddAnimalServer struct {
}

func (*UnimplementedAddAnimalServer) AddAnimal(ctx context.Context, req *AddAnimalMessage) (*AddAnimalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnimal not implemented")
}

func RegisterAddAnimalServer(s *grpc.Server, srv AddAnimalServer) {
	s.RegisterService(&_AddAnimal_serviceDesc, srv)
}

func _AddAnimal_AddAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnimalMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddAnimalServer).AddAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pupapi.AddAnimal/AddAnimal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddAnimalServer).AddAnimal(ctx, req.(*AddAnimalMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddAnimal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pupapi.AddAnimal",
	HandlerType: (*AddAnimalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAnimal",
			Handler:    _AddAnimal_AddAnimal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/addanimal.proto",
}
