// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content_service.proto

package proto // import "../core/service/proto"

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

type PagingArticleRequest struct {
	Cat                  string   `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
	Begin                int32    `protobuf:"zigzag32,2,opt,name=begin,proto3" json:"begin,omitempty"`
	Size                 int32    `protobuf:"zigzag32,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagingArticleRequest) Reset()         { *m = PagingArticleRequest{} }
func (m *PagingArticleRequest) String() string { return proto.CompactTextString(m) }
func (*PagingArticleRequest) ProtoMessage()    {}
func (*PagingArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_0bf1469e4721937d, []int{0}
}
func (m *PagingArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingArticleRequest.Unmarshal(m, b)
}
func (m *PagingArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingArticleRequest.Marshal(b, m, deterministic)
}
func (dst *PagingArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingArticleRequest.Merge(dst, src)
}
func (m *PagingArticleRequest) XXX_Size() int {
	return xxx_messageInfo_PagingArticleRequest.Size(m)
}
func (m *PagingArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PagingArticleRequest proto.InternalMessageInfo

func (m *PagingArticleRequest) GetCat() string {
	if m != nil {
		return m.Cat
	}
	return ""
}

func (m *PagingArticleRequest) GetBegin() int32 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *PagingArticleRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ArticlesResponse struct {
	List                 []*SArticle `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArticlesResponse) Reset()         { *m = ArticlesResponse{} }
func (m *ArticlesResponse) String() string { return proto.CompactTextString(m) }
func (*ArticlesResponse) ProtoMessage()    {}
func (*ArticlesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_0bf1469e4721937d, []int{1}
}
func (m *ArticlesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlesResponse.Unmarshal(m, b)
}
func (m *ArticlesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlesResponse.Marshal(b, m, deterministic)
}
func (dst *ArticlesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlesResponse.Merge(dst, src)
}
func (m *ArticlesResponse) XXX_Size() int {
	return xxx_messageInfo_ArticlesResponse.Size(m)
}
func (m *ArticlesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlesResponse proto.InternalMessageInfo

func (m *ArticlesResponse) GetList() []*SArticle {
	if m != nil {
		return m.List
	}
	return nil
}

// * 文章
type SArticle struct {
	// * 编号
	Id int32 `protobuf:"zigzag32,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// * 栏目编号
	CatId int32 `protobuf:"zigzag32,2,opt,name=CatId,proto3" json:"CatId,omitempty"`
	// * 标题
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	// * 小标题
	SmallTitle string `protobuf:"bytes,4,opt,name=SmallTitle,proto3" json:"SmallTitle,omitempty"`
	// * 文章附图
	Thumbnail string `protobuf:"bytes,5,opt,name=Thumbnail,proto3" json:"Thumbnail,omitempty"`
	// * 重定向URL
	PublisherId int32 `protobuf:"zigzag32,6,opt,name=PublisherId,proto3" json:"PublisherId,omitempty"`
	// * 重定向URL
	Location string `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	// * 优先级,优先级越高，则置顶
	Priority int32 `protobuf:"zigzag32,8,opt,name=Priority,proto3" json:"Priority,omitempty"`
	// * 浏览钥匙
	AccessKey string `protobuf:"bytes,9,opt,name=AccessKey,proto3" json:"AccessKey,omitempty"`
	// * 文档内容
	Content string `protobuf:"bytes,10,opt,name=Content,proto3" json:"Content,omitempty"`
	// * 标签（关键词）
	Tags string `protobuf:"bytes,11,opt,name=Tags,proto3" json:"Tags,omitempty"`
	// * 显示次数
	ViewCount int32 `protobuf:"zigzag32,12,opt,name=ViewCount,proto3" json:"ViewCount,omitempty"`
	// * 排序序号
	SortNum int32 `protobuf:"zigzag32,13,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	// * 创建时间
	CreateTime int32 `protobuf:"zigzag32,14,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// * 最后修改时间
	UpdateTime           int32    `protobuf:"zigzag32,15,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SArticle) Reset()         { *m = SArticle{} }
func (m *SArticle) String() string { return proto.CompactTextString(m) }
func (*SArticle) ProtoMessage()    {}
func (*SArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_0bf1469e4721937d, []int{2}
}
func (m *SArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SArticle.Unmarshal(m, b)
}
func (m *SArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SArticle.Marshal(b, m, deterministic)
}
func (dst *SArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SArticle.Merge(dst, src)
}
func (m *SArticle) XXX_Size() int {
	return xxx_messageInfo_SArticle.Size(m)
}
func (m *SArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_SArticle.DiscardUnknown(m)
}

var xxx_messageInfo_SArticle proto.InternalMessageInfo

func (m *SArticle) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SArticle) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *SArticle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SArticle) GetSmallTitle() string {
	if m != nil {
		return m.SmallTitle
	}
	return ""
}

func (m *SArticle) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *SArticle) GetPublisherId() int32 {
	if m != nil {
		return m.PublisherId
	}
	return 0
}

