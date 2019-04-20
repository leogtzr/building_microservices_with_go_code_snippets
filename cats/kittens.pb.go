// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kittens.proto

package bmigo_micro

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestEnvelope struct {
	ServiceMethod        string   `protobuf:"bytes,1,opt,name=service_method,json=serviceMethod,proto3" json:"service_method,omitempty"`
	Seq                  uint64   `protobuf:"fixed64,2,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestEnvelope) Reset()         { *m = RequestEnvelope{} }
func (m *RequestEnvelope) String() string { return proto.CompactTextString(m) }
func (*RequestEnvelope) ProtoMessage()    {}
func (*RequestEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_kittens_56e5e9406ab79eb5, []int{0}
}
func (m *RequestEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestEnvelope.Unmarshal(m, b)
}
func (m *RequestEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestEnvelope.Marshal(b, m, deterministic)
}
func (dst *RequestEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEnvelope.Merge(dst, src)
}
func (m *RequestEnvelope) XXX_Size() int {
	return xxx_messageInfo_RequestEnvelope.Size(m)
}
func (m *RequestEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEnvelope proto.InternalMessageInfo

func (m *RequestEnvelope) GetServiceMethod() string {
	if m != nil {
		return m.ServiceMethod
	}
	return ""
}

func (m *RequestEnvelope) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type ResponseEnvelope struct {
	ServiceMethod        string   `protobuf:"bytes,1,opt,name=service_method,json=serviceMethod,proto3" json:"service_method,omitempty"`
	Seq                  uint64   `protobuf:"fixed64,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseEnvelope) Reset()         { *m = ResponseEnvelope{} }
func (m *ResponseEnvelope) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelope) ProtoMessage()    {}
func (*ResponseEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_kittens_56e5e9406ab79eb5, []int{1}
}
func (m *ResponseEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseEnvelope.Unmarshal(m, b)
}
func (m *ResponseEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseEnvelope.Marshal(b, m, deterministic)
}
func (dst *ResponseEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEnvelope.Merge(dst, src)
}
func (m *ResponseEnvelope) XXX_Size() int {
	return xxx_messageInfo_ResponseEnvelope.Size(m)
}
func (m *ResponseEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEnvelope proto.InternalMessageInfo

func (m *ResponseEnvelope) GetServiceMethod() string {
	if m != nil {
		return m.ServiceMethod
	}
	return ""
}

func (m *ResponseEnvelope) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ResponseEnvelope) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_kittens_56e5e9406ab79eb5, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_kittens_56e5e9406ab79eb5, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestEnvelope)(nil), "bmigo.micro.RequestEnvelope")
	proto.RegisterType((*ResponseEnvelope)(nil), "bmigo.micro.ResponseEnvelope")
	proto.RegisterType((*Request)(nil), "bmigo.micro.Request")
	proto.RegisterType((*Response)(nil), "bmigo.micro.Response")
}

func init() { proto.RegisterFile("kittens.proto", fileDescriptor_kittens_56e5e9406ab79eb5) }

var fileDescriptor_kittens_56e5e9406ab79eb5 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x5d, 0xd7, 0xdd, 0xea, 0x48, 0xb5, 0x0c, 0x15, 0x8a, 0x28, 0x48, 0x40, 0xf0, 0x94,
	0x83, 0x82, 0x77, 0x0f, 0x82, 0x28, 0x5e, 0xf2, 0x07, 0xa4, 0xad, 0x43, 0x0d, 0x36, 0x99, 0x36,
	0x89, 0xfd, 0xfd, 0xda, 0x34, 0x82, 0x78, 0xdd, 0xdb, 0x9b, 0xf7, 0x86, 0x8f, 0xc7, 0x83, 0xfc,
	0x53, 0x87, 0x40, 0xd6, 0xcb, 0xc1, 0x71, 0x60, 0x3c, 0x6e, 0x8c, 0xee, 0x58, 0x1a, 0xdd, 0x3a,
	0x16, 0xcf, 0x70, 0xaa, 0x68, 0xfc, 0x22, 0x1f, 0x1e, 0xed, 0x44, 0x3d, 0x0f, 0x84, 0xd7, 0x70,
	0xe2, 0xc9, 0x4d, 0xba, 0xa5, 0x37, 0x43, 0xe1, 0x83, 0xdf, 0xab, 0xd5, 0xd5, 0xea, 0xe6, 0x48,
	0xe5, 0xc9, 0x7d, 0x8d, 0x26, 0x16, 0xb0, 0xf6, 0x34, 0x56, 0xfb, 0x3f, 0xd9, 0x56, 0xcd, 0x52,
	0xd4, 0x50, 0x28, 0xf2, 0x03, 0x5b, 0x4f, 0x3b, 0xc3, 0xb0, 0x84, 0x0d, 0x39, 0xc7, 0xae, 0x5a,
	0xc7, 0xff, 0xe5, 0x10, 0x97, 0x90, 0xa5, 0xba, 0x88, 0x70, 0x60, 0x6b, 0x43, 0x89, 0x17, 0xb5,
	0xb8, 0x80, 0xc3, 0xdf, 0x06, 0x33, 0xd2, 0xf8, 0x2e, 0xc5, 0xb3, 0xbc, 0x7d, 0x80, 0xec, 0x65,
	0x59, 0x02, 0xef, 0x61, 0xf3, 0x44, 0x7d, 0xcf, 0x58, 0xca, 0x3f, 0x6b, 0xc8, 0xc4, 0x3e, 0x3f,
	0xfb, 0xe7, 0x2e, 0x48, 0xb1, 0xd7, 0x6c, 0xe3, 0x84, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0x7b, 0xe5, 0xf3, 0x53, 0x01, 0x00, 0x00,
}
