// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crossconnect.proto

package crossconnect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import connection "github.com/ligato/networkservicemesh/controlplane/pkg/apis/local/connection"
import connection1 "github.com/ligato/networkservicemesh/controlplane/pkg/apis/remote/connection"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CrossConnectEventType int32

const (
	CrossConnectEventType_INITIAL_STATE_TRANSFER CrossConnectEventType = 0
	CrossConnectEventType_UPDATE                 CrossConnectEventType = 1
	CrossConnectEventType_DELETE                 CrossConnectEventType = 2
)

var CrossConnectEventType_name = map[int32]string{
	0: "INITIAL_STATE_TRANSFER",
	1: "UPDATE",
	2: "DELETE",
}
var CrossConnectEventType_value = map[string]int32{
	"INITIAL_STATE_TRANSFER": 0,
	"UPDATE":                 1,
	"DELETE":                 2,
}

func (x CrossConnectEventType) String() string {
	return proto.EnumName(CrossConnectEventType_name, int32(x))
}
func (CrossConnectEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_crossconnect_be9614887acd2496, []int{0}
}

type CrossConnectEvent struct {
	Type                 CrossConnectEventType    `protobuf:"varint,1,opt,name=type,proto3,enum=crossconnect.CrossConnectEventType" json:"type,omitempty"`
	CrossConnects        map[string]*CrossConnect `protobuf:"bytes,2,rep,name=cross_connects,json=crossConnects,proto3" json:"cross_connects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CrossConnectEvent) Reset()         { *m = CrossConnectEvent{} }
func (m *CrossConnectEvent) String() string { return proto.CompactTextString(m) }
func (*CrossConnectEvent) ProtoMessage()    {}
func (*CrossConnectEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_crossconnect_be9614887acd2496, []int{0}
}
func (m *CrossConnectEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossConnectEvent.Unmarshal(m, b)
}
func (m *CrossConnectEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossConnectEvent.Marshal(b, m, deterministic)
}
func (dst *CrossConnectEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossConnectEvent.Merge(dst, src)
}
func (m *CrossConnectEvent) XXX_Size() int {
	return xxx_messageInfo_CrossConnectEvent.Size(m)
}
func (m *CrossConnectEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossConnectEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CrossConnectEvent proto.InternalMessageInfo

func (m *CrossConnectEvent) GetType() CrossConnectEventType {
	if m != nil {
		return m.Type
	}
	return CrossConnectEventType_INITIAL_STATE_TRANSFER
}

func (m *CrossConnectEvent) GetCrossConnects() map[string]*CrossConnect {
	if m != nil {
		return m.CrossConnects
	}
	return nil
}

type CrossConnect struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Types that are valid to be assigned to Source:
	//	*CrossConnect_LocalSource
	//	*CrossConnect_RemoteSource
	Source isCrossConnect_Source `protobuf_oneof:"source"`
	// Types that are valid to be assigned to Destination:
	//	*CrossConnect_LocalDestination
	//	*CrossConnect_RemoteDestination
	Destination          isCrossConnect_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CrossConnect) Reset()         { *m = CrossConnect{} }
func (m *CrossConnect) String() string { return proto.CompactTextString(m) }
func (*CrossConnect) ProtoMessage()    {}
func (*CrossConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_crossconnect_be9614887acd2496, []int{1}
}
func (m *CrossConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrossConnect.Unmarshal(m, b)
}
func (m *CrossConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrossConnect.Marshal(b, m, deterministic)
}
func (dst *CrossConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossConnect.Merge(dst, src)
}
func (m *CrossConnect) XXX_Size() int {
	return xxx_messageInfo_CrossConnect.Size(m)
}
func (m *CrossConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossConnect.DiscardUnknown(m)
}

var xxx_messageInfo_CrossConnect proto.InternalMessageInfo

func (m *CrossConnect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CrossConnect) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type isCrossConnect_Source interface {
	isCrossConnect_Source()
}

type CrossConnect_LocalSource struct {
	LocalSource *connection.Connection `protobuf:"bytes,3,opt,name=local_source,json=localSource,proto3,oneof"`
}

type CrossConnect_RemoteSource struct {
	RemoteSource *connection1.Connection `protobuf:"bytes,4,opt,name=remote_source,json=remoteSource,proto3,oneof"`
}

func (*CrossConnect_LocalSource) isCrossConnect_Source() {}

func (*CrossConnect_RemoteSource) isCrossConnect_Source() {}

func (m *CrossConnect) GetSource() isCrossConnect_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *CrossConnect) GetLocalSource() *connection.Connection {
	if x, ok := m.GetSource().(*CrossConnect_LocalSource); ok {
		return x.LocalSource
	}
	return nil
}

func (m *CrossConnect) GetRemoteSource() *connection1.Connection {
	if x, ok := m.GetSource().(*CrossConnect_RemoteSource); ok {
		return x.RemoteSource
	}
	return nil
}

type isCrossConnect_Destination interface {
	isCrossConnect_Destination()
}

type CrossConnect_LocalDestination struct {
	LocalDestination *connection.Connection `protobuf:"bytes,5,opt,name=local_destination,json=localDestination,proto3,oneof"`
}

type CrossConnect_RemoteDestination struct {
	RemoteDestination *connection1.Connection `protobuf:"bytes,6,opt,name=remote_destination,json=remoteDestination,proto3,oneof"`
}

func (*CrossConnect_LocalDestination) isCrossConnect_Destination() {}

func (*CrossConnect_RemoteDestination) isCrossConnect_Destination() {}

func (m *CrossConnect) GetDestination() isCrossConnect_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *CrossConnect) GetLocalDestination() *connection.Connection {
	if x, ok := m.GetDestination().(*CrossConnect_LocalDestination); ok {
		return x.LocalDestination
	}
	return nil
}

func (m *CrossConnect) GetRemoteDestination() *connection1.Connection {
	if x, ok := m.GetDestination().(*CrossConnect_RemoteDestination); ok {
		return x.RemoteDestination
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CrossConnect) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CrossConnect_OneofMarshaler, _CrossConnect_OneofUnmarshaler, _CrossConnect_OneofSizer, []interface{}{
		(*CrossConnect_LocalSource)(nil),
		(*CrossConnect_RemoteSource)(nil),
		(*CrossConnect_LocalDestination)(nil),
		(*CrossConnect_RemoteDestination)(nil),
	}
}

func _CrossConnect_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CrossConnect)
	// source
	switch x := m.Source.(type) {
	case *CrossConnect_LocalSource:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LocalSource); err != nil {
			return err
		}
	case *CrossConnect_RemoteSource:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RemoteSource); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CrossConnect.Source has unexpected type %T", x)
	}
	// destination
	switch x := m.Destination.(type) {
	case *CrossConnect_LocalDestination:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LocalDestination); err != nil {
			return err
		}
	case *CrossConnect_RemoteDestination:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RemoteDestination); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CrossConnect.Destination has unexpected type %T", x)
	}
	return nil
}

func _CrossConnect_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CrossConnect)
	switch tag {
	case 3: // source.local_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(connection.Connection)
		err := b.DecodeMessage(msg)
		m.Source = &CrossConnect_LocalSource{msg}
		return true, err
	case 4: // source.remote_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(connection1.Connection)
		err := b.DecodeMessage(msg)
		m.Source = &CrossConnect_RemoteSource{msg}
		return true, err
	case 5: // destination.local_destination
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(connection.Connection)
		err := b.DecodeMessage(msg)
		m.Destination = &CrossConnect_LocalDestination{msg}
		return true, err
	case 6: // destination.remote_destination
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(connection1.Connection)
		err := b.DecodeMessage(msg)
		m.Destination = &CrossConnect_RemoteDestination{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CrossConnect_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CrossConnect)
	// source
	switch x := m.Source.(type) {
	case *CrossConnect_LocalSource:
		s := proto.Size(x.LocalSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CrossConnect_RemoteSource:
		s := proto.Size(x.RemoteSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// destination
	switch x := m.Destination.(type) {
	case *CrossConnect_LocalDestination:
		s := proto.Size(x.LocalDestination)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CrossConnect_RemoteDestination:
		s := proto.Size(x.RemoteDestination)
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
	proto.RegisterType((*CrossConnectEvent)(nil), "crossconnect.CrossConnectEvent")
	proto.RegisterMapType((map[string]*CrossConnect)(nil), "crossconnect.CrossConnectEvent.CrossConnectsEntry")
	proto.RegisterType((*CrossConnect)(nil), "crossconnect.CrossConnect")
	proto.RegisterEnum("crossconnect.CrossConnectEventType", CrossConnectEventType_name, CrossConnectEventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorCrossConnectClient is the client API for MonitorCrossConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorCrossConnectClient interface {
	MonitorCrossConnects(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MonitorCrossConnect_MonitorCrossConnectsClient, error)
}

type monitorCrossConnectClient struct {
	cc *grpc.ClientConn
}

func NewMonitorCrossConnectClient(cc *grpc.ClientConn) MonitorCrossConnectClient {
	return &monitorCrossConnectClient{cc}
}

func (c *monitorCrossConnectClient) MonitorCrossConnects(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (MonitorCrossConnect_MonitorCrossConnectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitorCrossConnect_serviceDesc.Streams[0], "/crossconnect.MonitorCrossConnect/MonitorCrossConnects", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorCrossConnectMonitorCrossConnectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitorCrossConnect_MonitorCrossConnectsClient interface {
	Recv() (*CrossConnectEvent, error)
	grpc.ClientStream
}

type monitorCrossConnectMonitorCrossConnectsClient struct {
	grpc.ClientStream
}

func (x *monitorCrossConnectMonitorCrossConnectsClient) Recv() (*CrossConnectEvent, error) {
	m := new(CrossConnectEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorCrossConnectServer is the server API for MonitorCrossConnect service.
type MonitorCrossConnectServer interface {
	MonitorCrossConnects(*empty.Empty, MonitorCrossConnect_MonitorCrossConnectsServer) error
}

func RegisterMonitorCrossConnectServer(s *grpc.Server, srv MonitorCrossConnectServer) {
	s.RegisterService(&_MonitorCrossConnect_serviceDesc, srv)
}

func _MonitorCrossConnect_MonitorCrossConnects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorCrossConnectServer).MonitorCrossConnects(m, &monitorCrossConnectMonitorCrossConnectsServer{stream})
}

type MonitorCrossConnect_MonitorCrossConnectsServer interface {
	Send(*CrossConnectEvent) error
	grpc.ServerStream
}

type monitorCrossConnectMonitorCrossConnectsServer struct {
	grpc.ServerStream
}

func (x *monitorCrossConnectMonitorCrossConnectsServer) Send(m *CrossConnectEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _MonitorCrossConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crossconnect.MonitorCrossConnect",
	HandlerType: (*MonitorCrossConnectServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorCrossConnects",
			Handler:       _MonitorCrossConnect_MonitorCrossConnects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crossconnect.proto",
}

func init() { proto.RegisterFile("crossconnect.proto", fileDescriptor_crossconnect_be9614887acd2496) }

var fileDescriptor_crossconnect_be9614887acd2496 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x74, 0x2b, 0xec, 0xf5, 0x87, 0x5a, 0x03, 0x53, 0x14, 0x81, 0xa8, 0xc6, 0xa5,
	0xe2, 0x10, 0x4f, 0xdd, 0x01, 0xc4, 0x2d, 0x6b, 0x83, 0xa8, 0x36, 0x2a, 0x70, 0xc3, 0x01, 0x69,
	0xa8, 0x4a, 0x53, 0x93, 0x46, 0x4d, 0xed, 0x28, 0x76, 0x8b, 0xf2, 0x07, 0xf0, 0x37, 0x73, 0x45,
	0xb1, 0x53, 0x48, 0xb5, 0x8d, 0x4d, 0xe2, 0x52, 0xbd, 0xbe, 0xf7, 0xfd, 0x7e, 0xde, 0x8b, 0x9f,
	0x0d, 0x28, 0xcc, 0xb8, 0x10, 0x21, 0x67, 0x8c, 0x86, 0xd2, 0x49, 0x33, 0x2e, 0x39, 0x6a, 0x56,
	0x73, 0xf6, 0x75, 0x14, 0xcb, 0xe5, 0x66, 0xee, 0x84, 0x7c, 0x8d, 0x93, 0x38, 0x0a, 0x24, 0xc7,
	0x8c, 0xca, 0x1f, 0x3c, 0x5b, 0x09, 0x9a, 0x6d, 0xe3, 0x90, 0xae, 0xa9, 0x58, 0xe2, 0x90, 0x33,
	0x99, 0xf1, 0x24, 0x4d, 0x02, 0x46, 0x71, 0xba, 0x8a, 0x70, 0x90, 0xc6, 0x02, 0x27, 0x3c, 0x0c,
	0x12, 0x5c, 0x92, 0x62, 0xce, 0x2a, 0xa1, 0xee, 0x65, 0x7f, 0xfb, 0x0f, 0x7a, 0x46, 0xd7, 0x5c,
	0xd2, 0x7f, 0xe2, 0xcf, 0x2b, 0xf8, 0x88, 0x27, 0x01, 0x8b, 0xb0, 0x2a, 0xcc, 0x37, 0xdf, 0x71,
	0x2a, 0xf3, 0x94, 0x0a, 0x4c, 0xd7, 0xa9, 0xcc, 0xf5, 0xaf, 0x36, 0x9d, 0xfe, 0x34, 0xa1, 0x3b,
	0x2c, 0x8e, 0x60, 0xa8, 0x71, 0xde, 0x96, 0x32, 0x89, 0xde, 0xc0, 0x61, 0x61, 0xb0, 0x8c, 0x9e,
	0xd1, 0x6f, 0x0f, 0x5e, 0x39, 0x7b, 0x07, 0x77, 0x43, 0xee, 0xe7, 0x29, 0x25, 0xca, 0x80, 0xbe,
	0x42, 0x5b, 0x69, 0x67, 0xa5, 0x58, 0x58, 0x66, 0xaf, 0xd6, 0x6f, 0x0c, 0x06, 0xf7, 0x20, 0xf6,
	0x32, 0xc2, 0x63, 0x32, 0xcb, 0x49, 0x2b, 0xac, 0xe6, 0xec, 0x6b, 0x40, 0x37, 0x45, 0xa8, 0x03,
	0xb5, 0x15, 0xcd, 0xd5, 0xa0, 0xc7, 0xa4, 0x08, 0xd1, 0x19, 0x1c, 0x6d, 0x83, 0x64, 0x43, 0x2d,
	0xb3, 0x67, 0xf4, 0x1b, 0x03, 0xfb, 0xee, 0xce, 0x44, 0x0b, 0xdf, 0x99, 0x6f, 0x8d, 0xd3, 0x5f,
	0x26, 0x34, 0xab, 0x35, 0xd4, 0x06, 0x33, 0x5e, 0x94, 0x5c, 0x33, 0x5e, 0x20, 0x0b, 0x1e, 0xa5,
	0x41, 0x9e, 0xf0, 0x60, 0xa1, 0xc0, 0xc7, 0x64, 0xf7, 0x17, 0xb9, 0xd0, 0x54, 0xbb, 0x9f, 0x09,
	0xbe, 0xc9, 0x42, 0x6a, 0xd5, 0x54, 0xdf, 0xe7, 0x8e, 0x4a, 0x3a, 0x95, 0x35, 0x0d, 0xff, 0x84,
	0x1f, 0x0e, 0x48, 0x43, 0x95, 0xa7, 0xca, 0x82, 0x46, 0xd0, 0xd2, 0x0b, 0xde, 0x31, 0x0e, 0x15,
	0xe3, 0x85, 0xa3, 0xb3, 0x77, 0x42, 0x9a, 0xba, 0x5e, 0x52, 0x2e, 0xa1, 0xab, 0x07, 0x59, 0x50,
	0x21, 0x63, 0x16, 0x14, 0x22, 0xeb, 0xe8, 0x01, 0xd3, 0x18, 0xa4, 0xa3, 0xca, 0xa3, 0xbf, 0x3e,
	0x34, 0x01, 0x54, 0x8e, 0x54, 0xa5, 0xd5, 0x1f, 0x32, 0x97, 0x41, 0xba, 0xba, 0x5e, 0xe1, 0x5d,
	0x3c, 0x86, 0xba, 0xfe, 0xb6, 0x8b, 0x16, 0x34, 0x2a, 0xc8, 0xd7, 0x97, 0xf0, 0xec, 0xd6, 0x1b,
	0x85, 0x6c, 0x38, 0x19, 0x4f, 0xc6, 0xfe, 0xd8, 0xbd, 0x9a, 0x4d, 0x7d, 0xd7, 0xf7, 0x66, 0x3e,
	0x71, 0x27, 0xd3, 0xf7, 0x1e, 0xe9, 0x1c, 0x20, 0x80, 0xfa, 0x97, 0x4f, 0x23, 0xd7, 0xf7, 0x3a,
	0x46, 0x11, 0x8f, 0xbc, 0x2b, 0xcf, 0xf7, 0x3a, 0xe6, 0x60, 0x09, 0x4f, 0x3e, 0x72, 0x16, 0x4b,
	0x9e, 0xed, 0x2d, 0xf3, 0x33, 0x3c, 0xbd, 0x25, 0x2d, 0xd0, 0x89, 0x13, 0x71, 0x1e, 0x25, 0xd4,
	0xd9, 0x3d, 0x14, 0xc7, 0x2b, 0xde, 0x86, 0xfd, 0xf2, 0x9e, 0xeb, 0x7a, 0x66, 0xcc, 0xeb, 0xca,
	0x72, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xf5, 0x82, 0xb5, 0x55, 0x04, 0x00, 0x00,
}
