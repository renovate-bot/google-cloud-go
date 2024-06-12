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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/cloud/retail/v2alpha/branch.proto

package retailpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A view that specifies different level of fields of a
// [Branch][google.cloud.retail.v2alpha.Branch] to show in responses.
type BranchView int32

const (
	// The value when it's unspecified. This defaults to the BASIC view.
	BranchView_BRANCH_VIEW_UNSPECIFIED BranchView = 0
	// Includes basic metadata about the branch, but not statistical fields.
	// See documentation of fields of [Branch][google.cloud.retail.v2alpha.Branch]
	// to find what fields are excluded from BASIC view.
	BranchView_BRANCH_VIEW_BASIC BranchView = 1
	// Includes all fields of a [Branch][google.cloud.retail.v2alpha.Branch].
	BranchView_BRANCH_VIEW_FULL BranchView = 2
)

// Enum value maps for BranchView.
var (
	BranchView_name = map[int32]string{
		0: "BRANCH_VIEW_UNSPECIFIED",
		1: "BRANCH_VIEW_BASIC",
		2: "BRANCH_VIEW_FULL",
	}
	BranchView_value = map[string]int32{
		"BRANCH_VIEW_UNSPECIFIED": 0,
		"BRANCH_VIEW_BASIC":       1,
		"BRANCH_VIEW_FULL":        2,
	}
)

func (x BranchView) Enum() *BranchView {
	p := new(BranchView)
	*p = x
	return p
}

func (x BranchView) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BranchView) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2alpha_branch_proto_enumTypes[0].Descriptor()
}

func (BranchView) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2alpha_branch_proto_enumTypes[0]
}

func (x BranchView) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BranchView.Descriptor instead.
func (BranchView) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_branch_proto_rawDescGZIP(), []int{0}
}

// Scope of what products are included for this count.
type Branch_ProductCountStatistic_ProductCountScope int32

const (
	// Default value for enum. This value is not used in the API response.
	Branch_ProductCountStatistic_PRODUCT_COUNT_SCOPE_UNSPECIFIED Branch_ProductCountStatistic_ProductCountScope = 0
	// Scope for all existing products in the branch. Useful for understanding
	// how many products there are in a branch.
	Branch_ProductCountStatistic_ALL_PRODUCTS Branch_ProductCountStatistic_ProductCountScope = 1
	// Scope for products created or updated in the last 24 hours.
	Branch_ProductCountStatistic_LAST_24_HOUR_UPDATE Branch_ProductCountStatistic_ProductCountScope = 2
)

// Enum value maps for Branch_ProductCountStatistic_ProductCountScope.
var (
	Branch_ProductCountStatistic_ProductCountScope_name = map[int32]string{
		0: "PRODUCT_COUNT_SCOPE_UNSPECIFIED",
		1: "ALL_PRODUCTS",
		2: "LAST_24_HOUR_UPDATE",
	}
	Branch_ProductCountStatistic_ProductCountScope_value = map[string]int32{
		"PRODUCT_COUNT_SCOPE_UNSPECIFIED": 0,
		"ALL_PRODUCTS":                    1,
		"LAST_24_HOUR_UPDATE":             2,
	}
)

func (x Branch_ProductCountStatistic_ProductCountScope) Enum() *Branch_ProductCountStatistic_ProductCountScope {
	p := new(Branch_ProductCountStatistic_ProductCountScope)
	*p = x
	return p
}

func (x Branch_ProductCountStatistic_ProductCountScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Branch_ProductCountStatistic_ProductCountScope) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2alpha_branch_proto_enumTypes[1].Descriptor()
}

func (Branch_ProductCountStatistic_ProductCountScope) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2alpha_branch_proto_enumTypes[1]
}

func (x Branch_ProductCountStatistic_ProductCountScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Branch_ProductCountStatistic_ProductCountScope.Descriptor instead.
func (Branch_ProductCountStatistic_ProductCountScope) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_branch_proto_rawDescGZIP(), []int{0, 0, 0}
}

