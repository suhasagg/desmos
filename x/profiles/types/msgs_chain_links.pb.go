// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/msgs_chain_links.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgLinkChainAccount represents a message to link an account to a profile.
type MsgLinkChainAccount struct {
	// ChainAddress contains the details of the external chain address to be
	// linked
	ChainAddress *types.Any `protobuf:"bytes,1,opt,name=chain_address,json=chainAddress,proto3" json:"chain_address,omitempty" yaml:"source_address"`
	// Proof contains the proof of ownership of the external chain address
	Proof Proof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof" yaml:"source_proof"`
	// ChainConfig contains the configuration of the external chain
	ChainConfig ChainConfig `protobuf:"bytes,3,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"source_chain_config"`
	// Signer represents the Desmos address associated with the
	// profile to which link the external account
	Signer string `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
}

func (m *MsgLinkChainAccount) Reset()         { *m = MsgLinkChainAccount{} }
func (m *MsgLinkChainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgLinkChainAccount) ProtoMessage()    {}
func (*MsgLinkChainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9916e62e3c000d, []int{0}
}
func (m *MsgLinkChainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkChainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkChainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkChainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkChainAccount.Merge(m, src)
}
func (m *MsgLinkChainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkChainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkChainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkChainAccount proto.InternalMessageInfo

func (m *MsgLinkChainAccount) GetChainAddress() *types.Any {
	if m != nil {
		return m.ChainAddress
	}
	return nil
}

func (m *MsgLinkChainAccount) GetProof() Proof {
	if m != nil {
		return m.Proof
	}
	return Proof{}
}

func (m *MsgLinkChainAccount) GetChainConfig() ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return ChainConfig{}
}

func (m *MsgLinkChainAccount) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

// LinkChainAccountResponse defines the Msg/LinkAccount response type.
type LinkChainAccountResponse struct {
}

func (m *LinkChainAccountResponse) Reset()         { *m = LinkChainAccountResponse{} }
func (m *LinkChainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountResponse) ProtoMessage()    {}
func (*LinkChainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9916e62e3c000d, []int{1}
}
func (m *LinkChainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountResponse.Merge(m, src)
}
func (m *LinkChainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountResponse proto.InternalMessageInfo

// MsgUnlinkChainAccount represents a message to unlink an account from a
// profile.
type MsgUnlinkChainAccount struct {
	// Owner represents the Desmos profile from which to remove the link
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// ChainName represents the name of the chain to which the link to remove is
	// associated
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty" yaml:"chain_name"`
	// Target represents the external address to be removed
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty" yaml:"target"`
}

func (m *MsgUnlinkChainAccount) Reset()         { *m = MsgUnlinkChainAccount{} }
func (m *MsgUnlinkChainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkChainAccount) ProtoMessage()    {}
func (*MsgUnlinkChainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9916e62e3c000d, []int{2}
}
func (m *MsgUnlinkChainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkChainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkChainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkChainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkChainAccount.Merge(m, src)
}
func (m *MsgUnlinkChainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkChainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkChainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkChainAccount proto.InternalMessageInfo

func (m *MsgUnlinkChainAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUnlinkChainAccount) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgUnlinkChainAccount) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

// UnlinkChainAccountResponse defines the Msg/UnlinkAccount response type.
type UnlinkChainAccountResponse struct {
}

func (m *UnlinkChainAccountResponse) Reset()         { *m = UnlinkChainAccountResponse{} }
func (m *UnlinkChainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*UnlinkChainAccountResponse) ProtoMessage()    {}
func (*UnlinkChainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9916e62e3c000d, []int{3}
}
func (m *UnlinkChainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnlinkChainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnlinkChainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnlinkChainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlinkChainAccountResponse.Merge(m, src)
}
func (m *UnlinkChainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *UnlinkChainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlinkChainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnlinkChainAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLinkChainAccount)(nil), "desmos.profiles.v1beta1.MsgLinkChainAccount")
	proto.RegisterType((*LinkChainAccountResponse)(nil), "desmos.profiles.v1beta1.LinkChainAccountResponse")
	proto.RegisterType((*MsgUnlinkChainAccount)(nil), "desmos.profiles.v1beta1.MsgUnlinkChainAccount")
	proto.RegisterType((*UnlinkChainAccountResponse)(nil), "desmos.profiles.v1beta1.UnlinkChainAccountResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/msgs_chain_links.proto", fileDescriptor_da9916e62e3c000d)
}

