// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bottos-project/core/common/types/block.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Block struct {
	Header       *Header        `protobuf:"bytes,1,opt,name=header" json:"header" `
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Block) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Header struct {
	Version         uint32   `protobuf:"varint,1,opt,name=version" json:"version"`
	PrevBlockHash   []byte   `protobuf:"bytes,2,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash"`
	Number          uint32   `protobuf:"varint,3,opt,name=number" json:"number"`
	Timestamp       uint64   `protobuf:"varint,4,opt,name=timestamp" json:"timestamp"`
	MerkleRoot      []byte   `protobuf:"bytes,5,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root"`
	Delegate        []byte   `protobuf:"bytes,6,opt,name=delegate,proto3" json:"delegate"`
	DelegateSign    []byte   `protobuf:"bytes,7,opt,name=delegate_sign,json=delegateSign,proto3" json:"delegate_sign"`
	DelegateChanges []string `protobuf:"bytes,8,rep,name=delegate_changes,json=delegateChanges" json:"delegate_changes"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Header) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *Header) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Header) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Header) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *Header) GetDelegate() []byte {
	if m != nil {
		return m.Delegate
	}
	return nil
}

func (m *Header) GetDelegateSign() []byte {
	if m != nil {
		return m.DelegateSign
	}
	return nil
}

func (m *Header) GetDelegateChanges() []string {
	if m != nil {
		return m.DelegateChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*Header)(nil), "types.Header")
}

func init() {
	proto.RegisterFile("github.com/bottos-project/core/common/types/block.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x33, 0x7e, 0x0c, 0x28, 0x2c, 0x98, 0x1e, 0x4c, 0x43, 0x4c, 0x5c, 0x30, 0x9a, 0x19,
	0xc3, 0x96, 0x60, 0xa2, 0x27, 0x2f, 0x78, 0xe1, 0x3c, 0x3d, 0x79, 0x59, 0xba, 0xf1, 0xdc, 0x26,
	0xb4, 0x6f, 0x69, 0x0b, 0x89, 0x7f, 0x84, 0xff, 0xb3, 0xa1, 0x63, 0xa0, 0x47, 0x6e, 0x7d, 0x9f,
	0xf7, 0x7d, 0x9f, 0xb6, 0x79, 0xe4, 0x39, 0x2f, 0x4d, 0xb1, 0x4d, 0xc3, 0x0c, 0x45, 0x94, 0xa2,
	0x31, 0xa8, 0x67, 0x95, 0xc2, 0x2f, 0xc8, 0x4c, 0x94, 0xa1, 0x82, 0x28, 0x43, 0x21, 0x50, 0x46,
	0xe6, 0xbb, 0x02, 0x1d, 0xa5, 0x1b, 0xcc, 0xd6, 0x61, 0xa5, 0xd0, 0x20, 0xed, 0x5a, 0x34, 0x79,
	0x39, 0x67, 0xde, 0x28, 0x2e, 0x35, 0xcf, 0x4c, 0x89, 0xb2, 0xb6, 0x4c, 0x3f, 0x49, 0x77, 0xb1,
	0x97, 0xd2, 0x5b, 0xe2, 0x16, 0xc0, 0x57, 0xa0, 0x98, 0xe3, 0x3b, 0xc1, 0x70, 0xee, 0x85, 0x76,
	0x24, 0x5c, 0x5a, 0x18, 0x1f, 0x9a, 0xf4, 0x89, 0x8c, 0xfe, 0x48, 0x34, 0x6b, 0xf9, 0xed, 0x60,
	0x38, 0xa7, 0x87, 0xf0, 0xfb, 0xa9, 0x15, 0xff, 0xcb, 0x4d, 0x7f, 0x5a, 0xc4, 0xad, 0x55, 0x94,
	0x91, 0xde, 0x0e, 0x94, 0x2e, 0x51, 0xda, 0xab, 0xbc, 0xb8, 0x29, 0xe9, 0x1d, 0x19, 0x57, 0x0a,
	0x76, 0x89, 0xfd, 0x66, 0x52, 0x70, 0x5d, 0xb0, 0x96, 0xef, 0x04, 0xa3, 0xd8, 0xdb, 0x63, 0xfb,
	0xce, 0x25, 0xd7, 0x05, 0xbd, 0x24, 0xae, 0xdc, 0x8a, 0x14, 0x14, 0x6b, 0x5b, 0xc1, 0xa1, 0xa2,
	0x57, 0x64, 0x60, 0x4a, 0x01, 0xda, 0x70, 0x51, 0xb1, 0x8e, 0xef, 0x04, 0x9d, 0xf8, 0x04, 0xe8,
	0x35, 0x19, 0x0a, 0x50, 0xeb, 0x0d, 0x24, 0x0a, 0xd1, 0xb0, 0xae, 0x35, 0x93, 0x1a, 0xc5, 0x88,
	0x86, 0x4e, 0x48, 0x7f, 0x05, 0x1b, 0xc8, 0xb9, 0x01, 0xe6, 0xda, 0xee, 0xb1, 0xa6, 0x37, 0xc4,
	0x6b, 0xce, 0x89, 0x2e, 0x73, 0xc9, 0x7a, 0x36, 0x30, 0x6a, 0xe0, 0x5b, 0x99, 0x4b, 0x7a, 0x4f,
	0x2e, 0x8e, 0xa1, 0xac, 0xe0, 0x32, 0x07, 0xcd, 0xfa, 0x7e, 0x3b, 0x18, 0xc4, 0xe3, 0x86, 0xbf,
	0xd6, 0x78, 0x31, 0xfb, 0x78, 0x38, 0x63, 0x71, 0xa9, 0x6b, 0xb7, 0xf5, 0xf8, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x5e, 0x28, 0x61, 0x2e, 0x02, 0x00, 0x00,
}
