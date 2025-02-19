// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/testeditions/testeditions_hybrid/test_import.hybrid.proto

//go:build protoopaque

package testeditions_hybrid

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
	unsafe "unsafe"
)

type ImportEnum int32

const (
	ImportEnum_IMPORT_ZERO ImportEnum = 0
)

// Enum value maps for ImportEnum.
var (
	ImportEnum_name = map[int32]string{
		0: "IMPORT_ZERO",
	}
	ImportEnum_value = map[string]int32{
		"IMPORT_ZERO": 0,
	}
)

func (x ImportEnum) Enum() *ImportEnum {
	p := new(ImportEnum)
	*p = x
	return p
}

func (x ImportEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_enumTypes[0].Descriptor()
}

func (ImportEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_enumTypes[0]
}

func (x ImportEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type ImportMessage struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportMessage) Reset() {
	*x = ImportMessage{}
	mi := &file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportMessage) ProtoMessage() {}

func (x *ImportMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type ImportMessage_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 ImportMessage_builder) Build() *ImportMessage {
	m0 := &ImportMessage{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var File_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto protoreflect.FileDescriptor

var file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_rawDesc = string([]byte{
	0x0a, 0x4d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1d, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x5a,
	0x45, 0x52, 0x4f, 0x10, 0x00, 0x42, 0x59, 0x5a, 0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x02,
	0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_goTypes = []any{
	(ImportEnum)(0),       // 0: hybrid.goproto.proto.testeditions.ImportEnum
	(*ImportMessage)(nil), // 1: hybrid.goproto.proto.testeditions.ImportMessage
}
var file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_init()
}
func file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_init() {
	if File_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_rawDesc), len(file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_msgTypes,
	}.Build()
	File_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto = out.File
	file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_goTypes = nil
	file_internal_testprotos_testeditions_testeditions_hybrid_test_import_hybrid_proto_depIdxs = nil
}
