// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: othello/v1/othello.proto

package othello

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAIMoveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Board         []int32                `protobuf:"varint,1,rep,packed,name=board,proto3" json:"board,omitempty"`
	Player        int32                  `protobuf:"varint,2,opt,name=player,proto3" json:"player,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAIMoveRequest) Reset() {
	*x = GetAIMoveRequest{}
	mi := &file_othello_v1_othello_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAIMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAIMoveRequest) ProtoMessage() {}

func (x *GetAIMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_othello_v1_othello_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAIMoveRequest.ProtoReflect.Descriptor instead.
func (*GetAIMoveRequest) Descriptor() ([]byte, []int) {
	return file_othello_v1_othello_proto_rawDescGZIP(), []int{0}
}

func (x *GetAIMoveRequest) GetBoard() []int32 {
	if x != nil {
		return x.Board
	}
	return nil
}

func (x *GetAIMoveRequest) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

type GetAIMoveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Y             int32                  `protobuf:"varint,1,opt,name=y,proto3" json:"y,omitempty"`
	X             int32                  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAIMoveResponse) Reset() {
	*x = GetAIMoveResponse{}
	mi := &file_othello_v1_othello_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAIMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAIMoveResponse) ProtoMessage() {}

func (x *GetAIMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_othello_v1_othello_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAIMoveResponse.ProtoReflect.Descriptor instead.
func (*GetAIMoveResponse) Descriptor() ([]byte, []int) {
	return file_othello_v1_othello_proto_rawDescGZIP(), []int{1}
}

func (x *GetAIMoveResponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *GetAIMoveResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

var File_othello_v1_othello_proto protoreflect.FileDescriptor

var file_othello_v1_othello_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x6f, 0x74, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x74, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x74, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x49, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x49, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x32, 0x5a, 0x0a, 0x0e, 0x4f, 0x74, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x41, 0x49, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x1c, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x49, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x49, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2d, 0x6f, 0x74, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x6f,
	0x74, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_othello_v1_othello_proto_rawDescOnce sync.Once
	file_othello_v1_othello_proto_rawDescData []byte
)

func file_othello_v1_othello_proto_rawDescGZIP() []byte {
	file_othello_v1_othello_proto_rawDescOnce.Do(func() {
		file_othello_v1_othello_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_othello_v1_othello_proto_rawDesc), len(file_othello_v1_othello_proto_rawDesc)))
	})
	return file_othello_v1_othello_proto_rawDescData
}

var file_othello_v1_othello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_othello_v1_othello_proto_goTypes = []any{
	(*GetAIMoveRequest)(nil),  // 0: othello.v1.GetAIMoveRequest
	(*GetAIMoveResponse)(nil), // 1: othello.v1.GetAIMoveResponse
}
var file_othello_v1_othello_proto_depIdxs = []int32{
	0, // 0: othello.v1.OthelloService.GetAIMove:input_type -> othello.v1.GetAIMoveRequest
	1, // 1: othello.v1.OthelloService.GetAIMove:output_type -> othello.v1.GetAIMoveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_othello_v1_othello_proto_init() }
func file_othello_v1_othello_proto_init() {
	if File_othello_v1_othello_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_othello_v1_othello_proto_rawDesc), len(file_othello_v1_othello_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_othello_v1_othello_proto_goTypes,
		DependencyIndexes: file_othello_v1_othello_proto_depIdxs,
		MessageInfos:      file_othello_v1_othello_proto_msgTypes,
	}.Build()
	File_othello_v1_othello_proto = out.File
	file_othello_v1_othello_proto_goTypes = nil
	file_othello_v1_othello_proto_depIdxs = nil
}
