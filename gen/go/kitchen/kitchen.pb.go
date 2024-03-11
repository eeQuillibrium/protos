// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: kitchen.proto

package nikita_kitchen1

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

type SendOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units  []*SendOrderReq_PieceUnitnum `protobuf:"bytes,1,rep,name=units,proto3" json:"units,omitempty"`
	Userid int64                        `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Price  int64                        `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SendOrderReq) Reset() {
	*x = SendOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderReq) ProtoMessage() {}

func (x *SendOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderReq.ProtoReflect.Descriptor instead.
func (*SendOrderReq) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{0}
}

func (x *SendOrderReq) GetUnits() []*SendOrderReq_PieceUnitnum {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *SendOrderReq) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *SendOrderReq) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type EmptyOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyOrderResp) Reset() {
	*x = EmptyOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyOrderResp) ProtoMessage() {}

func (x *EmptyOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyOrderResp.ProtoReflect.Descriptor instead.
func (*EmptyOrderResp) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{1}
}

type SendOrderReq_PieceUnitnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Piece   int64 `protobuf:"varint,1,opt,name=piece,proto3" json:"piece,omitempty"`
	Unitnum int64 `protobuf:"varint,2,opt,name=unitnum,proto3" json:"unitnum,omitempty"`
}

func (x *SendOrderReq_PieceUnitnum) Reset() {
	*x = SendOrderReq_PieceUnitnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kitchen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendOrderReq_PieceUnitnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderReq_PieceUnitnum) ProtoMessage() {}

func (x *SendOrderReq_PieceUnitnum) ProtoReflect() protoreflect.Message {
	mi := &file_kitchen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderReq_PieceUnitnum.ProtoReflect.Descriptor instead.
func (*SendOrderReq_PieceUnitnum) Descriptor() ([]byte, []int) {
	return file_kitchen_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SendOrderReq_PieceUnitnum) GetPiece() int64 {
	if x != nil {
		return x.Piece
	}
	return 0
}

func (x *SendOrderReq_PieceUnitnum) GetUnitnum() int64 {
	if x != nil {
		return x.Unitnum
	}
	return 0
}

var File_kitchen_proto protoreflect.FileDescriptor

var file_kitchen_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x38, 0x0a, 0x05, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x2e,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x1a, 0x3e, 0x0a, 0x0c, 0x50, 0x69, 0x65, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x6e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x69, 0x65, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x6e, 0x75,
	0x6d, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x32, 0x46, 0x0a, 0x07, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x12, 0x3b,
	0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6b, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x11, 0x5a, 0x0f, 0x6e,
	0x69, 0x6b, 0x69, 0x74, 0x61, 0x2e, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kitchen_proto_rawDescOnce sync.Once
	file_kitchen_proto_rawDescData = file_kitchen_proto_rawDesc
)

func file_kitchen_proto_rawDescGZIP() []byte {
	file_kitchen_proto_rawDescOnce.Do(func() {
		file_kitchen_proto_rawDescData = protoimpl.X.CompressGZIP(file_kitchen_proto_rawDescData)
	})
	return file_kitchen_proto_rawDescData
}

var file_kitchen_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kitchen_proto_goTypes = []interface{}{
	(*SendOrderReq)(nil),              // 0: kitchen.SendOrderReq
	(*EmptyOrderResp)(nil),            // 1: kitchen.EmptyOrderResp
	(*SendOrderReq_PieceUnitnum)(nil), // 2: kitchen.SendOrderReq.PieceUnitnum
}
var file_kitchen_proto_depIdxs = []int32{
	2, // 0: kitchen.SendOrderReq.units:type_name -> kitchen.SendOrderReq.PieceUnitnum
	0, // 1: kitchen.Kitchen.SendOrder:input_type -> kitchen.SendOrderReq
	1, // 2: kitchen.Kitchen.SendOrder:output_type -> kitchen.EmptyOrderResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kitchen_proto_init() }
func file_kitchen_proto_init() {
	if File_kitchen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kitchen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderReq); i {
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
		file_kitchen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyOrderResp); i {
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
		file_kitchen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendOrderReq_PieceUnitnum); i {
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
			RawDescriptor: file_kitchen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kitchen_proto_goTypes,
		DependencyIndexes: file_kitchen_proto_depIdxs,
		MessageInfos:      file_kitchen_proto_msgTypes,
	}.Build()
	File_kitchen_proto = out.File
	file_kitchen_proto_rawDesc = nil
	file_kitchen_proto_goTypes = nil
	file_kitchen_proto_depIdxs = nil
}
