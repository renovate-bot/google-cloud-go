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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/maps/addressvalidation/v1/usps_data.proto

package addressvalidationpb

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

// USPS representation of a US address.
type UspsAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// First address line.
	FirstAddressLine string `protobuf:"bytes,1,opt,name=first_address_line,json=firstAddressLine,proto3" json:"first_address_line,omitempty"`
	// Firm name.
	Firm string `protobuf:"bytes,2,opt,name=firm,proto3" json:"firm,omitempty"`
	// Second address line.
	SecondAddressLine string `protobuf:"bytes,3,opt,name=second_address_line,json=secondAddressLine,proto3" json:"second_address_line,omitempty"`
	// Puerto Rican urbanization name.
	Urbanization string `protobuf:"bytes,4,opt,name=urbanization,proto3" json:"urbanization,omitempty"`
	// City + state + postal code.
	CityStateZipAddressLine string `protobuf:"bytes,5,opt,name=city_state_zip_address_line,json=cityStateZipAddressLine,proto3" json:"city_state_zip_address_line,omitempty"`
	// City name.
	City string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	// 2 letter state code.
	State string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	// Postal code e.g. 10009.
	ZipCode string `protobuf:"bytes,8,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	// 4-digit postal code extension e.g. 5023.
	ZipCodeExtension string `protobuf:"bytes,9,opt,name=zip_code_extension,json=zipCodeExtension,proto3" json:"zip_code_extension,omitempty"`
}

func (x *UspsAddress) Reset() {
	*x = UspsAddress{}
	mi := &file_google_maps_addressvalidation_v1_usps_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UspsAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UspsAddress) ProtoMessage() {}

func (x *UspsAddress) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_addressvalidation_v1_usps_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UspsAddress.ProtoReflect.Descriptor instead.
func (*UspsAddress) Descriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_usps_data_proto_rawDescGZIP(), []int{0}
}

func (x *UspsAddress) GetFirstAddressLine() string {
	if x != nil {
		return x.FirstAddressLine
	}
	return ""
}

func (x *UspsAddress) GetFirm() string {
	if x != nil {
		return x.Firm
	}
	return ""
}

func (x *UspsAddress) GetSecondAddressLine() string {
	if x != nil {
		return x.SecondAddressLine
	}
	return ""
}

func (x *UspsAddress) GetUrbanization() string {
	if x != nil {
		return x.Urbanization
	}
	return ""
}

func (x *UspsAddress) GetCityStateZipAddressLine() string {
	if x != nil {
		return x.CityStateZipAddressLine
	}
	return ""
}

func (x *UspsAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UspsAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UspsAddress) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *UspsAddress) GetZipCodeExtension() string {
	if x != nil {
		return x.ZipCodeExtension
	}
	return ""
}

// The USPS data for the address. `uspsData` is not guaranteed to be fully
// populated for every US or PR address sent to the Address Validation API. It's
// recommended to integrate the backup address fields in the response if you
// utilize uspsData as the primary part of the response.
type UspsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// USPS standardized address.
	StandardizedAddress *UspsAddress `protobuf:"bytes,1,opt,name=standardized_address,json=standardizedAddress,proto3" json:"standardized_address,omitempty"`
	// 2 digit delivery point code
	DeliveryPointCode string `protobuf:"bytes,2,opt,name=delivery_point_code,json=deliveryPointCode,proto3" json:"delivery_point_code,omitempty"`
	// The delivery point check digit. This number is added to the end of the
	// delivery_point_barcode for mechanically scanned mail. Adding all the
	// digits of the delivery_point_barcode, delivery_point_check_digit, postal
	// code, and ZIP+4 together should yield a number divisible by 10.
	DeliveryPointCheckDigit string `protobuf:"bytes,3,opt,name=delivery_point_check_digit,json=deliveryPointCheckDigit,proto3" json:"delivery_point_check_digit,omitempty"`
	// The possible values for DPV confirmation. Returns a single character or
	// returns no value.
	//
	// * `N`: Primary and any secondary number information failed to
	// DPV confirm.
	// * `D`: Address was DPV confirmed for the primary number only, and the
	// secondary number information was missing.
	// * `S`: Address was DPV confirmed for the primary number only, and the
	// secondary number information was present but not confirmed.
	// * `Y`: Address was DPV confirmed for primary and any secondary numbers.
	// * Empty: If the response does not contain a `dpv_confirmation` value, the
	// address was not submitted for DPV confirmation.
	DpvConfirmation string `protobuf:"bytes,4,opt,name=dpv_confirmation,json=dpvConfirmation,proto3" json:"dpv_confirmation,omitempty"`
	// The footnotes from delivery point validation.
	// Multiple footnotes may be strung together in the same string.
	//
	// * `AA`: Input address matched to the ZIP+4 file
	// * `A1`: Input address was not matched to the ZIP+4 file
	// * `BB`: Matched to DPV (all components)
	// * `CC`: Secondary number not matched and not required
	// * `C1`: Secondary number not matched but required
	// * `N1`: High-rise address missing secondary number
	// * `M1`: Primary number missing
	// * `M3`: Primary number invalid
	// * `P1`: Input address PO, RR or HC box number missing
	// * `P3`: Input address PO, RR, or HC Box number invalid
	// * `F1`: Input address matched to a military address
	// * `G1`: Input address matched to a general delivery address
	// * `U1`: Input address matched to a unique ZIP code
	// * `PB`: Input address matched to PBSA record
	// * `RR`: DPV confirmed address with PMB information
	// * `R1`: DPV confirmed address without PMB information
	// * `R7`: Carrier Route R777 or R779 record
	// * `IA`: Informed Address identified
	// * `TA`: Primary number matched by dropping a trailing alpha
	DpvFootnote string `protobuf:"bytes,5,opt,name=dpv_footnote,json=dpvFootnote,proto3" json:"dpv_footnote,omitempty"`
	// Indicates if the address is a CMRA (Commercial Mail Receiving Agency)--a
	// private business receiving mail for clients. Returns a single character.
	//
	// * `Y`: The address is a CMRA
	// * `N`: The address is not a CMRA
	DpvCmra string `protobuf:"bytes,6,opt,name=dpv_cmra,json=dpvCmra,proto3" json:"dpv_cmra,omitempty"`
	// Is this place vacant?
	// Returns a single character.
	//
	// * `Y`: The address is vacant
	// * `N`: The address is not vacant
	DpvVacant string `protobuf:"bytes,7,opt,name=dpv_vacant,json=dpvVacant,proto3" json:"dpv_vacant,omitempty"`
	// Is this a no stat address or an active address?
	// No stat addresses are ones which are not continuously occupied or addresses
	// that the USPS does not service. Returns a single character.
	//
	// * `Y`: The address is not active
	// * `N`: The address is active
	DpvNoStat string `protobuf:"bytes,8,opt,name=dpv_no_stat,json=dpvNoStat,proto3" json:"dpv_no_stat,omitempty"`
	// Indicates the NoStat type. Returns a reason code as int.
	//
	// * `1`: IDA (Internal Drop Address) – Addresses that do not receive mail
	// directly from the USPS but are delivered to a drop address that services
	// them.
	// * `2`: CDS - Addresses that have not yet become deliverable. For example, a
	// new subdivision where lots and primary numbers have been determined, but no
	// structure exists yet for occupancy.
	// * `3`: Collision - Addresses that do not actually DPV confirm.
	// * `4`: CMZ (College, Military and Other Types) - ZIP + 4 records USPS has
	// incorporated into the data.
	// * `5`: Regular - Indicates addresses not receiving delivery and the
	// addresses are not counted as possible deliveries.
	// * `6`: Secondary Required - The address requires secondary information.
	DpvNoStatReasonCode int32 `protobuf:"varint,29,opt,name=dpv_no_stat_reason_code,json=dpvNoStatReasonCode,proto3" json:"dpv_no_stat_reason_code,omitempty"`
	// Flag indicates mail is delivered to a single receptable at a site.
	// Returns a single character.
	//
	// * `Y`: The mail is delivered to a single receptable at a site.
	// * `N`: The mail is not delivered to a single receptable at a site.
	DpvDrop string `protobuf:"bytes,30,opt,name=dpv_drop,json=dpvDrop,proto3" json:"dpv_drop,omitempty"`
	// Indicates that mail is not delivered to the street address.
	// Returns a single character.
	//
	// * `Y`: The mail is not delivered to the street address.
	// * `N`: The mail is delivered to the street address.
	DpvThrowback string `protobuf:"bytes,31,opt,name=dpv_throwback,json=dpvThrowback,proto3" json:"dpv_throwback,omitempty"`
	// Flag indicates mail delivery is not performed every day of the week.
	// Returns a single character.
	//
	// * `Y`: The mail delivery is not performed every day of the week.
	// * `N`: No indication the mail delivery is not performed every day of the
	// week.
	DpvNonDeliveryDays string `protobuf:"bytes,32,opt,name=dpv_non_delivery_days,json=dpvNonDeliveryDays,proto3" json:"dpv_non_delivery_days,omitempty"`
	// Integer identifying non-delivery days. It can be interrogated using bit
	// flags:
	// 0x40 – Sunday is a non-delivery day
	// 0x20 – Monday is a non-delivery day
	// 0x10 – Tuesday is a non-delivery day
	// 0x08 – Wednesday is a non-delivery day
	// 0x04 – Thursday is a non-delivery day
	// 0x02 – Friday is a non-delivery day
	// 0x01 – Saturday is a non-delivery day
	DpvNonDeliveryDaysValues int32 `protobuf:"varint,33,opt,name=dpv_non_delivery_days_values,json=dpvNonDeliveryDaysValues,proto3" json:"dpv_non_delivery_days_values,omitempty"`
	// Flag indicates door is accessible, but package will not be left due to
	// security concerns.
	// Returns a single character.
	//
	// * `Y`: The package will not be left due to security concerns.
	// * `N`: No indication the package will not be left due to security concerns.
	DpvNoSecureLocation string `protobuf:"bytes,34,opt,name=dpv_no_secure_location,json=dpvNoSecureLocation,proto3" json:"dpv_no_secure_location,omitempty"`
	// Indicates the address was matched to PBSA record.
	// Returns a single character.
	//
	// * `Y`: The address was matched to PBSA record.
	// * `N`: The address was not matched to PBSA record.
	DpvPbsa string `protobuf:"bytes,35,opt,name=dpv_pbsa,json=dpvPbsa,proto3" json:"dpv_pbsa,omitempty"`
	// Flag indicates addresses where USPS cannot knock on a door to deliver mail.
	// Returns a single character.
	//
	// * `Y`: The door is not accessible.
	// * `N`: No indication the door is not accessible.
	DpvDoorNotAccessible string `protobuf:"bytes,36,opt,name=dpv_door_not_accessible,json=dpvDoorNotAccessible,proto3" json:"dpv_door_not_accessible,omitempty"`
	// Indicates that more than one DPV return code is valid for the address.
	// Returns a single character.
	//
	// * `Y`: Address was DPV confirmed for primary and any secondary numbers.
	// * `N`: Primary and any secondary number information failed to
	// DPV confirm.
	// * `S`: Address was DPV confirmed for the primary number only, and the
	// secondary number information was present but not confirmed,  or a single
	// trailing alpha on a primary number was dropped to make a DPV match and
	// secondary information required.
	// * `D`: Address was DPV confirmed for the primary number only, and the
	// secondary number information was missing.
	// * `R`: Address confirmed but assigned to phantom route R777 and R779 and
	// USPS delivery is not provided.
	DpvEnhancedDeliveryCode string `protobuf:"bytes,37,opt,name=dpv_enhanced_delivery_code,json=dpvEnhancedDeliveryCode,proto3" json:"dpv_enhanced_delivery_code,omitempty"`
	// The carrier route code.
	// A four character code consisting of a one letter prefix and a three digit
	// route designator.
	//
	// Prefixes:
	//
	// * `C`: Carrier route (or city route)
	// * `R`: Rural route
	// * `H`: Highway Contract Route
	// * `B`: Post Office Box Section
	// * `G`: General delivery unit
	CarrierRoute string `protobuf:"bytes,9,opt,name=carrier_route,json=carrierRoute,proto3" json:"carrier_route,omitempty"`
	// Carrier route rate sort indicator.
	CarrierRouteIndicator string `protobuf:"bytes,10,opt,name=carrier_route_indicator,json=carrierRouteIndicator,proto3" json:"carrier_route_indicator,omitempty"`
	// The delivery address is matchable, but the EWS file indicates that an exact
	// match will be available soon.
	EwsNoMatch bool `protobuf:"varint,11,opt,name=ews_no_match,json=ewsNoMatch,proto3" json:"ews_no_match,omitempty"`
	// Main post office city.
	PostOfficeCity string `protobuf:"bytes,12,opt,name=post_office_city,json=postOfficeCity,proto3" json:"post_office_city,omitempty"`
	// Main post office state.
	PostOfficeState string `protobuf:"bytes,13,opt,name=post_office_state,json=postOfficeState,proto3" json:"post_office_state,omitempty"`
	// Abbreviated city.
	AbbreviatedCity string `protobuf:"bytes,14,opt,name=abbreviated_city,json=abbreviatedCity,proto3" json:"abbreviated_city,omitempty"`
	// FIPS county code.
	FipsCountyCode string `protobuf:"bytes,15,opt,name=fips_county_code,json=fipsCountyCode,proto3" json:"fips_county_code,omitempty"`
	// County name.
	County string `protobuf:"bytes,16,opt,name=county,proto3" json:"county,omitempty"`
	// Enhanced Line of Travel (eLOT) number.
	ElotNumber string `protobuf:"bytes,17,opt,name=elot_number,json=elotNumber,proto3" json:"elot_number,omitempty"`
	// eLOT Ascending/Descending Flag (A/D).
	ElotFlag string `protobuf:"bytes,18,opt,name=elot_flag,json=elotFlag,proto3" json:"elot_flag,omitempty"`
	// LACSLink return code.
	LacsLinkReturnCode string `protobuf:"bytes,19,opt,name=lacs_link_return_code,json=lacsLinkReturnCode,proto3" json:"lacs_link_return_code,omitempty"`
	// LACSLink indicator.
	LacsLinkIndicator string `protobuf:"bytes,20,opt,name=lacs_link_indicator,json=lacsLinkIndicator,proto3" json:"lacs_link_indicator,omitempty"`
	// PO Box only postal code.
	PoBoxOnlyPostalCode bool `protobuf:"varint,21,opt,name=po_box_only_postal_code,json=poBoxOnlyPostalCode,proto3" json:"po_box_only_postal_code,omitempty"`
	// Footnotes from matching a street or highrise record to suite information.
	// If business name match is found, the secondary number is returned.
	//
	// * `A`: SuiteLink record match, business address improved.
	// * `00`: No match, business address is not improved.
	SuitelinkFootnote string `protobuf:"bytes,22,opt,name=suitelink_footnote,json=suitelinkFootnote,proto3" json:"suitelink_footnote,omitempty"`
	// PMB (Private Mail Box) unit designator.
	PmbDesignator string `protobuf:"bytes,23,opt,name=pmb_designator,json=pmbDesignator,proto3" json:"pmb_designator,omitempty"`
	// PMB (Private Mail Box) number;
	PmbNumber string `protobuf:"bytes,24,opt,name=pmb_number,json=pmbNumber,proto3" json:"pmb_number,omitempty"`
	// Type of the address record that matches the input address.
	//
	// * `F`: FIRM. This is a match to a Firm Record, which is the finest level of
	// match available for an address.
	// * `G`: GENERAL DELIVERY. This is a match to a General Delivery record.
	// * `H`: BUILDING / APARTMENT. This is a match to a Building or Apartment
	// record.
	// * `P`: POST OFFICE BOX. This is a match to a Post Office Box.
	// * `R`: RURAL ROUTE or HIGHWAY CONTRACT: This is a match to either a Rural
	// Route or a Highway Contract record, both of which may have associated Box
	// Number ranges.
	// * `S`: STREET RECORD: This is a match to a Street record containing a valid
	// primary number range.
	AddressRecordType string `protobuf:"bytes,25,opt,name=address_record_type,json=addressRecordType,proto3" json:"address_record_type,omitempty"`
	// Indicator that a default address was found, but more specific addresses
	// exists.
	DefaultAddress bool `protobuf:"varint,26,opt,name=default_address,json=defaultAddress,proto3" json:"default_address,omitempty"`
	// Error message for USPS data retrieval. This is populated when USPS
	// processing is suspended because of the detection of artificially created
	// addresses.
	//
	// The USPS data fields might not be populated when this error is present.
	ErrorMessage string `protobuf:"bytes,27,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// Indicator that the request has been CASS processed.
	CassProcessed bool `protobuf:"varint,28,opt,name=cass_processed,json=cassProcessed,proto3" json:"cass_processed,omitempty"`
}

