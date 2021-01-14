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
// source: cisco_ios_xr_ipv4_ospf_oper/ospf/processes/process/default_vrf/route_information/route_table/route/ospf_sh_topology.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route

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

// OSPF Route Information
type OspfShTopology_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Prefix       string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
}

func (x *OspfShTopology_KEYS) Reset() {
	*x = OspfShTopology_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShTopology_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShTopology_KEYS) ProtoMessage() {}

func (x *OspfShTopology_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShTopology_KEYS.ProtoReflect.Descriptor instead.
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescGZIP(), []int{0}
}

func (x *OspfShTopology_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *OspfShTopology_KEYS) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *OspfShTopology_KEYS) GetPrefixLength() uint32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

type OspfShTopology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prefix
	RoutePrefix string `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix,proto3" json:"route_prefix,omitempty"`
	// Prefix length
	RoutePrefixLength uint32 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength,proto3" json:"route_prefix_length,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	// Route type
	RouteType string `protobuf:"bytes,53,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	// If true, connected route
	RouteConnected bool `protobuf:"varint,54,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	// Route information
	RouteInfo *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo,proto3" json:"route_info,omitempty"`
	// List of paths to this route
	RoutePathList []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path_list,json=routePathList,proto3" json:"route_path_list,omitempty"`
}

func (x *OspfShTopology) Reset() {
	*x = OspfShTopology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShTopology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShTopology) ProtoMessage() {}

func (x *OspfShTopology) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShTopology.ProtoReflect.Descriptor instead.
func (*OspfShTopology) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescGZIP(), []int{1}
}

func (x *OspfShTopology) GetRoutePrefix() string {
	if x != nil {
		return x.RoutePrefix
	}
	return ""
}

func (x *OspfShTopology) GetRoutePrefixLength() uint32 {
	if x != nil {
		return x.RoutePrefixLength
	}
	return 0
}

func (x *OspfShTopology) GetRouteMetric() uint32 {
	if x != nil {
		return x.RouteMetric
	}
	return 0
}

func (x *OspfShTopology) GetRouteType() string {
	if x != nil {
		return x.RouteType
	}
	return ""
}

func (x *OspfShTopology) GetRouteConnected() bool {
	if x != nil {
		return x.RouteConnected
	}
	return false
}

func (x *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if x != nil {
		return x.RouteInfo
	}
	return nil
}

func (x *OspfShTopology) GetRoutePathList() []*OspfShTopPath {
	if x != nil {
		return x.RoutePathList
	}
	return nil
}

type OspfShTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Second     uint32 `protobuf:"varint,1,opt,name=second,proto3" json:"second,omitempty"`
	Nanosecond uint32 `protobuf:"varint,2,opt,name=nanosecond,proto3" json:"nanosecond,omitempty"`
}

func (x *OspfShTime) Reset() {
	*x = OspfShTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShTime) ProtoMessage() {}

func (x *OspfShTime) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShTime.ProtoReflect.Descriptor instead.
func (*OspfShTime) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescGZIP(), []int{2}
}

func (x *OspfShTime) GetSecond() uint32 {
	if x != nil {
		return x.Second
	}
	return 0
}

func (x *OspfShTime) GetNanosecond() uint32 {
	if x != nil {
		return x.Nanosecond
	}
	return 0
}

// OSPF Common Route Information
type OspfShTopCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Area ID
	RouteAreaId uint32 `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId,proto3" json:"route_area_id,omitempty"`
	// TE metric
	RouteTeMetric uint32 `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric,proto3" json:"route_te_metric,omitempty"`
	// RIB version
	RouteRibVersion uint32 `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion,proto3" json:"route_rib_version,omitempty"`
	// SPF version
	RouteSpfVersion uint64 `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion,proto3" json:"route_spf_version,omitempty"`
	// Forward distance
	RouteForwardDistance uint32 `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance,proto3" json:"route_forward_distance,omitempty"`
	// Protocol source
	RouteSource uint32 `protobuf:"varint,6,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	// Last time updated
	RouteUpdateTime *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime,proto3" json:"route_update_time,omitempty"`
	// Last time update failed
	RouteFailTime *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime,proto3" json:"route_fail_time,omitempty"`
	// SPF priority
	RouteSpfPriority uint32 `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority,proto3" json:"route_spf_priority,omitempty"`
	// If true, exclude from TE paths
	RouteAutoExcluded bool `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded,proto3" json:"route_auto_excluded,omitempty"`
	// If true, SRTE registered prefix route
	RouteSrtePrefixRegistered bool `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered,proto3" json:"route_srte_prefix_registered,omitempty"`
	// SRTE registered neigbhor count on route
	RouteSrteNbrRegistered uint32 `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered,proto3" json:"route_srte_nbr_registered,omitempty"`
}

func (x *OspfShTopCommon) Reset() {
	*x = OspfShTopCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShTopCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShTopCommon) ProtoMessage() {}

func (x *OspfShTopCommon) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShTopCommon.ProtoReflect.Descriptor instead.
func (*OspfShTopCommon) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescGZIP(), []int{3}
}

