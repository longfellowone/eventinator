// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// represents a message with generic protobuf payload
// this is an internal type used for storing messages
// i.e. messages on the backend are stored in this form
type Event struct {
	// unique UUID
	// e.g. "123e4567-e89b-12d3-a456-426655440000"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// version
	// e.g. "1.0"
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// context of the message; either system or tenant
	//
	// Types that are valid to be assigned to Context:
	//	*Event_SystemContext
	//	*Event_TenantContext
	Context isEvent_Context `protobuf_oneof:"context"`
	// source/origin urn
	// e.g. "urn:system-x.org.com/service-a"
	Source string `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	// protocol buffer for the message payload
	// the "any" message embeds the message type via `type_url` and the data as byte[]
	Payload *any.Any `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	// correlation id for distributed tracing and tracking
	CorrelationId string `protobuf:"bytes,7,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// metadata for message, used for communicating contextual information
	MetaData map[string]string `protobuf:"bytes,8,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the MIME type of the data
	//
	// this will be derived from the provided payload, taking the form
	//
	//    'application/x-protobuf; messageType="x.y.Z"'
	//
	// where the messageType is the protobuf message type (corresponds with Any#type_url)
	//
	// see https://tools.ietf.org/html/draft-rfernando-protocol-buffers-00
	// https://www.charlesproxy.com/documentation/using-charles/protocol-buffers/
	// and https://prometheus.io/docs/instrumenting/exposition_formats/
	//
	ContentType string `protobuf:"bytes,9,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// UTC time the message was created
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_41f4a519b878ee3b, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type isEvent_Context interface {
	isEvent_Context()
}

type Event_SystemContext struct {
	SystemContext *SystemContext `protobuf:"bytes,3,opt,name=system_context,json=systemContext,proto3,oneof"`
}

type Event_TenantContext struct {
	TenantContext *TenantContext `protobuf:"bytes,4,opt,name=tenant_context,json=tenantContext,proto3,oneof"`
}

func (*Event_SystemContext) isEvent_Context() {}

func (*Event_TenantContext) isEvent_Context() {}

func (m *Event) GetContext() isEvent_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Event) GetSystemContext() *SystemContext {
	if x, ok := m.GetContext().(*Event_SystemContext); ok {
		return x.SystemContext
	}
	return nil
}

func (m *Event) GetTenantContext() *TenantContext {
	if x, ok := m.GetContext().(*Event_TenantContext); ok {
		return x.TenantContext
	}
	return nil
}

func (m *Event) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Event) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Event) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *Event) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Event) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_SystemContext)(nil),
		(*Event_TenantContext)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// context
	switch x := m.Context.(type) {
	case *Event_SystemContext:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SystemContext); err != nil {
			return err
		}
	case *Event_TenantContext:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TenantContext); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Context has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 3: // context.system_context
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SystemContext)
		err := b.DecodeMessage(msg)
		m.Context = &Event_SystemContext{msg}
		return true, err
	case 4: // context.tenant_context
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TenantContext)
		err := b.DecodeMessage(msg)
		m.Context = &Event_TenantContext{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// context
	switch x := m.Context.(type) {
	case *Event_SystemContext:
		s := proto.Size(x.SystemContext)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_TenantContext:
		s := proto.Size(x.TenantContext)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Event)(nil), "eventinator.protobuf.Event")
	proto.RegisterMapType((map[string]string)(nil), "eventinator.protobuf.Event.MetaDataEntry")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_41f4a519b878ee3b) }

var fileDescriptor_41f4a519b878ee3b = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x69, 0xbb, 0xfd, 0x93, 0x29, 0x8d, 0x90, 0x55, 0x21, 0x93, 0x0b, 0x05, 0x84, 0x54,
	0x2e, 0x59, 0x69, 0xb9, 0xf0, 0xe7, 0xb4, 0x0b, 0x8b, 0x40, 0x82, 0x4b, 0xe8, 0x89, 0x4b, 0x34,
	0x9b, 0x0c, 0x2b, 0x8b, 0xc4, 0x8e, 0x9c, 0x69, 0x85, 0xbf, 0x12, 0x9f, 0x12, 0xd5, 0x71, 0x20,
	0x40, 0xc5, 0x6d, 0xe6, 0xf9, 0xe9, 0x37, 0x33, 0x7e, 0x10, 0x2b, 0xcd, 0x64, 0x35, 0x56, 0x69,
	0x63, 0x0d, 0x1b, 0xb1, 0xa6, 0x03, 0x69, 0x56, 0x1a, 0xd9, 0xd8, 0x4e, 0xba, 0xd9, 0x7f, 0x4d,
	0x1e, 0xdc, 0x1a, 0x73, 0x5b, 0xd1, 0x79, 0x2f, 0x9c, 0xa3, 0x76, 0xdd, 0x6b, 0xf2, 0xf0, 0xef,
	0x27, 0x56, 0x35, 0xb5, 0x8c, 0x75, 0x13, 0x0c, 0x11, 0x36, 0xaa, 0x2b, 0x1f, 0xff, 0x38, 0x83,
	0xe9, 0xf5, 0x91, 0x2f, 0x62, 0x18, 0xab, 0x52, 0x8e, 0x36, 0xa3, 0x6d, 0x94, 0x8d, 0x55, 0x29,
	0x24, 0xcc, 0x0f, 0x64, 0x5b, 0x65, 0xb4, 0x1c, 0x7b, 0xb1, 0x6f, 0xc5, 0x47, 0x88, 0x5b, 0xd7,
	0x32, 0xd5, 0x79, 0x61, 0x34, 0xd3, 0x77, 0x96, 0x93, 0xcd, 0x68, 0xbb, 0xbc, 0x78, 0x92, 0x9e,
	0xda, 0x34, 0xfd, 0xec, 0xbd, 0x6f, 0x3a, 0xeb, 0xfb, 0x3b, 0xd9, 0xaa, 0x1d, 0x0a, 0x47, 0x1a,
	0x93, 0x46, 0xcd, 0xbf, 0x68, 0x67, 0xff, 0xa3, 0xed, 0xbc, 0x77, 0x40, 0xe3, 0xa1, 0x20, 0xee,
	0xc3, 0xac, 0x35, 0x7b, 0x5b, 0x90, 0x9c, 0xfa, 0xa5, 0x43, 0x27, 0x52, 0x98, 0x37, 0xe8, 0x2a,
	0x83, 0xa5, 0x9c, 0x79, 0xfc, 0x3a, 0xed, 0x7e, 0xe9, 0x37, 0xf9, 0x52, 0xbb, 0xac, 0x37, 0x89,
	0xa7, 0x10, 0x17, 0xc6, 0x5a, 0xaa, 0x90, 0x95, 0xd1, 0xb9, 0x2a, 0xe5, 0xdc, 0xf3, 0x56, 0x03,
	0xf5, 0x43, 0x29, 0xde, 0x41, 0x54, 0x13, 0x63, 0x5e, 0x22, 0xa3, 0x5c, 0x6c, 0x26, 0xdb, 0xe5,
	0xc5, 0xb3, 0xd3, 0x7b, 0xfb, 0x4f, 0x4e, 0x3f, 0x11, 0xe3, 0x5b, 0x64, 0xbc, 0xd6, 0x6c, 0x5d,
	0xb6, 0xa8, 0x43, 0x2b, 0x1e, 0xc1, 0x5d, 0x7f, 0xbd, 0xe6, 0x9c, 0x5d, 0x43, 0x32, 0xf2, 0xc3,
	0x96, 0x41, 0xdb, 0xb9, 0x86, 0xc4, 0x4b, 0x80, 0xc2, 0x12, 0x32, 0x95, 0x39, 0xb2, 0x04, 0x7f,
	0x44, 0xf2, 0xcf, 0x11, 0xbb, 0x3e, 0xea, 0x2c, 0x0a, 0xee, 0x4b, 0x4e, 0x5e, 0xc3, 0xea, 0x8f,
	0xc1, 0xe2, 0x1e, 0x4c, 0xbe, 0x91, 0x0b, 0x61, 0x1f, 0x4b, 0xb1, 0x86, 0xe9, 0x01, 0xab, 0x3d,
	0x85, 0xac, 0xbb, 0xe6, 0xd5, 0xf8, 0xc5, 0xe8, 0x2a, 0x82, 0x79, 0x08, 0xe6, 0x0a, 0xbe, 0x2c,
	0xfa, 0x41, 0x37, 0x33, 0x5f, 0x3d, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x96, 0xf6, 0x08, 0x9b,
	0xae, 0x02, 0x00, 0x00,
}