// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foundation_service.proto

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

type SmsApiSaveRequest struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Api                  *SSmsApi `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmsApiSaveRequest) Reset()         { *m = SmsApiSaveRequest{} }
func (m *SmsApiSaveRequest) String() string { return proto.CompactTextString(m) }
func (*SmsApiSaveRequest) ProtoMessage()    {}
func (*SmsApiSaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{0}
}
func (m *SmsApiSaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmsApiSaveRequest.Unmarshal(m, b)
}
func (m *SmsApiSaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmsApiSaveRequest.Marshal(b, m, deterministic)
}
func (dst *SmsApiSaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmsApiSaveRequest.Merge(dst, src)
}
func (m *SmsApiSaveRequest) XXX_Size() int {
	return xxx_messageInfo_SmsApiSaveRequest.Size(m)
}
func (m *SmsApiSaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SmsApiSaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SmsApiSaveRequest proto.InternalMessageInfo

func (m *SmsApiSaveRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *SmsApiSaveRequest) GetApi() *SSmsApi {
	if m != nil {
		return m.Api
	}
	return nil
}

type BoardHookSaveRequest struct {
	HookURL              string   `protobuf:"bytes,1,opt,name=hookURL,proto3" json:"hookURL,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoardHookSaveRequest) Reset()         { *m = BoardHookSaveRequest{} }
