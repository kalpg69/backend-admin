// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/protos/v1/customer.proto

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

type Customer struct {
	CustomerCode         string   `protobuf:"bytes,1,opt,name=customer_code,json=customerCode,proto3" json:"customer_code,omitempty"`
	CustomerName         string   `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	CustomerAddress      string   `protobuf:"bytes,3,opt,name=customer_address,json=customerAddress,proto3" json:"customer_address,omitempty"`
	CustomerEmail        string   `protobuf:"bytes,4,opt,name=customer_email,json=customerEmail,proto3" json:"customer_email,omitempty"`
	CustomerPhone        string   `protobuf:"bytes,5,opt,name=customer_phone,json=customerPhone,proto3" json:"customer_phone,omitempty"`
	LastUpdatedBy        string   `protobuf:"bytes,6,opt,name=last_updated_by,json=lastUpdatedBy,proto3" json:"last_updated_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetCustomerCode() string {
	if m != nil {
		return m.CustomerCode
	}
	return ""
}

func (m *Customer) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *Customer) GetCustomerAddress() string {
	if m != nil {
		return m.CustomerAddress
	}
	return ""
}

func (m *Customer) GetCustomerEmail() string {
	if m != nil {
		return m.CustomerEmail
	}
	return ""
}

func (m *Customer) GetCustomerPhone() string {
	if m != nil {
		return m.CustomerPhone
	}
	return ""
}

func (m *Customer) GetLastUpdatedBy() string {
	if m != nil {
		return m.LastUpdatedBy
	}
	return ""
}

type CreateCustomerRequest struct {
	Api                  string    `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Customer             *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerRequest) Reset()         { *m = CreateCustomerRequest{} }
func (m *CreateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerRequest) ProtoMessage()    {}
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{1}
}

func (m *CreateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerRequest.Unmarshal(m, b)
}
func (m *CreateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerRequest.Merge(m, src)
}
func (m *CreateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerRequest.Size(m)
}
func (m *CreateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerRequest proto.InternalMessageInfo

func (m *CreateCustomerRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateCustomerRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type CreateCustomerResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerResponse) Reset()         { *m = CreateCustomerResponse{} }
func (m *CreateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerResponse) ProtoMessage()    {}
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{2}
}

func (m *CreateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerResponse.Unmarshal(m, b)
}
func (m *CreateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *CreateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerResponse.Merge(m, src)
}
func (m *CreateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerResponse.Size(m)
}
func (m *CreateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerResponse proto.InternalMessageInfo

func (m *CreateCustomerResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type UpdateCustomerRequest struct {
	Api                  string    `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	CustomerCode         string    `protobuf:"bytes,2,opt,name=customer_code,json=customerCode,proto3" json:"customer_code,omitempty"`
	Customer             *Customer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateCustomerRequest) Reset()         { *m = UpdateCustomerRequest{} }
func (m *UpdateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerRequest) ProtoMessage()    {}
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{3}
}

