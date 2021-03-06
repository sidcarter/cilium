// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accesslog.proto

/*
Package envoy is a generated protocol buffer package.

It is generated from these files:
	accesslog.proto

It has these top-level messages:
	KeyValue
	HttpLogEntry
*/
package envoy

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

type Protocol int32

const (
	Protocol_HTTP10 Protocol = 0
	Protocol_HTTP11 Protocol = 1
	Protocol_HTTP2  Protocol = 2
)

var Protocol_name = map[int32]string{
	0: "HTTP10",
	1: "HTTP11",
	2: "HTTP2",
}
var Protocol_value = map[string]int32{
	"HTTP10": 0,
	"HTTP11": 1,
	"HTTP2":  2,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EntryType int32

const (
	EntryType_Request  EntryType = 0
	EntryType_Response EntryType = 1
	EntryType_Denied   EntryType = 2
)

var EntryType_name = map[int32]string{
	0: "Request",
	1: "Response",
	2: "Denied",
}
var EntryType_value = map[string]int32{
	"Request":  0,
	"Response": 1,
	"Denied":   2,
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}
func (EntryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HttpLogEntry struct {
	// The time that Cilium filter captured this log entry,
	// in, nanoseconds since 1/1/1970.
	Timestamp    uint64    `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	HttpProtocol Protocol  `protobuf:"varint,2,opt,name=http_protocol,json=httpProtocol,enum=pb.cilium.Protocol" json:"http_protocol,omitempty"`
	EntryType    EntryType `protobuf:"varint,3,opt,name=entry_type,json=entryType,enum=pb.cilium.EntryType" json:"entry_type,omitempty"`
	// Cilium Redirect resource name
	CiliumResourceName string `protobuf:"bytes,4,opt,name=cilium_resource_name,json=ciliumResourceName" json:"cilium_resource_name,omitempty"`
	// Cilium rule reference
	CiliumRuleRef string `protobuf:"bytes,5,opt,name=cilium_rule_ref,json=ciliumRuleRef" json:"cilium_rule_ref,omitempty"`
	// Cilium security ID of the source
	SourceSecurityId uint32 `protobuf:"varint,6,opt,name=source_security_id,json=sourceSecurityId" json:"source_security_id,omitempty"`
	// These fields record the original source and destination addresses,
	// stored in ipv4:port or [ipv6]:port format.
	SourceAddress      string `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	DestinationAddress string `protobuf:"bytes,8,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,9,opt,name=scheme" json:"scheme,omitempty"`
	Host   string `protobuf:"bytes,10,opt,name=host" json:"host,omitempty"`
	Path   string `protobuf:"bytes,11,opt,name=path" json:"path,omitempty"`
	Method string `protobuf:"bytes,12,opt,name=method" json:"method,omitempty"`
	// Response info
	Status uint32 `protobuf:"varint,13,opt,name=status" json:"status,omitempty"`
	// Request headers not included above
	Headers []*KeyValue `protobuf:"bytes,14,rep,name=headers" json:"headers,omitempty"`
}

func (m *HttpLogEntry) Reset()                    { *m = HttpLogEntry{} }
func (m *HttpLogEntry) String() string            { return proto.CompactTextString(m) }
func (*HttpLogEntry) ProtoMessage()               {}
func (*HttpLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HttpLogEntry) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HttpLogEntry) GetHttpProtocol() Protocol {
	if m != nil {
		return m.HttpProtocol
	}
	return Protocol_HTTP10
}

func (m *HttpLogEntry) GetEntryType() EntryType {
	if m != nil {
		return m.EntryType
	}
	return EntryType_Request
}

func (m *HttpLogEntry) GetCiliumResourceName() string {
	if m != nil {
		return m.CiliumResourceName
	}
	return ""
}

func (m *HttpLogEntry) GetCiliumRuleRef() string {
	if m != nil {
		return m.CiliumRuleRef
	}
	return ""
}

func (m *HttpLogEntry) GetSourceSecurityId() uint32 {
	if m != nil {
		return m.SourceSecurityId
	}
	return 0
}

func (m *HttpLogEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *HttpLogEntry) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *HttpLogEntry) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *HttpLogEntry) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HttpLogEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HttpLogEntry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpLogEntry) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpLogEntry) GetHeaders() []*KeyValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyValue)(nil), "pb.cilium.KeyValue")
	proto.RegisterType((*HttpLogEntry)(nil), "pb.cilium.HttpLogEntry")
	proto.RegisterEnum("pb.cilium.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("pb.cilium.EntryType", EntryType_name, EntryType_value)
}