// A data branch that stores [Product][google.cloud.retail.v2alpha.Product]s.
type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Full resource name of the branch, such as
	// `projects/*/locations/global/catalogs/default_catalog/branches/branch_id`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Human readable name of the branch to display in the UI.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Indicates whether this branch is set as the default branch of
	// its parent catalog.
	IsDefault bool `protobuf:"varint,3,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	// Output only. Timestamp of last import through
	// [ProductService.ImportProducts][google.cloud.retail.v2alpha.ProductService.ImportProducts].
	// Empty value means no import has been made to this branch.
	LastProductImportTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_product_import_time,json=lastProductImportTime,proto3" json:"last_product_import_time,omitempty"`
	// Output only. Statistics for number of products in the branch, provided for
	// different
	// [scopes][google.cloud.retail.v2alpha.Branch.ProductCountStatistic.ProductCountScope].
	//
	// This field is not populated in [BranchView.BASIC][] view.
	ProductCountStats []*Branch_ProductCountStatistic `protobuf:"bytes,7,rep,name=product_count_stats,json=productCountStats,proto3" json:"product_count_stats,omitempty"`
	// Output only. The quality metrics measured among products of this branch.
	//
	// See
	// [QualityMetric.requirement_key][google.cloud.retail.v2alpha.Branch.QualityMetric.requirement_key]
	// for supported metrics. Metrics could be missing if failed to retrieve.
	//
	// This field is not populated in [BranchView.BASIC][] view.
	QualityMetrics []*Branch_QualityMetric `protobuf:"bytes,6,rep,name=quality_metrics,json=qualityMetrics,proto3" json:"quality_metrics,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_branch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_branch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_branch_proto_rawDescGZIP(), []int{0}
}

func (x *Branch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Branch) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Branch) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *Branch) GetLastProductImportTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastProductImportTime
	}
	return nil
}

func (x *Branch) GetProductCountStats() []*Branch_ProductCountStatistic {
	if x != nil {
		return x.ProductCountStats
	}
	return nil
}

func (x *Branch) GetQualityMetrics() []*Branch_QualityMetric {
	if x != nil {
		return x.QualityMetrics
	}
	return nil
}

// A statistic about the number of products in a branch.
type Branch_ProductCountStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [ProductCountScope] of the [counts].
	Scope Branch_ProductCountStatistic_ProductCountScope `protobuf:"varint,1,opt,name=scope,proto3,enum=google.cloud.retail.v2alpha.Branch_ProductCountStatistic_ProductCountScope" json:"scope,omitempty"`
	// The number of products in
	// [scope][google.cloud.retail.v2alpha.Branch.ProductCountStatistic.scope]
	// broken down into different groups.
	//
	// The key is a group representing a set of products, and the value is the
	// number of products in that group.
	// Note: keys in this map may change over time.
	//
	// Possible keys:
	// * "primary-in-stock", products have
	// [Product.Type.PRIMARY][google.cloud.retail.v2alpha.Product.Type.PRIMARY]
	// type and
	// [Product.Availability.IN_STOCK][google.cloud.retail.v2alpha.Product.Availability.IN_STOCK]
	// availability.
	//
	// * "primary-out-of-stock", products have
	// [Product.Type.PRIMARY][google.cloud.retail.v2alpha.Product.Type.PRIMARY]
	// type and
	// [Product.Availability.OUT_OF_STOCK][google.cloud.retail.v2alpha.Product.Availability.OUT_OF_STOCK]
	// availability.
	//
	// * "primary-preorder", products have
	// [Product.Type.PRIMARY][google.cloud.retail.v2alpha.Product.Type.PRIMARY]
	// type and
	// [Product.Availability.PREORDER][google.cloud.retail.v2alpha.Product.Availability.PREORDER]
	// availability.
	//
	// * "primary-backorder", products have
	// [Product.Type.PRIMARY][google.cloud.retail.v2alpha.Product.Type.PRIMARY]
	// type and
	// [Product.Availability.BACKORDER][google.cloud.retail.v2alpha.Product.Availability.BACKORDER]
	// availability.
	//
	// * "variant-in-stock", products have
	// [Product.Type.VARIANT][google.cloud.retail.v2alpha.Product.Type.VARIANT]
	// type and
	// [Product.Availability.IN_STOCK][google.cloud.retail.v2alpha.Product.Availability.IN_STOCK]
	// availability.
	//
	// * "variant-out-of-stock", products have
	// [Product.Type.VARIANT][google.cloud.retail.v2alpha.Product.Type.VARIANT]
	// type and
	// [Product.Availability.OUT_OF_STOCK][google.cloud.retail.v2alpha.Product.Availability.OUT_OF_STOCK]
	// availability.
	//
	// * "variant-preorder", products have
	// [Product.Type.VARIANT][google.cloud.retail.v2alpha.Product.Type.VARIANT]
	// type and
	// [Product.Availability.PREORDER][google.cloud.retail.v2alpha.Product.Availability.PREORDER]
	// availability.
	//
	// * "variant-backorder", products have
	// [Product.Type.VARIANT][google.cloud.retail.v2alpha.Product.Type.VARIANT]
	// type and
	// [Product.Availability.BACKORDER][google.cloud.retail.v2alpha.Product.Availability.BACKORDER]
	// availability.
	//
	// * "price-discounted", products have [Product.price_info.price] <
	// [Product.price_info.original_price].
	Counts map[string]int64 `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Branch_ProductCountStatistic) Reset() {
	*x = Branch_ProductCountStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_branch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch_ProductCountStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch_ProductCountStatistic) ProtoMessage() {}