func (m *UpdateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerRequest.Unmarshal(m, b)
}
func (m *UpdateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerRequest.Merge(m, src)
}
func (m *UpdateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerRequest.Size(m)
}
func (m *UpdateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerRequest proto.InternalMessageInfo

func (m *UpdateCustomerRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateCustomerRequest) GetCustomerCode() string {
	if m != nil {
		return m.CustomerCode
	}
	return ""
}

func (m *UpdateCustomerRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type UpdateCustomerResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateCustomerResponse) Reset()         { *m = UpdateCustomerResponse{} }
func (m *UpdateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCustomerResponse) ProtoMessage()    {}
func (*UpdateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{4}
}

func (m *UpdateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCustomerResponse.Unmarshal(m, b)
}
func (m *UpdateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCustomerResponse.Merge(m, src)
}
func (m *UpdateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCustomerResponse.Size(m)
}
func (m *UpdateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCustomerResponse proto.InternalMessageInfo

func (m *UpdateCustomerResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type DeleteCustomerRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	CustomerCode         string   `protobuf:"bytes,2,opt,name=customer_code,json=customerCode,proto3" json:"customer_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCustomerRequest) Reset()         { *m = DeleteCustomerRequest{} }
func (m *DeleteCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerRequest) ProtoMessage()    {}
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{5}
}

func (m *DeleteCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerRequest.Unmarshal(m, b)
}
func (m *DeleteCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerRequest.Merge(m, src)
}
func (m *DeleteCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerRequest.Size(m)
}
func (m *DeleteCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerRequest proto.InternalMessageInfo

func (m *DeleteCustomerRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteCustomerRequest) GetCustomerCode() string {
	if m != nil {
		return m.CustomerCode
	}
	return ""
}

type DeleteCustomerResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteCustomerResponse) Reset()         { *m = DeleteCustomerResponse{} }
func (m *DeleteCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCustomerResponse) ProtoMessage()    {}
func (*DeleteCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_885660cbdaa161bf, []int{6}
}

func (m *DeleteCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCustomerResponse.Unmarshal(m, b)
}
func (m *DeleteCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCustomerResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCustomerResponse.Merge(m, src)
}
func (m *DeleteCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCustomerResponse.Size(m)
}
func (m *DeleteCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCustomerResponse proto.InternalMessageInfo

func (m *DeleteCustomerResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "v1.Customer")
	proto.RegisterType((*CreateCustomerRequest)(nil), "v1.CreateCustomerRequest")
	proto.RegisterType((*CreateCustomerResponse)(nil), "v1.CreateCustomerResponse")
	proto.RegisterType((*UpdateCustomerRequest)(nil), "v1.UpdateCustomerRequest")
	proto.RegisterType((*UpdateCustomerResponse)(nil), "v1.UpdateCustomerResponse")
	proto.RegisterType((*DeleteCustomerRequest)(nil), "v1.DeleteCustomerRequest")
	proto.RegisterType((*DeleteCustomerResponse)(nil), "v1.DeleteCustomerResponse")
}

func init() { proto.RegisterFile("src/protos/v1/customer.proto", fileDescriptor_885660cbdaa161bf) }

var fileDescriptor_885660cbdaa161bf = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x15, 0x87, 0x96, 0x32, 0x04, 0xa7, 0xac, 0x14, 0x30, 0x56, 0x84, 0x8a, 0xf9, 0x57,
	0x7a, 0xb0, 0xe5, 0x72, 0x83, 0x13, 0x09, 0x5c, 0x2b, 0x94, 0x8a, 0x0b, 0x17, 0xb3, 0xb5, 0x47,
	0xc5, 0x92, 0xed, 0x35, 0xde, 0x8d, 0x45, 0x85, 0xb8, 0xf0, 0x0a, 0x3c, 0x1a, 0x3c, 0x02, 0xaf,
	0xc0, 0x1d, 0xed, 0xae, 0xd7, 0xca, 0x06, 0x8b, 0xaa, 0x12, 0xb7, 0xe8, 0x9b, 0x6f, 0xe7, 0xf7,
	0xcd, 0x4c, 0x12, 0x98, 0xf3, 0x26, 0x8d, 0xea, 0x86, 0x09, 0xc6, 0xa3, 0x36, 0x8e, 0xd2, 0x35,
	0x17, 0xac, 0xc4, 0x26, 0x54, 0x12, 0x71, 0xda, 0xd8, 0x3f, 0xb0, 0x1d, 0xb4, 0xce, 0x93, 0x06,
	0x79, 0xcd, 0x2a, 0x8e, 0xda, 0xe5, 0xcf, 0xcf, 0x19, 0x3b, 0x2f, 0x50, 0x96, 0x22, 0x5a, 0x55,
	0x4c, 0x50, 0x91, 0xb3, 0x8a, 0xeb, 0x6a, 0xf0, 0x7b, 0x04, 0x7b, 0xcb, 0xae, 0x2d, 0x79, 0x08,
	0xb7, 0x0c, 0x22, 0x49, 0x59, 0x86, 0xde, 0xe8, 0x60, 0x74, 0x78, 0x63, 0x35, 0x31, 0xe2, 0x92,
	0x65, 0x68, 0x99, 0x2a, 0x5a, 0xa2, 0xe7, 0xd8, 0xa6, 0x13, 0x5a, 0x22, 0x79, 0x06, 0xfb, 0xbd,
	0x89, 0x66, 0x59, 0x83, 0x9c, 0x7b, 0x63, 0xe5, 0x9b, 0x1a, 0xfd, 0x95, 0x96, 0xc9, 0x63, 0x70,
	0x7b, 0x2b, 0x96, 0x34, 0x2f, 0xbc, 0x6b, 0xca, 0xd8, 0x53, 0xde, 0x48, 0xd1, 0xb2, 0xd5, 0x1f,
	0x59, 0x85, 0xde, 0x8e, 0x6d, 0x7b, 0x2b, 0x45, 0xf2, 0x04, 0xa6, 0x05, 0xe5, 0x22, 0x59, 0xd7,
	0x19, 0x15, 0x98, 0x25, 0x67, 0x17, 0xde, 0xae, 0xf6, 0x49, 0xf9, 0x9d, 0x56, 0x17, 0x17, 0xc1,
	0x29, 0xcc, 0x96, 0x0d, 0x52, 0x81, 0x66, 0xf8, 0x15, 0x7e, 0x5a, 0x23, 0x17, 0x64, 0x1f, 0xc6,
	0xb4, 0xce, 0xbb, 0xc9, 0xe5, 0x47, 0x72, 0x08, 0x7b, 0x86, 0xa1, 0x66, 0xbd, 0x79, 0x3c, 0x09,
	0xdb, 0x38, 0xec, 0x1f, 0xf6, 0xd5, 0x60, 0x01, 0x77, 0xb6, 0x9b, 0xea, 0x53, 0xc8, 0x1e, 0xe6,
	0x2c, 0xaa, 0x75, 0xd7, 0xc3, 0xd4, 0x57, 0x7d, 0x35, 0xf8, 0x0c, 0x33, 0x9d, 0xf2, 0xf2, 0x60,
	0x7f, 0x9d, 0xcb, 0x19, 0x38, 0xd7, 0x66, 0xfa, 0xf1, 0x65, 0xe9, 0xb7, 0xc9, 0x57, 0x4e, 0x7f,
	0x02, 0xb3, 0xd7, 0x58, 0xe0, 0xff, 0x4a, 0x2f, 0x33, 0x6d, 0xf7, 0xbb, 0x6a, 0xa6, 0xe3, 0x9f,
	0x0e, 0x4c, 0xcd, 0xf3, 0x53, 0x6c, 0xda, 0x3c, 0x45, 0xf2, 0x01, 0x5c, 0xfb, 0x52, 0xe4, 0x9e,
	0xda, 0xca, 0xd0, 0x57, 0xc2, 0xf7, 0x87, 0x4a, 0xdd, 0xc0, 0x77, 0xbf, 0xfd, 0xf8, 0xf5, 0xdd,
	0xb9, 0x1d, 0x4c, 0x36, 0x7f, 0x9f, 0x2f, 0x46, 0x47, 0xa4, 0x06, 0xd7, 0xde, 0xa6, 0x26, 0x0c,
	0xde, 0x56, 0x13, 0x86, 0x97, 0x1f, 0x3c, 0x55, 0x84, 0x07, 0xfe, 0x7c, 0x93, 0x10, 0x7d, 0xb1,
	0x76, 0xf7, 0x55, 0x12, 0x4b, 0x70, 0xed, 0x5d, 0x69, 0xe2, 0xe0, 0x3d, 0x34, 0x71, 0x78, 0xb5,
	0xc1, 0x23, 0x45, 0xbc, 0x7f, 0xf4, 0x4f, 0xe2, 0xe2, 0xfa, 0xfb, 0x9d, 0x30, 0x7a, 0xd9, 0xc6,
	0x67, 0xbb, 0xea, 0x9f, 0xe4, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0xce, 0x19, 0x43,
	0xad, 0x04, 0x00, 0x00,
}
