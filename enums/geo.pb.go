// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: api/enums/geo.proto

package enums

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

type GeoSegment int32

const (
	GeoSegment_WORLD        GeoSegment = 0
	GeoSegment_CONTINENT    GeoSegment = 1
	GeoSegment_COUNTRY      GeoSegment = 2
	GeoSegment_STATE        GeoSegment = 3
	GeoSegment_COUNTY       GeoSegment = 4
	GeoSegment_MUNICIPALITY GeoSegment = 5
	GeoSegment_NEIGHBORHOOD GeoSegment = 6
)

// Enum value maps for GeoSegment.
var (
	GeoSegment_name = map[int32]string{
		0: "WORLD",
		1: "CONTINENT",
		2: "COUNTRY",
		3: "STATE",
		4: "COUNTY",
		5: "MUNICIPALITY",
		6: "NEIGHBORHOOD",
	}
	GeoSegment_value = map[string]int32{
		"WORLD":        0,
		"CONTINENT":    1,
		"COUNTRY":      2,
		"STATE":        3,
		"COUNTY":       4,
		"MUNICIPALITY": 5,
		"NEIGHBORHOOD": 6,
	}
)

func (x GeoSegment) Enum() *GeoSegment {
	p := new(GeoSegment)
	*p = x
	return p
}

func (x GeoSegment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeoSegment) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_geo_proto_enumTypes[0].Descriptor()
}

func (GeoSegment) Type() protoreflect.EnumType {
	return &file_api_enums_geo_proto_enumTypes[0]
}

func (x GeoSegment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeoSegment.Descriptor instead.
func (GeoSegment) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_geo_proto_rawDescGZIP(), []int{0}
}

type GeoSegmentGroup int32

const (
	GeoSegmentGroup_REGION GeoSegmentGroup = 0
	GeoSegmentGroup_ZONE   GeoSegmentGroup = 1
)

// Enum value maps for GeoSegmentGroup.
var (
	GeoSegmentGroup_name = map[int32]string{
		0: "REGION",
		1: "ZONE",
	}
	GeoSegmentGroup_value = map[string]int32{
		"REGION": 0,
		"ZONE":   1,
	}
)

func (x GeoSegmentGroup) Enum() *GeoSegmentGroup {
	p := new(GeoSegmentGroup)
	*p = x
	return p
}

func (x GeoSegmentGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeoSegmentGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_api_enums_geo_proto_enumTypes[1].Descriptor()
}

func (GeoSegmentGroup) Type() protoreflect.EnumType {
	return &file_api_enums_geo_proto_enumTypes[1]
}

func (x GeoSegmentGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeoSegmentGroup.Descriptor instead.
func (GeoSegmentGroup) EnumDescriptor() ([]byte, []int) {
	return file_api_enums_geo_proto_rawDescGZIP(), []int{1}
}

var File_api_enums_geo_proto protoreflect.FileDescriptor

var file_api_enums_geo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x67, 0x65, 0x6f, 0x2a, 0x6e,
	0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x09, 0x0a, 0x05,
	0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x54, 0x49,
	0x4e, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52,
	0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x59, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x55,
	0x4e, 0x49, 0x43, 0x49, 0x50, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c,
	0x4e, 0x45, 0x49, 0x47, 0x48, 0x42, 0x4f, 0x52, 0x48, 0x4f, 0x4f, 0x44, 0x10, 0x06, 0x2a, 0x27,
	0x0a, 0x0f, 0x47, 0x65, 0x6f, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x42, 0x53, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2d, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_enums_geo_proto_rawDescOnce sync.Once
	file_api_enums_geo_proto_rawDescData = file_api_enums_geo_proto_rawDesc
)

func file_api_enums_geo_proto_rawDescGZIP() []byte {
	file_api_enums_geo_proto_rawDescOnce.Do(func() {
		file_api_enums_geo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_enums_geo_proto_rawDescData)
	})
	return file_api_enums_geo_proto_rawDescData
}

var file_api_enums_geo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_enums_geo_proto_goTypes = []interface{}{
	(GeoSegment)(0),      // 0: google.retail.enums.geo.GeoSegment
	(GeoSegmentGroup)(0), // 1: google.retail.enums.geo.GeoSegmentGroup
}
var file_api_enums_geo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_enums_geo_proto_init() }
func file_api_enums_geo_proto_init() {
	if File_api_enums_geo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_enums_geo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_enums_geo_proto_goTypes,
		DependencyIndexes: file_api_enums_geo_proto_depIdxs,
		EnumInfos:         file_api_enums_geo_proto_enumTypes,
	}.Build()
	File_api_enums_geo_proto = out.File
	file_api_enums_geo_proto_rawDesc = nil
	file_api_enums_geo_proto_goTypes = nil
	file_api_enums_geo_proto_depIdxs = nil
}