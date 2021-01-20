// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1beta1/polls.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   string `protobuf:"bytes,1,opt,name=answer_id,json=answerId,proto3" json:"answer_id" yaml:"id"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text" yaml:"text"`
}

func (m *PollAnswer) Reset()         { *m = PollAnswer{} }
func (m *PollAnswer) String() string { return proto.CompactTextString(m) }
func (*PollAnswer) ProtoMessage()    {}
func (*PollAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{0}
}
func (m *PollAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollAnswer.Merge(m, src)
}
func (m *PollAnswer) XXX_Size() int {
	return m.Size()
}
func (m *PollAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_PollAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_PollAnswer proto.InternalMessageInfo

func (m *PollAnswer) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PollAnswer) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// PollAnswer contains the data of a single poll answer inserted by the creator
// inside a PollData object
type PollData struct {
	Question              string       `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	ProvidedAnswers       []PollAnswer `protobuf:"bytes,2,rep,name=provided_answers,json=providedAnswers,proto3" json:"provided_answers" yaml:"provided_answers"`
	EndDate               time.Time    `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3,stdtime" json:"end_date" yaml:"end_date"`
	AllowsMultipleAnswers bool         `protobuf:"varint,4,opt,name=allows_multiple_answers,json=allowsMultipleAnswers,proto3" json:"allows_multiple_answers" yaml:"allows_multiple_answers"`
	AllowsAnswerEdits     bool         `protobuf:"varint,5,opt,name=allows_answer_edits,json=allowsAnswerEdits,proto3" json:"allows_answer_edits" yaml:"allows_answer_edits"`
}

func (m *PollData) Reset()         { *m = PollData{} }
func (m *PollData) String() string { return proto.CompactTextString(m) }
func (*PollData) ProtoMessage()    {}
func (*PollData) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{1}
}
func (m *PollData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollData.Merge(m, src)
}
func (m *PollData) XXX_Size() int {
	return m.Size()
}
func (m *PollData) XXX_DiscardUnknown() {
	xxx_messageInfo_PollData.DiscardUnknown(m)
}

var xxx_messageInfo_PollData proto.InternalMessageInfo

func (m *PollData) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *PollData) GetProvidedAnswers() []PollAnswer {
	if m != nil {
		return m.ProvidedAnswers
	}
	return nil
}

func (m *PollData) GetEndDate() time.Time {
	if m != nil {
		return m.EndDate
	}
	return time.Time{}
}

func (m *PollData) GetAllowsMultipleAnswers() bool {
	if m != nil {
		return m.AllowsMultipleAnswers
	}
	return false
}

func (m *PollData) GetAllowsAnswerEdits() bool {
	if m != nil {
		return m.AllowsAnswerEdits
	}
	return false
}

// UserAnswer contains the data of a user's answer to a poll
type UserAnswer struct {
	User    string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Answers []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (m *UserAnswer) Reset()         { *m = UserAnswer{} }
func (m *UserAnswer) String() string { return proto.CompactTextString(m) }
func (*UserAnswer) ProtoMessage()    {}
func (*UserAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{2}
}
func (m *UserAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAnswer.Merge(m, src)
}
func (m *UserAnswer) XXX_Size() int {
	return m.Size()
}
func (m *UserAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_UserAnswer proto.InternalMessageInfo

func (m *UserAnswer) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserAnswer) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

// UserAnswers wraps a list of UserAnswer
type UserAnswers struct {
	Answers []UserAnswer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers"`
}

func (m *UserAnswers) Reset()         { *m = UserAnswers{} }
func (m *UserAnswers) String() string { return proto.CompactTextString(m) }
func (*UserAnswers) ProtoMessage()    {}
func (*UserAnswers) Descriptor() ([]byte, []int) {
	return fileDescriptor_397fcd2757705694, []int{3}
}
func (m *UserAnswers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserAnswers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserAnswers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserAnswers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAnswers.Merge(m, src)
}
func (m *UserAnswers) XXX_Size() int {
	return m.Size()
}
func (m *UserAnswers) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAnswers.DiscardUnknown(m)
}