func (x *Branch_ProductCountStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_branch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch_ProductCountStatistic.ProtoReflect.Descriptor instead.
func (*Branch_ProductCountStatistic) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_branch_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Branch_ProductCountStatistic) GetScope() Branch_ProductCountStatistic_ProductCountScope {
	if x != nil {
		return x.Scope
	}
	return Branch_ProductCountStatistic_PRODUCT_COUNT_SCOPE_UNSPECIFIED
}

func (x *Branch_ProductCountStatistic) GetCounts() map[string]int64 {
	if x != nil {
		return x.Counts
	}
	return nil
}

// Metric measured on a group of
// [Product][google.cloud.retail.v2alpha.Product]s against a certain quality
// requirement. Contains the number of products that pass the check and the
// number of products that don't.
type Branch_QualityMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key that represents a quality requirement rule.
	//
	// Supported keys:
	// * "has-valid-uri": product has a valid and accessible
	// [uri][google.cloud.retail.v2alpha.Product.uri].
	//
	// * "available-expire-time-conformance":
	// [Product.available_time][google.cloud.retail.v2alpha.Product.available_time]
	// is early than "now", and
	// [Product.expire_time][google.cloud.retail.v2alpha.Product.expire_time] is
	// greater than "now".
	//
	// * "has-searchable-attributes": product has at least one
	// [attribute][google.cloud.retail.v2alpha.Product.attributes] set to
	// searchable.
	//
	// * "has-description": product has non-empty
	// [description][google.cloud.retail.v2alpha.Product.description].
	//
	// * "has-at-least-bigram-title": Product
	// [title][google.cloud.retail.v2alpha.Product.title] has at least two
	// words. A comprehensive title helps to improve search quality.
	//
	// * "variant-has-image": the
	// [variant][google.cloud.retail.v2alpha.Product.Type.VARIANT] products has
	// at least one [image][google.cloud.retail.v2alpha.Product.images]. You may
	// ignore this metric if all your products are at
	// [primary][google.cloud.retail.v2alpha.Product.Type.PRIMARY] level.
	//
	// * "variant-has-price-info": the
	// [variant][google.cloud.retail.v2alpha.Product.Type.VARIANT] products has
	// [price_info][google.cloud.retail.v2alpha.Product.price_info] set. You may
	// ignore this metric if all your products are at
	// [primary][google.cloud.retail.v2alpha.Product.Type.PRIMARY] level.
	//
	// * "has-publish-time": product has non-empty
	// [publish_time][google.cloud.retail.v2alpha.Product.publish_time].
	RequirementKey string `protobuf:"bytes,1,opt,name=requirement_key,json=requirementKey,proto3" json:"requirement_key,omitempty"`
	// Number of products passing the quality requirement check. We only check
	// searchable products.
	QualifiedProductCount int32 `protobuf:"varint,2,opt,name=qualified_product_count,json=qualifiedProductCount,proto3" json:"qualified_product_count,omitempty"`
	// Number of products failing the quality requirement check. We only check
	// searchable products.
	UnqualifiedProductCount int32 `protobuf:"varint,3,opt,name=unqualified_product_count,json=unqualifiedProductCount,proto3" json:"unqualified_product_count,omitempty"`
	// Value from 0 to 100 representing the suggested percentage of products
	// that meet the quality requirements to get good search and recommendation
	// performance. 100 * (qualified_product_count) /
	// (qualified_product_count + unqualified_product_count) should be greater
	// or equal to this suggestion.
	SuggestedQualityPercentThreshold float64 `protobuf:"fixed64,4,opt,name=suggested_quality_percent_threshold,json=suggestedQualityPercentThreshold,proto3" json:"suggested_quality_percent_threshold,omitempty"`
	// A list of a maximum of 100 sample products that do not qualify for
	// this requirement.
	//
	// This field is only populated in the response to
	// [BranchService.GetBranch][google.cloud.retail.v2alpha.BranchService.GetBranch]
	// API, and is always empty for
	// [BranchService.ListBranches][google.cloud.retail.v2alpha.BranchService.ListBranches].
	//
	// Only the following fields are set in the
	// [Product][google.cloud.retail.v2alpha.Product].
	//
	// * [Product.name][google.cloud.retail.v2alpha.Product.name]
	// * [Product.id][google.cloud.retail.v2alpha.Product.id]
	// * [Product.title][google.cloud.retail.v2alpha.Product.title]
	UnqualifiedSampleProducts []*Product `protobuf:"bytes,5,rep,name=unqualified_sample_products,json=unqualifiedSampleProducts,proto3" json:"unqualified_sample_products,omitempty"`
}

