// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_service.proto

package proto // import "."

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

// * 消息方式
type EMessageChannel int32

const (
	EMessageChannel__2 EMessageChannel = 0
	// * 站内信
	EMessageChannel_SiteMessage  EMessageChannel = 1
	EMessageChannel_EmailMessage EMessageChannel = 2
	EMessageChannel_SmsMessage   EMessageChannel = 3
)

var EMessageChannel_name = map[int32]string{
	0: "_2",
	1: "SiteMessage",
	2: "EmailMessage",
	3: "SmsMessage",
}
var EMessageChannel_value = map[string]int32{
	"_2":           0,
	"SiteMessage":  1,
	"EmailMessage": 2,
	"SmsMessage":   3,
}

func (x EMessageChannel) String() string {
	return proto.EnumName(EMessageChannel_name, int32(x))
}
func (EMessageChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_service_398b69cdc84f9ae3, []int{0}
}

type SendMessageRequest struct {
	Account              string            `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_service_398b69cdc84f9ae3, []int{0}
}
func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (dst *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(dst, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SendMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SendMessageRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

// * 通知项
type SNotifyItem struct {
	// * 键
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// * 发送方式
	NotifyBy int32 `protobuf:"zigzag32,2,opt,name=NotifyBy,proto3" json:"NotifyBy,omitempty"`
	// * 不允许修改发送方式
	ReadonlyBy bool `protobuf:"varint,3,opt,name=ReadonlyBy,proto3" json:"ReadonlyBy,omitempty"`
	// * 模板编号
	TplId int32 `protobuf:"zigzag32,4,opt,name=TplId,proto3" json:"TplId,omitempty"`
	// * 内容
	Content string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	// * 模板包含的标签
	Tags                 map[string]string `protobuf:"bytes,6,rep,name=Tags,proto3" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SNotifyItem) Reset()         { *m = SNotifyItem{} }
func (m *SNotifyItem) String() string { return proto.CompactTextString(m) }
func (*SNotifyItem) ProtoMessage()    {}
func (*SNotifyItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_service_398b69cdc84f9ae3, []int{1}
}
func (m *SNotifyItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SNotifyItem.Unmarshal(m, b)
}
func (m *SNotifyItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SNotifyItem.Marshal(b, m, deterministic)
}
func (dst *SNotifyItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SNotifyItem.Merge(dst, src)
}
func (m *SNotifyItem) XXX_Size() int {
	return xxx_messageInfo_SNotifyItem.Size(m)
}
func (m *SNotifyItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SNotifyItem.DiscardUnknown(m)
}

var xxx_messageInfo_SNotifyItem proto.InternalMessageInfo

func (m *SNotifyItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SNotifyItem) GetNotifyBy() int32 {
	if m != nil {
		return m.NotifyBy
	}
	return 0
}

func (m *SNotifyItem) GetReadonlyBy() bool {
	if m != nil {
		return m.ReadonlyBy
	}
	return false
}

func (m *SNotifyItem) GetTplId() int32 {
	if m != nil {
		return m.TplId
	}
	return 0
}