var xxx_messageInfo_UserAnswers proto.InternalMessageInfo

func (m *UserAnswers) GetAnswers() []UserAnswer {
	if m != nil {
		return m.Answers
	}
	return nil
}

func init() {
	proto.RegisterType((*PollAnswer)(nil), "desmos.posts.v1beta1.PollAnswer")
	proto.RegisterType((*PollData)(nil), "desmos.posts.v1beta1.PollData")
	proto.RegisterType((*UserAnswer)(nil), "desmos.posts.v1beta1.UserAnswer")
	proto.RegisterType((*UserAnswers)(nil), "desmos.posts.v1beta1.UserAnswers")
}

func init() { proto.RegisterFile("desmos/posts/v1beta1/polls.proto", fileDescriptor_397fcd2757705694) }

var fileDescriptor_397fcd2757705694 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xbd, 0x6e, 0xdb, 0x3c,
	0x14, 0x35, 0x6d, 0x7f, 0x5f, 0x64, 0x7a, 0x70, 0xab, 0xa6, 0xb0, 0xe0, 0x41, 0x14, 0xb4, 0xd4,
	0x40, 0x50, 0x09, 0x71, 0xd1, 0x25, 0x40, 0x80, 0x56, 0x70, 0x86, 0x0c, 0x45, 0x03, 0xa1, 0x5d,
	0x3a, 0xd4, 0xa0, 0x4b, 0xd6, 0x15, 0x40, 0x99, 0xaa, 0x49, 0xe5, 0x67, 0x2a, 0xd0, 0x27, 0xc8,
	0xd8, 0x31, 0x8f, 0x93, 0x31, 0x63, 0x27, 0xb6, 0xb0, 0x97, 0xc2, 0x4b, 0x01, 0x3d, 0x41, 0x21,
	0x52, 0xfe, 0x49, 0x91, 0x74, 0xbb, 0xf7, 0xdc, 0x73, 0xcf, 0x21, 0xef, 0x25, 0xa1, 0x47, 0xa8,
	0x48, 0xb9, 0x08, 0x33, 0x2e, 0xa4, 0x08, 0x4f, 0xf7, 0xc7, 0x54, 0xe2, 0xfd, 0x30, 0xe3, 0x8c,
	0x89, 0x20, 0x9b, 0x71, 0xc9, 0xed, 0x5d, 0xc3, 0x08, 0x34, 0x23, 0xa8, 0x18, 0xbd, 0xdd, 0x09,
	0x9f, 0x70, 0x4d, 0x08, 0xcb, 0xc8, 0x70, 0x7b, 0x68, 0xc2, 0xf9, 0x84, 0xd1, 0x50, 0x67, 0xe3,
	0xfc, 0x63, 0x28, 0x93, 0x94, 0x0a, 0x89, 0xd3, 0xcc, 0x10, 0xfc, 0xaf, 0x00, 0xc2, 0x13, 0xce,
	0xd8, 0xcb, 0xa9, 0x38, 0xa3, 0x33, 0xfb, 0x10, 0xb6, 0xb0, 0x8e, 0x46, 0x09, 0x71, 0x80, 0x07,
	0xfa, 0xad, 0xc8, 0x9b, 0x2b, 0x54, 0x3f, 0x1e, 0x2e, 0x15, 0xda, 0x94, 0x0a, 0x85, 0x5a, 0x17,
	0x38, 0x65, 0x07, 0x7e, 0x42, 0xfc, 0xd8, 0x32, 0xf8, 0x31, 0xb1, 0xf7, 0x60, 0x53, 0xd2, 0x73,
	0xe9, 0xd4, 0x75, 0x67, 0x77, 0xa9, 0x90, 0xce, 0x0b, 0x85, 0xda, 0x86, 0x5e, 0x66, 0x7e, 0xac,
	0xc1, 0x03, 0xeb, 0xdb, 0x15, 0x02, 0xbf, 0xae, 0x10, 0xf0, 0x7f, 0x37, 0xa0, 0x55, 0x1e, 0x62,
	0x88, 0x25, 0xb6, 0x7b, 0xd0, 0xfa, 0x9c, 0x53, 0x21, 0x13, 0x3e, 0x35, 0x27, 0x88, 0xd7, 0xb9,
	0xfd, 0x05, 0x3e, 0xc8, 0x66, 0xfc, 0x34, 0x21, 0x94, 0x8c, 0x8c, 0xa9, 0x70, 0xea, 0x5e, 0xa3,
	0xdf, 0x1e, 0x78, 0xc1, 0x5d, 0x53, 0x09, 0x36, 0x57, 0x8b, 0x06, 0xd7, 0x0a, 0xd5, 0xe6, 0x0a,
	0x75, 0x4e, 0x2a, 0x05, 0x83, 0x8b, 0x42, 0xa1, 0xae, 0x39, 0xdc, 0xdf, 0xd2, 0x7e, 0xdc, 0xc9,
	0x6e, 0x73, 0xed, 0xf7, 0xd0, 0xa2, 0x53, 0x32, 0x22, 0x58, 0x52, 0xa7, 0xe1, 0x81, 0x7e, 0x7b,
	0xd0, 0x0b, 0xcc, 0x88, 0x83, 0xd5, 0x88, 0x83, 0x37, 0xab, 0x11, 0x47, 0x4f, 0x4a, 0xcb, 0xa5,
	0x42, 0xeb, 0x9e, 0x42, 0xa1, 0x8e, 0xf1, 0x5a, 0x21, 0xfe, 0xe5, 0x0f, 0x04, 0xe2, 0x1d, 0x3a,
	0x25, 0x43, 0x2c, 0xa9, 0x9d, 0xc3, 0x2e, 0x66, 0x8c, 0x9f, 0x89, 0x51, 0x9a, 0x33, 0x99, 0x64,
	0x8c, 0xae, 0xef, 0xd9, 0xf4, 0x40, 0xdf, 0x8a, 0x0e, 0x97, 0x0a, 0xdd, 0x47, 0x29, 0x14, 0x72,
	0x8d, 0xfa, 0x3d, 0x04, 0x3f, 0x7e, 0x6c, 0x2a, 0xaf, 0xaa, 0xc2, 0xea, 0x5a, 0x14, 0x3e, 0xaa,
	0x5a, 0xaa, 0x15, 0x53, 0x92, 0x48, 0xe1, 0xfc, 0xa7, 0x2d, 0x9f, 0x2f, 0x15, 0xba, 0xab, 0x5c,
	0x28, 0xd4, 0xbb, 0x65, 0xb7, 0x5d, 0xf4, 0xe3, 0x87, 0x06, 0x35, 0x16, 0x47, 0x25, 0xb6, 0xb5,
	0xf1, 0x21, 0x84, 0x6f, 0x05, 0x9d, 0x55, 0xaf, 0xce, 0x86, 0xcd, 0x5c, 0xd0, 0x59, 0xb5, 0x6e,
	0x1d, 0xdb, 0x0e, 0xdc, 0xd9, 0xde, 0x70, 0x2b, 0x5e, 0xa5, 0x5b, 0x2a, 0xaf, 0x61, 0x7b, 0xa3,
	0x22, 0xec, 0x17, 0x9b, 0x16, 0xf0, 0xaf, 0x47, 0xb1, 0xe9, 0x89, 0x9a, 0xe5, 0x86, 0xd6, 0xd2,
	0xd1, 0xd1, 0xf5, 0xdc, 0x05, 0x37, 0x73, 0x17, 0xfc, 0x9c, 0xbb, 0xe0, 0x72, 0xe1, 0xd6, 0x6e,
	0x16, 0x6e, 0xed, 0xfb, 0xc2, 0xad, 0xbd, 0xdb, 0x9b, 0x24, 0xf2, 0x53, 0x3e, 0x0e, 0x3e, 0xf0,
	0x34, 0x34, 0xa2, 0x4f, 0x19, 0x1e, 0x8b, 0x2a, 0x0e, 0xcf, 0xab, 0xff, 0x2a, 0x2f, 0x32, 0x2a,
	0xc6, 0xff, 0xeb, 0xb7, 0xf0, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xdc, 0x18, 0xa0,
	0xcc, 0x03, 0x00, 0x00,
}

