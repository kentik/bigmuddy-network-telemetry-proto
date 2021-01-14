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
// source: cisco_ios_xr_ipv6_ospfv3_oper/ospfv3/processes/process/default_vrf/areas/area/retransmission_list_table/retransmission/ospfv3_edm_retrans.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission

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

// OSPFv3 retransmission list information
type Ospfv3EdmRetrans_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName     string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId          uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName   string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
}

func (x *Ospfv3EdmRetrans_KEYS) Reset() {
	*x = Ospfv3EdmRetrans_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmRetrans_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmRetrans_KEYS) ProtoMessage() {}

func (x *Ospfv3EdmRetrans_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmRetrans_KEYS.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmRetrans_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescGZIP(), []int{0}
}

func (x *Ospfv3EdmRetrans_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *Ospfv3EdmRetrans_KEYS) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Ospfv3EdmRetrans_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *Ospfv3EdmRetrans_KEYS) GetNeighborAddress() string {
	if x != nil {
		return x.NeighborAddress
	}
	return ""
}

type Ospfv3EdmRetrans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Neighbor IP address
	RetransmissionNeighborAddress string `protobuf:"bytes,50,opt,name=retransmission_neighbor_address,json=retransmissionNeighborAddress,proto3" json:"retransmission_neighbor_address,omitempty"`
	// If true, virtual link is retransmitted
	IsRetransmissionVirtualLink bool `protobuf:"varint,51,opt,name=is_retransmission_virtual_link,json=isRetransmissionVirtualLink,proto3" json:"is_retransmission_virtual_link,omitempty"`
	// Retransmission virtual link ID
	RetransmissionVirtualLinkId uint32 `protobuf:"varint,52,opt,name=retransmission_virtual_link_id,json=retransmissionVirtualLinkId,proto3" json:"retransmission_virtual_link_id,omitempty"`
	// If true, sham link is retransmitted
	IsRetransmissionShamLink bool `protobuf:"varint,53,opt,name=is_retransmission_sham_link,json=isRetransmissionShamLink,proto3" json:"is_retransmission_sham_link,omitempty"`
	// Retransmission sham link ID
	RetransmissionShamLinkId uint32 `protobuf:"varint,54,opt,name=retransmission_sham_link_id,json=retransmissionShamLinkId,proto3" json:"retransmission_sham_link_id,omitempty"`
	// Amount of time remaining on retransmission timer (ms)
	RetransmissionTimer uint32 `protobuf:"varint,55,opt,name=retransmission_timer,json=retransmissionTimer,proto3" json:"retransmission_timer,omitempty"`
	// Retransmission queue length
	RetransmissionLength uint32 `protobuf:"varint,56,opt,name=retransmission_length,json=retransmissionLength,proto3" json:"retransmission_length,omitempty"`
	// List of virtual link scope entries
	RetransmissionVirtualLinkDbList []*Ospfv3EdmLsaSum `protobuf:"bytes,57,rep,name=retransmission_virtual_link_db_list,json=retransmissionVirtualLinkDbList,proto3" json:"retransmission_virtual_link_db_list,omitempty"`
	// List of area scope entries
	RetransmissionAreaDbList []*Ospfv3EdmLsaSum `protobuf:"bytes,58,rep,name=retransmission_area_db_list,json=retransmissionAreaDbList,proto3" json:"retransmission_area_db_list,omitempty"`
	// List of AS scope entries
	RetransmissionAsdbList []*Ospfv3EdmLsaSum `protobuf:"bytes,59,rep,name=retransmission_asdb_list,json=retransmissionAsdbList,proto3" json:"retransmission_asdb_list,omitempty"`
}

func (x *Ospfv3EdmRetrans) Reset() {
	*x = Ospfv3EdmRetrans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmRetrans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmRetrans) ProtoMessage() {}

func (x *Ospfv3EdmRetrans) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmRetrans.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmRetrans) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescGZIP(), []int{1}
}

func (x *Ospfv3EdmRetrans) GetRetransmissionNeighborAddress() string {
	if x != nil {
		return x.RetransmissionNeighborAddress
	}
	return ""
}

func (x *Ospfv3EdmRetrans) GetIsRetransmissionVirtualLink() bool {
	if x != nil {
		return x.IsRetransmissionVirtualLink
	}
	return false
}