func (x *Branch_QualityMetric) Reset() {
	*x = Branch_QualityMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_branch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch_QualityMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch_QualityMetric) ProtoMessage() {}

func (x *Branch_QualityMetric) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_branch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch_QualityMetric.ProtoReflect.Descriptor instead.
func (*Branch_QualityMetric) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_branch_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Branch_QualityMetric) GetRequirementKey() string {
	if x != nil {
		return x.RequirementKey
	}
	return ""
}

func (x *Branch_QualityMetric) GetQualifiedProductCount() int32 {
	if x != nil {
		return x.QualifiedProductCount
	}
	return 0
}

func (x *Branch_QualityMetric) GetUnqualifiedProductCount() int32 {
	if x != nil {
		return x.UnqualifiedProductCount
	}
	return 0
}

func (x *Branch_QualityMetric) GetSuggestedQualityPercentThreshold() float64 {
	if x != nil {
		return x.SuggestedQualityPercentThreshold
	}
	return 0
}

func (x *Branch_QualityMetric) GetUnqualifiedSampleProducts() []*Product {
	if x != nil {
		return x.UnqualifiedSampleProducts
	}
	return nil
}

var File_google_cloud_retail_v2alpha_branch_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_branch_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe9, 0x09, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x58, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x6e, 0x0a, 0x13, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x0f, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0xf9, 0x02, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x12, 0x61, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x63, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x4c, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x53, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x32, 0x34, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x1a, 0xe1, 0x02, 0x0a, 0x0d, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x36, 0x0a, 0x17, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x75, 0x6e, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x75, 0x6e,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x23, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x20, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x12, 0x64, 0x0a, 0x1b, 0x75, 0x6e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x19, 0x75, 0x6e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x3a, 0x6f, 0xea, 0x41, 0x6c, 0x0a,
	0x1c, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x4c, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x7d, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0x2a, 0x56, 0x0a, 0x0a, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x56, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x52, 0x41,
	0x4e, 0x43, 0x48, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48,
	0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x42, 0x52, 0x41, 0x4e, 0x43, 0x48, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x46, 0x55, 0x4c,
	0x4c, 0x10, 0x02, 0x42, 0xcf, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0b, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2,
	0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56,
	0x32, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_branch_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_branch_proto_rawDescData = file_google_cloud_retail_v2alpha_branch_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_branch_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_branch_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_branch_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_branch_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_branch_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_retail_v2alpha_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_retail_v2alpha_branch_proto_goTypes = []interface{}{
	(BranchView)(0), // 0: google.cloud.retail.v2alpha.BranchView
	(Branch_ProductCountStatistic_ProductCountScope)(0), // 1: google.cloud.retail.v2alpha.Branch.ProductCountStatistic.ProductCountScope
	(*Branch)(nil),                       // 2: google.cloud.retail.v2alpha.Branch
	(*Branch_ProductCountStatistic)(nil), // 3: google.cloud.retail.v2alpha.Branch.ProductCountStatistic
	(*Branch_QualityMetric)(nil),         // 4: google.cloud.retail.v2alpha.Branch.QualityMetric
	nil,                                  // 5: google.cloud.retail.v2alpha.Branch.ProductCountStatistic.CountsEntry
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
	(*Product)(nil),                      // 7: google.cloud.retail.v2alpha.Product
}
var file_google_cloud_retail_v2alpha_branch_proto_depIdxs = []int32{
	6, // 0: google.cloud.retail.v2alpha.Branch.last_product_import_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.cloud.retail.v2alpha.Branch.product_count_stats:type_name -> google.cloud.retail.v2alpha.Branch.ProductCountStatistic
	4, // 2: google.cloud.retail.v2alpha.Branch.quality_metrics:type_name -> google.cloud.retail.v2alpha.Branch.QualityMetric
	1, // 3: google.cloud.retail.v2alpha.Branch.ProductCountStatistic.scope:type_name -> google.cloud.retail.v2alpha.Branch.ProductCountStatistic.ProductCountScope
	5, // 4: google.cloud.retail.v2alpha.Branch.ProductCountStatistic.counts:type_name -> google.cloud.retail.v2alpha.Branch.ProductCountStatistic.CountsEntry
	7, // 5: google.cloud.retail.v2alpha.Branch.QualityMetric.unqualified_sample_products:type_name -> google.cloud.retail.v2alpha.Product
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_branch_proto_init() }
func file_google_cloud_retail_v2alpha_branch_proto_init() {
	if File_google_cloud_retail_v2alpha_branch_proto != nil {
		return
	}
	file_google_cloud_retail_v2alpha_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_branch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
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
		file_google_cloud_retail_v2alpha_branch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch_ProductCountStatistic); i {
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
		file_google_cloud_retail_v2alpha_branch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch_QualityMetric); i {
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
			RawDescriptor: file_google_cloud_retail_v2alpha_branch_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_branch_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_branch_proto_depIdxs,
		EnumInfos:         file_google_cloud_retail_v2alpha_branch_proto_enumTypes,
		MessageInfos:      file_google_cloud_retail_v2alpha_branch_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_branch_proto = out.File
	file_google_cloud_retail_v2alpha_branch_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_branch_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_branch_proto_depIdxs = nil
}