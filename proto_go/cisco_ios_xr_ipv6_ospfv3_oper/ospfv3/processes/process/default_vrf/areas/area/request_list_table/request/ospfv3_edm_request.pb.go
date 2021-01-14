// Copyright (c) 2015, Cisco Systems
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived
//    from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
// TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
// EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// This file is autogenerated
//
// The following edits are possible, without affecting the validity of the
// file:
//
//  * Fields may be renamed.
//  * Fields may be deleted.
//  * The unique numbered tag for a field may be changed, provided that
//    the ordering of tags for fields within a message is preserved.
//  * Message types may be renamed.
//  * Message types may be deleted (if all fields that reference them
//    have been deleted).
//
// All Cisco message and field extensions must be preserved (except when the
// field itself is being deleted).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: cisco_ios_xr_ipv6_ospfv3_oper/ospfv3/processes/process/default_vrf/areas/area/request_list_table/request/ospfv3_edm_request.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// OSPFv3 request list information
type Ospfv3EdmRequest_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName     string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId          uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName   string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
}

func (x *Ospfv3EdmRequest_KEYS) Reset() {
	*x = Ospfv3EdmRequest_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmRequest_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmRequest_KEYS) ProtoMessage() {}

func (x *Ospfv3EdmRequest_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmRequest_KEYS.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmRequest_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescGZIP(), []int{0}
}

func (x *Ospfv3EdmRequest_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *Ospfv3EdmRequest_KEYS) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Ospfv3EdmRequest_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *Ospfv3EdmRequest_KEYS) GetNeighborAddress() string {
	if x != nil {
		return x.NeighborAddress
	}
	return ""
}

type Ospfv3EdmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Neighbor IP address
	RequestNeighborAddress string `protobuf:"bytes,50,opt,name=request_neighbor_address,json=requestNeighborAddress,proto3" json:"request_neighbor_address,omitempty"`
	// If true, virtual link is requested
	IsRequestVirtualLink bool `protobuf:"varint,51,opt,name=is_request_virtual_link,json=isRequestVirtualLink,proto3" json:"is_request_virtual_link,omitempty"`
	// Request virtual link ID
	RequestVirtualLinkId uint32 `protobuf:"varint,52,opt,name=request_virtual_link_id,json=requestVirtualLinkId,proto3" json:"request_virtual_link_id,omitempty"`
	// If true, sham link is requested
	IsRequestShamLink bool `protobuf:"varint,53,opt,name=is_request_sham_link,json=isRequestShamLink,proto3" json:"is_request_sham_link,omitempty"`
	// Request sham link ID
	RequestShamLinkId uint32 `protobuf:"varint,54,opt,name=request_sham_link_id,json=requestShamLinkId,proto3" json:"request_sham_link_id,omitempty"`
	// List of request list entries
	RequestList []*Ospfv3EdmLsaSum `protobuf:"bytes,55,rep,name=request_list,json=requestList,proto3" json:"request_list,omitempty"`
}

func (x *Ospfv3EdmRequest) Reset() {
	*x = Ospfv3EdmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmRequest) ProtoMessage() {}

func (x *Ospfv3EdmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmRequest.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmRequest) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescGZIP(), []int{1}
}

func (x *Ospfv3EdmRequest) GetRequestNeighborAddress() string {
	if x != nil {
		return x.RequestNeighborAddress
	}
	return ""
}

func (x *Ospfv3EdmRequest) GetIsRequestVirtualLink() bool {
	if x != nil {
		return x.IsRequestVirtualLink
	}
	return false
}

func (x *Ospfv3EdmRequest) GetRequestVirtualLinkId() uint32 {
	if x != nil {
		return x.RequestVirtualLinkId
	}
	return 0
}

func (x *Ospfv3EdmRequest) GetIsRequestShamLink() bool {
	if x != nil {
		return x.IsRequestShamLink
	}
	return false
}

func (x *Ospfv3EdmRequest) GetRequestShamLinkId() uint32 {
	if x != nil {
		return x.RequestShamLinkId
	}
	return 0
}

func (x *Ospfv3EdmRequest) GetRequestList() []*Ospfv3EdmLsaSum {
	if x != nil {
		return x.RequestList
	}
	return nil
}

// LSA summary entry
type Ospfv3EdmLsaSum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LSA type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	// Age of the LSA (seconds)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsaId string `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId,proto3" json:"header_lsa_id,omitempty"`
	// Router ID of the advertising router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber int32 `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
}

func (x *Ospfv3EdmLsaSum) Reset() {
	*x = Ospfv3EdmLsaSum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmLsaSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmLsaSum) ProtoMessage() {}

func (x *Ospfv3EdmLsaSum) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmLsaSum.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescGZIP(), []int{2}
}

func (x *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if x != nil {
		return x.HeaderLsaType
	}
	return ""
}

func (x *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if x != nil {
		return x.HeaderLsaAge
	}
	return 0
}

func (x *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if x != nil {
		return x.HeaderLsaId
	}
	return ""
}

func (x *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if x != nil {
		return x.HeaderAdvertisingRouter
	}
	return ""
}

func (x *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if x != nil {
		return x.HeaderSequenceNumber
	}
	return 0
}

var File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDesc = []byte{
	0x0a, 0x81, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x76, 0x72, 0x66, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2f, 0x61, 0x72, 0x65, 0x61,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x73, 0x70, 0x66,
	0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x68, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f,
	0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x61,
	0x72, 0x65, 0x61, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa7,
	0x01, 0x0a, 0x17, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc0, 0x03, 0x0a, 0x12, 0x6f, 0x73, 0x70,
	0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62,
	0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x33, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x35, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x34, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x35, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x68, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x68, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x9f, 0x01, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x37, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x7c, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x65, 0x61,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x73, 0x70, 0x66,
	0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73, 0x75, 0x6d, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xf8, 0x01, 0x0a, 0x12,
	0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73,
	0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x41, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c,
	0x73, 0x61, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x16, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x14, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescData = file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDesc
)

func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDescData
}

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_goTypes = []interface{}{
	(*Ospfv3EdmRequest_KEYS)(nil), // 0: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.request_list_table.request.ospfv3_edm_request_KEYS
	(*Ospfv3EdmRequest)(nil),      // 1: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.request_list_table.request.ospfv3_edm_request
	(*Ospfv3EdmLsaSum)(nil),       // 2: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.request_list_table.request.ospfv3_edm_lsa_sum
}
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.request_list_table.request.ospfv3_edm_request.request_list:type_name -> cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.request_list_table.request.ospfv3_edm_lsa_sum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_init()
}
func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_init() {
	if File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmRequest_KEYS); i {
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
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmRequest); i {
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
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmLsaSum); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto = out.File
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_rawDesc = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_goTypes = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_request_list_table_request_ospfv3_edm_request_proto_depIdxs = nil
}