func (x *Ospfv3EdmRetrans) GetRetransmissionVirtualLinkId() uint32 {
	if x != nil {
		return x.RetransmissionVirtualLinkId
	}
	return 0
}

func (x *Ospfv3EdmRetrans) GetIsRetransmissionShamLink() bool {
	if x != nil {
		return x.IsRetransmissionShamLink
	}
	return false
}

func (x *Ospfv3EdmRetrans) GetRetransmissionShamLinkId() uint32 {
	if x != nil {
		return x.RetransmissionShamLinkId
	}
	return 0
}

func (x *Ospfv3EdmRetrans) GetRetransmissionTimer() uint32 {
	if x != nil {
		return x.RetransmissionTimer
	}
	return 0
}

func (x *Ospfv3EdmRetrans) GetRetransmissionLength() uint32 {
	if x != nil {
		return x.RetransmissionLength
	}
	return 0
}

func (x *Ospfv3EdmRetrans) GetRetransmissionVirtualLinkDbList() []*Ospfv3EdmLsaSum {
	if x != nil {
		return x.RetransmissionVirtualLinkDbList
	}
	return nil
}

func (x *Ospfv3EdmRetrans) GetRetransmissionAreaDbList() []*Ospfv3EdmLsaSum {
	if x != nil {
		return x.RetransmissionAreaDbList
	}
	return nil
}

func (x *Ospfv3EdmRetrans) GetRetransmissionAsdbList() []*Ospfv3EdmLsaSum {
	if x != nil {
		return x.RetransmissionAsdbList
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
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmLsaSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmLsaSum) ProtoMessage() {}

func (x *Ospfv3EdmLsaSum) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[2]
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
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescGZIP(), []int{2}
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

var File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDesc = []byte{
	0x0a, 0x8f, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x76, 0x72, 0x66, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2f, 0x61, 0x72, 0x65, 0x61,
	0x2f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33,
	0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x76, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72,
	0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x65,
	0x61, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x6f,
	0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6e, 0x65, 0x69,
	0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0xbd, 0x08, 0x0a, 0x12, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f,
	0x65, 0x64, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x1f, 0x72,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x65,
	0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x33, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x69, 0x73, 0x52,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x43, 0x0a, 0x1e, 0x72, 0x65, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x1b, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x1b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x35, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x18, 0x69, 0x73, 0x52, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x3d, 0x0a, 0x1b,
	0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x68, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x36, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x18, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x68, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x72,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x72, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x72, 0x65, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x33,
	0x0a, 0x15, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x72,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0xd9, 0x01, 0x0a, 0x23, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x64, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x39, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x8a, 0x01, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78,
	0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x61, 0x72,
	0x65, 0x61, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x73, 0x70, 0x66,
	0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73, 0x75, 0x6d, 0x52, 0x1f,
	0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0xca, 0x01, 0x0a, 0x1b, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x64, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x3a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x8a, 0x01, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69,
	0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76,
	0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61,
	0x73, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73,
	0x75, 0x6d, 0x52, 0x18, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x44, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xc5, 0x01, 0x0a,
	0x18, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x73, 0x64, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x3b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x8a, 0x01, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x65, 0x61,
	0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33,
	0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73, 0x75, 0x6d, 0x52, 0x16, 0x72, 0x65,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x64, 0x62,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0xf8, 0x01, 0x0a, 0x12, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f,
	0x65, 0x64, 0x6d, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73,
	0x61, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x41, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x49, 0x64, 0x12, 0x3a, 0x0a,
	0x19, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x14, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescData = file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDesc
)

func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDescData
}

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_goTypes = []interface{}{
	(*Ospfv3EdmRetrans_KEYS)(nil), // 0: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans_KEYS
	(*Ospfv3EdmRetrans)(nil),      // 1: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans
	(*Ospfv3EdmLsaSum)(nil),       // 2: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_lsa_sum
}
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans.retransmission_virtual_link_db_list:type_name -> cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_lsa_sum
	2, // 1: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans.retransmission_area_db_list:type_name -> cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_lsa_sum
	2, // 2: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_retrans.retransmission_asdb_list:type_name -> cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.retransmission_list_table.retransmission.ospfv3_edm_lsa_sum
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_init()
}
func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_init() {
	if File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmRetrans_KEYS); i {
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
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmRetrans); i {
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
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto = out.File
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_rawDesc = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_goTypes = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_retransmission_list_table_retransmission_ospfv3_edm_retrans_proto_depIdxs = nil
}
