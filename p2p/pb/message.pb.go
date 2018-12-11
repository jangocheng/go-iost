// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p/pb/message.proto

package p2pb

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

type RoutingQuery struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutingQuery) Reset()         { *m = RoutingQuery{} }
func (m *RoutingQuery) String() string { return proto.CompactTextString(m) }
func (*RoutingQuery) ProtoMessage()    {}
func (*RoutingQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_05b036241fc763b9, []int{0}
}
func (m *RoutingQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingQuery.Unmarshal(m, b)
}
func (m *RoutingQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingQuery.Marshal(b, m, deterministic)
}
func (dst *RoutingQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingQuery.Merge(dst, src)
}
func (m *RoutingQuery) XXX_Size() int {
	return xxx_messageInfo_RoutingQuery.Size(m)
}
func (m *RoutingQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingQuery proto.InternalMessageInfo

func (m *RoutingQuery) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type PeerInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addrs                []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_05b036241fc763b9, []int{1}
}
func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerInfo.Unmarshal(m, b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
}
func (dst *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(dst, src)
}
func (m *PeerInfo) XXX_Size() int {
	return xxx_messageInfo_PeerInfo.Size(m)
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PeerInfo) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type RoutingResponse struct {
	Peers                []*PeerInfo `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoutingResponse) Reset()         { *m = RoutingResponse{} }
func (m *RoutingResponse) String() string { return proto.CompactTextString(m) }
func (*RoutingResponse) ProtoMessage()    {}
func (*RoutingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_05b036241fc763b9, []int{2}
}
func (m *RoutingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingResponse.Unmarshal(m, b)
}
func (m *RoutingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingResponse.Marshal(b, m, deterministic)
}
func (dst *RoutingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingResponse.Merge(dst, src)
}
func (m *RoutingResponse) XXX_Size() int {
	return xxx_messageInfo_RoutingResponse.Size(m)
}
func (m *RoutingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingResponse proto.InternalMessageInfo

func (m *RoutingResponse) GetPeers() []*PeerInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*RoutingQuery)(nil), "p2pb.RoutingQuery")
	proto.RegisterType((*PeerInfo)(nil), "p2pb.PeerInfo")
	proto.RegisterType((*RoutingResponse)(nil), "p2pb.RoutingResponse")
}

func init() { proto.RegisterFile("p2p/pb/message.proto", fileDescriptor_message_05b036241fc763b9) }

var fileDescriptor_message_05b036241fc763b9 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x31, 0x0f, 0x82, 0x30,
	0x14, 0x84, 0x03, 0x88, 0x91, 0xa7, 0x41, 0xd3, 0x30, 0x30, 0x12, 0xe2, 0xc0, 0x04, 0x06, 0x07,
	0x7f, 0x83, 0x9b, 0xf6, 0x1f, 0xd0, 0xf4, 0x49, 0x3a, 0xd8, 0xbe, 0xf4, 0xc1, 0xe0, 0xbf, 0x37,
	0x50, 0xdd, 0xee, 0x72, 0x97, 0xef, 0x0e, 0x0a, 0xea, 0xa9, 0x23, 0xd5, 0xbd, 0x91, 0x79, 0x18,
	0xb1, 0x25, 0xef, 0x26, 0x27, 0x36, 0xd4, 0x93, 0xaa, 0x2b, 0x38, 0x48, 0x37, 0x4f, 0xc6, 0x8e,
	0xcf, 0x19, 0xfd, 0x47, 0x9c, 0x20, 0x31, 0x9a, 0xcb, 0xa8, 0x4a, 0x9a, 0x4c, 0x2e, 0xb2, 0xbe,
	0xc0, 0xee, 0x81, 0xe8, 0xef, 0xf6, 0xe5, 0x44, 0x0e, 0xb1, 0xd1, 0x65, 0x54, 0x45, 0x4d, 0x26,
	0x63, 0xa3, 0x45, 0x01, 0xe9, 0xa0, 0xb5, 0xe7, 0x32, 0x5e, 0xfb, 0xc1, 0xd4, 0x37, 0x38, 0xfe,
	0x98, 0x12, 0x99, 0x9c, 0x65, 0x14, 0x67, 0x48, 0x09, 0xd1, 0x07, 0xf0, 0xbe, 0xcf, 0xdb, 0x65,
	0xbc, 0xfd, 0x73, 0x65, 0x08, 0xd5, 0x76, 0x7d, 0x76, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x1f, 0xbf, 0xe5, 0xb1, 0x00, 0x00, 0x00,
}