func (x *UspsData) Reset() {
	*x = UspsData{}
	mi := &file_google_maps_addressvalidation_v1_usps_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UspsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UspsData) ProtoMessage() {}

func (x *UspsData) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_addressvalidation_v1_usps_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UspsData.ProtoReflect.Descriptor instead.
func (*UspsData) Descriptor() ([]byte, []int) {
	return file_google_maps_addressvalidation_v1_usps_data_proto_rawDescGZIP(), []int{1}
}

func (x *UspsData) GetStandardizedAddress() *UspsAddress {
	if x != nil {
		return x.StandardizedAddress
	}
	return nil
}

func (x *UspsData) GetDeliveryPointCode() string {
	if x != nil {
		return x.DeliveryPointCode
	}
	return ""
}

func (x *UspsData) GetDeliveryPointCheckDigit() string {
	if x != nil {
		return x.DeliveryPointCheckDigit
	}
	return ""
}

func (x *UspsData) GetDpvConfirmation() string {
	if x != nil {
		return x.DpvConfirmation
	}
	return ""
}

func (x *UspsData) GetDpvFootnote() string {
	if x != nil {
		return x.DpvFootnote
	}
	return ""
}

func (x *UspsData) GetDpvCmra() string {
	if x != nil {
		return x.DpvCmra
	}
	return ""
}

