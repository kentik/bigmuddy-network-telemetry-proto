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
// source: cisco_ios_xr_ipv6_ospfv3_oper/ospfv3/processes/process/vrfs/vrf/route_summary/ospfv3_edm_route_sum.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary

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

// OSPFv3 route summary information
type Ospfv3EdmRouteSum_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName     string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
}

func (x *Ospfv3EdmRouteSum_KEYS) Reset() {
	*x = Ospfv3EdmRouteSum_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmRouteSum_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmRouteSum_KEYS) ProtoMessage() {}

func (x *Ospfv3EdmRouteSum_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmRouteSum_KEYS.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmRouteSum_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescGZIP(), []int{0}
}

func (x *Ospfv3EdmRouteSum_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *Ospfv3EdmRouteSum_KEYS) GetVrfName() string {
	if x != nil {
		return x.VrfName
	}
	return ""
}

type Ospfv3EdmRouteSum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Route summary of a route ID
	RouteId string `protobuf:"bytes,50,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	// Intra route summary
	IntraAreaRoute uint32 `protobuf:"varint,51,opt,name=intra_area_route,json=intraAreaRoute,proto3" json:"intra_area_route,omitempty"`
	// Inter route summary
	InterAreaRoute uint32 `protobuf:"varint,52,opt,name=inter_area_route,json=interAreaRoute,proto3" json:"inter_area_route,omitempty"`
	// Extern 1 route summary
	ExternOneRoute uint32 `protobuf:"varint,53,opt,name=extern_one_route,json=externOneRoute,proto3" json:"extern_one_route,omitempty"`
	// Extern 2 route summary
	ExternTwoRoute uint32 `protobuf:"varint,54,opt,name=extern_two_route,json=externTwoRoute,proto3" json:"extern_two_route,omitempty"`
	// NSSA 1 route summary
	NssaOneRoute uint32 `protobuf:"varint,55,opt,name=nssa_one_route,json=nssaOneRoute,proto3" json:"nssa_one_route,omitempty"`
	// NSSA 2 route summary
	NssaTwoRoute uint32 `protobuf:"varint,56,opt,name=nssa_two_route,json=nssaTwoRoute,proto3" json:"nssa_two_route,omitempty"`
	// Total route summary
	TotalSentRoute uint32 `protobuf:"varint,57,opt,name=total_sent_route,json=totalSentRoute,proto3" json:"total_sent_route,omitempty"`
	// Route connected
	RouteConnected uint32 `protobuf:"varint,58,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	// Redistribution route summary
	RedistributionRoute uint32 `protobuf:"varint,59,opt,name=redistribution_route,json=redistributionRoute,proto3" json:"redistribution_route,omitempty"`
	// Total route received summary
	TotalReceivedRoute uint32 `protobuf:"varint,60,opt,name=total_received_route,json=totalReceivedRoute,proto3" json:"total_received_route,omitempty"`
}

func (x *Ospfv3EdmRouteSum) Reset() {
	*x = Ospfv3EdmRouteSum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ospfv3EdmRouteSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ospfv3EdmRouteSum) ProtoMessage() {}

func (x *Ospfv3EdmRouteSum) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ospfv3EdmRouteSum.ProtoReflect.Descriptor instead.
func (*Ospfv3EdmRouteSum) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescGZIP(), []int{1}
}

func (x *Ospfv3EdmRouteSum) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

func (x *Ospfv3EdmRouteSum) GetIntraAreaRoute() uint32 {
	if x != nil {
		return x.IntraAreaRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetInterAreaRoute() uint32 {
	if x != nil {
		return x.InterAreaRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetExternOneRoute() uint32 {
	if x != nil {
		return x.ExternOneRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetExternTwoRoute() uint32 {
	if x != nil {
		return x.ExternTwoRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetNssaOneRoute() uint32 {
	if x != nil {
		return x.NssaOneRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetNssaTwoRoute() uint32 {
	if x != nil {
		return x.NssaTwoRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetTotalSentRoute() uint32 {
	if x != nil {
		return x.TotalSentRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetRouteConnected() uint32 {
	if x != nil {
		return x.RouteConnected
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetRedistributionRoute() uint32 {
	if x != nil {
		return x.RedistributionRoute
	}
	return 0
}

func (x *Ospfv3EdmRouteSum) GetTotalReceivedRoute() uint32 {
	if x != nil {
		return x.TotalReceivedRoute
	}
	return 0
}

var File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDesc = []byte{
	0x0a, 0x68, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f,
	0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x72, 0x66, 0x73, 0x2f, 0x76, 0x72,
	0x66, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f,
	0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4d, 0x63, 0x69, 0x73, 0x63,
	0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6f, 0x73,
	0x70, 0x66, 0x76, 0x33, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x72, 0x66, 0x73, 0x2e, 0x76, 0x72, 0x66, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x59, 0x0a, 0x19, 0x6f, 0x73, 0x70,
	0x66, 0x76, 0x33, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x75,
	0x6d, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x72, 0x66,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x72, 0x66,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdd, 0x03, 0x0a, 0x14, 0x6f, 0x73, 0x70, 0x66, 0x76, 0x33, 0x5f,
	0x65, 0x64, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x12, 0x19, 0x0a,
	0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x72,
	0x61, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x41, 0x72, 0x65, 0x61, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x35, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x4f, 0x6e,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x54, 0x77, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x73, 0x73, 0x61, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x73, 0x73, 0x61, 0x4f, 0x6e,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x73, 0x73, 0x61, 0x5f, 0x74,
	0x77, 0x6f, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6e, 0x73, 0x73, 0x61, 0x54, 0x77, 0x6f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x18, 0x39, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x6e,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x31, 0x0a, 0x14, 0x72, 0x65, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescData = file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDesc
)

func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDescData
}

var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_goTypes = []interface{}{
	(*Ospfv3EdmRouteSum_KEYS)(nil), // 0: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.route_summary.ospfv3_edm_route_sum_KEYS
	(*Ospfv3EdmRouteSum)(nil),      // 1: cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.route_summary.ospfv3_edm_route_sum
}
var file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_init()
}
func file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_init() {
	if File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmRouteSum_KEYS); i {
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
		file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ospfv3EdmRouteSum); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto = out.File
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_rawDesc = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_goTypes = nil
	file_cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_route_summary_ospfv3_edm_route_sum_proto_depIdxs = nil
}
