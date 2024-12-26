// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/lazy/lazy_normalized_wire_test.proto

package lazy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type FSub struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	B             *uint32                `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	C             *uint32                `protobuf:"varint,3,opt,name=c" json:"c,omitempty"`
	Grandchild    *FSub                  `protobuf:"bytes,4,opt,name=grandchild" json:"grandchild,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FSub) Reset() {
	*x = FSub{}
	mi := &file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FSub) ProtoMessage() {}

func (x *FSub) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FSub.ProtoReflect.Descriptor instead.
func (*FSub) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescGZIP(), []int{0}
}

func (x *FSub) GetB() uint32 {
	if x != nil && x.B != nil {
		return *x.B
	}
	return 0
}

func (x *FSub) GetC() uint32 {
	if x != nil && x.C != nil {
		return *x.C
	}
	return 0
}

func (x *FSub) GetGrandchild() *FSub {
	if x != nil {
		return x.Grandchild
	}
	return nil
}

type FTop struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             *uint32                `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	Child         *FSub                  `protobuf:"bytes,2,opt,name=child" json:"child,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FTop) Reset() {
	*x = FTop{}
	mi := &file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FTop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FTop) ProtoMessage() {}

func (x *FTop) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FTop.ProtoReflect.Descriptor instead.
func (*FTop) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescGZIP(), []int{1}
}

func (x *FTop) GetA() uint32 {
	if x != nil && x.A != nil {
		return *x.A
	}
	return 0
}

func (x *FTop) GetChild() *FSub {
	if x != nil {
		return x.Child
	}
	return nil
}

var File_internal_testprotos_lazy_lazy_normalized_wire_test_proto protoreflect.FileDescriptor

var file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDesc = []byte{
	0x0a, 0x38, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x72, 0x65, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6c, 0x61, 0x7a, 0x79,
	0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x72, 0x65,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x04, 0x46, 0x53, 0x75, 0x62, 0x12, 0x0c, 0x0a,
	0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x63, 0x12, 0x43, 0x0a, 0x0a, 0x67, 0x72, 0x61,
	0x6e, 0x64, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x77, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x53, 0x75, 0x62, 0x42, 0x02,
	0x28, 0x01, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x64, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x22, 0x4b,
	0x0a, 0x04, 0x46, 0x54, 0x6f, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x01, 0x61, 0x12, 0x35, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x46, 0x53, 0x75, 0x62, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61,
	0x7a, 0x79,
}

var (
	file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescData = file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDesc
)

func file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescData)
	})
	return file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDescData
}

var file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_goTypes = []any{
	(*FSub)(nil), // 0: lazy_normalized_wire_test.FSub
	(*FTop)(nil), // 1: lazy_normalized_wire_test.FTop
}
var file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_depIdxs = []int32{
	0, // 0: lazy_normalized_wire_test.FSub.grandchild:type_name -> lazy_normalized_wire_test.FSub
	0, // 1: lazy_normalized_wire_test.FTop.child:type_name -> lazy_normalized_wire_test.FSub
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_init() }
func file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_init() {
	if File_internal_testprotos_lazy_lazy_normalized_wire_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_msgTypes,
	}.Build()
	File_internal_testprotos_lazy_lazy_normalized_wire_test_proto = out.File
	file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_rawDesc = nil
	file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_goTypes = nil
	file_internal_testprotos_lazy_lazy_normalized_wire_test_proto_depIdxs = nil
}
