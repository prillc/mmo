// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: msg.proto

package __

import (
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

// 广播的类型
type BroadcastType int32

const (
	// 聊天
	BroadcastType_TALK BroadcastType = 0
	// 玩家坐标
	BroadcastType_POSITION BroadcastType = 1
)

// Enum value maps for BroadcastType.
var (
	BroadcastType_name = map[int32]string{
		0: "TALK",
		1: "POSITION",
	}
	BroadcastType_value = map[string]int32{
		"TALK":     0,
		"POSITION": 1,
	}
)

func (x BroadcastType) Enum() *BroadcastType {
	p := new(BroadcastType)
	*p = x
	return p
}

func (x BroadcastType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BroadcastType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (BroadcastType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x BroadcastType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BroadcastType.Descriptor instead.
func (BroadcastType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

// 同步玩家数据，玩家上线了
type SyncPid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
}

func (x *SyncPid) Reset() {
	*x = SyncPid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPid) ProtoMessage() {}

func (x *SyncPid) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPid.ProtoReflect.Descriptor instead.
func (*SyncPid) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SyncPid) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// 聊天
type Talk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Talk) Reset() {
	*x = Talk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Talk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Talk) ProtoMessage() {}

func (x *Talk) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Talk.ProtoReflect.Descriptor instead.
func (*Talk) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Talk) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 玩家位置信息
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V float32 `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *Position) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Position) GetV() float32 {
	if x != nil {
		return x.V
	}
	return 0
}

// 服务器广播
type Broadcast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  int32         `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Tp   int32         `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	Type BroadcastType `protobuf:"varint,3,opt,name=type,proto3,enum=pb.BroadcastType" json:"type,omitempty"` // 0聊天 1玩家坐标
	// Types that are assignable to Data:
	//	*Broadcast_Content
	//	*Broadcast_Position
	Data isBroadcast_Data `protobuf_oneof:"Data"`
}

func (x *Broadcast) Reset() {
	*x = Broadcast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Broadcast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Broadcast) ProtoMessage() {}

func (x *Broadcast) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Broadcast.ProtoReflect.Descriptor instead.
func (*Broadcast) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *Broadcast) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Broadcast) GetTp() int32 {
	if x != nil {
		return x.Tp
	}
	return 0
}

func (x *Broadcast) GetType() BroadcastType {
	if x != nil {
		return x.Type
	}
	return BroadcastType_TALK
}

func (m *Broadcast) GetData() isBroadcast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Broadcast) GetContent() string {
	if x, ok := x.GetData().(*Broadcast_Content); ok {
		return x.Content
	}
	return ""
}

func (x *Broadcast) GetPosition() *Position {
	if x, ok := x.GetData().(*Broadcast_Position); ok {
		return x.Position
	}
	return nil
}

type isBroadcast_Data interface {
	isBroadcast_Data()
}

type Broadcast_Content struct {
	// 聊天数据
	Content string `protobuf:"bytes,4,opt,name=Content,proto3,oneof"`
}

type Broadcast_Position struct {
	// 位置数据
	Position *Position `protobuf:"bytes,5,opt,name=position,proto3,oneof"`
}

func (*Broadcast_Content) isBroadcast_Data() {}

func (*Broadcast_Position) isBroadcast_Data() {}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x1b, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x04,
	0x54, 0x61, 0x6c, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x42,
	0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x5a, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x56, 0x22, 0xa4, 0x01, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x54, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x27, 0x0a, 0x0d, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x41,
	0x4c, 0x4b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x01, 0x42, 0x08, 0x5a, 0x01, 0x2f, 0xaa, 0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msg_proto_goTypes = []interface{}{
	(BroadcastType)(0), // 0: pb.BroadcastType
	(*SyncPid)(nil),    // 1: pb.SyncPid
	(*Talk)(nil),       // 2: pb.Talk
	(*Position)(nil),   // 3: pb.Position
	(*Broadcast)(nil),  // 4: pb.Broadcast
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: pb.Broadcast.type:type_name -> pb.BroadcastType
	3, // 1: pb.Broadcast.position:type_name -> pb.Position
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPid); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Talk); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Broadcast); i {
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
	file_msg_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Broadcast_Content)(nil),
		(*Broadcast_Position)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