func (x *OspfShTopCommon) GetRouteAreaId() uint32 {
	if x != nil {
		return x.RouteAreaId
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if x != nil {
		return x.RouteTeMetric
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if x != nil {
		return x.RouteRibVersion
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if x != nil {
		return x.RouteSpfVersion
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if x != nil {
		return x.RouteForwardDistance
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteSource() uint32 {
	if x != nil {
		return x.RouteSource
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if x != nil {
		return x.RouteUpdateTime
	}
	return nil
}

func (x *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if x != nil {
		return x.RouteFailTime
	}
	return nil
}

func (x *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if x != nil {
		return x.RouteSpfPriority
	}
	return 0
}

func (x *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if x != nil {
		return x.RouteAutoExcluded
	}
	return false
}

func (x *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if x != nil {
		return x.RouteSrtePrefixRegistered
	}
	return false
}

func (x *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if x != nil {
		return x.RouteSrteNbrRegistered
	}
	return 0
}

// OSPF Route Path Information
type OspfShTopPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Next hop Interface
	RouteInterfaceName string `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName,proto3" json:"route_interface_name,omitempty"`
	// Nexthop IP address
	RouteNextHopAddress string `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress,proto3" json:"route_next_hop_address,omitempty"`
	// IP address of source of route
	RouteSource string `protobuf:"bytes,3,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	// LSA ID, see RFC2328
	RouteLsaid string `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid,proto3" json:"route_lsaid,omitempty"`
	// Multicast-intact path
	RoutePathIsMcastIntact bool `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact,proto3" json:"route_path_is_mcast_intact,omitempty"`
	// UCMP path
	RoutePathIsUcmpPath bool `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath,proto3" json:"route_path_is_ucmp_path,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,7,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	// LSA type, see RFC2328 etc.
	LsaType uint32 `protobuf:"varint,8,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	// Area ID
	AreaId uint32 `protobuf:"varint,9,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	// Area format IP or uint32
	AreaFormat bool `protobuf:"varint,10,opt,name=area_format,json=areaFormat,proto3" json:"area_format,omitempty"`
}

func (x *OspfShTopPath) Reset() {
	*x = OspfShTopPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShTopPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShTopPath) ProtoMessage() {}

func (x *OspfShTopPath) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShTopPath.ProtoReflect.Descriptor instead.
func (*OspfShTopPath) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescGZIP(), []int{4}
}

func (x *OspfShTopPath) GetRouteInterfaceName() string {
	if x != nil {
		return x.RouteInterfaceName
	}
	return ""
}

func (x *OspfShTopPath) GetRouteNextHopAddress() string {
	if x != nil {
		return x.RouteNextHopAddress
	}
	return ""
}

func (x *OspfShTopPath) GetRouteSource() string {
	if x != nil {
		return x.RouteSource
	}
	return ""
}

func (x *OspfShTopPath) GetRouteLsaid() string {
	if x != nil {
		return x.RouteLsaid
	}
	return ""
}

func (x *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if x != nil {
		return x.RoutePathIsMcastIntact
	}
	return false
}

func (x *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if x != nil {
		return x.RoutePathIsUcmpPath
	}
	return false
}

func (x *OspfShTopPath) GetRouteMetric() uint32 {
	if x != nil {
		return x.RouteMetric
	}
	return 0
}

func (x *OspfShTopPath) GetLsaType() uint32 {
	if x != nil {
		return x.LsaType
	}
	return 0
}

func (x *OspfShTopPath) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *OspfShTopPath) GetAreaFormat() bool {
	if x != nil {
		return x.AreaFormat
	}
	return false
}

var File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDesc = []byte{
	0x0a, 0x79, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x6f, 0x73,
	0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x62, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f,
	0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22,
	0x77, 0x0a, 0x15, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x87, 0x04, 0x0a, 0x10, 0x6f, 0x73, 0x70,
	0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x2e, 0x0a, 0x13, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x36, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x95, 0x01, 0x0a, 0x0a,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x76, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f,
	0x73, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72,
	0x66, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x9c, 0x01, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x38, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x74, 0x2e,
	0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76,
	0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x46, 0x0a, 0x0c, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0xa5, 0x06, 0x0a, 0x12, 0x6f,
	0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x41,
	0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x2a, 0x0a,
	0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x69, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x73, 0x70, 0x66, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x70, 0x66, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x9c,
	0x01, 0x0a, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x70, 0x2e, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f,
	0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x0f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x98, 0x01,
	0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x70, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f,
	0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x6f, 0x73, 0x70,
	0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x46, 0x61, 0x69, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x73, 0x70, 0x66, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x70, 0x66, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x45, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x73, 0x72, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x53, 0x72, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x73, 0x72, 0x74, 0x65, 0x5f, 0x6e, 0x62, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x53, 0x72, 0x74, 0x65, 0x4e, 0x62, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x22, 0xa7, 0x03, 0x0a, 0x10, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6c, 0x73, 0x61, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4c, 0x73, 0x61,
	0x69, 0x64, 0x12, 0x3a, 0x0a, 0x1a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x69, 0x73, 0x5f, 0x6d, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x49, 0x73, 0x4d, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x34,
	0x0a, 0x17, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x69, 0x73, 0x5f,
	0x75, 0x63, 0x6d, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x49, 0x73, 0x55, 0x63, 0x6d, 0x70,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x73, 0x61, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x73, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x72, 0x65, 0x61, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescData = file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDesc
)

func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDescData
}

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_goTypes = []interface{}{
	(*OspfShTopology_KEYS)(nil), // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_topology_KEYS
	(*OspfShTopology)(nil),      // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_topology
	(*OspfShTime)(nil),          // 2: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_time
	(*OspfShTopCommon)(nil),     // 3: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_common
	(*OspfShTopPath)(nil),       // 4: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_path
}
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_depIdxs = []int32{
	3, // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_topology.route_info:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_common
	4, // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_topology.route_path_list:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_path
	2, // 2: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_common.route_update_time:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_time
	2, // 3: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_top_common.route_fail_time:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_table.route.ospf_sh_time
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_init()
}
func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_init() {
	if File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShTopology_KEYS); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShTopology); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShTime); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShTopCommon); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShTopPath); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto = out.File
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_rawDesc = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_goTypes = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_table_route_ospf_sh_topology_proto_depIdxs = nil
}
