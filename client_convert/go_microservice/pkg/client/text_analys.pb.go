// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: text_analys.proto

package protos

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

type ResultParsingPB_Mood int32

const (
	ResultParsingPB_sad      ResultParsingPB_Mood = 0
	ResultParsingPB_happy    ResultParsingPB_Mood = 1
	ResultParsingPB_lovely   ResultParsingPB_Mood = 2
	ResultParsingPB_terrible ResultParsingPB_Mood = 3
	ResultParsingPB_boring   ResultParsingPB_Mood = 4
)

// Enum value maps for ResultParsingPB_Mood.
var (
	ResultParsingPB_Mood_name = map[int32]string{
		0: "sad",
		1: "happy",
		2: "lovely",
		3: "terrible",
		4: "boring",
	}
	ResultParsingPB_Mood_value = map[string]int32{
		"sad":      0,
		"happy":    1,
		"lovely":   2,
		"terrible": 3,
		"boring":   4,
	}
)

func (x ResultParsingPB_Mood) Enum() *ResultParsingPB_Mood {
	p := new(ResultParsingPB_Mood)
	*p = x
	return p
}

func (x ResultParsingPB_Mood) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultParsingPB_Mood) Descriptor() protoreflect.EnumDescriptor {
	return file_text_analys_proto_enumTypes[0].Descriptor()
}

func (ResultParsingPB_Mood) Type() protoreflect.EnumType {
	return &file_text_analys_proto_enumTypes[0]
}

func (x ResultParsingPB_Mood) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultParsingPB_Mood.Descriptor instead.
func (ResultParsingPB_Mood) EnumDescriptor() ([]byte, []int) {
	return file_text_analys_proto_rawDescGZIP(), []int{1, 0}
}

type SettingsTextPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SettingsTextPB) Reset() {
	*x = SettingsTextPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_analys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsTextPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsTextPB) ProtoMessage() {}

func (x *SettingsTextPB) ProtoReflect() protoreflect.Message {
	mi := &file_text_analys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsTextPB.ProtoReflect.Descriptor instead.
func (*SettingsTextPB) Descriptor() ([]byte, []int) {
	return file_text_analys_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsTextPB) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ResultParsingPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WaterValue  int32                `protobuf:"varint,1,opt,name=water_value,json=waterValue,proto3" json:"water_value,omitempty"`
	Mood        ResultParsingPB_Mood `protobuf:"varint,2,opt,name=mood,proto3,enum=Text_analys.ResultParsingPB_Mood" json:"mood,omitempty"`
	HardReading int32                `protobuf:"varint,3,opt,name=hard_reading,json=hardReading,proto3" json:"hard_reading,omitempty"`
}

func (x *ResultParsingPB) Reset() {
	*x = ResultParsingPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_analys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultParsingPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultParsingPB) ProtoMessage() {}

func (x *ResultParsingPB) ProtoReflect() protoreflect.Message {
	mi := &file_text_analys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultParsingPB.ProtoReflect.Descriptor instead.
func (*ResultParsingPB) Descriptor() ([]byte, []int) {
	return file_text_analys_proto_rawDescGZIP(), []int{1}
}

func (x *ResultParsingPB) GetWaterValue() int32 {
	if x != nil {
		return x.WaterValue
	}
	return 0
}

func (x *ResultParsingPB) GetMood() ResultParsingPB_Mood {
	if x != nil {
		return x.Mood
	}
	return ResultParsingPB_sad
}

func (x *ResultParsingPB) GetHardReading() int32 {
	if x != nil {
		return x.HardReading
	}
	return 0
}

var File_text_analys_proto protoreflect.FileDescriptor

var file_text_analys_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x54, 0x65, 0x78, 0x74,
	0x50, 0x42, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xce, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x42, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x77, 0x61, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6d,
	0x6f, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x54, 0x65, 0x78, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61,
	0x72, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x42, 0x2e, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x6d, 0x6f,
	0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x68, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x40, 0x0a, 0x04, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x07, 0x0a,
	0x03, 0x73, 0x61, 0x64, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x68, 0x61, 0x70, 0x70, 0x79, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x6c, 0x6f, 0x76, 0x65, 0x6c, 0x79, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x74, 0x65, 0x72, 0x72, 0x69, 0x62, 0x6c, 0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x62,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x32, 0x5d, 0x0a, 0x11, 0x54, 0x65, 0x78, 0x74, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09,
	0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x2e, 0x54, 0x65, 0x78, 0x74,
	0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x54, 0x65, 0x78, 0x74, 0x50, 0x42, 0x1a, 0x1c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x72, 0x73, 0x69,
	0x6e, 0x67, 0x50, 0x42, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_text_analys_proto_rawDescOnce sync.Once
	file_text_analys_proto_rawDescData = file_text_analys_proto_rawDesc
)

func file_text_analys_proto_rawDescGZIP() []byte {
	file_text_analys_proto_rawDescOnce.Do(func() {
		file_text_analys_proto_rawDescData = protoimpl.X.CompressGZIP(file_text_analys_proto_rawDescData)
	})
	return file_text_analys_proto_rawDescData
}

var file_text_analys_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_text_analys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_text_analys_proto_goTypes = []interface{}{
	(ResultParsingPB_Mood)(0), // 0: Text_analys.ResultParsingPB.Mood
	(*SettingsTextPB)(nil),    // 1: Text_analys.SettingsTextPB
	(*ResultParsingPB)(nil),   // 2: Text_analys.ResultParsingPB
}
var file_text_analys_proto_depIdxs = []int32{
	0, // 0: Text_analys.ResultParsingPB.mood:type_name -> Text_analys.ResultParsingPB.Mood
	1, // 1: Text_analys.TextAnalysService.getResult:input_type -> Text_analys.SettingsTextPB
	2, // 2: Text_analys.TextAnalysService.getResult:output_type -> Text_analys.ResultParsingPB
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_text_analys_proto_init() }
func file_text_analys_proto_init() {
	if File_text_analys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_text_analys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsTextPB); i {
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
		file_text_analys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultParsingPB); i {
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
			RawDescriptor: file_text_analys_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_text_analys_proto_goTypes,
		DependencyIndexes: file_text_analys_proto_depIdxs,
		EnumInfos:         file_text_analys_proto_enumTypes,
		MessageInfos:      file_text_analys_proto_msgTypes,
	}.Build()
	File_text_analys_proto = out.File
	file_text_analys_proto_rawDesc = nil
	file_text_analys_proto_goTypes = nil
	file_text_analys_proto_depIdxs = nil
}