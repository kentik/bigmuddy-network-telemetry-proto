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
// source: cisco_ios_xr_ipv6_ospfv3_oper/ospfv3/processes/process/default_vrf/areas/area/interface_brief_table/interface_brief/ospfv3_edm_intf_brief.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief

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

// OSPFv3 brief interface information
type Ospfv3EdmIntfBrief_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId        uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *Ospfv3EdmIntfBrief_KEYS) Reset() {
	*x = Ospfv3EdmIntfBrief_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmIntfBrief_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmIntfBrief_KEYS) ProtoMessage() {}

func (x *Ospfv3EdmIntfBrief_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmIntfBrief_KEYS.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmIntfBrief_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescGZIP(), []int{0}
}

func (x *Ospfv3EdmIntfBrief_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *Ospfv3EdmIntfBrief_KEYS) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *Ospfv3EdmIntfBrief_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

type Ospfv3EdmIntfBrief struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interface IP address
	InterfaceAddress string `protobuf:"bytes,50,opt,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,51,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,52,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	// Total number of neighbors
	InterfaceNeighbors uint32 `protobuf:"varint,53,opt,name=interface_neighbors,json=interfaceNeighbors,proto3" json:"interface_neighbors,omitempty"`
	// Total number of adjacent neighbors
	InterfaceAdjacentNeighbors uint32 `protobuf:"varint,54,opt,name=interface_adjacent_neighbors,json=interfaceAdjacentNeighbors,proto3" json:"interface_adjacent_neighbors,omitempty"`
	// Network type
	NetworkType string `protobuf:"bytes,55,opt,name=network_type,json=networkType,proto3" json:"network_type,omitempty"`
}

func (x *Ospfv3EdmIntfBrief) Reset() {
	*x = Ospfv3EdmIntfBrief{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmIntfBrief) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmIntfBrief) ProtoMessage() {}

func (x *Ospfv3EdmIntfBrief) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmIntfBrief.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmIntfBrief) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescGZIP(), []int{1}
}

func (x *Ospfv3EdmIntfBrief) GetInterfaceAddress() string {
	if x != nil {
		return x.InterfaceAddress
	}
	return ""
}

func (x *Ospfv3EdmIntfBrief) GetInterfaceLinkCost() uint32 {
	if x != nil {
		return x.InterfaceLinkCost
	}
	return 0
}

func (x *Ospfv3EdmIntfBrief) GetOspfInterfaceState() string {
	if x != nil {
		return x.OspfInterfaceState
	}
	return ""
}

func (x *Ospfv3EdmIntfBrief) GetInterfaceNeighbors() uint32 {
	if x != nil {
		return x.InterfaceNeighbors
	}
	return 0
}

func (x *Ospfv3EdmIntfBrief) GetInterfaceAdjacentNeighbors() uint32 {
	if x != nil {
		return x.InterfaceAdjacentNeighbors
	}
	return 0
}

func (x *Ospfv3EdmIntfBrief) GetNetworkType() string {
	if x != nil {
		return x.NetworkType
	}
	return ""
}

var File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDesc = []byte{
	0x0a, 0x8f, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x76, 0x72, 0x66, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2f, 0x61, 0x72, 0x65, 0x61,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x2f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x65, 0x64,
	0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x66, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x73, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72,
	0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x61, 0x72, 0x65, 0x61, 0x73, 0x2e, 0x61, 0x72, 0x65,
	0x61, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x69, 0x65,
	0x66, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x22, 0x7f, 0x0a, 0x1a, 0x6f, 0x73, 0x70, 0x66, 0x76,
	0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x66, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66,
	0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbc, 0x02, 0x0a, 0x15, 0x6f, 0x73, 0x70,
	0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x66, 0x5f, 0x62, 0x72, 0x69,
	0x65, 0x66, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x14, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6f,
	0x73, 0x70, 0x66, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x73, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x73, 0x12, 0x40, 0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x72, 0x73, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x41, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescData = file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDesc
)

func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDescData
}

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_goTypes = []interface{}{
	(*Ospfv3EdmIntfBrief_KEYS)(nil), // 0: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.interface_brief_table.interface_brief.ospfv3_edm_intf_brief_KEYS
	(*Ospfv3EdmIntfBrief)(nil),      // 1: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.areas.area.interface_brief_table.interface_brief.ospfv3_edm_intf_brief
}
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_init()
}
func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_init() {
	if File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmIntfBrief_KEYS); i {
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
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmIntfBrief); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto = out.File
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_rawDesc = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_goTypes = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_areas_area_interface_brief_table_interface_brief_ospfv3_edm_intf_brief_proto_depIdxs = nil
}
