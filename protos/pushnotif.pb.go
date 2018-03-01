// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pushnotif.proto

/*
Package pushnotif is a generated protocol buffer package.

It is generated from these files:
	pushnotif.proto

It has these top-level messages:
	Alert
	Mode
	Temp
	RegistrationRequest
	RegistrationResponse
	Topic
	Notification
*/
package pushnotif

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type TopicType int32

const (
	TopicType_ALERT TopicType = 0
	TopicType_MODE  TopicType = 1
	TopicType_TEMP  TopicType = 2
)

var TopicType_name = map[int32]string{
	0: "ALERT",
	1: "MODE",
	2: "TEMP",
}
var TopicType_value = map[string]int32{
	"ALERT": 0,
	"MODE":  1,
	"TEMP":  2,
}

func (x TopicType) String() string {
	return proto.EnumName(TopicType_name, int32(x))
}
func (TopicType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Alert struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Alert) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Mode struct {
	NewMode string `protobuf:"bytes,1,opt,name=newMode" json:"newMode,omitempty"`
}

func (m *Mode) Reset()                    { *m = Mode{} }
func (m *Mode) String() string            { return proto.CompactTextString(m) }
func (*Mode) ProtoMessage()               {}
func (*Mode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Mode) GetNewMode() string {
	if m != nil {
		return m.NewMode
	}
	return ""
}

type Temp struct {
	NewTemp uint32 `protobuf:"varint,1,opt,name=newTemp" json:"newTemp,omitempty"`
}

func (m *Temp) Reset()                    { *m = Temp{} }
func (m *Temp) String() string            { return proto.CompactTextString(m) }
func (*Temp) ProtoMessage()               {}
func (*Temp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Temp) GetNewTemp() uint32 {
	if m != nil {
		return m.NewTemp
	}
	return 0
}

type RegistrationRequest struct {
	ClientName string `protobuf:"bytes,1,opt,name=clientName" json:"clientName,omitempty"`
}