func (x *UspsData) GetDpvVacant() string {
	if x != nil {
		return x.DpvVacant
	}
	return ""
}

func (x *UspsData) GetDpvNoStat() string {
	if x != nil {
		return x.DpvNoStat
	}
	return ""
}

func (x *UspsData) GetDpvNoStatReasonCode() int32 {
	if x != nil {
		return x.DpvNoStatReasonCode
	}
	return 0
}

func (x *UspsData) GetDpvDrop() string {
	if x != nil {
		return x.DpvDrop
	}
	return ""
}

func (x *UspsData) GetDpvThrowback() string {
	if x != nil {
		return x.DpvThrowback
	}
	return ""
}

func (x *UspsData) GetDpvNonDeliveryDays() string {
	if x != nil {
		return x.DpvNonDeliveryDays
	}
	return ""
}

func (x *UspsData) GetDpvNonDeliveryDaysValues() int32 {
	if x != nil {
		return x.DpvNonDeliveryDaysValues
	}
	return 0
}

func (x *UspsData) GetDpvNoSecureLocation() string {
	if x != nil {
		return x.DpvNoSecureLocation
	}
	return ""
}

func (x *UspsData) GetDpvPbsa() string {
	if x != nil {
		return x.DpvPbsa
	}
	return ""
}

