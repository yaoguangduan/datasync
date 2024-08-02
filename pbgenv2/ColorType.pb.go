// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: ColorType.proto

package pbgenv2

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

type ColorType int32

const (
	ColorType_Red   ColorType = 0 //
	ColorType_Blue  ColorType = 1 //
	ColorType_Green ColorType = 2 //
)

// Enum value maps for ColorType.
var (
	ColorType_name = map[int32]string{
		0: "Red",
		1: "Blue",
		2: "Green",
	}
	ColorType_value = map[string]int32{
		"Red":   0,
		"Blue":  1,
		"Green": 2,
	}
)

func (x ColorType) Enum() *ColorType {
	p := new(ColorType)
	*p = x
	return p
}

func (x ColorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColorType) Descriptor() protoreflect.EnumDescriptor {
	return file_ColorType_proto_enumTypes[0].Descriptor()
}

func (ColorType) Type() protoreflect.EnumType {
	return &file_ColorType_proto_enumTypes[0]
}

func (x ColorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColorType.Descriptor instead.
func (ColorType) EnumDescriptor() ([]byte, []int) {
	return file_ColorType_proto_rawDescGZIP(), []int{0}
}

var File_ColorType_proto protoreflect.FileDescriptor

var file_ColorType_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x29, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x52, 0x65, 0x64, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x6c, 0x75, 0x65, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ColorType_proto_rawDescOnce sync.Once
	file_ColorType_proto_rawDescData = file_ColorType_proto_rawDesc
)

func file_ColorType_proto_rawDescGZIP() []byte {
	file_ColorType_proto_rawDescOnce.Do(func() {
		file_ColorType_proto_rawDescData = protoimpl.X.CompressGZIP(file_ColorType_proto_rawDescData)
	})
	return file_ColorType_proto_rawDescData
}

var file_ColorType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ColorType_proto_goTypes = []any{
	(ColorType)(0), // 0: ColorType
}
var file_ColorType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ColorType_proto_init() }
func file_ColorType_proto_init() {
	if File_ColorType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ColorType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ColorType_proto_goTypes,
		DependencyIndexes: file_ColorType_proto_depIdxs,
		EnumInfos:         file_ColorType_proto_enumTypes,
	}.Build()
	File_ColorType_proto = out.File
	file_ColorType_proto_rawDesc = nil
	file_ColorType_proto_goTypes = nil
	file_ColorType_proto_depIdxs = nil
}