// Copyright 2025 Google LLC
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/cloud/managedidentities/v1/resource.proto

package managedidentitiespb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the different states of a managed domain.
type Domain_State int32

const (
	// Not set.
	Domain_STATE_UNSPECIFIED Domain_State = 0
	// The domain is being created.
	Domain_CREATING Domain_State = 1
	// The domain has been created and is fully usable.
	Domain_READY Domain_State = 2
	// The domain's configuration is being updated.
	Domain_UPDATING Domain_State = 3
	// The domain is being deleted.
	Domain_DELETING Domain_State = 4
	// The domain is being repaired and may be unusable. Details
	// can be found in the `status_message` field.
	Domain_REPAIRING Domain_State = 5
	// The domain is undergoing maintenance.
	Domain_PERFORMING_MAINTENANCE Domain_State = 6
	// The domain is not serving requests.
	Domain_UNAVAILABLE Domain_State = 7
)

// Enum value maps for Domain_State.
var (
	Domain_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "UPDATING",
		4: "DELETING",
		5: "REPAIRING",
		6: "PERFORMING_MAINTENANCE",
		7: "UNAVAILABLE",
	}
	Domain_State_value = map[string]int32{
		"STATE_UNSPECIFIED":      0,
		"CREATING":               1,
		"READY":                  2,
		"UPDATING":               3,
		"DELETING":               4,
		"REPAIRING":              5,
		"PERFORMING_MAINTENANCE": 6,
		"UNAVAILABLE":            7,
	}
)

func (x Domain_State) Enum() *Domain_State {
	p := new(Domain_State)
	*p = x
	return p
}

func (x Domain_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Domain_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_managedidentities_v1_resource_proto_enumTypes[0].Descriptor()
}

func (Domain_State) Type() protoreflect.EnumType {
	return &file_google_cloud_managedidentities_v1_resource_proto_enumTypes[0]
}

func (x Domain_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Domain_State.Descriptor instead.
func (Domain_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP(), []int{0, 0}
}

// Represents the different states of a domain trust.
type Trust_State int32

const (
	// Not set.
	Trust_STATE_UNSPECIFIED Trust_State = 0
	// The domain trust is being created.
	Trust_CREATING Trust_State = 1
	// The domain trust is being updated.
	Trust_UPDATING Trust_State = 2
	// The domain trust is being deleted.
	Trust_DELETING Trust_State = 3
	// The domain trust is connected.
	Trust_CONNECTED Trust_State = 4
	// The domain trust is disconnected.
	Trust_DISCONNECTED Trust_State = 5
)

// Enum value maps for Trust_State.
var (
	Trust_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "UPDATING",
		3: "DELETING",
		4: "CONNECTED",
		5: "DISCONNECTED",
	}
	Trust_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"UPDATING":          2,
		"DELETING":          3,
		"CONNECTED":         4,
		"DISCONNECTED":      5,
	}
)

func (x Trust_State) Enum() *Trust_State {
	p := new(Trust_State)
	*p = x
	return p
}

func (x Trust_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Trust_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_managedidentities_v1_resource_proto_enumTypes[1].Descriptor()
}

func (Trust_State) Type() protoreflect.EnumType {
	return &file_google_cloud_managedidentities_v1_resource_proto_enumTypes[1]
}