func (m *SArticle) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SArticle) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SArticle) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *SArticle) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SArticle) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *SArticle) GetViewCount() int32 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *SArticle) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SArticle) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SArticle) GetUpdateTime() int32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*PagingArticleRequest)(nil), "PagingArticleRequest")
	proto.RegisterType((*ArticlesResponse)(nil), "ArticlesResponse")
	proto.RegisterType((*SArticle)(nil), "SArticle")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentServiceClient interface {
	// * 获取置顶的文章,cat
	QueryTopArticles(ctx context.Context, in *String, opts ...grpc.CallOption) (*ArticlesResponse, error)
	// * 获取分页文章
	QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*SPagingResult, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) QueryTopArticles(ctx context.Context, in *String, opts ...grpc.CallOption) (*ArticlesResponse, error) {
	out := new(ArticlesResponse)
	err := c.cc.Invoke(ctx, "/ContentService/QueryTopArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*SPagingResult, error) {
	out := new(SPagingResult)
	err := c.cc.Invoke(ctx, "/ContentService/QueryPagingArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
type ContentServiceServer interface {
	// * 获取置顶的文章,cat
	QueryTopArticles(context.Context, *String) (*ArticlesResponse, error)
	// * 获取分页文章
	QueryPagingArticles(context.Context, *PagingArticleRequest) (*SPagingResult, error)
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_QueryTopArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryTopArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_QueryPagingArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryPagingArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, req.(*PagingArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryTopArticles",
			Handler:    _ContentService_QueryTopArticles_Handler,
		},
		{
			MethodName: "QueryPagingArticles",
			Handler:    _ContentService_QueryPagingArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content_service.proto",
}

func init() {
	proto.RegisterFile("content_service.proto", fileDescriptor_content_service_0bf1469e4721937d)
}

var fileDescriptor_content_service_0bf1469e4721937d = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8f, 0xd2, 0x40,
	0x14, 0xc6, 0x97, 0x85, 0x5d, 0xe8, 0x43, 0x11, 0xc6, 0xdd, 0x64, 0x82, 0xae, 0x21, 0x3d, 0x71,
	0x2a, 0xba, 0x1e, 0x4d, 0x4c, 0x56, 0x4e, 0xc4, 0x8d, 0xc1, 0xb6, 0x7a, 0xf0, 0x62, 0x4a, 0xfb,
	0xd2, 0x9d, 0xa4, 0xcc, 0xd4, 0x99, 0xa9, 0x06, 0x8f, 0xfe, 0x29, 0xfe, 0xa5, 0x66, 0xde, 0x50,
	0x40, 0xe3, 0x05, 0xe6, 0xfb, 0x7d, 0x79, 0xef, 0x6b, 0x66, 0x3e, 0xb8, 0xce, 0x95, 0xb4, 0x28,
	0xed, 0x57, 0x83, 0xfa, 0xbb, 0xc8, 0x31, 0xaa, 0xb5, 0xb2, 0x6a, 0x3a, 0xb4, 0x76, 0x57, 0xef,
	0x45, 0x18, 0xc3, 0xd5, 0x3a, 0x2b, 0x85, 0x2c, 0xef, 0xb4, 0x15, 0x79, 0x85, 0x31, 0x7e, 0x6b,
	0xd0, 0x58, 0x36, 0x86, 0x6e, 0x9e, 0x59, 0xde, 0x99, 0x75, 0xe6, 0x41, 0xec, 0x8e, 0xec, 0x0a,
	0x2e, 0x36, 0x58, 0x0a, 0xc9, 0xcf, 0x67, 0x9d, 0xf9, 0x24, 0xf6, 0x82, 0x31, 0xe8, 0x19, 0xf1,
	0x13, 0x79, 0x97, 0x20, 0x9d, 0xc3, 0x57, 0x30, 0xde, 0x6f, 0x33, 0x31, 0x9a, 0x5a, 0x49, 0x83,
	0xec, 0x06, 0x7a, 0xf7, 0xc2, 0xb8, 0x85, 0xdd, 0xf9, 0xf0, 0x36, 0x88, 0x92, 0x36, 0x8f, 0x70,
	0xf8, 0xbb, 0x0b, 0x83, 0x16, 0xb1, 0x11, 0x9c, 0xaf, 0x0a, 0x8a, 0x9e, 0xc4, 0xe7, 0xab, 0xc2,
	0x25, 0x2f, 0x33, 0xbb, 0x2a, 0xda, 0x64, 0x12, 0x8e, 0xa6, 0xc2, 0x56, 0x3e, 0x3a, 0x88, 0xbd,
	0x60, 0x2f, 0x00, 0x92, 0x6d, 0x56, 0x55, 0xde, 0xea, 0x91, 0x75, 0x42, 0xd8, 0x73, 0x08, 0xd2,
	0x87, 0x66, 0xbb, 0x91, 0x99, 0xa8, 0xf8, 0x05, 0xd9, 0x47, 0xc0, 0x66, 0x30, 0x5c, 0x37, 0x9b,
	0x4a, 0x98, 0x07, 0xd4, 0xab, 0x82, 0x5f, 0x52, 0xde, 0x29, 0x62, 0x53, 0x18, 0xdc, 0xab, 0x3c,
	0xb3, 0x42, 0x49, 0xde, 0xa7, 0xf1, 0x83, 0x76, 0xde, 0x5a, 0x0b, 0xa5, 0x85, 0xdd, 0xf1, 0x01,
	0x8d, 0x1e, 0xb4, 0xcb, 0xbd, 0xcb, 0x73, 0x34, 0xe6, 0x3d, 0xee, 0x78, 0xe0, 0x73, 0x0f, 0x80,
	0x71, 0xe8, 0x2f, 0xfd, 0x5b, 0x71, 0x20, 0xaf, 0x95, 0xee, 0x7e, 0xd3, 0xac, 0x34, 0x7c, 0x48,
	0x98, 0xce, 0x6e, 0xd7, 0x67, 0x81, 0x3f, 0x96, 0xaa, 0x91, 0x96, 0x3f, 0xa2, 0xa0, 0x23, 0x70,
	0xbb, 0x12, 0xa5, 0xed, 0x87, 0x66, 0xcb, 0x1f, 0x93, 0xd7, 0x4a, 0x77, 0x37, 0x4b, 0x8d, 0x99,
	0xc5, 0x54, 0x6c, 0x91, 0x8f, 0xc8, 0x3c, 0x21, 0xce, 0xff, 0x54, 0x17, 0xad, 0xff, 0xc4, 0xfb,
	0x47, 0x72, 0xfb, 0xab, 0x03, 0xa3, 0xfd, 0x77, 0x25, 0xbe, 0x51, 0xec, 0x25, 0x8c, 0x3f, 0x36,
	0xa8, 0x77, 0xa9, 0xaa, 0xdb, 0x27, 0x67, 0xfd, 0x28, 0xb1, 0x5a, 0xc8, 0x72, 0x3a, 0x89, 0xfe,
	0xad, 0x41, 0x78, 0xc6, 0xde, 0xc2, 0x53, 0x9a, 0xf8, 0xab, 0x75, 0x86, 0x5d, 0x47, 0xff, 0xab,
	0xe1, 0x74, 0x14, 0x25, 0x9e, 0xc7, 0x68, 0x9a, 0xca, 0x86, 0x67, 0xef, 0x6e, 0xbe, 0x3c, 0x8b,
	0xa2, 0x45, 0xae, 0x34, 0x2e, 0xf6, 0xb5, 0x5e, 0x50, 0x93, 0xdf, 0xd0, 0xef, 0xe6, 0x92, 0xfe,
	0x5e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x83, 0x0b, 0x25, 0x42, 0xfc, 0x02, 0x00, 0x00,
}