// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/committee/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState defines the committee module's genesis state.
type GenesisState struct {
	VotingStartHeight  uint64      `protobuf:"varint,1,opt,name=voting_start_height,json=votingStartHeight,proto3" json:"voting_start_height,omitempty"`
	VotingPeriod       uint64      `protobuf:"varint,2,opt,name=voting_period,json=votingPeriod,proto3" json:"voting_period,omitempty"`
	CurrentCommitteeID uint64      `protobuf:"varint,3,opt,name=current_committee_id,json=currentCommitteeId,proto3" json:"current_committee_id,omitempty"`
	Committees         []Committee `protobuf:"bytes,4,rep,name=committees,proto3" json:"committees"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ecc9b7bccadc11, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

type Committee struct {
	ID                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VotingStartHeight uint64 `protobuf:"varint,2,opt,name=voting_start_height,json=votingStartHeight,proto3" json:"voting_start_height,omitempty"`
	StartHeight       uint64 `protobuf:"varint,3,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight         uint64 `protobuf:"varint,4,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	Votes             []Vote `protobuf:"bytes,5,rep,name=votes,proto3" json:"votes"`
}

func (m *Committee) Reset()         { *m = Committee{} }
func (m *Committee) String() string { return proto.CompactTextString(m) }
func (*Committee) ProtoMessage()    {}
func (*Committee) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ecc9b7bccadc11, []int{1}
}
func (m *Committee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Committee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Committee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Committee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Committee.Merge(m, src)
}
func (m *Committee) XXX_Size() int {
	return m.Size()
}
func (m *Committee) XXX_DiscardUnknown() {
	xxx_messageInfo_Committee.DiscardUnknown(m)
}

var xxx_messageInfo_Committee proto.InternalMessageInfo

func (m *Committee) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Committee) GetVotingStartHeight() uint64 {
	if m != nil {
		return m.VotingStartHeight
	}
	return 0
}

func (m *Committee) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *Committee) GetEndHeight() uint64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

func (m *Committee) GetVotes() []Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type Vote struct {
	CommitteeID uint64                                        `protobuf:"varint,1,opt,name=committee_id,json=committeeId,proto3" json:"committee_id,omitempty"`
	Voter       github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"voter,omitempty"`
	Ballots     []*Ballot                                     `protobuf:"bytes,3,rep,name=ballots,proto3" json:"ballots,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ecc9b7bccadc11, []int{2}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

type Ballot struct {
	ID      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Ballot) Reset()         { *m = Ballot{} }
func (m *Ballot) String() string { return proto.CompactTextString(m) }
func (*Ballot) ProtoMessage()    {}
func (*Ballot) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ecc9b7bccadc11, []int{3}
}
func (m *Ballot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ballot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ballot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ballot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ballot.Merge(m, src)
}
func (m *Ballot) XXX_Size() int {
	return m.Size()
}
func (m *Ballot) XXX_DiscardUnknown() {
	xxx_messageInfo_Ballot.DiscardUnknown(m)
}

var xxx_messageInfo_Ballot proto.InternalMessageInfo

func (m *Ballot) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Ballot) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "evmos.committee.v1.GenesisState")
	proto.RegisterType((*Committee)(nil), "evmos.committee.v1.Committee")
	proto.RegisterType((*Vote)(nil), "evmos.committee.v1.Vote")
	proto.RegisterType((*Ballot)(nil), "evmos.committee.v1.Ballot")
}

func init() { proto.RegisterFile("evmos/committee/v1/genesis.proto", fileDescriptor_57ecc9b7bccadc11) }