func (x *UspsData) GetDpvDoorNotAccessible() string {
	if x != nil {
		return x.DpvDoorNotAccessible
	}
	return ""
}

func (x *UspsData) GetDpvEnhancedDeliveryCode() string {
	if x != nil {
		return x.DpvEnhancedDeliveryCode
	}
	return ""
}

func (x *UspsData) GetCarrierRoute() string {
	if x != nil {
		return x.CarrierRoute
	}
	return ""
}

func (x *UspsData) GetCarrierRouteIndicator() string {
	if x != nil {
		return x.CarrierRouteIndicator
	}
	return ""
}

func (x *UspsData) GetEwsNoMatch() bool {
	if x != nil {
		return x.EwsNoMatch
	}
	return false
}

func (x *UspsData) GetPostOfficeCity() string {
	if x != nil {
		return x.PostOfficeCity
	}
	return ""
}

func (x *UspsData) GetPostOfficeState() string {
	if x != nil {
		return x.PostOfficeState
	}
	return ""
}

func (x *UspsData) GetAbbreviatedCity() string {
	if x != nil {
		return x.AbbreviatedCity
	}
	return ""
}

func (x *UspsData) GetFipsCountyCode() string {
	if x != nil {
		return x.FipsCountyCode
	}
	return ""
}