func init() { proto.RegisterFile("accesslog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcd, 0x6b, 0xd4, 0x40,
	0x14, 0xef, 0x7e, 0x6f, 0xde, 0x7e, 0x34, 0xbc, 0x2e, 0x32, 0x07, 0x0f, 0x4b, 0x41, 0x59, 0x8a,
	0x5d, 0xdb, 0xed, 0xc5, 0xab, 0xa2, 0x50, 0x51, 0xa4, 0x8c, 0x8b, 0x07, 0x2f, 0x61, 0x9a, 0xbc,
	0x36, 0xc1, 0x24, 0x13, 0x33, 0x2f, 0x85, 0xfc, 0x2b, 0xfe, 0xb5, 0x92, 0x99, 0x24, 0xd6, 0xdb,
	0xef, 0xf3, 0xcd, 0x0c, 0x6f, 0xe0, 0x54, 0x85, 0x21, 0x19, 0x93, 0xea, 0xc7, 0x7d, 0x51, 0x6a,
	0xd6, 0xe8, 0x15, 0xf7, 0xfb, 0x30, 0x49, 0x93, 0x2a, 0x3b, 0x3f, 0xc0, 0xfc, 0x0b, 0xd5, 0x3f,
	0x54, 0x5a, 0x11, 0xfa, 0x30, 0xfa, 0x45, 0xb5, 0x18, 0x6c, 0x07, 0x3b, 0x4f, 0x36, 0x10, 0x37,
	0x30, 0x79, 0x6a, 0x2c, 0x31, 0xb4, 0x9a, 0x23, 0xe7, 0x7f, 0xc6, 0xb0, 0xbc, 0x65, 0x2e, 0xbe,
	0xea, 0xc7, 0x4f, 0x39, 0x97, 0x35, 0xbe, 0x04, 0x8f, 0x93, 0x8c, 0x0c, 0xab, 0xac, 0xb0, 0xf5,
	0xb1, 0xfc, 0x27, 0xe0, 0x3b, 0x58, 0xc5, 0xcc, 0x45, 0x60, 0xcf, 0x0e, 0x75, 0x6a, 0x87, 0xad,
	0x0f, 0x67, 0xfb, 0xfe, 0x16, 0xfb, 0xbb, 0xd6, 0x92, 0xcb, 0x26, 0xd9, 0x31, 0xbc, 0x01, 0xa0,
	0xe6, 0x80, 0x80, 0xeb, 0x82, 0xc4, 0xc8, 0xd6, 0x36, 0xcf, 0x6a, 0xf6, 0xf4, 0x63, 0x5d, 0x90,
	0xf4, 0xa8, 0x83, 0x78, 0x05, 0x1b, 0x67, 0x07, 0x25, 0x19, 0x5d, 0x95, 0x21, 0x05, 0xb9, 0xca,
	0x48, 0x8c, 0xed, 0x13, 0xd0, 0x79, 0xb2, 0xb5, 0xbe, 0xa9, 0x8c, 0xf0, 0x35, 0x9c, 0x76, 0x8d,
	0x2a, 0xa5, 0xa0, 0xa4, 0x07, 0x31, 0xb1, 0xe1, 0x55, 0x1b, 0xae, 0x52, 0x92, 0xf4, 0x80, 0x6f,
	0x00, 0xdb, 0x81, 0x86, 0xc2, 0xaa, 0x4c, 0xb8, 0x0e, 0x92, 0x48, 0x4c, 0xb7, 0x83, 0xdd, 0x4a,
	0xfa, 0xce, 0xf9, 0xde, 0x1a, 0x9f, 0x23, 0x7c, 0x05, 0xeb, 0x36, 0xad, 0xa2, 0xa8, 0x24, 0x63,
	0xc4, 0xcc, 0x0d, 0x75, 0xea, 0x7b, 0x27, 0xe2, 0x5b, 0x38, 0x8b, 0xc8, 0x70, 0x92, 0x2b, 0x4e,
	0x74, 0xde, 0x67, 0xe7, 0xee, 0xb6, 0xcf, 0xac, 0xae, 0xf0, 0x02, 0xa6, 0x26, 0x8c, 0x29, 0x23,
	0xe1, 0xd9, 0x4c, 0xcb, 0x10, 0x61, 0x1c, 0x6b, 0xc3, 0x02, 0xac, 0x6a, 0x71, 0xa3, 0x15, 0x8a,
	0x63, 0xb1, 0x70, 0x5a, 0x83, 0x9b, 0x7e, 0x46, 0x1c, 0xeb, 0x48, 0x2c, 0x5d, 0xdf, 0x31, 0x3b,
	0x97, 0x15, 0x57, 0x46, 0xac, 0xec, 0x8b, 0x5a, 0x86, 0x97, 0x30, 0x8b, 0x49, 0x45, 0x54, 0x1a,
	0xb1, 0xde, 0x8e, 0x76, 0x8b, 0xff, 0x16, 0xd7, 0xfd, 0x1d, 0xd9, 0x65, 0x2e, 0x2e, 0x61, 0xde,
	0xef, 0x0f, 0x60, 0x7a, 0x7b, 0x3c, 0xde, 0x5d, 0x5f, 0xf9, 0x27, 0x3d, 0xbe, 0xf6, 0x07, 0xe8,
	0xc1, 0xa4, 0xc1, 0x07, 0x7f, 0x78, 0x71, 0x00, 0xaf, 0xdf, 0x22, 0x2e, 0x60, 0x26, 0xe9, 0x77,
	0x45, 0x86, 0xfd, 0x13, 0x5c, 0xc2, 0x5c, 0x92, 0x29, 0x74, 0x6e, 0xc8, 0x1f, 0x34, 0xf5, 0x8f,
	0x94, 0x27, 0x14, 0xf9, 0xc3, 0x0f, 0xb3, 0x9f, 0x13, 0xca, 0x9f, 0x74, 0x7d, 0x3f, 0xb5, 0x5f,
	0xea, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xbd, 0xf0, 0x13, 0xe1, 0x02, 0x00, 0x00,
}
