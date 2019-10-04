// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checker.proto

package checker

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

type KV struct {
	// key is the key of this key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value of this key-value pair.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd9e8c6cb3e925d8, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CheckRequest struct {
	// data is an arbitrary list of key-value pairs that the
	// Checker may require in order to execute successfully.
	Data                 []*KV    `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd9e8c6cb3e925d8, []int{1}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetData() []*KV {
	if m != nil {
		return m.Data
	}
	return nil
}

type CheckResponse struct {
	// percent is the percentage of possible points earned,
	// represented as a floating-point value between 0.0 (0%)
	// and 1.0 (100%). Any values less than 0.0 will be
	// truncated to 0.0 and any values greater than 1.0 will
	// be truncated to 1.0.
	Percent float32 `protobuf:"fixed32,1,opt,name=percent,proto3" json:"percent,omitempty"`
	// weight is the total amount of points possible, if
	// 100% of the points are earned.
	// For example, if a Checker has a weight of 10 and a
	// percent value of 0.5, then a total of 5 points are
	// earned for that execution of the check.
	Weight int32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// data is an arbitrary list of key-value pairs that the
	// Checker may return. It is used by the server to model
	// the state of a Checker's last execution against each
	// unique set of request data.
	Data                 []*KV    `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd9e8c6cb3e925d8, []int{2}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetPercent() float32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *CheckResponse) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *CheckResponse) GetData() []*KV {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*KV)(nil), "KV")
	proto.RegisterType((*CheckRequest)(nil), "CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "CheckResponse")
}

func init() { proto.RegisterFile("checker.proto", fileDescriptor_dd9e8c6cb3e925d8) }

var fileDescriptor_dd9e8c6cb3e925d8 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0x4d,
	0xce, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe1, 0x62, 0xf2, 0x0e, 0x13,
	0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85,
	0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x92, 0x3a,
	0x17, 0x8f, 0x33, 0x48, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x38, 0x17, 0x4b,
	0x4a, 0x62, 0x49, 0xa2, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0xb3, 0x9e, 0x77, 0x58, 0x10,
	0x58, 0x40, 0x29, 0x8a, 0x8b, 0x17, 0xaa, 0xb0, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82,
	0x8b, 0xbd, 0x20, 0xb5, 0x28, 0x39, 0x35, 0xaf, 0x04, 0x6c, 0x0b, 0x53, 0x10, 0x8c, 0x2b, 0x24,
	0xc6, 0xc5, 0x56, 0x9e, 0x9a, 0x99, 0x9e, 0x51, 0x02, 0xb6, 0x8a, 0x35, 0x08, 0xca, 0x83, 0x9b,
	0xcd, 0x8c, 0x66, 0xb6, 0x91, 0x31, 0x17, 0xbb, 0x33, 0xc4, 0x0f, 0x42, 0x1a, 0x5c, 0xac, 0x60,
	0xa6, 0x10, 0xaf, 0x1e, 0xb2, 0xbb, 0xa4, 0xf8, 0xf4, 0x50, 0x6c, 0x57, 0x62, 0x48, 0x62, 0x03,
	0x7b, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x88, 0xc3, 0x85, 0x54, 0xff, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckerClient is the client API for Checker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckerClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type checkerClient struct {
	cc *grpc.ClientConn
}

func NewCheckerClient(cc *grpc.ClientConn) CheckerClient {
	return &checkerClient{cc}
}

func (c *checkerClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/Checker/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServer is the server API for Checker service.
type CheckerServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

// UnimplementedCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedCheckerServer struct {
}

func (*UnimplementedCheckerServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterCheckerServer(s *grpc.Server, srv CheckerServer) {
	s.RegisterService(&_Checker_serviceDesc, srv)
}

func _Checker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Checker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Checker",
	HandlerType: (*CheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Checker_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checker.proto",
}