func (x *UspsData) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *UspsData) GetElotNumber() string {
	if x != nil {
		return x.ElotNumber
	}
	return ""
}

func (x *UspsData) GetElotFlag() string {
	if x != nil {
		return x.ElotFlag
	}
	return ""
}

func (x *UspsData) GetLacsLinkReturnCode() string {
	if x != nil {
		return x.LacsLinkReturnCode
	}
	return ""
}

func (x *UspsData) GetLacsLinkIndicator() string {
	if x != nil {
		return x.LacsLinkIndicator
	}
	return ""
}

func (x *UspsData) GetPoBoxOnlyPostalCode() bool {
	if x != nil {
		return x.PoBoxOnlyPostalCode
	}
	return false
}

func (x *UspsData) GetSuitelinkFootnote() string {
	if x != nil {
		return x.SuitelinkFootnote
	}
	return ""
}

func (x *UspsData) GetPmbDesignator() string {
	if x != nil {
		return x.PmbDesignator
	}
	return ""
}

func (x *UspsData) GetPmbNumber() string {
	if x != nil {
		return x.PmbNumber
	}
	return ""
}

func (x *UspsData) GetAddressRecordType() string {
	if x != nil {
		return x.AddressRecordType
	}
	return ""
}

func (x *UspsData) GetDefaultAddress() bool {
	if x != nil {
		return x.DefaultAddress
	}
	return false
}