func (m *RegistrationRequest) Reset()                    { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string            { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()               {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegistrationRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type RegistrationResponse struct {
	ClientName string `protobuf:"bytes,1,opt,name=clientName" json:"clientName,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=serverName" json:"serverName,omitempty"`
}

func (m *RegistrationResponse) Reset()                    { *m = RegistrationResponse{} }
func (m *RegistrationResponse) String() string            { return proto.CompactTextString(m) }
func (*RegistrationResponse) ProtoMessage()               {}
func (*RegistrationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegistrationResponse) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *RegistrationResponse) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

type Topic struct {
	ClientName string `protobuf:"bytes,1,opt,name=clientName" json:"clientName,omitempty"`
	// uint64 clientCookie = 1;
	Type TopicType `protobuf:"varint,2,opt,name=type,enum=TopicType" json:"type,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Topic) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Topic) GetType() TopicType {
	if m != nil {
		return m.Type
	}
	return TopicType_ALERT
}

type Notification struct {
	ClientName string    `protobuf:"bytes,1,opt,name=clientName" json:"clientName,omitempty"`
	ServerName string    `protobuf:"bytes,2,opt,name=serverName" json:"serverName,omitempty"`
	Type       TopicType `protobuf:"varint,3,opt,name=type,enum=TopicType" json:"type,omitempty"`
	A          *Alert    `protobuf:"bytes,6,opt,name=a" json:"a,omitempty"`
	M          *Mode     `protobuf:"bytes,4,opt,name=m" json:"m,omitempty"`
	T          *Temp     `protobuf:"bytes,5,opt,name=t" json:"t,omitempty"`
}

func (m *Notification) Reset()                    { *m = Notification{} }
func (m *Notification) String() string            { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()               {}
func (*Notification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Notification) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Notification) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *Notification) GetType() TopicType {
	if m != nil {
		return m.Type
	}
	return TopicType_ALERT
}

func (m *Notification) GetA() *Alert {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Notification) GetM() *Mode {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Notification) GetT() *Temp {
	if m != nil {
		return m.T
	}
	return nil
}

func init() {
	proto.RegisterType((*Alert)(nil), "Alert")
	proto.RegisterType((*Mode)(nil), "Mode")
	proto.RegisterType((*Temp)(nil), "Temp")
	proto.RegisterType((*RegistrationRequest)(nil), "RegistrationRequest")
	proto.RegisterType((*RegistrationResponse)(nil), "RegistrationResponse")
	proto.RegisterType((*Topic)(nil), "Topic")
	proto.RegisterType((*Notification)(nil), "Notification")
	proto.RegisterEnum("TopicType", TopicType_name, TopicType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PushNotif service

type PushNotifClient interface {
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	Alert(ctx context.Context, in *Topic, opts ...grpc.CallOption) (PushNotif_AlertClient, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (PushNotif_SubscribeClient, error)
}

type pushNotifClient struct {
	cc *grpc.ClientConn
}

func NewPushNotifClient(cc *grpc.ClientConn) PushNotifClient {
	return &pushNotifClient{cc}
}

func (c *pushNotifClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := grpc.Invoke(ctx, "/pushNotif/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushNotifClient) Alert(ctx context.Context, in *Topic, opts ...grpc.CallOption) (PushNotif_AlertClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PushNotif_serviceDesc.Streams[0], c.cc, "/pushNotif/Alert", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushNotifAlertClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushNotif_AlertClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type pushNotifAlertClient struct {
	grpc.ClientStream
}

func (x *pushNotifAlertClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushNotifClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (PushNotif_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PushNotif_serviceDesc.Streams[1], c.cc, "/pushNotif/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushNotifSubscribeClient{stream}
	return x, nil
}

type PushNotif_SubscribeClient interface {
	Send(*Topic) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type pushNotifSubscribeClient struct {
	grpc.ClientStream
}

func (x *pushNotifSubscribeClient) Send(m *Topic) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pushNotifSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PushNotif service

type PushNotifServer interface {
	Register(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	Alert(*Topic, PushNotif_AlertServer) error
	Subscribe(PushNotif_SubscribeServer) error
}

func RegisterPushNotifServer(s *grpc.Server, srv PushNotifServer) {
	s.RegisterService(&_PushNotif_serviceDesc, srv)
}

func _PushNotif_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushNotifServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pushNotif/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushNotifServer).Register(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushNotif_Alert_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushNotifServer).Alert(m, &pushNotifAlertServer{stream})
}

type PushNotif_AlertServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type pushNotifAlertServer struct {
	grpc.ServerStream
}

func (x *pushNotifAlertServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _PushNotif_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PushNotifServer).Subscribe(&pushNotifSubscribeServer{stream})
}

type PushNotif_SubscribeServer interface {
	Send(*Notification) error
	Recv() (*Topic, error)
	grpc.ServerStream
}

type pushNotifSubscribeServer struct {
	grpc.ServerStream
}

func (x *pushNotifSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pushNotifSubscribeServer) Recv() (*Topic, error) {
	m := new(Topic)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PushNotif_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pushNotif",
	HandlerType: (*PushNotifServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PushNotif_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Alert",
			Handler:       _PushNotif_Alert_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _PushNotif_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pushnotif.proto",
}

func init() { proto.RegisterFile("pushnotif.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x5d, 0x4f, 0xea, 0x40,
	0x10, 0xbd, 0xc3, 0x6d, 0x7b, 0xe9, 0x5c, 0x51, 0xb2, 0x60, 0xd2, 0xf0, 0x60, 0x6a, 0x1f, 0x0c,
	0xe1, 0x61, 0x43, 0x30, 0xc6, 0x67, 0x12, 0x89, 0x2f, 0x82, 0x66, 0x6d, 0x7c, 0x2f, 0x75, 0xc4,
	0x26, 0xf4, 0xc3, 0xee, 0xa2, 0xe1, 0x4f, 0xf8, 0x53, 0xfc, 0x8d, 0x66, 0x17, 0x8a, 0x98, 0x60,
	0x34, 0xf1, 0x6d, 0xe6, 0xcc, 0x99, 0xb3, 0xb3, 0x73, 0x06, 0x0f, 0x8a, 0x85, 0x7c, 0xcc, 0x72,
	0x95, 0x3c, 0xf0, 0xa2, 0xcc, 0x55, 0x1e, 0x1c, 0xa3, 0x3d, 0x9c, 0x53, 0xa9, 0x98, 0x87, 0xff,
	0x52, 0x92, 0x32, 0x9a, 0x91, 0x07, 0x3e, 0x74, 0x5d, 0x51, 0xa5, 0x81, 0x8f, 0xd6, 0x38, 0xbf,
	0x27, 0xcd, 0xc8, 0xe8, 0x45, 0x87, 0x15, 0x63, 0x9d, 0x6a, 0x46, 0x48, 0x69, 0xb1, 0x66, 0xe8,
	0xd0, 0x30, 0x1a, 0xa2, 0x4a, 0x83, 0x33, 0x6c, 0x09, 0x9a, 0x25, 0x52, 0x95, 0x91, 0x4a, 0xf2,
	0x4c, 0xd0, 0xd3, 0x82, 0xa4, 0x62, 0x47, 0x88, 0xf1, 0x3c, 0xa1, 0x4c, 0x4d, 0xa2, 0xb4, 0x52,
	0xdd, 0x42, 0x82, 0x3b, 0x6c, 0x7f, 0x6e, 0x93, 0x45, 0x9e, 0x49, 0xfa, 0xae, 0x4f, 0xd7, 0x25,
	0x95, 0xcf, 0x54, 0x9a, 0x7a, 0x6d, 0x55, 0xff, 0x40, 0x82, 0x4b, 0xb4, 0xc3, 0xbc, 0x48, 0xe2,
	0x1f, 0x08, 0x59, 0x6a, 0x59, 0xac, 0x24, 0xf6, 0x07, 0xc8, 0x95, 0xee, 0x0a, 0x97, 0x05, 0x09,
	0x83, 0x07, 0x6f, 0x80, 0x7b, 0x13, 0xbd, 0xce, 0x24, 0x36, 0x13, 0xfe, 0x76, 0xb2, 0xcd, 0x83,
	0x7f, 0x77, 0x3f, 0xc8, 0xda, 0x08, 0x91, 0xe7, 0xf8, 0xd0, 0xfd, 0x3f, 0x70, 0xb8, 0x71, 0x4e,
	0x40, 0xc4, 0x5a, 0x08, 0xa9, 0x67, 0x19, 0xd4, 0xe6, 0xda, 0x12, 0x01, 0xa9, 0x06, 0x95, 0x67,
	0xaf, 0x41, 0xed, 0x82, 0x00, 0xd5, 0xeb, 0xa1, 0xbb, 0x91, 0x64, 0x2e, 0xda, 0xc3, 0xab, 0x91,
	0x08, 0x9b, 0x7f, 0x58, 0x1d, 0xad, 0xf1, 0xf5, 0xc5, 0xa8, 0x09, 0x3a, 0x0a, 0x47, 0xe3, 0x9b,
	0x66, 0x6d, 0xf0, 0x0a, 0xe8, 0xea, 0x7b, 0x31, 0x1f, 0x64, 0xe7, 0x58, 0x5f, 0x79, 0x41, 0x25,
	0x6b, 0xf3, 0x1d, 0x6e, 0x76, 0x0e, 0xf9, 0x4e, 0xb3, 0xfc, 0xea, 0xc4, 0x1c, 0x6e, 0x96, 0xde,
	0x69, 0xf0, 0xed, 0x95, 0xf5, 0x81, 0x9d, 0xa0, 0x7b, 0xbb, 0x98, 0xca, 0xb8, 0x4c, 0xa6, 0xf4,
	0x05, 0xab, 0x0b, 0x7d, 0x98, 0x3a, 0xe6, 0x66, 0x4f, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0x28, 0x33, 0x40, 0xc6, 0x02, 0x00, 0x00,
}
