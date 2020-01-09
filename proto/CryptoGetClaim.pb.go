// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoGetClaim.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Get a single claim attached to an account, or return null if it does not exist.
type CryptoGetClaimQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Hash                 []byte       `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CryptoGetClaimQuery) Reset()         { *m = CryptoGetClaimQuery{} }
func (m *CryptoGetClaimQuery) String() string { return proto.CompactTextString(m) }
func (*CryptoGetClaimQuery) ProtoMessage()    {}
func (*CryptoGetClaimQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9317107d4ba2c40, []int{0}
}

func (m *CryptoGetClaimQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetClaimQuery.Unmarshal(m, b)
}
func (m *CryptoGetClaimQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetClaimQuery.Marshal(b, m, deterministic)
}
func (m *CryptoGetClaimQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetClaimQuery.Merge(m, src)
}
func (m *CryptoGetClaimQuery) XXX_Size() int {
	return xxx_messageInfo_CryptoGetClaimQuery.Size(m)
}
func (m *CryptoGetClaimQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetClaimQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetClaimQuery proto.InternalMessageInfo

func (m *CryptoGetClaimQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetClaimQuery) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *CryptoGetClaimQuery) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// Response when the client sends the node CryptoGetClaimQuery. If the claim exists, there can be a state proof for that single claim. If the claim doesn't exist, then the state proof must be obtained for the account as a whole, which lists all the attached claims, which then proves that any claim not on the list must not exist.
type CryptoGetClaimResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Claim                *Claim          `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CryptoGetClaimResponse) Reset()         { *m = CryptoGetClaimResponse{} }
func (m *CryptoGetClaimResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoGetClaimResponse) ProtoMessage()    {}
func (*CryptoGetClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9317107d4ba2c40, []int{1}
}

func (m *CryptoGetClaimResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetClaimResponse.Unmarshal(m, b)
}
func (m *CryptoGetClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetClaimResponse.Marshal(b, m, deterministic)
}
func (m *CryptoGetClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetClaimResponse.Merge(m, src)
}
func (m *CryptoGetClaimResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoGetClaimResponse.Size(m)
}
func (m *CryptoGetClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetClaimResponse proto.InternalMessageInfo

func (m *CryptoGetClaimResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetClaimResponse) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoGetClaimQuery)(nil), "proto.CryptoGetClaimQuery")
	proto.RegisterType((*CryptoGetClaimResponse)(nil), "proto.CryptoGetClaimResponse")
}

func init() { proto.RegisterFile("CryptoGetClaim.proto", fileDescriptor_a9317107d4ba2c40) }

var fileDescriptor_a9317107d4ba2c40 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x71, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0x77, 0x4f, 0x2d, 0x71, 0xce, 0x49, 0xcc, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x52, 0x02, 0x4e, 0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9, 0xc5,
	0x10, 0x09, 0x29, 0xc1, 0xc0, 0xd2, 0xd4, 0xa2, 0x4a, 0x8f, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0xa8,
	0x90, 0x48, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xaa, 0x28, 0xc4, 0x5c, 0xc7, 0x94,
	0x14, 0x24, 0x73, 0x95, 0x5a, 0x19, 0xb9, 0x84, 0x51, 0x2d, 0x04, 0x9b, 0x27, 0xa4, 0xc5, 0xc5,
	0x96, 0x01, 0xd6, 0x2d, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0x04, 0x51, 0xaf, 0x87, 0x64,
	0x5b, 0x10, 0x54, 0x85, 0x90, 0x1e, 0x17, 0x67, 0x62, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x89, 0xa7,
	0x8b, 0x04, 0x13, 0x58, 0xb9, 0x00, 0x54, 0xb9, 0x23, 0x4c, 0x3c, 0x08, 0xa1, 0x44, 0x48, 0x88,
	0x8b, 0x25, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x56, 0xca,
	0xe6, 0x12, 0x43, 0x75, 0x06, 0xcc, 0x0f, 0x42, 0xba, 0x68, 0x2e, 0x11, 0x85, 0x1a, 0x8d, 0xea,
	0x49, 0xb8, 0x63, 0x94, 0xb8, 0x58, 0x93, 0x41, 0xfa, 0xa1, 0x0e, 0xe1, 0x81, 0xaa, 0x86, 0x98,
	0x09, 0x91, 0x72, 0x92, 0xe3, 0x92, 0x4a, 0xce, 0xcf, 0xd5, 0xcb, 0x48, 0x4d, 0x49, 0x2d, 0x4a,
	0xd4, 0x03, 0xd9, 0x9f, 0x5e, 0x94, 0x58, 0x90, 0x01, 0x51, 0x1a, 0xc0, 0x98, 0xc4, 0x06, 0x66,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0x48, 0xb2, 0x14, 0x8b, 0x01, 0x00, 0x00,
}