func (x *UspsData) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *UspsData) GetCassProcessed() bool {
	if x != nil {
		return x.CassProcessed
	}
	return false
}

var File_google_maps_addressvalidation_v1_usps_data_proto protoreflect.FileDescriptor

var file_google_maps_addressvalidation_v1_usps_data_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x70, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0xd4, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x70, 0x73, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x66, 0x69, 0x72, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x72,
	0x62, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x1b, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x7a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x63, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5a, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x7a, 0x69, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe1, 0x0c, 0x0a, 0x08,
	0x55, 0x73, 0x70, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x60, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x70, 0x73, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x13, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x69,
	0x7a, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x44, 0x69, 0x67, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x70, 0x76, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x64, 0x70, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x70, 0x76, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x6e, 0x6f,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x70, 0x76, 0x46, 0x6f, 0x6f,
	0x74, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x70, 0x76, 0x5f, 0x63, 0x6d, 0x72,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x70, 0x76, 0x43, 0x6d, 0x72, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x70, 0x76, 0x5f, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x70, 0x76, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0b, 0x64, 0x70, 0x76, 0x5f, 0x6e, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x70, 0x76, 0x4e, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x34, 0x0a, 0x17, 0x64, 0x70, 0x76, 0x5f, 0x6e, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x13, 0x64, 0x70, 0x76, 0x4e, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x70, 0x76, 0x5f, 0x64, 0x72, 0x6f,
	0x70, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x70, 0x76, 0x44, 0x72, 0x6f, 0x70,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x70, 0x76, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x62, 0x61, 0x63,
	0x6b, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x70, 0x76, 0x54, 0x68, 0x72, 0x6f,
	0x77, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x15, 0x64, 0x70, 0x76, 0x5f, 0x6e, 0x6f, 0x6e,
	0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x70, 0x76, 0x4e, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x70, 0x76, 0x5f,
	0x6e, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x21, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18,
	0x64, 0x70, 0x76, 0x4e, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61,
	0x79, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x64, 0x70, 0x76, 0x5f,
	0x6e, 0x6f, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x70, 0x76, 0x4e, 0x6f, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x70, 0x76, 0x5f, 0x70, 0x62, 0x73, 0x61, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x70, 0x76, 0x50, 0x62, 0x73, 0x61, 0x12, 0x35, 0x0a, 0x17, 0x64, 0x70, 0x76, 0x5f,
	0x64, 0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x70, 0x76, 0x44, 0x6f,
	0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12,
	0x3b, 0x0a, 0x1a, 0x64, 0x70, 0x76, 0x5f, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x25, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x70, 0x76, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x65, 0x77, 0x73,
	0x5f, 0x6e, 0x6f, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x65, 0x77, 0x73, 0x4e, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x62, 0x62,
	0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x10,
	0x66, 0x69, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x70, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6c, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x6c, 0x6f, 0x74, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6c, 0x6f, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x31, 0x0a, 0x15,
	0x6c, 0x61, 0x63, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x63,
	0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x63, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61,
	0x63, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x34, 0x0a, 0x17, 0x70, 0x6f, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x70, 0x6f, 0x42, 0x6f, 0x78, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x75, 0x69, 0x74, 0x65, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x73, 0x75, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x6f, 0x6f, 0x74,
	0x6e, 0x6f, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6d, 0x62, 0x5f, 0x64, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6d,
	0x62, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x6d, 0x62, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x63, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x42,
	0x87, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x55, 0x73, 0x70, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d,
	0x61, 0x70, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0xa2, 0x02, 0x07, 0x47, 0x4d, 0x50, 0x41, 0x56, 0x56, 0x31, 0xaa, 0x02, 0x20,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x4d, 0x61,
	0x70, 0x73, 0x3a, 0x3a, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_maps_addressvalidation_v1_usps_data_proto_rawDescOnce sync.Once
	file_google_maps_addressvalidation_v1_usps_data_proto_rawDescData = file_google_maps_addressvalidation_v1_usps_data_proto_rawDesc
)