func (x Trust_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Trust_State.Descriptor instead.
func (Trust_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP(), []int{1, 0}
}

// Represents the different inter-forest trust types.
type Trust_TrustType int32

const (
	// Not set.
	Trust_TRUST_TYPE_UNSPECIFIED Trust_TrustType = 0
	// The forest trust.
	Trust_FOREST Trust_TrustType = 1
	// The external domain trust.
	Trust_EXTERNAL Trust_TrustType = 2
)

// Enum value maps for Trust_TrustType.
var (
	Trust_TrustType_name = map[int32]string{
		0: "TRUST_TYPE_UNSPECIFIED",
		1: "FOREST",
		2: "EXTERNAL",
	}
	Trust_TrustType_value = map[string]int32{
		"TRUST_TYPE_UNSPECIFIED": 0,
		"FOREST":                 1,
		"EXTERNAL":               2,
	}
)

func (x Trust_TrustType) Enum() *Trust_TrustType {
	p := new(Trust_TrustType)
	*p = x
	return p
}

func (x Trust_TrustType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Trust_TrustType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_managedidentities_v1_resource_proto_enumTypes[2].Descriptor()
}

func (Trust_TrustType) Type() protoreflect.EnumType {
	return &file_google_cloud_managedidentities_v1_resource_proto_enumTypes[2]
}

func (x Trust_TrustType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Trust_TrustType.Descriptor instead.
func (Trust_TrustType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP(), []int{1, 1}
}

// Represents the direction of trust.
// See
// [System.DirectoryServices.ActiveDirectory.TrustDirection](https://docs.microsoft.com/en-us/dotnet/api/system.directoryservices.activedirectory.trustdirection?view=netframework-4.7.2)
// for more information.
type Trust_TrustDirection int32

const (
	// Not set.
	Trust_TRUST_DIRECTION_UNSPECIFIED Trust_TrustDirection = 0
	// The inbound direction represents the trusting side.
	Trust_INBOUND Trust_TrustDirection = 1
	// The outboud direction represents the trusted side.
	Trust_OUTBOUND Trust_TrustDirection = 2
	// The bidirectional direction represents the trusted / trusting side.
	Trust_BIDIRECTIONAL Trust_TrustDirection = 3
)

// Enum value maps for Trust_TrustDirection.
var (
	Trust_TrustDirection_name = map[int32]string{
		0: "TRUST_DIRECTION_UNSPECIFIED",
		1: "INBOUND",
		2: "OUTBOUND",
		3: "BIDIRECTIONAL",
	}
	Trust_TrustDirection_value = map[string]int32{
		"TRUST_DIRECTION_UNSPECIFIED": 0,
		"INBOUND":                     1,
		"OUTBOUND":                    2,
		"BIDIRECTIONAL":               3,
	}
)

func (x Trust_TrustDirection) Enum() *Trust_TrustDirection {
	p := new(Trust_TrustDirection)
	*p = x
	return p
}

func (x Trust_TrustDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Trust_TrustDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_managedidentities_v1_resource_proto_enumTypes[3].Descriptor()
}

func (Trust_TrustDirection) Type() protoreflect.EnumType {
	return &file_google_cloud_managedidentities_v1_resource_proto_enumTypes[3]
}

func (x Trust_TrustDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Trust_TrustDirection.Descriptor instead.
func (Trust_TrustDirection) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP(), []int{1, 2}
}

// Represents a managed Microsoft Active Directory domain.
type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The unique name of the domain using the form:
	// `projects/{project_id}/locations/global/domains/{domain_name}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Resource labels that can contain user-provided metadata.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. The full names of the Google Compute Engine
	// [networks](/compute/docs/networks-and-firewalls#networks) the domain
	// instance is connected to. Networks can be added using UpdateDomain.
	// The domain is only available on networks listed in `authorized_networks`.
	// If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks []string `protobuf:"bytes,3,rep,name=authorized_networks,json=authorizedNetworks,proto3" json:"authorized_networks,omitempty"`
	// Required. The CIDR range of internal addresses that are reserved for this
	// domain. Reserved networks must be /24 or larger. Ranges must be
	// unique and non-overlapping with existing subnets in
	// [Domain].[authorized_networks].
	ReservedIpRange string `protobuf:"bytes,4,opt,name=reserved_ip_range,json=reservedIpRange,proto3" json:"reserved_ip_range,omitempty"`
	// Required. Locations where domain needs to be provisioned.
	// [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4
	// Service supports up to 4 locations at once. Each location will use a /26
	// block.
	Locations []string `protobuf:"bytes,5,rep,name=locations,proto3" json:"locations,omitempty"`
	// Optional. The name of delegated administrator account used to perform
	// Active Directory operations. If not specified, `setupadmin` will be used.
	Admin string `protobuf:"bytes,6,opt,name=admin,proto3" json:"admin,omitempty"`
	// Output only. The fully-qualified domain name of the exposed domain used by
	// clients to connect to the service. Similar to what would be chosen for an
	// Active Directory set up on an internal network.
	Fqdn string `protobuf:"bytes,10,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Output only. The time the instance was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last update time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. The current state of this domain.
	State Domain_State `protobuf:"varint,13,opt,name=state,proto3,enum=google.cloud.managedidentities.v1.Domain_State" json:"state,omitempty"`
	// Output only. Additional information about the current status of this
	// domain, if available.
	StatusMessage string `protobuf:"bytes,14,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	// Output only. The current trusts associated with the domain.
	Trusts []*Trust `protobuf:"bytes,15,rep,name=trusts,proto3" json:"trusts,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	mi := &file_google_cloud_managedidentities_v1_resource_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_managedidentities_v1_resource_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Domain) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Domain) GetAuthorizedNetworks() []string {
	if x != nil {
		return x.AuthorizedNetworks
	}
	return nil
}

func (x *Domain) GetReservedIpRange() string {
	if x != nil {
		return x.ReservedIpRange
	}
	return ""
}

func (x *Domain) GetLocations() []string {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *Domain) GetAdmin() string {
	if x != nil {
		return x.Admin
	}
	return ""
}

func (x *Domain) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *Domain) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Domain) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Domain) GetState() Domain_State {
	if x != nil {
		return x.State
	}
	return Domain_STATE_UNSPECIFIED
}

func (x *Domain) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *Domain) GetTrusts() []*Trust {
	if x != nil {
		return x.Trusts
	}
	return nil
}

// Represents a relationship between two domains. This allows a controller in
// one domain to authenticate a user in another domain.
type Trust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The fully qualified target domain name which will be in trust with the
	// current domain.
	TargetDomainName string `protobuf:"bytes,1,opt,name=target_domain_name,json=targetDomainName,proto3" json:"target_domain_name,omitempty"`
	// Required. The type of trust represented by the trust resource.
	TrustType Trust_TrustType `protobuf:"varint,2,opt,name=trust_type,json=trustType,proto3,enum=google.cloud.managedidentities.v1.Trust_TrustType" json:"trust_type,omitempty"`
	// Required. The trust direction, which decides if the current domain is trusted,
	// trusting, or both.
	TrustDirection Trust_TrustDirection `protobuf:"varint,3,opt,name=trust_direction,json=trustDirection,proto3,enum=google.cloud.managedidentities.v1.Trust_TrustDirection" json:"trust_direction,omitempty"`
	// Optional. The trust authentication type, which decides whether the trusted side has
	// forest/domain wide access or selective access to an approved set of
	// resources.
	SelectiveAuthentication bool `protobuf:"varint,4,opt,name=selective_authentication,json=selectiveAuthentication,proto3" json:"selective_authentication,omitempty"`
	// Required. The target DNS server IP addresses which can resolve the remote domain
	// involved in the trust.
	TargetDnsIpAddresses []string `protobuf:"bytes,5,rep,name=target_dns_ip_addresses,json=targetDnsIpAddresses,proto3" json:"target_dns_ip_addresses,omitempty"`
	// Required. The trust secret used for the handshake with the target domain. This will
	// not be stored.
	TrustHandshakeSecret string `protobuf:"bytes,6,opt,name=trust_handshake_secret,json=trustHandshakeSecret,proto3" json:"trust_handshake_secret,omitempty"`
	// Output only. The time the instance was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last update time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. The current state of the trust.
	State Trust_State `protobuf:"varint,9,opt,name=state,proto3,enum=google.cloud.managedidentities.v1.Trust_State" json:"state,omitempty"`
	// Output only. Additional information about the current state of the trust, if available.
	StateDescription string `protobuf:"bytes,11,opt,name=state_description,json=stateDescription,proto3" json:"state_description,omitempty"`
	// Output only. The last heartbeat time when the trust was known to be connected.
	LastTrustHeartbeatTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=last_trust_heartbeat_time,json=lastTrustHeartbeatTime,proto3" json:"last_trust_heartbeat_time,omitempty"`
}

func (x *Trust) Reset() {
	*x = Trust{}
	mi := &file_google_cloud_managedidentities_v1_resource_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trust) ProtoMessage() {}

func (x *Trust) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_managedidentities_v1_resource_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trust.ProtoReflect.Descriptor instead.
func (*Trust) Descriptor() ([]byte, []int) {
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Trust) GetTargetDomainName() string {
	if x != nil {
		return x.TargetDomainName
	}
	return ""
}

func (x *Trust) GetTrustType() Trust_TrustType {
	if x != nil {
		return x.TrustType
	}
	return Trust_TRUST_TYPE_UNSPECIFIED
}

func (x *Trust) GetTrustDirection() Trust_TrustDirection {
	if x != nil {
		return x.TrustDirection
	}
	return Trust_TRUST_DIRECTION_UNSPECIFIED
}

func (x *Trust) GetSelectiveAuthentication() bool {
	if x != nil {
		return x.SelectiveAuthentication
	}
	return false
}

func (x *Trust) GetTargetDnsIpAddresses() []string {
	if x != nil {
		return x.TargetDnsIpAddresses
	}
	return nil
}

func (x *Trust) GetTrustHandshakeSecret() string {
	if x != nil {
		return x.TrustHandshakeSecret
	}
	return ""
}

func (x *Trust) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Trust) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Trust) GetState() Trust_State {
	if x != nil {
		return x.State
	}
	return Trust_STATE_UNSPECIFIED
}

