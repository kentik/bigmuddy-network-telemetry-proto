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
// source: cisco_ios_xr_ipv4_ospf_oper/ospf/processes/process/default_vrf/interface_information/interface_briefs/interface_brief/ospf_sh_if_brief.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief

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

// OSPF Interface Brief Information
type OspfShIfBrief_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *OspfShIfBrief_KEYS) Reset() {
	*x = OspfShIfBrief_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShIfBrief_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShIfBrief_KEYS) ProtoMessage() {}

func (x *OspfShIfBrief_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShIfBrief_KEYS.ProtoReflect.Descriptor instead.
func (*OspfShIfBrief_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescGZIP(), []int{0}
}

func (x *OspfShIfBrief_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *OspfShIfBrief_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

type OspfShIfBrief struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interface name
	InterfaceName string `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,51,opt,name=interface_area,json=interfaceArea,proto3" json:"interface_area,omitempty"`
	// Interface IP Address
	InterfaceAddress string `protobuf:"bytes,52,opt,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	// Interface IP Mask
	InterfaceMask uint32 `protobuf:"varint,53,opt,name=interface_mask,json=interfaceMask,proto3" json:"interface_mask,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,54,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,55,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	// Interface in fast detect hold down state
	InterfaceFastDetectHoldDown bool `protobuf:"varint,56,opt,name=interface_fast_detect_hold_down,json=interfaceFastDetectHoldDown,proto3" json:"interface_fast_detect_hold_down,omitempty"`
	// Total number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,57,opt,name=interface_neighbor_count,json=interfaceNeighborCount,proto3" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,58,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount,proto3" json:"interface_adj_neighbor_count,omitempty"`
	// If true, interface is multi-area
	InterfaceisMadj bool `protobuf:"varint,59,opt,name=interfaceis_madj,json=interfaceisMadj,proto3" json:"interfaceis_madj,omitempty"`
	// Total number of multi-area
	InterfaceMadjCount uint32 `protobuf:"varint,60,opt,name=interface_madj_count,json=interfaceMadjCount,proto3" json:"interface_madj_count,omitempty"`
	// Information for multi-area on the interface
	InterfaceMadjList []*OspfShInterfaceMadj `protobuf:"bytes,61,rep,name=interface_madj_list,json=interfaceMadjList,proto3" json:"interface_madj_list,omitempty"`
}

func (x *OspfShIfBrief) Reset() {
	*x = OspfShIfBrief{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShIfBrief) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShIfBrief) ProtoMessage() {}

func (x *OspfShIfBrief) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShIfBrief.ProtoReflect.Descriptor instead.
func (*OspfShIfBrief) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescGZIP(), []int{1}
}

func (x *OspfShIfBrief) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *OspfShIfBrief) GetInterfaceArea() string {
	if x != nil {
		return x.InterfaceArea
	}
	return ""
}

func (x *OspfShIfBrief) GetInterfaceAddress() string {
	if x != nil {
		return x.InterfaceAddress
	}
	return ""
}

func (x *OspfShIfBrief) GetInterfaceMask() uint32 {
	if x != nil {
		return x.InterfaceMask
	}
	return 0
}

func (x *OspfShIfBrief) GetInterfaceLinkCost() uint32 {
	if x != nil {
		return x.InterfaceLinkCost
	}
	return 0
}

func (x *OspfShIfBrief) GetOspfInterfaceState() string {
	if x != nil {
		return x.OspfInterfaceState
	}
	return ""
}

func (x *OspfShIfBrief) GetInterfaceFastDetectHoldDown() bool {
	if x != nil {
		return x.InterfaceFastDetectHoldDown
	}
	return false
}

func (x *OspfShIfBrief) GetInterfaceNeighborCount() uint32 {
	if x != nil {
		return x.InterfaceNeighborCount
	}
	return 0
}

func (x *OspfShIfBrief) GetInterfaceAdjNeighborCount() uint32 {
	if x != nil {
		return x.InterfaceAdjNeighborCount
	}
	return 0
}

func (x *OspfShIfBrief) GetInterfaceisMadj() bool {
	if x != nil {
		return x.InterfaceisMadj
	}
	return false
}

func (x *OspfShIfBrief) GetInterfaceMadjCount() uint32 {
	if x != nil {
		return x.InterfaceMadjCount
	}
	return 0
}

