// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: example_proto/example.proto

package example_proto

import (
	_ "github.com/dotdak/protoc-gen-gorm-dao/gorm"
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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GormTestPrimary     string          `protobuf:"bytes,2,opt,name=gorm_test_primary,json=gormTestPrimary,proto3" json:"gorm_test_primary,omitempty"`
	GormTestInt         int32           `protobuf:"varint,3,opt,name=gorm_test_int,json=gormTestInt,proto3" json:"gorm_test_int,omitempty"`
	GormTestBigInt      int64           `protobuf:"varint,4,opt,name=gorm_test_big_int,json=gormTestBigInt,proto3" json:"gorm_test_big_int,omitempty"`
	GormTestString      string          `protobuf:"bytes,5,opt,name=gorm_test_string,json=gormTestString,proto3" json:"gorm_test_string,omitempty"`
	GormTestFloat       float32         `protobuf:"fixed32,6,opt,name=gorm_test_float,json=gormTestFloat,proto3" json:"gorm_test_float,omitempty"`
	GormTestDouble      float64         `protobuf:"fixed64,7,opt,name=gorm_test_double,json=gormTestDouble,proto3" json:"gorm_test_double,omitempty"`
	GormTestIndex       float64         `protobuf:"fixed64,8,opt,name=gorm_test_index,json=gormTestIndex,proto3" json:"gorm_test_index,omitempty"`
	GormTestUniqueIndex float64         `protobuf:"fixed64,9,opt,name=gorm_test_unique_index,json=gormTestUniqueIndex,proto3" json:"gorm_test_unique_index,omitempty"`
	Child               *ChildExample   `protobuf:"bytes,10,opt,name=child,proto3" json:"child,omitempty"`
	ListChilds          []*ChildExample `protobuf:"bytes,11,rep,name=list_childs,json=listChilds,proto3" json:"list_childs,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Example) GetGormTestPrimary() string {
	if x != nil {
		return x.GormTestPrimary
	}
	return ""
}

func (x *Example) GetGormTestInt() int32 {
	if x != nil {
		return x.GormTestInt
	}
	return 0
}

func (x *Example) GetGormTestBigInt() int64 {
	if x != nil {
		return x.GormTestBigInt
	}
	return 0
}

func (x *Example) GetGormTestString() string {
	if x != nil {
		return x.GormTestString
	}
	return ""
}

func (x *Example) GetGormTestFloat() float32 {
	if x != nil {
		return x.GormTestFloat
	}
	return 0
}

func (x *Example) GetGormTestDouble() float64 {
	if x != nil {
		return x.GormTestDouble
	}
	return 0
}

func (x *Example) GetGormTestIndex() float64 {
	if x != nil {
		return x.GormTestIndex
	}
	return 0
}

func (x *Example) GetGormTestUniqueIndex() float64 {
	if x != nil {
		return x.GormTestUniqueIndex
	}
	return 0
}

func (x *Example) GetChild() *ChildExample {
	if x != nil {
		return x.Child
	}
	return nil
}

func (x *Example) GetListChilds() []*ChildExample {
	if x != nil {
		return x.ListChilds
	}
	return nil
}

type ChildExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChildExample) Reset() {
	*x = ChildExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChildExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChildExample) ProtoMessage() {}

func (x *ChildExample) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChildExample.ProtoReflect.Descriptor instead.
func (*ChildExample) Descriptor() ([]byte, []int) {
	return file_example_proto_example_proto_rawDescGZIP(), []int{1}
}

func (x *ChildExample) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_example_proto_example_proto protoreflect.FileDescriptor

var file_example_proto_example_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x67, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x04, 0x0a, 0x07, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x05, 0xd2, 0x7e, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x11, 0x67,
	0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xd2, 0x7e, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x67,
	0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x29,
	0x0a, 0x0d, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x05, 0xd2, 0x7e, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x67, 0x6f,
	0x72, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x67, 0x6f, 0x72,
	0x6d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x67, 0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x67, 0x49, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x67, 0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26,
	0x0a, 0x0f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x67, 0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73,
	0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0e, 0x67, 0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x12, 0x2d, 0x0a, 0x0f, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x42, 0x05, 0xd2, 0x7e, 0x02, 0x18, 0x01,
	0x52, 0x0d, 0x67, 0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x3a, 0x0a, 0x16, 0x67, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x05, 0xd2, 0x7e, 0x02, 0x28, 0x01, 0x52, 0x13, 0x67, 0x6f, 0x72, 0x6d, 0x54, 0x65, 0x73, 0x74,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x31, 0x0a, 0x05, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x3c,
	0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x3a, 0x05, 0xca, 0x7e,
	0x02, 0x08, 0x01, 0x22, 0x25, 0x0a, 0x0c, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x3a, 0x05, 0xca, 0x7e, 0x02, 0x08, 0x01, 0x42, 0xad, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6f, 0x74, 0x64, 0x61, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2d, 0x64, 0x61, 0x6f, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x45,
	0x58, 0x58, 0xaa, 0x02, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0xca, 0x02, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0xe2, 0x02, 0x18, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_example_proto_example_proto_rawDescOnce sync.Once
	file_example_proto_example_proto_rawDescData = file_example_proto_example_proto_rawDesc
)

func file_example_proto_example_proto_rawDescGZIP() []byte {
	file_example_proto_example_proto_rawDescOnce.Do(func() {
		file_example_proto_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_example_proto_rawDescData)
	})
	return file_example_proto_example_proto_rawDescData
}

var file_example_proto_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_proto_example_proto_goTypes = []interface{}{
	(*Example)(nil),      // 0: example_proto.Example
	(*ChildExample)(nil), // 1: example_proto.ChildExample
}
var file_example_proto_example_proto_depIdxs = []int32{
	1, // 0: example_proto.Example.child:type_name -> example_proto.ChildExample
	1, // 1: example_proto.Example.list_childs:type_name -> example_proto.ChildExample
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_proto_example_proto_init() }
func file_example_proto_example_proto_init() {
	if File_example_proto_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_example_proto_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChildExample); i {
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
			RawDescriptor: file_example_proto_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_example_proto_goTypes,
		DependencyIndexes: file_example_proto_example_proto_depIdxs,
		MessageInfos:      file_example_proto_example_proto_msgTypes,
	}.Build()
	File_example_proto_example_proto = out.File
	file_example_proto_example_proto_rawDesc = nil
	file_example_proto_example_proto_goTypes = nil
	file_example_proto_example_proto_depIdxs = nil
}