func (m *BoardHookSaveRequest) String() string { return proto.CompactTextString(m) }
func (*BoardHookSaveRequest) ProtoMessage()    {}
func (*BoardHookSaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{1}
}
func (m *BoardHookSaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardHookSaveRequest.Unmarshal(m, b)
}
func (m *BoardHookSaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardHookSaveRequest.Marshal(b, m, deterministic)
}
func (dst *BoardHookSaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardHookSaveRequest.Merge(dst, src)
}
func (m *BoardHookSaveRequest) XXX_Size() int {
	return xxx_messageInfo_BoardHookSaveRequest.Size(m)
}
func (m *BoardHookSaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardHookSaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BoardHookSaveRequest proto.InternalMessageInfo

func (m *BoardHookSaveRequest) GetHookURL() string {
	if m != nil {
		return m.HookURL
	}
	return ""
}

func (m *BoardHookSaveRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AreaListResponse struct {
	Value                []*SArea `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaListResponse) Reset()         { *m = AreaListResponse{} }
func (m *AreaListResponse) String() string { return proto.CompactTextString(m) }
func (*AreaListResponse) ProtoMessage()    {}
func (*AreaListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{2}
}
func (m *AreaListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaListResponse.Unmarshal(m, b)
}
func (m *AreaListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaListResponse.Marshal(b, m, deterministic)
}
func (dst *AreaListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaListResponse.Merge(dst, src)
}
func (m *AreaListResponse) XXX_Size() int {
	return xxx_messageInfo_AreaListResponse.Size(m)
}
func (m *AreaListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AreaListResponse proto.InternalMessageInfo

func (m *AreaListResponse) GetValue() []*SArea {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringListResponse struct {
	Value                []string `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringListResponse) Reset()         { *m = StringListResponse{} }
func (m *StringListResponse) String() string { return proto.CompactTextString(m) }
func (*StringListResponse) ProtoMessage()    {}
func (*StringListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{3}
}
func (m *StringListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringListResponse.Unmarshal(m, b)
}
func (m *StringListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringListResponse.Marshal(b, m, deterministic)
}
func (dst *StringListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringListResponse.Merge(dst, src)
}
func (m *StringListResponse) XXX_Size() int {
	return xxx_messageInfo_StringListResponse.Size(m)
}
func (m *StringListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringListResponse proto.InternalMessageInfo

func (m *StringListResponse) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetAreaNamesRequest struct {
	Value                []int32  `protobuf:"zigzag32,1,rep,packed,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAreaNamesRequest) Reset()         { *m = GetAreaNamesRequest{} }
func (m *GetAreaNamesRequest) String() string { return proto.CompactTextString(m) }
func (*GetAreaNamesRequest) ProtoMessage()    {}
func (*GetAreaNamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{4}
}
func (m *GetAreaNamesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAreaNamesRequest.Unmarshal(m, b)
}
func (m *GetAreaNamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAreaNamesRequest.Marshal(b, m, deterministic)
}
func (dst *GetAreaNamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAreaNamesRequest.Merge(dst, src)
}
func (m *GetAreaNamesRequest) XXX_Size() int {
	return xxx_messageInfo_GetAreaNamesRequest.Size(m)
}
func (m *GetAreaNamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAreaNamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAreaNamesRequest proto.InternalMessageInfo

func (m *GetAreaNamesRequest) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type UserPwd struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPwd) Reset()         { *m = UserPwd{} }
func (m *UserPwd) String() string { return proto.CompactTextString(m) }
func (*UserPwd) ProtoMessage()    {}
func (*UserPwd) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{5}
}
func (m *UserPwd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPwd.Unmarshal(m, b)
}
func (m *UserPwd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPwd.Marshal(b, m, deterministic)
}
func (dst *UserPwd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPwd.Merge(dst, src)
}
func (m *UserPwd) XXX_Size() int {
	return xxx_messageInfo_UserPwd.Size(m)
}
func (m *UserPwd) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPwd.DiscardUnknown(m)
}

var xxx_messageInfo_UserPwd proto.InternalMessageInfo

func (m *UserPwd) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserPwd) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

// 单点登录应用
type SSsoApp struct {
	// 编号
	ID int32 `protobuf:"zigzag32,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 应用名称
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// API地址
	ApiUrl string `protobuf:"bytes,3,opt,name=ApiUrl,proto3" json:"ApiUrl,omitempty"`
	// 密钥
	Token                string   `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSsoApp) Reset()         { *m = SSsoApp{} }
func (m *SSsoApp) String() string { return proto.CompactTextString(m) }
func (*SSsoApp) ProtoMessage()    {}
func (*SSsoApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{6}
}
func (m *SSsoApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSsoApp.Unmarshal(m, b)
}
func (m *SSsoApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSsoApp.Marshal(b, m, deterministic)
}
func (dst *SSsoApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSsoApp.Merge(dst, src)
}
func (m *SSsoApp) XXX_Size() int {
	return xxx_messageInfo_SSsoApp.Size(m)
}
func (m *SSsoApp) XXX_DiscardUnknown() {
	xxx_messageInfo_SSsoApp.DiscardUnknown(m)
}

var xxx_messageInfo_SSsoApp proto.InternalMessageInfo

func (m *SSsoApp) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SSsoApp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SSsoApp) GetApiUrl() string {
	if m != nil {
		return m.ApiUrl
	}
	return ""
}

func (m *SSsoApp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// * 行政区域
type SArea struct {
	Code                 int32    `protobuf:"zigzag32,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Parent               int32    `protobuf:"zigzag32,2,opt,name=Parent,proto3" json:"Parent,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SArea) Reset()         { *m = SArea{} }
func (m *SArea) String() string { return proto.CompactTextString(m) }
func (*SArea) ProtoMessage()    {}
func (*SArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{7}
}
func (m *SArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SArea.Unmarshal(m, b)
}
func (m *SArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SArea.Marshal(b, m, deterministic)
}
func (dst *SArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SArea.Merge(dst, src)
}
func (m *SArea) XXX_Size() int {
	return xxx_messageInfo_SArea.Size(m)
}
func (m *SArea) XXX_DiscardUnknown() {
	xxx_messageInfo_SArea.DiscardUnknown(m)
}

var xxx_messageInfo_SArea proto.InternalMessageInfo

func (m *SArea) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SArea) GetParent() int32 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *SArea) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// * 短信接口信息
type SSmsApi struct {
	// * 接口地址
	ApiUrl string `protobuf:"bytes,1,opt,name=ApiUrl,proto3" json:"ApiUrl,omitempty"`
	// * 接口KEY
	Key string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	// * 接口密钥
	Secret string `protobuf:"bytes,3,opt,name=Secret,proto3" json:"Secret,omitempty"`
	// * 请求数据,如: phone={phone}&content={content}
	Params string `protobuf:"bytes,4,opt,name=Params,proto3" json:"Params,omitempty"`
	// * 请求方式, GET或POST
	Method string `protobuf:"bytes,5,opt,name=Method,proto3" json:"Method,omitempty"`
	// * 编码
	Charset string `protobuf:"bytes,6,opt,name=Charset,proto3" json:"Charset,omitempty"`
	// * 签名
	Signature string `protobuf:"bytes,7,opt,name=Signature,proto3" json:"Signature,omitempty"`
	// * 发送成功，包含的字符，用于检测是否发送成功
	SuccessChar          string   `protobuf:"bytes,8,opt,name=SuccessChar,proto3" json:"SuccessChar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSmsApi) Reset()         { *m = SSmsApi{} }
func (m *SSmsApi) String() string { return proto.CompactTextString(m) }
func (*SSmsApi) ProtoMessage()    {}
func (*SSmsApi) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_f8c40d04930833fb, []int{8}
}
func (m *SSmsApi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSmsApi.Unmarshal(m, b)
}
func (m *SSmsApi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSmsApi.Marshal(b, m, deterministic)
}
func (dst *SSmsApi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSmsApi.Merge(dst, src)
}
func (m *SSmsApi) XXX_Size() int {
	return xxx_messageInfo_SSmsApi.Size(m)
}
func (m *SSmsApi) XXX_DiscardUnknown() {
	xxx_messageInfo_SSmsApi.DiscardUnknown(m)
}

var xxx_messageInfo_SSmsApi proto.InternalMessageInfo

func (m *SSmsApi) GetApiUrl() string {
	if m != nil {
		return m.ApiUrl
	}
	return ""
}

func (m *SSmsApi) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SSmsApi) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *SSmsApi) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *SSmsApi) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SSmsApi) GetCharset() string {
	if m != nil {
		return m.Charset
	}
	return ""
}

func (m *SSmsApi) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *SSmsApi) GetSuccessChar() string {
	if m != nil {
		return m.SuccessChar
	}
	return ""
}

func init() {
	proto.RegisterType((*SmsApiSaveRequest)(nil), "SmsApiSaveRequest")
	proto.RegisterType((*BoardHookSaveRequest)(nil), "BoardHookSaveRequest")
	proto.RegisterType((*AreaListResponse)(nil), "AreaListResponse")
	proto.RegisterType((*StringListResponse)(nil), "StringListResponse")
	proto.RegisterType((*GetAreaNamesRequest)(nil), "GetAreaNamesRequest")
	proto.RegisterType((*UserPwd)(nil), "UserPwd")
	proto.RegisterType((*SSsoApp)(nil), "SSsoApp")
	proto.RegisterType((*SArea)(nil), "SArea")
	proto.RegisterType((*SSmsApi)(nil), "SSmsApi")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FoundationServiceClient is the client API for FoundationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoundationServiceClient interface {
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   -  1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error)
	// 获取应用信息,name
	GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Bool, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Empty, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 获取地区名称
	GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error)
	// 获取下级区域,code
	GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error)
}

type foundationServiceClient struct {
	cc *grpc.ClientConn
}

func NewFoundationServiceClient(cc *grpc.ClientConn) FoundationServiceClient {
	return &foundationServiceClient{cc}
}

func (c *foundationServiceClient) GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error) {
	out := new(SSmsApi)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveBoardHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ResourceUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/RegisterApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error) {
	out := new(SSsoApp)
	err := c.cc.Invoke(ctx, "/FoundationService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAllSsoApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/FoundationService/SuperValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/FoundationService/FlushSuperPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSyncLoginUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error) {
	out := new(AreaListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetChildAreas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoundationServiceServer is the server API for FoundationService service.
type FoundationServiceServer interface {
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(context.Context, *String) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(context.Context, *SmsApiSaveRequest) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(context.Context, *BoardHookSaveRequest) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(context.Context, *String) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   -  1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(context.Context, *SSsoApp) (*String, error)
	// 获取应用信息,name
	GetApp(context.Context, *String) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(context.Context, *Empty) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(context.Context, *UserPwd) (*Bool, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(context.Context, *UserPwd) (*Empty, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(context.Context, *String) (*String, error)
	// 获取地区名称
	GetAreaNames(context.Context, *GetAreaNamesRequest) (*StringListResponse, error)
	// 获取下级区域,code
	GetChildAreas(context.Context, *Int32) (*AreaListResponse, error)
}

func RegisterFoundationServiceServer(s *grpc.Server, srv FoundationServiceServer) {
	s.RegisterService(&_FoundationService_serviceDesc, srv)
}

func _FoundationService_GetSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsApiSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, req.(*SmsApiSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveBoardHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardHookSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveBoardHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, req.(*BoardHookSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ResourceUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ResourceUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_RegisterApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSsoApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).RegisterApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/RegisterApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).RegisterApp(ctx, req.(*SSsoApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetApp(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAllSsoApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAllSsoApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SuperValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SuperValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SuperValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SuperValidate(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_FlushSuperPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/FlushSuperPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSyncLoginUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSyncLoginUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, req.(*GetAreaNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetChildAreas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetChildAreas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

var _FoundationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FoundationService",
	HandlerType: (*FoundationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSmsApi",
			Handler:    _FoundationService_GetSmsApi_Handler,
		},
		{
			MethodName: "SaveSmsApi",
			Handler:    _FoundationService_SaveSmsApi_Handler,
		},
		{
			MethodName: "SaveBoardHook",
			Handler:    _FoundationService_SaveBoardHook_Handler,
		},
		{
			MethodName: "ResourceUrl",
			Handler:    _FoundationService_ResourceUrl_Handler,
		},
		{
			MethodName: "RegisterApp",
			Handler:    _FoundationService_RegisterApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _FoundationService_GetApp_Handler,
		},
		{
			MethodName: "GetAllSsoApp",
			Handler:    _FoundationService_GetAllSsoApp_Handler,
		},
		{
			MethodName: "SuperValidate",
			Handler:    _FoundationService_SuperValidate_Handler,
		},
		{
			MethodName: "FlushSuperPwd",
			Handler:    _FoundationService_FlushSuperPwd_Handler,
		},
		{
			MethodName: "GetSyncLoginUrl",
			Handler:    _FoundationService_GetSyncLoginUrl_Handler,
		},
		{
			MethodName: "GetAreaNames",
			Handler:    _FoundationService_GetAreaNames_Handler,
		},
		{
			MethodName: "GetChildAreas",
			Handler:    _FoundationService_GetChildAreas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foundation_service.proto",
}

func init() {
	proto.RegisterFile("foundation_service.proto", fileDescriptor_foundation_service_f8c40d04930833fb)
}

var fileDescriptor_foundation_service_f8c40d04930833fb = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdf, 0x4f, 0x13, 0x41,
	0x10, 0x6e, 0x29, 0xfd, 0x35, 0xa5, 0x4a, 0x97, 0x6a, 0x2e, 0x17, 0x4c, 0xea, 0x1a, 0x13, 0x22,
	0xba, 0x6a, 0x79, 0xe4, 0xa9, 0x80, 0x54, 0x02, 0x1a, 0x72, 0x27, 0x3c, 0xe8, 0x83, 0x39, 0x7a,
	0x63, 0x7b, 0xe1, 0x7a, 0x7b, 0xee, 0xee, 0x41, 0xfa, 0x07, 0xf8, 0x27, 0xfa, 0xff, 0x98, 0xdd,
	0xdb, 0x3b, 0x5a, 0x81, 0xa7, 0xee, 0x37, 0xfb, 0xcd, 0xcc, 0xb7, 0xd3, 0x6f, 0x0e, 0x9c, 0x5f,
	0x3c, 0x4b, 0xc2, 0x40, 0x45, 0x3c, 0xf9, 0x29, 0x51, 0xdc, 0x44, 0x13, 0x64, 0xa9, 0xe0, 0x8a,
	0xbb, 0x1b, 0xd3, 0x98, 0x5f, 0x05, 0x71, 0x8e, 0xe8, 0x29, 0xf4, 0xfc, 0xb9, 0x1c, 0xa5, 0x91,
	0x1f, 0xdc, 0xa0, 0x87, 0xbf, 0x33, 0x94, 0x8a, 0xb8, 0xd0, 0x4a, 0x05, 0xbf, 0x89, 0x42, 0x14,
	0x4e, 0x75, 0x50, 0xdd, 0x69, 0x7b, 0x25, 0x26, 0x2e, 0xd4, 0x82, 0x34, 0x72, 0xd6, 0x06, 0xd5,
	0x9d, 0xce, 0xb0, 0xc5, 0xfc, 0x3c, 0xdb, 0xd3, 0x41, 0x7a, 0x0c, 0xfd, 0x03, 0x1e, 0x88, 0xf0,
	0x33, 0xe7, 0xd7, 0xcb, 0xf5, 0x1c, 0x68, 0xce, 0x38, 0xbf, 0xbe, 0xf0, 0xce, 0x6c, 0xb9, 0x02,
	0x92, 0x3e, 0xd4, 0x15, 0xbf, 0xc6, 0xc4, 0xd4, 0x6b, 0x7b, 0x39, 0xa0, 0x1f, 0x60, 0x73, 0x24,
	0x30, 0x38, 0x8b, 0xa4, 0xf2, 0x50, 0xa6, 0x3c, 0x91, 0x48, 0xb6, 0xa1, 0x7e, 0x19, 0xc4, 0x19,
	0x3a, 0xd5, 0x41, 0x6d, 0xa7, 0x33, 0x6c, 0x30, 0x5f, 0x53, 0xbc, 0x3c, 0x48, 0xdf, 0x00, 0xf1,
	0x95, 0x88, 0x92, 0xe9, 0x4a, 0x4e, 0x7f, 0x39, 0xa7, 0x5d, 0x70, 0x77, 0x61, 0x6b, 0x8c, 0x4a,
	0x67, 0x7f, 0x0d, 0xe6, 0x28, 0x0b, 0x91, 0x2b, 0xe4, 0x5e, 0x41, 0x7e, 0x0f, 0xcd, 0x0b, 0x89,
	0xe2, 0xfc, 0x36, 0x24, 0x04, 0xd6, 0x33, 0x59, 0x4e, 0xc4, 0x9c, 0xc9, 0x26, 0xd4, 0xd2, 0xdb,
	0xd0, 0xaa, 0xd7, 0x47, 0xfa, 0x03, 0x9a, 0xbe, 0x2f, 0xf9, 0x28, 0x4d, 0xc9, 0x13, 0x58, 0x3b,
	0x39, 0x32, 0xf4, 0x9e, 0xb7, 0x76, 0x72, 0xa4, 0x0b, 0xe8, 0x8e, 0x96, 0x6d, 0xce, 0xe4, 0x39,
	0x34, 0x46, 0x69, 0x74, 0x21, 0x62, 0xa7, 0x66, 0xa2, 0x16, 0x69, 0x35, 0xdf, 0xcc, 0x60, 0xd6,
	0xf3, 0xc1, 0x18, 0x40, 0xc7, 0x50, 0x37, 0xcf, 0xd6, 0xa5, 0x0e, 0x79, 0x88, 0xb6, 0xb8, 0x39,
	0xeb, 0x52, 0xe7, 0x81, 0xc0, 0x44, 0x99, 0x06, 0x3d, 0xcf, 0xa2, 0xb2, 0x6d, 0xed, 0xae, 0x2d,
	0xfd, 0x5b, 0xd5, 0x32, 0xcd, 0x5f, 0xb7, 0x24, 0xa1, 0xba, 0x22, 0x61, 0x13, 0x6a, 0xa7, 0xb8,
	0x28, 0xde, 0x76, 0x8a, 0x0b, 0xcd, 0xf4, 0x71, 0x22, 0x50, 0x15, 0x62, 0x73, 0x64, 0x3b, 0x07,
	0x73, 0x69, 0xd5, 0x5a, 0xa4, 0xe3, 0x5f, 0x50, 0xcd, 0x78, 0xe8, 0xd4, 0xf3, 0x78, 0x8e, 0xb4,
	0x1f, 0x0e, 0x67, 0x81, 0x90, 0xa8, 0x9c, 0x46, 0xee, 0x07, 0x0b, 0xc9, 0x36, 0xb4, 0xfd, 0x68,
	0x9a, 0x04, 0x2a, 0x13, 0xe8, 0x34, 0xcd, 0xdd, 0x5d, 0x80, 0x0c, 0xa0, 0xe3, 0x67, 0x93, 0x09,
	0x4a, 0xa9, 0xf9, 0x4e, 0xcb, 0xdc, 0x2f, 0x87, 0x86, 0x7f, 0xd6, 0xa1, 0x77, 0x5c, 0x3a, 0xdf,
	0xcf, 0x8d, 0x4f, 0x06, 0xd0, 0x1e, 0xa3, 0xb2, 0xcf, 0x6d, 0xb2, 0xdc, 0x29, 0x6e, 0x69, 0x5e,
	0x5a, 0x21, 0xbb, 0x00, 0xda, 0xb0, 0x96, 0x42, 0xd8, 0xbd, 0x9d, 0x70, 0x9b, 0xcc, 0x43, 0x99,
	0xc5, 0x8a, 0x56, 0xc8, 0x47, 0xe8, 0xea, 0x9b, 0xd2, 0xea, 0xe4, 0x19, 0x7b, 0xc8, 0xf6, 0xcb,
	0x29, 0x2f, 0xa1, 0xe3, 0xa1, 0xe4, 0x99, 0x98, 0xa0, 0x1e, 0x6d, 0xa9, 0xa1, 0x38, 0xd0, 0x0a,
	0xa1, 0x9a, 0x32, 0x8d, 0xa4, 0x42, 0xa1, 0xcd, 0xa3, 0xd5, 0x19, 0x1b, 0x2d, 0x73, 0x5e, 0x40,
	0x43, 0x5b, 0x37, 0x4d, 0x57, 0x5f, 0x61, 0x78, 0xb4, 0x42, 0xde, 0xc1, 0x86, 0xbe, 0x8e, 0x63,
	0x6b, 0xc0, 0x06, 0xfb, 0x34, 0x4f, 0xd5, 0xc2, 0xdd, 0x62, 0xf7, 0x97, 0xc3, 0x74, 0xec, 0xfa,
	0x59, 0x8a, 0xe2, 0x32, 0x88, 0xa3, 0x30, 0x50, 0x48, 0x5a, 0xcc, 0x7a, 0xdd, 0xad, 0xb3, 0x03,
	0xce, 0x63, 0x5a, 0x21, 0xaf, 0xa0, 0x7b, 0x1c, 0x67, 0x72, 0x66, 0x88, 0x7a, 0x0b, 0xee, 0x38,
	0xb6, 0x3a, 0xad, 0x90, 0xd7, 0xf0, 0x54, 0xcf, 0x77, 0x91, 0x4c, 0xce, 0xf8, 0x34, 0x4a, 0x1e,
	0x7b, 0xe1, 0x7e, 0x2e, 0xaf, 0x58, 0x3c, 0xd2, 0x67, 0x0f, 0xec, 0xe1, 0x63, 0x62, 0xdf, 0x42,
	0x77, 0x8c, 0xea, 0x70, 0x16, 0xc5, 0xa1, 0x4e, 0x91, 0xa4, 0xc1, 0x4e, 0x12, 0xb5, 0x37, 0x74,
	0x7b, 0xec, 0xff, 0x6f, 0x05, 0xad, 0x1c, 0xb4, 0xbf, 0x37, 0xd9, 0xbe, 0xf9, 0xc2, 0x5d, 0x35,
	0xcc, 0xcf, 0xde, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x91, 0xcf, 0xae, 0x12, 0x05, 0x00,
	0x00,
}