func (m *SNotifyItem) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SNotifyItem) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "SendMessageRequest")
	proto.RegisterMapType((map[string]string)(nil), "SendMessageRequest.DataEntry")
	proto.RegisterType((*SNotifyItem)(nil), "SNotifyItem")
	proto.RegisterMapType((map[string]string)(nil), "SNotifyItem.TagsEntry")
	proto.RegisterEnum("EMessageChannel", EMessageChannel_name, EMessageChannel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	// * 获取通知项,key
	GetNotifyItem(ctx context.Context, in *String, opts ...grpc.CallOption) (*SNotifyItem, error)
	// * 发送短信
	SendPhoneMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Result, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) GetNotifyItem(ctx context.Context, in *String, opts ...grpc.CallOption) (*SNotifyItem, error) {
	out := new(SNotifyItem)
	err := c.cc.Invoke(ctx, "/MessageService/GetNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendPhoneMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SendPhoneMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	// * 获取通知项,key
	GetNotifyItem(context.Context, *String) (*SNotifyItem, error)
	// * 发送短信
	SendPhoneMessage(context.Context, *SendMessageRequest) (*Result, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_GetNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetNotifyItem(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendPhoneMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendPhoneMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SendPhoneMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendPhoneMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifyItem",
			Handler:    _MessageService_GetNotifyItem_Handler,
		},
		{
			MethodName: "SendPhoneMessage",
			Handler:    _MessageService_SendPhoneMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_service.proto",
}

func init() {
	proto.RegisterFile("message_service.proto", fileDescriptor_message_service_398b69cdc84f9ae3)
}

var fileDescriptor_message_service_398b69cdc84f9ae3 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x37, 0xc9, 0x76, 0xd3, 0x9d, 0x5d, 0xda, 0x60, 0xfe, 0x28, 0x8a, 0x04, 0x5a, 0xe5,
	0x14, 0xf5, 0x60, 0x89, 0x70, 0x00, 0xc1, 0xad, 0x65, 0x85, 0xaa, 0x02, 0x42, 0x49, 0x4f, 0x5c,
	0x2a, 0x77, 0x33, 0xa4, 0x11, 0x8e, 0x5d, 0x12, 0x67, 0xa5, 0xbc, 0x17, 0x0f, 0xc6, 0x23, 0x20,
	0xc7, 0xce, 0x12, 0x09, 0x2e, 0x3d, 0xc5, 0xbf, 0xf9, 0x3c, 0xfe, 0xf2, 0xd9, 0x03, 0xcf, 0x6a,
	0x6c, 0x5b, 0x56, 0xe2, 0x4d, 0x8b, 0xcd, 0xbe, 0xda, 0x21, 0xbd, 0x6f, 0xa4, 0x92, 0xd1, 0xba,
	0xe4, 0xf2, 0x96, 0x71, 0x43, 0xf1, 0x2f, 0x07, 0x48, 0x8e, 0xa2, 0xf8, 0x6c, 0xf6, 0x66, 0xf8,
	0xb3, 0xc3, 0x56, 0x91, 0x10, 0x7c, 0xb6, 0xdb, 0xc9, 0x4e, 0xa8, 0xd0, 0xd9, 0x38, 0xc9, 0x32,
	0x1b, 0x51, 0x2b, 0xf6, 0xdc, 0xd0, 0x35, 0x8a, 0x45, 0xf2, 0x0a, 0xe6, 0x05, 0x53, 0x2c, 0xf4,
	0x36, 0x5e, 0xb2, 0x4a, 0x5f, 0xd0, 0x7f, 0x8f, 0xa5, 0x1f, 0x98, 0x62, 0x5b, 0xa1, 0x9a, 0x3e,
	0x1b, 0xb6, 0x46, 0x6f, 0x60, 0x79, 0x28, 0x91, 0x00, 0xbc, 0x1f, 0xd8, 0x5b, 0x3f, 0xbd, 0x24,
	0x4f, 0xe1, 0x68, 0xcf, 0x78, 0x37, 0x3a, 0x19, 0x78, 0xe7, 0xbe, 0x75, 0xe2, 0xdf, 0x0e, 0xac,
	0xf2, 0x2f, 0x52, 0x55, 0xdf, 0xfb, 0x4b, 0x85, 0xb5, 0xee, 0xbd, 0xfa, 0xdb, 0x7b, 0x85, 0x3d,
	0x89, 0xe0, 0xd8, 0xe8, 0xe7, 0xfd, 0xd0, 0xfe, 0x38, 0x3b, 0x30, 0x79, 0x09, 0x90, 0x21, 0x2b,
	0xa4, 0xe0, 0x5a, 0xf5, 0x36, 0x4e, 0x72, 0x9c, 0x4d, 0x2a, 0xda, 0xf7, 0xfa, 0x9e, 0x5f, 0x16,
	0xe1, 0x7c, 0x68, 0x34, 0xa0, 0x93, 0x5f, 0x48, 0xa1, 0x50, 0xa8, 0xf0, 0xc8, 0x24, 0xb7, 0x48,
	0xce, 0x60, 0x7e, 0xcd, 0xca, 0x36, 0x5c, 0x0c, 0xc9, 0x9f, 0xd3, 0xc9, 0x9f, 0x51, 0x2d, 0xd8,
	0xc8, 0x7a, 0xa9, 0x23, 0x1f, 0x4a, 0x0f, 0x89, 0x7c, 0xf6, 0x09, 0x4e, 0xb7, 0xf6, 0x3a, 0x2f,
	0xee, 0x98, 0x10, 0xc8, 0xc9, 0x02, 0xdc, 0x9b, 0x34, 0x98, 0x91, 0x53, 0x58, 0xe5, 0x95, 0x42,
	0xab, 0x06, 0x0e, 0x09, 0x60, 0xbd, 0xad, 0x59, 0xc5, 0xc7, 0x8a, 0x4b, 0x4e, 0x00, 0xf2, 0xba,
	0x1d, 0xd9, 0x4b, 0x05, 0x9c, 0x58, 0xc8, 0xcd, 0x74, 0x90, 0x04, 0x1e, 0x7d, 0x44, 0x35, 0xb9,
	0x53, 0x9f, 0xe6, 0xaa, 0xa9, 0x44, 0x19, 0xad, 0xa7, 0x81, 0xe2, 0x19, 0x49, 0x21, 0xd0, 0x6f,
	0xfb, 0xf5, 0x4e, 0x8a, 0xd1, 0x93, 0x3c, 0xf9, 0xcf, 0x73, 0x47, 0x3e, 0xcd, 0xb0, 0xed, 0xb8,
	0x8a, 0x67, 0xe7, 0xcb, 0x6f, 0x3e, 0x7d, 0x3f, 0x8c, 0xdc, 0xed, 0x62, 0xf8, 0xbc, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xa4, 0x58, 0x66, 0x00, 0xa0, 0x02, 0x00, 0x00,
}
