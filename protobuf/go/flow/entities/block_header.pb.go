// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/block_header.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type BlockHeader struct {
	Id                   []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             []byte                 `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Height               uint64                 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PayloadHash          []byte                 `protobuf:"bytes,5,opt,name=payload_hash,json=payloadHash,proto3" json:"payload_hash,omitempty"`
	View                 uint64                 `protobuf:"varint,6,opt,name=view,proto3" json:"view,omitempty"`
	ParentVoterIds       [][]byte               `protobuf:"bytes,7,rep,name=parent_voter_ids,json=parentVoterIds,proto3" json:"parent_voter_ids,omitempty"`
	ParentVoterSigData   []byte                 `protobuf:"bytes,8,opt,name=parent_voter_sig_data,json=parentVoterSigData,proto3" json:"parent_voter_sig_data,omitempty"`
	ProposerId           []byte                 `protobuf:"bytes,9,opt,name=proposer_id,json=proposerId,proto3" json:"proposer_id,omitempty"`
	ProposerSigData      []byte                 `protobuf:"bytes,10,opt,name=proposer_sig_data,json=proposerSigData,proto3" json:"proposer_sig_data,omitempty"`
	ChainId              string                 `protobuf:"bytes,11,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ParentVoterIndices   []byte                 `protobuf:"bytes,12,opt,name=parent_voter_indices,json=parentVoterIndices,proto3" json:"parent_voter_indices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9d8363da0276a74, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BlockHeader) GetParentId() []byte {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BlockHeader) GetPayloadHash() []byte {
	if m != nil {
		return m.PayloadHash
	}
	return nil
}

func (m *BlockHeader) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *BlockHeader) GetParentVoterIds() [][]byte {
	if m != nil {
		return m.ParentVoterIds
	}
	return nil
}

func (m *BlockHeader) GetParentVoterSigData() []byte {
	if m != nil {
		return m.ParentVoterSigData
	}
	return nil
}

func (m *BlockHeader) GetProposerId() []byte {
	if m != nil {
		return m.ProposerId
	}
	return nil
}

func (m *BlockHeader) GetProposerSigData() []byte {
	if m != nil {
		return m.ProposerSigData
	}
	return nil
}

func (m *BlockHeader) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *BlockHeader) GetParentVoterIndices() []byte {
	if m != nil {
		return m.ParentVoterIndices
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "flow.entities.BlockHeader")
}

func init() { proto.RegisterFile("flow/entities/block_header.proto", fileDescriptor_b9d8363da0276a74) }

var fileDescriptor_b9d8363da0276a74 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6f, 0x9b, 0x30,
	0x18, 0xc5, 0x45, 0x92, 0x25, 0xe1, 0x83, 0x65, 0x9b, 0xb5, 0x4d, 0x5e, 0x36, 0x29, 0x6c, 0x27,
	0xb4, 0x03, 0x64, 0xdb, 0x65, 0xe7, 0xa8, 0x87, 0x70, 0xab, 0x68, 0xd5, 0x43, 0x2f, 0xc8, 0x60,
	0xc7, 0xb6, 0x4a, 0x30, 0xc2, 0x4e, 0xa2, 0xfe, 0x6b, 0xfd, 0xeb, 0x2a, 0x4c, 0xa0, 0x8d, 0x7a,
	0x41, 0xf8, 0x7d, 0x3f, 0xbd, 0xf7, 0xc0, 0x1f, 0x04, 0xbb, 0x52, 0x9d, 0x62, 0x56, 0x19, 0x69,
	0x24, 0xd3, 0x71, 0x5e, 0xaa, 0xe2, 0x21, 0x13, 0x8c, 0x50, 0xd6, 0x44, 0x75, 0xa3, 0x8c, 0x42,
	0xef, 0x5b, 0x22, 0xea, 0x89, 0xe5, 0x8a, 0x2b, 0xc5, 0x4b, 0x16, 0xdb, 0x61, 0x7e, 0xd8, 0xc5,
	0x46, 0xee, 0x99, 0x36, 0x64, 0x5f, 0x77, 0xfc, 0xaf, 0xa7, 0x31, 0x78, 0x9b, 0xd6, 0x66, 0x6b,
	0x5d, 0xd0, 0x02, 0x46, 0x92, 0x62, 0x27, 0x70, 0x42, 0x3f, 0x1d, 0x49, 0x8a, 0xbe, 0x83, 0x5b,
	0x93, 0x86, 0x55, 0x26, 0x93, 0x14, 0x8f, 0xac, 0x3c, 0xef, 0x84, 0x84, 0xa2, 0xaf, 0x30, 0x15,
	0x4c, 0x72, 0x61, 0xf0, 0x38, 0x70, 0xc2, 0x49, 0x7a, 0x3e, 0xa1, 0xff, 0xe0, 0x0e, 0x39, 0x78,
	0x12, 0x38, 0xa1, 0xf7, 0x77, 0x19, 0x75, 0x4d, 0xa2, 0xbe, 0x49, 0x74, 0xdb, 0x13, 0xe9, 0x0b,
	0x8c, 0x7e, 0x82, 0x5f, 0x93, 0xc7, 0x52, 0x11, 0x9a, 0x09, 0xa2, 0x05, 0x7e, 0x67, 0x13, 0xbd,
	0xb3, 0xb6, 0x25, 0x5a, 0x20, 0x04, 0x93, 0xa3, 0x64, 0x27, 0x3c, 0xb5, 0x91, 0xf6, 0x1d, 0x85,
	0xf0, 0xf1, 0xdc, 0xf2, 0xa8, 0x0c, 0x6b, 0x32, 0x49, 0x35, 0x9e, 0x05, 0xe3, 0xd0, 0x4f, 0x17,
	0x9d, 0x7e, 0xd7, 0xca, 0x09, 0xd5, 0xe8, 0x0f, 0x7c, 0xb9, 0x20, 0xb5, 0xe4, 0x19, 0x25, 0x86,
	0xe0, 0xb9, 0x4d, 0x42, 0xaf, 0xf0, 0x1b, 0xc9, 0xaf, 0x88, 0x21, 0x68, 0x05, 0x5e, 0xdd, 0xa8,
	0x5a, 0x69, 0x6b, 0x8c, 0x5d, 0x0b, 0x42, 0x2f, 0x25, 0x14, 0xfd, 0x86, 0x4f, 0x03, 0x30, 0xf8,
	0x81, 0xc5, 0x3e, 0xf4, 0x83, 0xde, 0xec, 0x1b, 0xcc, 0x0b, 0x41, 0x64, 0xd5, 0x3a, 0x79, 0x81,
	0x13, 0xba, 0xe9, 0xcc, 0x9e, 0x13, 0x8a, 0xd6, 0xf0, 0xf9, 0xf2, 0x23, 0x2a, 0x2a, 0x0b, 0xa6,
	0xb1, 0xff, 0xa6, 0x59, 0xd2, 0x4d, 0x36, 0xd7, 0xf0, 0x43, 0x35, 0x3c, 0x52, 0x95, 0xbd, 0xf4,
	0xe1, 0xcf, 0xf6, 0xb7, 0x7f, 0xbf, 0xe6, 0xd2, 0x88, 0x43, 0x1e, 0x15, 0x6a, 0x1f, 0x77, 0x50,
	0x6c, 0x1f, 0xc3, 0x36, 0x70, 0x15, 0x5f, 0x6c, 0x54, 0x3e, 0xb5, 0xa3, 0x7f, 0xcf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x43, 0x25, 0xe0, 0xc8, 0x69, 0x02, 0x00, 0x00,
}
