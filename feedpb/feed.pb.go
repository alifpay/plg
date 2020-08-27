// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: feed.proto

package feedpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed    string `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
	ReplyTo string `protobuf:"bytes,2,opt,name=replyTo,proto3" json:"replyTo,omitempty"`
}

func (x *FeedRequest) Reset() {
	*x = FeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRequest) ProtoMessage() {}

func (x *FeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRequest.ProtoReflect.Descriptor instead.
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{0}
}

func (x *FeedRequest) GetFeed() string {
	if x != nil {
		return x.Feed
	}
	return ""
}

func (x *FeedRequest) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

type FeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feed    string `protobuf:"bytes,1,opt,name=feed,proto3" json:"feed,omitempty"`
	ReplyTo string `protobuf:"bytes,2,opt,name=replyTo,proto3" json:"replyTo,omitempty"`
}

func (x *FeedResponse) Reset() {
	*x = FeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResponse) ProtoMessage() {}

func (x *FeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResponse.ProtoReflect.Descriptor instead.
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{1}
}

func (x *FeedResponse) GetFeed() string {
	if x != nil {
		return x.Feed
	}
	return ""
}

func (x *FeedResponse) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

var File_feed_proto protoreflect.FileDescriptor

var file_feed_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x65,
	0x65, 0x64, 0x70, 0x62, 0x22, 0x3b, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x22, 0x3c, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x32,
	0x45, 0x0a, 0x05, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x70, 0x62, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x66, 0x65, 0x65, 0x64, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feed_proto_rawDescOnce sync.Once
	file_feed_proto_rawDescData = file_feed_proto_rawDesc
)

func file_feed_proto_rawDescGZIP() []byte {
	file_feed_proto_rawDescOnce.Do(func() {
		file_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_proto_rawDescData)
	})
	return file_feed_proto_rawDescData
}

var file_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feed_proto_goTypes = []interface{}{
	(*FeedRequest)(nil),  // 0: feedpb.FeedRequest
	(*FeedResponse)(nil), // 1: feedpb.FeedResponse
}
var file_feed_proto_depIdxs = []int32{
	0, // 0: feedpb.Feeds.Broadcast:input_type -> feedpb.FeedRequest
	1, // 1: feedpb.Feeds.Broadcast:output_type -> feedpb.FeedResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feed_proto_init() }
func file_feed_proto_init() {
	if File_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_proto_goTypes,
		DependencyIndexes: file_feed_proto_depIdxs,
		MessageInfos:      file_feed_proto_msgTypes,
	}.Build()
	File_feed_proto = out.File
	file_feed_proto_rawDesc = nil
	file_feed_proto_goTypes = nil
	file_feed_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedsClient is the client API for Feeds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedsClient interface {
	// bi-directional rpc
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (Feeds_BroadcastClient, error)
}

type feedsClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedsClient(cc grpc.ClientConnInterface) FeedsClient {
	return &feedsClient{cc}
}

func (c *feedsClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (Feeds_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Feeds_serviceDesc.Streams[0], "/feedpb.Feeds/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedsBroadcastClient{stream}
	return x, nil
}

type Feeds_BroadcastClient interface {
	Send(*FeedRequest) error
	Recv() (*FeedResponse, error)
	grpc.ClientStream
}

type feedsBroadcastClient struct {
	grpc.ClientStream
}

func (x *feedsBroadcastClient) Send(m *FeedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *feedsBroadcastClient) Recv() (*FeedResponse, error) {
	m := new(FeedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FeedsServer is the server API for Feeds service.
type FeedsServer interface {
	// bi-directional rpc
	Broadcast(Feeds_BroadcastServer) error
}

// UnimplementedFeedsServer can be embedded to have forward compatible implementations.
type UnimplementedFeedsServer struct {
}

func (*UnimplementedFeedsServer) Broadcast(Feeds_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterFeedsServer(s *grpc.Server, srv FeedsServer) {
	s.RegisterService(&_Feeds_serviceDesc, srv)
}

func _Feeds_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FeedsServer).Broadcast(&feedsBroadcastServer{stream})
}

type Feeds_BroadcastServer interface {
	Send(*FeedResponse) error
	Recv() (*FeedRequest, error)
	grpc.ServerStream
}

type feedsBroadcastServer struct {
	grpc.ServerStream
}

func (x *feedsBroadcastServer) Send(m *FeedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *feedsBroadcastServer) Recv() (*FeedRequest, error) {
	m := new(FeedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Feeds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feedpb.Feeds",
	HandlerType: (*FeedsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _Feeds_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "feed.proto",
}
