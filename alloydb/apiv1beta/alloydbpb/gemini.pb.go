// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/alloydb/v1beta/gemini.proto

package alloydbpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Cluster level configuration parameters related to the Gemini in Databases
// add-on.
type GeminiClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Whether the Gemini in Databases add-on is enabled for the
	// cluster. It will be true only if the add-on has been enabled for the
	// billing account corresponding to the cluster. Its status is toggled from
	// the Admin Control Center (ACC) and cannot be toggled using AlloyDB's APIs.
	Entitled bool `protobuf:"varint,1,opt,name=entitled,proto3" json:"entitled,omitempty"`
}

func (x *GeminiClusterConfig) Reset() {
	*x = GeminiClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeminiClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeminiClusterConfig) ProtoMessage() {}

func (x *GeminiClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeminiClusterConfig.ProtoReflect.Descriptor instead.
func (*GeminiClusterConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_v1beta_gemini_proto_rawDescGZIP(), []int{0}
}

func (x *GeminiClusterConfig) GetEntitled() bool {
	if x != nil {
		return x.Entitled
	}
	return false
}

// Instance level configuration parameters related to the Gemini in Databases
// add-on.
type GeminiInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Whether the Gemini in Databases add-on is enabled for the
	// instance. It will be true only if the add-on has been enabled for the
	// billing account corresponding to the instance. Its status is toggled from
	// the Admin Control Center (ACC) and cannot be toggled using AlloyDB's APIs.
	Entitled bool `protobuf:"varint,1,opt,name=entitled,proto3" json:"entitled,omitempty"`
}

func (x *GeminiInstanceConfig) Reset() {
	*x = GeminiInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeminiInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeminiInstanceConfig) ProtoMessage() {}

func (x *GeminiInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeminiInstanceConfig.ProtoReflect.Descriptor instead.
func (*GeminiInstanceConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_v1beta_gemini_proto_rawDescGZIP(), []int{1}
}

func (x *GeminiInstanceConfig) GetEntitled() bool {
	if x != nil {
		return x.Entitled
	}
	return false
}

var File_google_cloud_alloydb_v1beta_gemini_proto protoreflect.FileDescriptor

var file_google_cloud_alloydb_v1beta_gemini_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x67, 0x65,
	0x6d, 0x69, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x13, 0x47, 0x65, 0x6d, 0x69,
	0x6e, 0x69, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1f, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x64,
	0x22, 0x37, 0x0a, 0x14, 0x47, 0x65, 0x6d, 0x69, 0x6e, 0x69, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x64, 0x42, 0xc8, 0x01, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0b, 0x47,
	0x65, 0x6d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x70, 0x62, 0x3b, 0x61, 0x6c,
	0x6c, 0x6f, 0x79, 0x64, 0x62, 0x70, 0x62, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44, 0x62, 0x2e, 0x56,
	0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44, 0x62, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44, 0x42, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_alloydb_v1beta_gemini_proto_rawDescOnce sync.Once
	file_google_cloud_alloydb_v1beta_gemini_proto_rawDescData = file_google_cloud_alloydb_v1beta_gemini_proto_rawDesc
)

func file_google_cloud_alloydb_v1beta_gemini_proto_rawDescGZIP() []byte {
	file_google_cloud_alloydb_v1beta_gemini_proto_rawDescOnce.Do(func() {
		file_google_cloud_alloydb_v1beta_gemini_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_alloydb_v1beta_gemini_proto_rawDescData)
	})
	return file_google_cloud_alloydb_v1beta_gemini_proto_rawDescData
}

var file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_alloydb_v1beta_gemini_proto_goTypes = []any{
	(*GeminiClusterConfig)(nil),  // 0: google.cloud.alloydb.v1beta.GeminiClusterConfig
	(*GeminiInstanceConfig)(nil), // 1: google.cloud.alloydb.v1beta.GeminiInstanceConfig
}
var file_google_cloud_alloydb_v1beta_gemini_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_alloydb_v1beta_gemini_proto_init() }
func file_google_cloud_alloydb_v1beta_gemini_proto_init() {
	if File_google_cloud_alloydb_v1beta_gemini_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GeminiClusterConfig); i {
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
		file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GeminiInstanceConfig); i {
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
			RawDescriptor: file_google_cloud_alloydb_v1beta_gemini_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_alloydb_v1beta_gemini_proto_goTypes,
		DependencyIndexes: file_google_cloud_alloydb_v1beta_gemini_proto_depIdxs,
		MessageInfos:      file_google_cloud_alloydb_v1beta_gemini_proto_msgTypes,
	}.Build()
	File_google_cloud_alloydb_v1beta_gemini_proto = out.File
	file_google_cloud_alloydb_v1beta_gemini_proto_rawDesc = nil
	file_google_cloud_alloydb_v1beta_gemini_proto_goTypes = nil
	file_google_cloud_alloydb_v1beta_gemini_proto_depIdxs = nil
}