var fileDescriptor_da9916e62e3c000d = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0xe3, 0x42, 0x2b, 0x65, 0xd2, 0x4a, 0xd4, 0x6d, 0x44, 0x30, 0xc8, 0xae, 0x46, 0x08,
	0x85, 0x45, 0xc7, 0x14, 0x58, 0xb1, 0x8b, 0xcb, 0x82, 0x05, 0x45, 0x60, 0x89, 0x0d, 0x9b, 0x68,
	0x3c, 0x99, 0x4c, 0xad, 0xd8, 0xf3, 0x2c, 0x8f, 0x03, 0xe4, 0x16, 0x1c, 0x81, 0x43, 0x70, 0x88,
	0x8a, 0x55, 0x97, 0xac, 0x2c, 0x94, 0x88, 0x0b, 0xf8, 0x04, 0xc8, 0x33, 0xb6, 0x42, 0x4a, 0xb3,
	0x7b, 0x7f, 0xbe, 0xf7, 0x9b, 0x6f, 0xde, 0x0c, 0x22, 0x13, 0xae, 0x52, 0x50, 0x7e, 0x96, 0xc3,
	0x34, 0x4e, 0xb8, 0xf2, 0x3f, 0x9f, 0x45, 0xbc, 0xa0, 0x67, 0x7e, 0xaa, 0x84, 0x1a, 0xb3, 0x4b,
	0x1a, 0xcb, 0x71, 0x12, 0xcb, 0x99, 0x22, 0x59, 0x0e, 0x05, 0xd8, 0xf7, 0x8d, 0x9e, 0xb4, 0x7a,
	0xd2, 0xe8, 0x9d, 0x63, 0x01, 0x02, 0xb4, 0xc6, 0xaf, 0x23, 0x23, 0x77, 0x1e, 0x08, 0x00, 0x91,
	0x70, 0x5f, 0x67, 0xd1, 0x7c, 0xea, 0x53, 0xb9, 0x68, 0x5b, 0x0c, 0x6a, 0xd2, 0xd8, 0xcc, 0x98,
	0xa4, 0x69, 0x3d, 0xdb, 0x6a, 0x0a, 0x26, 0x3c, 0xb9, 0xc5, 0x16, 0xfe, 0xb3, 0x83, 0x8e, 0x2e,
	0x94, 0x78, 0x1b, 0xcb, 0xd9, 0x79, 0xdd, 0x1c, 0x31, 0x06, 0x73, 0x59, 0xd8, 0x0c, 0x1d, 0x18,
	0x31, 0x9d, 0x4c, 0x72, 0xae, 0xd4, 0xc0, 0x3a, 0xb1, 0x86, 0xbd, 0xe7, 0xc7, 0xc4, 0xf8, 0x22,
	0xad, 0x2f, 0x32, 0x92, 0x8b, 0x60, 0x58, 0x95, 0x5e, 0x7f, 0x41, 0xd3, 0xe4, 0x15, 0x56, 0x30,
	0xcf, 0x19, 0x6f, 0xa7, 0xf0, 0xcf, 0x1f, 0xa7, 0xbd, 0x91, 0x89, 0x5f, 0xd3, 0x82, 0x86, 0xfb,
	0x1a, 0xda, 0x54, 0xec, 0x0f, 0x68, 0x37, 0xcb, 0x01, 0xa6, 0x83, 0x1d, 0x0d, 0x77, 0xc9, 0x96,
	0x1d, 0x91, 0xf7, 0xb5, 0x2a, 0x78, 0x78, 0x55, 0x7a, 0x9d, 0xaa, 0xf4, 0x8e, 0x36, 0x8e, 0xd2,
	0x04, 0x1c, 0x1a, 0x92, 0x3d, 0x43, 0xe6, 0x88, 0x31, 0x03, 0x39, 0x8d, 0xc5, 0xe0, 0x8e, 0x26,
	0x3f, 0xde, 0x4a, 0xd6, 0x97, 0x3e, 0xd7, 0xda, 0x00, 0x37, 0x7c, 0x67, 0x83, 0xff, 0x2f, 0x0e,
	0x87, 0x3d, 0xb6, 0x1e, 0xb0, 0x9f, 0xa2, 0x3d, 0x15, 0x0b, 0xc9, 0xf3, 0xc1, 0xdd, 0x13, 0x6b,
	0xd8, 0x0d, 0x0e, 0xab, 0xd2, 0x3b, 0x68, 0x86, 0x75, 0x1d, 0x87, 0x8d, 0x00, 0x3b, 0x68, 0x70,
	0x73, 0xc7, 0x21, 0x57, 0x19, 0x48, 0xc5, 0xf1, 0x77, 0x0b, 0xf5, 0x2f, 0x94, 0xf8, 0x28, 0x93,
	0x9b, 0xaf, 0xf0, 0x04, 0xed, 0xc2, 0x97, 0x9a, 0x6f, 0x69, 0xfe, 0xbd, 0xaa, 0xf4, 0xf6, 0x0d,
	0x5f, 0x97, 0x71, 0x68, 0xda, 0xf6, 0x4b, 0x84, 0x8c, 0x4d, 0x49, 0x53, 0xae, 0xb7, 0xd9, 0x0d,
	0xfa, 0x55, 0xe9, 0x1d, 0x1a, 0xf1, 0xba, 0x87, 0xc3, 0xae, 0x4e, 0xde, 0xd1, 0x94, 0xd7, 0xf6,
	0x0b, 0x9a, 0x0b, 0x5e, 0xe8, 0x2d, 0x6d, 0xd8, 0x37, 0x75, 0x1c, 0x36, 0x02, 0xfc, 0x08, 0x39,
	0xff, 0xdb, 0x6b, 0x2f, 0x10, 0xbc, 0xb9, 0x5a, 0xba, 0xd6, 0xf5, 0xd2, 0xb5, 0x7e, 0x2f, 0x5d,
	0xeb, 0xdb, 0xca, 0xed, 0x5c, 0xaf, 0xdc, 0xce, 0xaf, 0x95, 0xdb, 0xf9, 0x44, 0x44, 0x5c, 0x5c,
	0xce, 0x23, 0xc2, 0x20, 0xf5, 0xcd, 0x13, 0x9c, 0x26, 0x34, 0x52, 0x4d, 0xec, 0x7f, 0x5d, 0xff,
	0xd4, 0x62, 0x91, 0x71, 0x15, 0xed, 0xe9, 0x7f, 0xf5, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x09, 0x7f, 0xa7, 0xb4, 0x5e, 0x03, 0x00, 0x00,
}

func (m *MsgLinkChainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkChainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkChainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ChainAddress != nil {
		{
			size, err := m.ChainAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LinkChainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkChainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkChainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkChainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnlinkChainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnlinkChainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnlinkChainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsChainLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsChainLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLinkChainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainAddress != nil {
		l = m.ChainAddress.Size()
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = m.Proof.Size()
	n += 1 + l + sovMsgsChainLinks(uint64(l))
	l = m.ChainConfig.Size()
	n += 1 + l + sovMsgsChainLinks(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *LinkChainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnlinkChainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *UnlinkChainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsChainLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsChainLinks(x uint64) (n int) {
	return sovMsgsChainLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLinkChainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgLinkChainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkChainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainAddress == nil {
				m.ChainAddress = &types.Any{}
			}
			if err := m.ChainAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LinkChainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LinkChainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnlinkChainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnlinkChainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkChainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnlinkChainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnlinkChainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnlinkChainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgsChainLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsgsChainLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsChainLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsChainLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsChainLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsChainLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsChainLinks = fmt.Errorf("proto: unexpected end of group")
)