func (this *PollAnswer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PollAnswer)
	if !ok {
		that2, ok := that.(PollAnswer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.Text != that1.Text {
		return false
	}
	return true
}
func (this *PollData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PollData)
	if !ok {
		that2, ok := that.(PollData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Question != that1.Question {
		return false
	}
	if len(this.ProvidedAnswers) != len(that1.ProvidedAnswers) {
		return false
	}
	for i := range this.ProvidedAnswers {
		if !this.ProvidedAnswers[i].Equal(&that1.ProvidedAnswers[i]) {
			return false
		}
	}
	if !this.EndDate.Equal(that1.EndDate) {
		return false
	}
	if this.AllowsMultipleAnswers != that1.AllowsMultipleAnswers {
		return false
	}
	if this.AllowsAnswerEdits != that1.AllowsAnswerEdits {
		return false
	}
	return true
}
func (this *UserAnswer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserAnswer)
	if !ok {
		that2, ok := that.(UserAnswer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.User != that1.User {
		return false
	}
	if len(this.Answers) != len(that1.Answers) {
		return false
	}
	for i := range this.Answers {
		if this.Answers[i] != that1.Answers[i] {
			return false
		}
	}
	return true
}
func (m *PollAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PollData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowsAnswerEdits {
		i--
		if m.AllowsAnswerEdits {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.AllowsMultipleAnswers {
		i--
		if m.AllowsMultipleAnswers {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndDate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPolls(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.ProvidedAnswers) > 0 {
		for iNdEx := len(m.ProvidedAnswers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProvidedAnswers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolls(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Answers) > 0 {
		for iNdEx := len(m.Answers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Answers[iNdEx])
			copy(dAtA[i:], m.Answers[iNdEx])
			i = encodeVarintPolls(dAtA, i, uint64(len(m.Answers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintPolls(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserAnswers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserAnswers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserAnswers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Answers) > 0 {
		for iNdEx := len(m.Answers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Answers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolls(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolls(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolls(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PollAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	return n
}

func (m *PollData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	if len(m.ProvidedAnswers) > 0 {
		for _, e := range m.ProvidedAnswers {
			l = e.Size()
			n += 1 + l + sovPolls(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndDate)
	n += 1 + l + sovPolls(uint64(l))
	if m.AllowsMultipleAnswers {
		n += 2
	}
	if m.AllowsAnswerEdits {
		n += 2
	}
	return n
}

func (m *UserAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovPolls(uint64(l))
	}
	if len(m.Answers) > 0 {
		for _, s := range m.Answers {
			l = len(s)
			n += 1 + l + sovPolls(uint64(l))
		}
	}
	return n
}

func (m *UserAnswers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Answers) > 0 {
		for _, e := range m.Answers {
			l = e.Size()
			n += 1 + l + sovPolls(uint64(l))
		}
	}
	return n
}

func sovPolls(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolls(x uint64) (n int) {
	return sovPolls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PollAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: PollAnswer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollAnswer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func (m *PollData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: PollData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvidedAnswers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvidedAnswers = append(m.ProvidedAnswers, PollAnswer{})
			if err := m.ProvidedAnswers[len(m.ProvidedAnswers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowsMultipleAnswers", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowsMultipleAnswers = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowsAnswerEdits", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowsAnswerEdits = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func (m *UserAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: UserAnswer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAnswer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answers = append(m.Answers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func (m *UserAnswers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolls
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
			return fmt.Errorf("proto: UserAnswers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserAnswers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolls
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
				return ErrInvalidLengthPolls
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answers = append(m.Answers, UserAnswer{})
			if err := m.Answers[len(m.Answers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPolls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPolls
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
func skipPolls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolls
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
					return 0, ErrIntOverflowPolls
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
					return 0, ErrIntOverflowPolls
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
				return 0, ErrInvalidLengthPolls
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolls
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolls
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolls        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolls          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolls = fmt.Errorf("proto: unexpected end of group")
)
