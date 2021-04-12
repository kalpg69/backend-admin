// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/protos/v1/user.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type CreateUserRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type UpdateUserRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{4}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type DeleteUserRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{5}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type DeleteUserResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{6}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{7}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20232d306d4d80df, []int{8}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "v1.User")
	proto.RegisterType((*CreateUserRequest)(nil), "v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "v1.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "v1.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "v1.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "v1.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "v1.DeleteUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "v1.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "v1.GetUserResponse")
}

func init() { proto.RegisterFile("src/protos/v1/user.proto", fileDescriptor_20232d306d4d80df) }

var fileDescriptor_20232d306d4d80df = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdd, 0x8a, 0xda, 0x40,
	0x14, 0xc7, 0x31, 0xda, 0x9a, 0x1e, 0xfb, 0xe5, 0x88, 0x25, 0xa4, 0x52, 0x64, 0xae, 0x44, 0x4a,
	0x42, 0xec, 0x5d, 0x0b, 0x05, 0x6b, 0xa1, 0x97, 0x2d, 0x29, 0x42, 0x69, 0x2f, 0x64, 0xaa, 0x07,
	0x09, 0xd8, 0x4c, 0x3a, 0x13, 0xd3, 0x8b, 0xb2, 0x37, 0xfb, 0x0a, 0xfb, 0x68, 0xfb, 0x0a, 0x0b,
	0xfb, 0x1a, 0xcb, 0xcc, 0x24, 0x9a, 0x18, 0x77, 0x17, 0xd9, 0xbb, 0x39, 0x5f, 0xbf, 0xf3, 0x3f,
	0x73, 0x66, 0xc0, 0x91, 0x62, 0xe9, 0x27, 0x82, 0xa7, 0x5c, 0xfa, 0x59, 0xe0, 0x6f, 0x25, 0x0a,
	0x4f, 0x9b, 0xc4, 0xca, 0x02, 0x77, 0xb0, 0xe6, 0x7c, 0xbd, 0x41, 0x9f, 0x25, 0x91, 0xcf, 0xe2,
	0x98, 0xa7, 0x2c, 0x8d, 0x78, 0x2c, 0x4d, 0x86, 0x3b, 0xac, 0xd6, 0xb2, 0x24, 0x5a, 0x08, 0x94,
	0x09, 0x8f, 0x25, 0x9a, 0x0c, 0x1a, 0x42, 0x6b, 0x2e, 0x51, 0x10, 0x17, 0x6c, 0x45, 0x8e, 0xd9,
	0x1f, 0x74, 0x1a, 0xc3, 0xc6, 0xe8, 0x49, 0xb8, 0xb3, 0x55, 0x2c, 0x61, 0x52, 0xfe, 0xe3, 0x62,
	0xe5, 0x58, 0x26, 0x56, 0xd8, 0x84, 0x40, 0x4b, 0xf0, 0x0d, 0x3a, 0x4d, 0xed, 0xd7, 0x67, 0x3a,
	0x83, 0xee, 0x4c, 0x20, 0x4b, 0x51, 0x91, 0x43, 0xfc, 0xbb, 0x45, 0x99, 0x92, 0x97, 0xd0, 0x64,
	0x49, 0x94, 0xb3, 0xd5, 0x91, 0x0c, 0xa0, 0xa5, 0x5a, 0x68, 0x64, 0x67, 0x62, 0x7b, 0x59, 0xe0,
	0xe9, 0x02, 0xed, 0xa5, 0x1f, 0x81, 0x94, 0x21, 0x46, 0x34, 0x19, 0x81, 0x5d, 0x0c, 0xa0, 0x51,
	0x9d, 0xc9, 0x53, 0x55, 0x57, 0xc4, 0xc3, 0x5d, 0x94, 0x2e, 0xa0, 0x3b, 0x4f, 0x56, 0xf7, 0x8a,
	0x28, 0xcf, 0x6d, 0x1d, 0xcc, 0x5d, 0x08, 0x6c, 0xde, 0x26, 0xb0, 0xdc, 0xe0, 0x64, 0x81, 0x53,
	0xe8, 0x7e, 0xc6, 0x0d, 0x3e, 0x40, 0xa0, 0x92, 0x50, 0x46, 0x9c, 0x2c, 0xe1, 0x2d, 0x3c, 0xff,
	0x82, 0x69, 0xb9, 0xff, 0x1d, 0xcf, 0x80, 0x4e, 0xe1, 0xc5, 0x2e, 0x3b, 0x6f, 0x75, 0xe2, 0x52,
	0x27, 0xd7, 0x16, 0x74, 0x94, 0xf9, 0x1d, 0x45, 0x16, 0x2d, 0x91, 0x7c, 0x03, 0xd8, 0x2f, 0x99,
	0xf4, 0x55, 0x76, 0xed, 0xe5, 0xb8, 0xaf, 0x0e, 0xdd, 0xb9, 0xfa, 0xde, 0xf9, 0xe5, 0xd5, 0x85,
	0xf5, 0x8c, 0xda, 0xc5, 0xb7, 0x78, 0xdf, 0x18, 0x93, 0x5f, 0x00, 0xfb, 0xad, 0x18, 0x62, 0xed,
	0x19, 0x18, 0x62, 0x7d, 0x79, 0xf4, 0x8d, 0x26, 0x3a, 0x6e, 0xaf, 0x20, 0xfa, 0xff, 0x8b, 0xe9,
	0xcf, 0x14, 0xfc, 0x07, 0xc0, 0xfe, 0xbe, 0x0d, 0xbc, 0xb6, 0x42, 0x03, 0xaf, 0xaf, 0x85, 0xbe,
	0xd6, 0xf0, 0xfe, 0xf8, 0x18, 0x9c, 0x7c, 0x85, 0x76, 0x7e, 0xb7, 0x84, 0xa8, 0xfa, 0xea, 0x5a,
	0xdc, 0x5e, 0xc5, 0x57, 0x05, 0x92, 0x63, 0xc0, 0x4f, 0xed, 0x9f, 0x8f, 0x3c, 0xff, 0x43, 0x16,
	0xfc, 0x7e, 0xac, 0xff, 0xf9, 0xbb, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0xf6, 0xaf, 0x0b,
	0x47, 0x04, 0x00, 0x00,
}