func (x *Trust) GetStateDescription() string {
	if x != nil {
		return x.StateDescription
	}
	return ""
}

func (x *Trust) GetLastTrustHeartbeatTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTrustHeartbeatTime
	}
	return nil
}

var File_google_cloud_managedidentities_v1_resource_proto protoreflect.FileDescriptor

var file_google_cloud_managedidentities_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x07, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x13, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x12, 0x2f, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x49, 0x70, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x17, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x06, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x45, 0x50, 0x41, 0x49, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x50,
	0x45, 0x52, 0x46, 0x4f, 0x52, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x45,
	0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x07, 0x3a, 0x66, 0xea, 0x41, 0x63, 0x0a, 0x27, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x38, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x7d,
	0x22, 0x9c, 0x08, 0x0a, 0x05, 0x54, 0x72, 0x75, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x12, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x56, 0x0a,
	0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x65, 0x0a, 0x0f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x18,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x17, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x17,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x6e, 0x73, 0x49, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x16, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x14, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x30, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x19, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x69, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x05, 0x22, 0x41, 0x0a, 0x09, 0x54,
	0x72, 0x75, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52, 0x55, 0x53,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x52, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x22, 0x5f,
	0x0a, 0x0e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x4f, 0x55, 0x54, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x42, 0x49, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x42,
	0xfc, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x70, 0x62, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_managedidentities_v1_resource_proto_rawDescOnce sync.Once
	file_google_cloud_managedidentities_v1_resource_proto_rawDescData = file_google_cloud_managedidentities_v1_resource_proto_rawDesc
)