func (x *OspfShIfBrief) GetInterfaceMadjList() []*OspfShInterfaceMadj {
	if x != nil {
		return x.InterfaceMadjList
	}
	return nil
}

// OSPF Interface Multi-Area Information
type OspfShInterfaceMadj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,1,opt,name=interface_area,json=interfaceArea,proto3" json:"interface_area,omitempty"`
	// Area ID
	MadjAreaId uint32 `protobuf:"varint,2,opt,name=madj_area_id,json=madjAreaId,proto3" json:"madj_area_id,omitempty"`
	// Number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,3,opt,name=interface_neighbor_count,json=interfaceNeighborCount,proto3" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,4,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount,proto3" json:"interface_adj_neighbor_count,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,5,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,6,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
}

func (x *OspfShInterfaceMadj) Reset() {
	*x = OspfShInterfaceMadj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShInterfaceMadj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShInterfaceMadj) ProtoMessage() {}

func (x *OspfShInterfaceMadj) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShInterfaceMadj.ProtoReflect.Descriptor instead.
func (*OspfShInterfaceMadj) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescGZIP(), []int{2}
}

func (x *OspfShInterfaceMadj) GetInterfaceArea() string {
	if x != nil {
		return x.InterfaceArea
	}
	return ""
}

func (x *OspfShInterfaceMadj) GetMadjAreaId() uint32 {
	if x != nil {
		return x.MadjAreaId
	}
	return 0
}

func (x *OspfShInterfaceMadj) GetInterfaceNeighborCount() uint32 {
	if x != nil {
		return x.InterfaceNeighborCount
	}
	return 0
}

func (x *OspfShInterfaceMadj) GetInterfaceAdjNeighborCount() uint32 {
	if x != nil {
		return x.InterfaceAdjNeighborCount
	}
	return 0
}

func (x *OspfShInterfaceMadj) GetInterfaceLinkCost() uint32 {
	if x != nil {
		return x.InterfaceLinkCost
	}
	return 0
}

func (x *OspfShInterfaceMadj) GetOspfInterfaceState() string {
	if x != nil {
		return x.OspfInterfaceState
	}
	return ""
}

var File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDesc = []byte{
	0x0a, 0x8c, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x6f,
	0x73, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72,
	0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x2f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68,
	0x5f, 0x69, 0x66, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x75, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70,
	0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x62, 0x72, 0x69, 0x65, 0x66, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x22, 0x61, 0x0a, 0x15, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73,
	0x68, 0x5f, 0x69, 0x66, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf5, 0x05, 0x0a, 0x10, 0x6f, 0x73,
	0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x69, 0x66, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x2b, 0x0a, 0x11,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x35, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6f, 0x73, 0x70, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x44, 0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x66, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x5f, 0x68, 0x6f, 0x6c, 0x64,
	0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x38, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x46, 0x61, 0x73, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x48, 0x6f, 0x6c, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x39, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x61, 0x64, 0x6a, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x69, 0x73, 0x5f, 0x6d, 0x61, 0x64, 0x6a, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x69, 0x73, 0x4d, 0x61, 0x64, 0x6a, 0x12, 0x30,
	0x0a, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x64, 0x6a,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x64, 0x6a, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0xbe, 0x01, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6d,
	0x61, 0x64, 0x6a, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x3d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x8d,
	0x01, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73,
	0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x64, 0x6a, 0x52, 0x11,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x64, 0x6a, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0xbe, 0x02, 0x0a, 0x16, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x64, 0x6a, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x41,
	0x72, 0x65, 0x61, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x64, 0x6a, 0x5f, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x64, 0x6a, 0x41,
	0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3f, 0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6a,
	0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x41, 0x64, 0x6a, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6f, 0x73, 0x70, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescData = file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDesc
)

func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDescData
}

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_goTypes = []interface{}{
	(*OspfShIfBrief_KEYS)(nil),  // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief_KEYS
	(*OspfShIfBrief)(nil),       // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief
	(*OspfShInterfaceMadj)(nil), // 2: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_interface_madj
}
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief.interface_madj_list:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_interface_madj
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_init()
}
func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_init() {
	if File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShIfBrief_KEYS); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShIfBrief); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShInterfaceMadj); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto = out.File
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_rawDesc = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_goTypes = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief_ospf_sh_if_brief_proto_depIdxs = nil
}
