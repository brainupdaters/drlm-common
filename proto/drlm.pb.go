// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drlm.proto

package drlm

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

type UserLoginRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Usr                  string   `protobuf:"bytes,2,opt,name=usr,proto3" json:"usr,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{0}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UserLoginRequest) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

func (m *UserLoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserLoginResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{1}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

type UserAddRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Usr                  string   `protobuf:"bytes,2,opt,name=usr,proto3" json:"usr,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddRequest) Reset()         { *m = UserAddRequest{} }
func (m *UserAddRequest) String() string { return proto.CompactTextString(m) }
func (*UserAddRequest) ProtoMessage()    {}
func (*UserAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{2}
}

func (m *UserAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddRequest.Unmarshal(m, b)
}
func (m *UserAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddRequest.Marshal(b, m, deterministic)
}
func (m *UserAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddRequest.Merge(m, src)
}
func (m *UserAddRequest) XXX_Size() int {
	return xxx_messageInfo_UserAddRequest.Size(m)
}
func (m *UserAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddRequest proto.InternalMessageInfo

func (m *UserAddRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UserAddRequest) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

func (m *UserAddRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserAddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddResponse) Reset()         { *m = UserAddResponse{} }
func (m *UserAddResponse) String() string { return proto.CompactTextString(m) }
func (*UserAddResponse) ProtoMessage()    {}
func (*UserAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{3}
}

func (m *UserAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddResponse.Unmarshal(m, b)
}
func (m *UserAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddResponse.Marshal(b, m, deterministic)
}
func (m *UserAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddResponse.Merge(m, src)
}
func (m *UserAddResponse) XXX_Size() int {
	return xxx_messageInfo_UserAddResponse.Size(m)
}
func (m *UserAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddResponse proto.InternalMessageInfo

type UserDeleteRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Usr                  string   `protobuf:"bytes,2,opt,name=usr,proto3" json:"usr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteRequest) Reset()         { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()    {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{4}
}

func (m *UserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRequest.Unmarshal(m, b)
}
func (m *UserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRequest.Merge(m, src)
}
func (m *UserDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRequest.Size(m)
}
func (m *UserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRequest proto.InternalMessageInfo

func (m *UserDeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UserDeleteRequest) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

type UserDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteResponse) Reset()         { *m = UserDeleteResponse{} }
func (m *UserDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*UserDeleteResponse) ProtoMessage()    {}
func (*UserDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{5}
}

func (m *UserDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteResponse.Unmarshal(m, b)
}
func (m *UserDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteResponse.Marshal(b, m, deterministic)
}
func (m *UserDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteResponse.Merge(m, src)
}
func (m *UserDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_UserDeleteResponse.Size(m)
}
func (m *UserDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteResponse proto.InternalMessageInfo

type UserListRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRequest) Reset()         { *m = UserListRequest{} }
func (m *UserListRequest) String() string { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()    {}
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{6}
}

func (m *UserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest.Unmarshal(m, b)
}
func (m *UserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest.Marshal(b, m, deterministic)
}
func (m *UserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest.Merge(m, src)
}
func (m *UserListRequest) XXX_Size() int {
	return xxx_messageInfo_UserListRequest.Size(m)
}
func (m *UserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest proto.InternalMessageInfo

func (m *UserListRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type UserListResponse struct {
	Users                []*UserListResponse_User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{7}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUsers() []*UserListResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserListResponse_User struct {
	Usr                  string   `protobuf:"bytes,1,opt,name=usr,proto3" json:"usr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListResponse_User) Reset()         { *m = UserListResponse_User{} }
func (m *UserListResponse_User) String() string { return proto.CompactTextString(m) }
func (*UserListResponse_User) ProtoMessage()    {}
func (*UserListResponse_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{7, 0}
}

func (m *UserListResponse_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse_User.Unmarshal(m, b)
}
func (m *UserListResponse_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse_User.Marshal(b, m, deterministic)
}
func (m *UserListResponse_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse_User.Merge(m, src)
}
func (m *UserListResponse_User) XXX_Size() int {
	return xxx_messageInfo_UserListResponse_User.Size(m)
}
func (m *UserListResponse_User) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse_User.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse_User proto.InternalMessageInfo

func (m *UserListResponse_User) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

func init() {
	proto.RegisterType((*UserLoginRequest)(nil), "drlm.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "drlm.UserLoginResponse")
	proto.RegisterType((*UserAddRequest)(nil), "drlm.UserAddRequest")
	proto.RegisterType((*UserAddResponse)(nil), "drlm.UserAddResponse")
	proto.RegisterType((*UserDeleteRequest)(nil), "drlm.UserDeleteRequest")
	proto.RegisterType((*UserDeleteResponse)(nil), "drlm.UserDeleteResponse")
	proto.RegisterType((*UserListRequest)(nil), "drlm.UserListRequest")
	proto.RegisterType((*UserListResponse)(nil), "drlm.UserListResponse")
	proto.RegisterType((*UserListResponse_User)(nil), "drlm.UserListResponse.User")
}

func init() { proto.RegisterFile("drlm.proto", fileDescriptor_a4bd9cd91f607bb1) }

var fileDescriptor_a4bd9cd91f607bb1 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd7, 0xb5, 0xfe, 0xd8, 0x13, 0x74, 0x7b, 0x6e, 0x33, 0xd4, 0xcb, 0x88, 0x97, 0x9d,
	0x06, 0xce, 0x83, 0x82, 0x20, 0x0c, 0x86, 0x78, 0xa8, 0x97, 0x82, 0x67, 0x51, 0x1a, 0xa4, 0x30,
	0xd7, 0x98, 0xd7, 0xb2, 0xab, 0x7f, 0xba, 0xa4, 0xc9, 0xd2, 0xd8, 0xe9, 0x41, 0xbc, 0xa5, 0xdf,
	0xbe, 0xf7, 0x79, 0xc9, 0x27, 0x01, 0xc8, 0xd4, 0xea, 0x7d, 0x26, 0x55, 0x51, 0x16, 0x18, 0xe9,
	0x35, 0x7f, 0x80, 0xfe, 0x13, 0x09, 0x95, 0x14, 0x6f, 0xf9, 0x3a, 0x15, 0x1f, 0x95, 0xa0, 0x12,
	0xfb, 0x10, 0xbe, 0xc8, 0x9c, 0x05, 0x93, 0x60, 0xda, 0x4b, 0xf5, 0x52, 0x27, 0x15, 0x29, 0xd6,
	0x35, 0x49, 0x45, 0x4a, 0x27, 0x72, 0x93, 0xb1, 0xd0, 0x24, 0x72, 0x93, 0xf1, 0x53, 0x18, 0x78,
	0x24, 0x92, 0xc5, 0x9a, 0x04, 0xbf, 0x87, 0x63, 0x1d, 0x2e, 0xb2, 0xec, 0x7f, 0xf0, 0x01, 0x9c,
	0x38, 0x8e, 0x45, 0x5f, 0x9b, 0x79, 0x4b, 0xb1, 0x12, 0xa5, 0xf8, 0x03, 0x9d, 0x0f, 0x01, 0xfd,
	0x46, 0x8b, 0xbb, 0x30, 0x13, 0x92, 0x9c, 0xca, 0x5f, 0x61, 0xfc, 0xd9, 0xda, 0xaa, 0x8b, 0x4c,
	0x23, 0x5e, 0xc2, 0x5e, 0x45, 0x42, 0x11, 0x0b, 0x26, 0xe1, 0xf4, 0x68, 0x7e, 0x3e, 0xab, 0x1d,
	0xb7, 0xcb, 0xea, 0x20, 0x35, 0x95, 0x31, 0x83, 0x48, 0x7f, 0x6e, 0xf7, 0x16, 0xb8, 0xbd, 0xcd,
	0x3f, 0xbb, 0x10, 0x2d, 0xd3, 0xe4, 0x11, 0xef, 0xa0, 0xe7, 0x6c, 0xe2, 0xd8, 0x63, 0x7a, 0x17,
	0x15, 0x9f, 0xed, 0xe4, 0xf6, 0x30, 0x1d, 0xbc, 0x81, 0x03, 0x2b, 0x0c, 0x87, 0x4d, 0x55, 0x73,
	0x0f, 0xf1, 0xa8, 0x95, 0xba, 0xce, 0x05, 0x40, 0xa3, 0x07, 0xbd, 0x11, 0xdf, 0x4c, 0xc7, 0x6c,
	0xf7, 0x87, 0x43, 0xdc, 0xc2, 0xe1, 0xf6, 0xfc, 0x38, 0x6a, 0xfb, 0x30, 0xed, 0xe3, 0x9f, 0x35,
	0xf1, 0xce, 0xeb, 0x7e, 0xfd, 0x3c, 0xaf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x89, 0xb0,
	0x59, 0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DRLMClient is the client API for DRLM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DRLMClient interface {
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserAdd(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserAddResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
}

type dRLMClient struct {
	cc *grpc.ClientConn
}

func NewDRLMClient(cc *grpc.ClientConn) DRLMClient {
	return &dRLMClient{cc}
}

func (c *dRLMClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserAdd(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserAddResponse, error) {
	out := new(UserAddResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DRLMServer is the server API for DRLM service.
type DRLMServer interface {
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserAdd(context.Context, *UserAddRequest) (*UserAddResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
}

// UnimplementedDRLMServer can be embedded to have forward compatible implementations.
type UnimplementedDRLMServer struct {
}

func (*UnimplementedDRLMServer) UserLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedDRLMServer) UserAdd(ctx context.Context, req *UserAddRequest) (*UserAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (*UnimplementedDRLMServer) UserDelete(ctx context.Context, req *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (*UnimplementedDRLMServer) UserList(ctx context.Context, req *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}

func RegisterDRLMServer(s *grpc.Server, srv DRLMServer) {
	s.RegisterService(&_DRLM_serviceDesc, srv)
}

func _DRLM_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserAdd(ctx, req.(*UserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DRLM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drlm.DRLM",
	HandlerType: (*DRLMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _DRLM_UserLogin_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _DRLM_UserAdd_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _DRLM_UserDelete_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _DRLM_UserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drlm.proto",
}
