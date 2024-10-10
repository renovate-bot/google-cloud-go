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
// source: google/cloud/retail/v2alpha/generative_question.proto

package retailpb

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

// Configuration for overall generative question feature state.
type GenerativeQuestionsFeatureConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the affected catalog.
	// Format: projects/{project}/locations/{location}/catalogs/{catalog}
	Catalog string `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	// Optional. Determines whether questions will be used at serving time.
	// Note: This feature cannot be enabled until initial data requirements are
	// satisfied.
	FeatureEnabled bool `protobuf:"varint,2,opt,name=feature_enabled,json=featureEnabled,proto3" json:"feature_enabled,omitempty"`
	// Optional. Minimum number of products in the response to trigger follow-up
	// questions. Value must be 0 or positive.
	MinimumProducts int32 `protobuf:"varint,3,opt,name=minimum_products,json=minimumProducts,proto3" json:"minimum_products,omitempty"`
}

func (x *GenerativeQuestionsFeatureConfig) Reset() {
	*x = GenerativeQuestionsFeatureConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerativeQuestionsFeatureConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerativeQuestionsFeatureConfig) ProtoMessage() {}

func (x *GenerativeQuestionsFeatureConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerativeQuestionsFeatureConfig.ProtoReflect.Descriptor instead.
func (*GenerativeQuestionsFeatureConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_generative_question_proto_rawDescGZIP(), []int{0}
}

func (x *GenerativeQuestionsFeatureConfig) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *GenerativeQuestionsFeatureConfig) GetFeatureEnabled() bool {
	if x != nil {
		return x.FeatureEnabled
	}
	return false
}

func (x *GenerativeQuestionsFeatureConfig) GetMinimumProducts() int32 {
	if x != nil {
		return x.MinimumProducts
	}
	return 0
}

// Configuration for a single generated question.
type GenerativeQuestionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the catalog.
	// Format: projects/{project}/locations/{location}/catalogs/{catalog}
	Catalog string `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	// Required. The facet to which the question is associated.
	Facet string `protobuf:"bytes,2,opt,name=facet,proto3" json:"facet,omitempty"`
	// Output only. The LLM generated question.
	GeneratedQuestion string `protobuf:"bytes,3,opt,name=generated_question,json=generatedQuestion,proto3" json:"generated_question,omitempty"`
	// Optional. The question that will be used at serving time.
	// Question can have a max length of 300 bytes.
	// When not populated, generated_question should be used.
	FinalQuestion string `protobuf:"bytes,4,opt,name=final_question,json=finalQuestion,proto3" json:"final_question,omitempty"`
	// Output only. Values that can be used to answer the question.
	ExampleValues []string `protobuf:"bytes,5,rep,name=example_values,json=exampleValues,proto3" json:"example_values,omitempty"`
	// Output only. The ratio of how often a question was asked.
	Frequency float32 `protobuf:"fixed32,6,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Optional. Whether the question is asked at serving time.
	AllowedInConversation bool `protobuf:"varint,7,opt,name=allowed_in_conversation,json=allowedInConversation,proto3" json:"allowed_in_conversation,omitempty"`
}

func (x *GenerativeQuestionConfig) Reset() {
	*x = GenerativeQuestionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerativeQuestionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerativeQuestionConfig) ProtoMessage() {}

func (x *GenerativeQuestionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerativeQuestionConfig.ProtoReflect.Descriptor instead.
func (*GenerativeQuestionConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_generative_question_proto_rawDescGZIP(), []int{1}
}

func (x *GenerativeQuestionConfig) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *GenerativeQuestionConfig) GetFacet() string {
	if x != nil {
		return x.Facet
	}
	return ""
}

func (x *GenerativeQuestionConfig) GetGeneratedQuestion() string {
	if x != nil {
		return x.GeneratedQuestion
	}
	return ""
}

func (x *GenerativeQuestionConfig) GetFinalQuestion() string {
	if x != nil {
		return x.FinalQuestion
	}
	return ""
}

func (x *GenerativeQuestionConfig) GetExampleValues() []string {
	if x != nil {
		return x.ExampleValues
	}
	return nil
}

func (x *GenerativeQuestionConfig) GetFrequency() float32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *GenerativeQuestionConfig) GetAllowedInConversation() bool {
	if x != nil {
		return x.AllowedInConversation
	}
	return false
}

var File_google_cloud_retail_v2alpha_generative_question_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_generative_question_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x2c, 0x0a, 0x0f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xc0, 0x02, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x66, 0x61, 0x63, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x66, 0x61, 0x63, 0x65, 0x74, 0x12, 0x32,
	0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a,
	0x0a, 0x0e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3b, 0x0a,
	0x17, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x6e, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xdb, 0x01, 0x0a, 0x1f, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x17,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x70, 0x62, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c,
	0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a,
	0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_generative_question_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_generative_question_proto_rawDescData = file_google_cloud_retail_v2alpha_generative_question_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_generative_question_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_generative_question_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_generative_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_generative_question_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_generative_question_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_retail_v2alpha_generative_question_proto_goTypes = []any{
	(*GenerativeQuestionsFeatureConfig)(nil), // 0: google.cloud.retail.v2alpha.GenerativeQuestionsFeatureConfig
	(*GenerativeQuestionConfig)(nil),         // 1: google.cloud.retail.v2alpha.GenerativeQuestionConfig
}
var file_google_cloud_retail_v2alpha_generative_question_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_generative_question_proto_init() }
func file_google_cloud_retail_v2alpha_generative_question_proto_init() {
	if File_google_cloud_retail_v2alpha_generative_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GenerativeQuestionsFeatureConfig); i {
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
		file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerativeQuestionConfig); i {
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
			RawDescriptor: file_google_cloud_retail_v2alpha_generative_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_generative_question_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_generative_question_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2alpha_generative_question_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_generative_question_proto = out.File
	file_google_cloud_retail_v2alpha_generative_question_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_generative_question_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_generative_question_proto_depIdxs = nil
}