var fileDescriptor_57ecc9b7bccadc11 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0x9b, 0x30,
	0x1c, 0xc6, 0x43, 0x4a, 0x53, 0xd5, 0x61, 0x9a, 0xe6, 0x56, 0x15, 0x8d, 0x54, 0xc8, 0xb2, 0x4b,
	0x2f, 0x01, 0xa5, 0xab, 0x76, 0xe8, 0x6d, 0xa4, 0xd2, 0x92, 0xdb, 0x44, 0xa5, 0x1e, 0x76, 0x18,
	0x22, 0xd8, 0x23, 0xd6, 0x02, 0x8e, 0xb0, 0x83, 0x96, 0x37, 0xd8, 0x71, 0x8f, 0xb0, 0x87, 0xd8,
	0x43, 0x54, 0xda, 0xa5, 0xda, 0x61, 0xda, 0x09, 0x4d, 0xe4, 0x25, 0xa6, 0x9d, 0x26, 0x6c, 0x40,
	0xa4, 0x6b, 0x2e, 0x80, 0xbf, 0xef, 0x67, 0xec, 0xef, 0x6f, 0xff, 0x41, 0x1f, 0xa7, 0x11, 0x65,
	0x76, 0x40, 0xa3, 0x88, 0x70, 0x8e, 0xb1, 0x9d, 0x8e, 0xec, 0x10, 0xc7, 0x98, 0x11, 0x66, 0x2d,
	0x13, 0xca, 0x29, 0x84, 0x82, 0xb0, 0x6a, 0xc2, 0x4a, 0x47, 0xbd, 0xd3, 0x80, 0xb2, 0x88, 0x32,
	0x4f, 0x10, 0xb6, 0x1c, 0x48, 0xbc, 0x77, 0x1c, 0xd2, 0x90, 0x4a, 0xbd, 0xf8, 0x2a, 0xd5, 0xd3,
	0x90, 0xd2, 0x70, 0x81, 0x6d, 0x31, 0x9a, 0xad, 0x3e, 0xd8, 0x7e, 0xbc, 0x2e, 0x2d, 0xf3, 0xa1,
	0xc5, 0x49, 0x84, 0x19, 0xf7, 0xa3, 0xa5, 0x04, 0x06, 0x7f, 0x14, 0xa0, 0xbd, 0x91, 0x5b, 0xba,
	0xe1, 0x3e, 0xc7, 0xd0, 0x02, 0x47, 0x29, 0xe5, 0x24, 0x0e, 0x3d, 0xc6, 0xfd, 0x84, 0x7b, 0x73,
	0x4c, 0xc2, 0x39, 0xd7, 0x95, 0xbe, 0x72, 0xae, 0xba, 0xcf, 0xa4, 0x75, 0x53, 0x38, 0x13, 0x61,
	0xc0, 0x17, 0xe0, 0x49, 0xc9, 0x2f, 0x71, 0x42, 0x28, 0xd2, 0xdb, 0x82, 0xd4, 0xa4, 0xf8, 0x56,
	0x68, 0x70, 0x02, 0x8e, 0x83, 0x55, 0x92, 0xe0, 0x98, 0x7b, 0x75, 0x54, 0x8f, 0x20, 0x7d, 0xaf,
	0x60, 0x9d, 0x93, 0x3c, 0x33, 0xe1, 0x58, 0xfa, 0xe3, 0xca, 0x9e, 0x5e, 0xbb, 0x30, 0x78, 0xa8,
	0x21, 0x38, 0x06, 0xa0, 0xfe, 0x03, 0xd3, 0xd5, 0xfe, 0xde, 0x79, 0xf7, 0xe2, 0xcc, 0xfa, 0xbf,
	0x8a, 0x56, 0x3d, 0xc9, 0x51, 0xef, 0x32, 0xb3, 0xe5, 0x36, 0xa6, 0x5d, 0xa9, 0x9f, 0xbf, 0x9a,
	0xad, 0xc1, 0x77, 0x05, 0x1c, 0xd6, 0x14, 0x3c, 0x01, 0x6d, 0x82, 0x64, 0x4c, 0xa7, 0x93, 0x67,
	0x66, 0x7b, 0x7a, 0xed, 0xb6, 0x09, 0xda, 0x55, 0x8f, 0xf6, 0xae, 0x7a, 0x3c, 0x07, 0xda, 0x16,
	0x28, 0x22, 0xba, 0x5d, 0xd6, 0x40, 0xce, 0x00, 0xc0, 0x31, 0xaa, 0x00, 0x55, 0x00, 0x87, 0x38,
	0x46, 0xa5, 0x7d, 0x09, 0xf6, 0x53, 0xca, 0x31, 0xd3, 0xf7, 0x45, 0x3a, 0xfd, 0xb1, 0x74, 0xb7,
	0x94, 0x57, 0xc1, 0x24, 0x3c, 0xf8, 0xa9, 0x00, 0xb5, 0x50, 0xe1, 0x05, 0xd0, 0xb6, 0x6a, 0x2c,
	0x23, 0x3d, 0xcd, 0x33, 0xb3, 0xdb, 0x2c, 0x6e, 0x37, 0x68, 0x54, 0xf5, 0xbd, 0x5c, 0x32, 0x11,
	0xb1, 0x34, 0x67, 0xf2, 0x37, 0x33, 0x87, 0x21, 0xe1, 0xf3, 0xd5, 0xac, 0x58, 0xb6, 0xbc, 0x83,
	0xe5, 0x6b, 0xc8, 0xd0, 0x47, 0x9b, 0xaf, 0x97, 0x98, 0x59, 0xb7, 0xfe, 0xe2, 0x35, 0x42, 0x09,
	0x66, 0xec, 0xc7, 0xb7, 0xe1, 0x51, 0x79, 0x53, 0x4b, 0xc5, 0x59, 0x73, 0xcc, 0xe4, 0xe6, 0x12,
	0x78, 0x09, 0x0e, 0x66, 0xfe, 0x62, 0x41, 0x39, 0xd3, 0xf7, 0x44, 0xa8, 0xde, 0x63, 0xa1, 0x1c,
	0x81, 0xb8, 0x15, 0x5a, 0x1e, 0xd3, 0x15, 0xe8, 0x48, 0x63, 0xe7, 0x11, 0xe9, 0xe0, 0x20, 0xa0,
	0x31, 0xc7, 0xb1, 0x3c, 0x16, 0xcd, 0xad, 0x86, 0xce, 0xf4, 0x2e, 0x37, 0x94, 0xfb, 0xdc, 0x50,
	0x7e, 0xe7, 0x86, 0xf2, 0x65, 0x63, 0xb4, 0xee, 0x37, 0x46, 0xeb, 0xd7, 0xc6, 0x68, 0xbd, 0xb3,
	0x1b, 0xf1, 0x64, 0x97, 0xca, 0x67, 0x3a, 0x7a, 0x65, 0x7f, 0xda, 0xee, 0x58, 0x91, 0x75, 0xd6,
	0x11, 0xfd, 0xf2, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x41, 0x1a, 0x92, 0xd4, 0x03,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Committees) > 0 {
		for iNdEx := len(m.Committees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Committees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.CurrentCommitteeID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CurrentCommitteeID))
		i--
		dAtA[i] = 0x18
	}
	if m.VotingPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VotingPeriod))
		i--
		dAtA[i] = 0x10
	}
	if m.VotingStartHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VotingStartHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Committee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Committee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Committee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.EndHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EndHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.StartHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.VotingStartHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.VotingStartHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ballots) > 0 {
		for iNdEx := len(m.Ballots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Ballots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.CommitteeID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CommitteeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Ballot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ballot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ballot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VotingStartHeight != 0 {
		n += 1 + sovGenesis(uint64(m.VotingStartHeight))
	}
	if m.VotingPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.VotingPeriod))
	}
	if m.CurrentCommitteeID != 0 {
		n += 1 + sovGenesis(uint64(m.CurrentCommitteeID))
	}
	if len(m.Committees) > 0 {
		for _, e := range m.Committees {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Committee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovGenesis(uint64(m.ID))
	}
	if m.VotingStartHeight != 0 {
		n += 1 + sovGenesis(uint64(m.VotingStartHeight))
	}
	if m.StartHeight != 0 {
		n += 1 + sovGenesis(uint64(m.StartHeight))
	}
	if m.EndHeight != 0 {
		n += 1 + sovGenesis(uint64(m.EndHeight))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommitteeID != 0 {
		n += 1 + sovGenesis(uint64(m.CommitteeID))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Ballots) > 0 {
		for _, e := range m.Ballots {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Ballot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovGenesis(uint64(m.ID))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingStartHeight", wireType)
			}
			m.VotingStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingStartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPeriod", wireType)
			}
			m.VotingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentCommitteeID", wireType)
			}
			m.CurrentCommitteeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentCommitteeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Committees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Committees = append(m.Committees, Committee{})
			if err := m.Committees[len(m.Committees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Committee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Committee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Committee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingStartHeight", wireType)
			}
			m.VotingStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingStartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
			}
			m.EndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitteeID", wireType)
			}
			m.CommitteeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitteeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ballots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ballots = append(m.Ballots, &Ballot{})
			if err := m.Ballots[len(m.Ballots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Ballot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Ballot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ballot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