func file_google_cloud_managedidentities_v1_resource_proto_rawDescGZIP() []byte {
	file_google_cloud_managedidentities_v1_resource_proto_rawDescOnce.Do(func() {
		file_google_cloud_managedidentities_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_managedidentities_v1_resource_proto_rawDescData)
	})
	return file_google_cloud_managedidentities_v1_resource_proto_rawDescData
}

var file_google_cloud_managedidentities_v1_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_google_cloud_managedidentities_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_managedidentities_v1_resource_proto_goTypes = []any{
	(Domain_State)(0),             // 0: google.cloud.managedidentities.v1.Domain.State
	(Trust_State)(0),              // 1: google.cloud.managedidentities.v1.Trust.State
	(Trust_TrustType)(0),          // 2: google.cloud.managedidentities.v1.Trust.TrustType
	(Trust_TrustDirection)(0),     // 3: google.cloud.managedidentities.v1.Trust.TrustDirection
	(*Domain)(nil),                // 4: google.cloud.managedidentities.v1.Domain
	(*Trust)(nil),                 // 5: google.cloud.managedidentities.v1.Trust
	nil,                           // 6: google.cloud.managedidentities.v1.Domain.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_google_cloud_managedidentities_v1_resource_proto_depIdxs = []int32{
	6,  // 0: google.cloud.managedidentities.v1.Domain.labels:type_name -> google.cloud.managedidentities.v1.Domain.LabelsEntry
	7,  // 1: google.cloud.managedidentities.v1.Domain.create_time:type_name -> google.protobuf.Timestamp
	7,  // 2: google.cloud.managedidentities.v1.Domain.update_time:type_name -> google.protobuf.Timestamp
	0,  // 3: google.cloud.managedidentities.v1.Domain.state:type_name -> google.cloud.managedidentities.v1.Domain.State
	5,  // 4: google.cloud.managedidentities.v1.Domain.trusts:type_name -> google.cloud.managedidentities.v1.Trust
	2,  // 5: google.cloud.managedidentities.v1.Trust.trust_type:type_name -> google.cloud.managedidentities.v1.Trust.TrustType
	3,  // 6: google.cloud.managedidentities.v1.Trust.trust_direction:type_name -> google.cloud.managedidentities.v1.Trust.TrustDirection
	7,  // 7: google.cloud.managedidentities.v1.Trust.create_time:type_name -> google.protobuf.Timestamp
	7,  // 8: google.cloud.managedidentities.v1.Trust.update_time:type_name -> google.protobuf.Timestamp
	1,  // 9: google.cloud.managedidentities.v1.Trust.state:type_name -> google.cloud.managedidentities.v1.Trust.State
	7,  // 10: google.cloud.managedidentities.v1.Trust.last_trust_heartbeat_time:type_name -> google.protobuf.Timestamp
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_cloud_managedidentities_v1_resource_proto_init() }
func file_google_cloud_managedidentities_v1_resource_proto_init() {
	if File_google_cloud_managedidentities_v1_resource_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_managedidentities_v1_resource_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_managedidentities_v1_resource_proto_goTypes,
		DependencyIndexes: file_google_cloud_managedidentities_v1_resource_proto_depIdxs,
		EnumInfos:         file_google_cloud_managedidentities_v1_resource_proto_enumTypes,
		MessageInfos:      file_google_cloud_managedidentities_v1_resource_proto_msgTypes,
	}.Build()
	File_google_cloud_managedidentities_v1_resource_proto = out.File
	file_google_cloud_managedidentities_v1_resource_proto_rawDesc = nil
	file_google_cloud_managedidentities_v1_resource_proto_goTypes = nil
	file_google_cloud_managedidentities_v1_resource_proto_depIdxs = nil
}
