// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: message.proto

package pb

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

// 同步消息
type SyncPid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务器将 ID 发送给玩家
	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
}

func (x *SyncPid) Reset() {
	*x = SyncPid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPid) ProtoMessage() {}

func (x *SyncPid) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
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
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *SyncPid) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// 玩家坐标
type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"` // 表示 X 轴
	Y float32 `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"` // 表示 Z 轴
	Z float32 `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"` // 表示 Y 轴
	V float32 `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"` // 表示方向
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
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
	return file_message_proto_rawDescGZIP(), []int{1}
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

// 广播消息
type BroadCast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"` // 接收消息的玩家
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`   // 发送的消息类型: TP=1 表示世界聊天 TP=2 表示玩家位置 TP=3 玩家动作 TP=4 表示玩家移动后的坐标信息更新
	// Types that are assignable to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data isBroadCast_Data `protobuf_oneof:"Data"`
}

func (x *BroadCast) Reset() {
	*x = BroadCast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadCast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadCast) ProtoMessage() {}

func (x *BroadCast) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadCast.ProtoReflect.Descriptor instead.
func (*BroadCast) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *BroadCast) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *BroadCast) GetTp() int32 {
	if x != nil {
		return x.Tp
	}
	return 0
}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *BroadCast) GetContent() string {
	if x, ok := x.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (x *BroadCast) GetP() *Position {
	if x, ok := x.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (x *BroadCast) GetActionData() int32 {
	if x, ok := x.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"` // 玩家聊天信息
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"` // 玩家位置
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"` // 玩家具体行为
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

// 聊天消息
type Talk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Talk) Reset() {
	*x = Talk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Talk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Talk) ProtoMessage() {}

func (x *Talk) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
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
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Talk) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 同步玩家位置
type SyncPlayers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=Players,proto3" json:"Players,omitempty"`
}

func (x *SyncPlayers) Reset() {
	*x = SyncPlayers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayers) ProtoMessage() {}

func (x *SyncPlayers) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayers.ProtoReflect.Descriptor instead.
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *SyncPlayers) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32     `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	P   *Position `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Player) GetP() *Position {
	if x != nil {
		return x.P
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x1b, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64,
	0x22, 0x42, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01,
	0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x5a, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x56, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x50, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x54, 0x70, 0x12, 0x1a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x01, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x01, 0x50, 0x12, 0x20,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x04, 0x54, 0x61, 0x6c, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x0b, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22,
	0x36, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x01, 0x50,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x01, 0x50, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0xaa,
	0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_message_proto_goTypes = []interface{}{
	(*SyncPid)(nil),     // 0: pb.SyncPid
	(*Position)(nil),    // 1: pb.Position
	(*BroadCast)(nil),   // 2: pb.BroadCast
	(*Talk)(nil),        // 3: pb.Talk
	(*SyncPlayers)(nil), // 4: pb.SyncPlayers
	(*Player)(nil),      // 5: pb.Player
}
var file_message_proto_depIdxs = []int32{
	1, // 0: pb.BroadCast.P:type_name -> pb.Position
	5, // 1: pb.SyncPlayers.Players:type_name -> pb.Player
	1, // 2: pb.Player.P:type_name -> pb.Position
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadCast); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayers); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
	file_message_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