func file_google_maps_addressvalidation_v1_usps_data_proto_rawDescGZIP() []byte {
	file_google_maps_addressvalidation_v1_usps_data_proto_rawDescOnce.Do(func() {
		file_google_maps_addressvalidation_v1_usps_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_addressvalidation_v1_usps_data_proto_rawDescData)
	})
	return file_google_maps_addressvalidation_v1_usps_data_proto_rawDescData
}

var file_google_maps_addressvalidation_v1_usps_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_addressvalidation_v1_usps_data_proto_goTypes = []any{
	(*UspsAddress)(nil), // 0: google.maps.addressvalidation.v1.UspsAddress
	(*UspsData)(nil),    // 1: google.maps.addressvalidation.v1.UspsData
}
var file_google_maps_addressvalidation_v1_usps_data_proto_depIdxs = []int32{
	0, // 0: google.maps.addressvalidation.v1.UspsData.standardized_address:type_name -> google.maps.addressvalidation.v1.UspsAddress
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_maps_addressvalidation_v1_usps_data_proto_init() }
func file_google_maps_addressvalidation_v1_usps_data_proto_init() {
	if File_google_maps_addressvalidation_v1_usps_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_addressvalidation_v1_usps_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_addressvalidation_v1_usps_data_proto_goTypes,
		DependencyIndexes: file_google_maps_addressvalidation_v1_usps_data_proto_depIdxs,
		MessageInfos:      file_google_maps_addressvalidation_v1_usps_data_proto_msgTypes,
	}.Build()
	File_google_maps_addressvalidation_v1_usps_data_proto = out.File
	file_google_maps_addressvalidation_v1_usps_data_proto_rawDesc = nil
	file_google_maps_addressvalidation_v1_usps_data_proto_goTypes = nil
	file_google_maps_addressvalidation_v1_usps_data_proto_depIdxs = nil
}
