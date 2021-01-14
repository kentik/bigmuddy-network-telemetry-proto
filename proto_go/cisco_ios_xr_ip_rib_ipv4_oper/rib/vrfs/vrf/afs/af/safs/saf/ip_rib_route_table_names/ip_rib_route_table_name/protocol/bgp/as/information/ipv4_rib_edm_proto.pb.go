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
// source: cisco_ios_xr_ip_rib_ipv4_oper/rib/vrfs/vrf/afs/af/safs/saf/ip_rib_route_table_names/ip_rib_route_table_name/protocol/bgp/as/information/ipv4_rib_edm_proto.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information

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

// Information of a rib protocol
type Ipv4RibEdmProto_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	As             string `protobuf:"bytes,5,opt,name=as,proto3" json:"as,omitempty"`
}

func (x *Ipv4RibEdmProto_KEYS) Reset() {
	*x = Ipv4RibEdmProto_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv4RibEdmProto_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv4RibEdmProto_KEYS) ProtoMessage() {}

func (x *Ipv4RibEdmProto_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv4RibEdmProto_KEYS.ProtoReflect.Descriptor instead.
func (*Ipv4RibEdmProto_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Ipv4RibEdmProto_KEYS) GetVrfName() string {
	if x != nil {
		return x.VrfName
	}
	return ""
}

func (x *Ipv4RibEdmProto_KEYS) GetAfName() string {
	if x != nil {
		return x.AfName
	}
	return ""
}

func (x *Ipv4RibEdmProto_KEYS) GetSafName() string {
	if x != nil {
		return x.SafName
	}
	return ""
}

func (x *Ipv4RibEdmProto_KEYS) GetRouteTableName() string {
	if x != nil {
		return x.RouteTableName
	}
	return ""
}

func (x *Ipv4RibEdmProto_KEYS) GetAs() string {
	if x != nil {
		return x.As
	}
	return ""
}

type Ipv4RibEdmProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames,proto3" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance,proto3" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version,proto3" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount,proto3" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount,proto3" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts,proto3" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount,proto3" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount,proto3" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount,proto3" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory,proto3" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount,proto3" json:"backup_routes_count,omitempty"`
}

func (x *Ipv4RibEdmProto) Reset() {
	*x = Ipv4RibEdmProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv4RibEdmProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv4RibEdmProto) ProtoMessage() {}

func (x *Ipv4RibEdmProto) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv4RibEdmProto.ProtoReflect.Descriptor instead.
func (*Ipv4RibEdmProto) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescGZIP(), []int{1}
}

func (x *Ipv4RibEdmProto) GetProtocolNames() string {
	if x != nil {
		return x.ProtocolNames
	}
	return ""
}

func (x *Ipv4RibEdmProto) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *Ipv4RibEdmProto) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetRedistributionClientCount() uint32 {
	if x != nil {
		return x.RedistributionClientCount
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetProtocolClientsCount() uint32 {
	if x != nil {
		return x.ProtocolClientsCount
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetRoutesCounts() uint32 {
	if x != nil {
		return x.RoutesCounts
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetActiveRoutesCount() uint32 {
	if x != nil {
		return x.ActiveRoutesCount
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetDeletedRoutesCount() uint32 {
	if x != nil {
		return x.DeletedRoutesCount
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetPathsCount() uint32 {
	if x != nil {
		return x.PathsCount
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetProtocolRouteMemory() uint32 {
	if x != nil {
		return x.ProtocolRouteMemory
	}
	return 0
}

func (x *Ipv4RibEdmProto) GetBackupRoutesCount() uint32 {
	if x != nil {
		return x.BackupRoutesCount
	}
	return 0
}

var File_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDesc = []byte{
	0x0a, 0xa0, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2f, 0x72, 0x69, 0x62, 0x2f, 0x76, 0x72, 0x66, 0x73, 0x2f, 0x76, 0x72, 0x66, 0x2f, 0x61, 0x66,
	0x73, 0x2f, 0x61, 0x66, 0x2f, 0x73, 0x61, 0x66, 0x73, 0x2f, 0x73, 0x61, 0x66, 0x2f, 0x69, 0x70,
	0x5f, 0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x62, 0x67, 0x70, 0x2f, 0x61, 0x73, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x70, 0x76, 0x34, 0x5f,
	0x72, 0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x87, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f,
	0x78, 0x72, 0x5f, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x72, 0x69, 0x62, 0x2e, 0x76, 0x72, 0x66, 0x73, 0x2e, 0x76, 0x72, 0x66,
	0x2e, 0x61, 0x66, 0x73, 0x2e, 0x61, 0x66, 0x2e, 0x73, 0x61, 0x66, 0x73, 0x2e, 0x73, 0x61, 0x66,
	0x2e, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x62, 0x67, 0x70, 0x2e, 0x61,
	0x73, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01,
	0x0a, 0x17, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x72, 0x66,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x72, 0x66,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x61, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x61, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x61, 0x73, 0x22, 0xf3, 0x03, 0x0a, 0x12, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f,
	0x65, 0x64, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x1b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x37, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x39, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x3b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescData = file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDesc
)

func file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescData)
	})
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDescData
}

var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_goTypes = []interface{}{
	(*Ipv4RibEdmProto_KEYS)(nil), // 0: cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.bgp.as.information.ipv4_rib_edm_proto_KEYS
	(*Ipv4RibEdmProto)(nil),      // 1: cisco_ios_xr_ip_rib_ipv4_oper.rib.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.bgp.as.information.ipv4_rib_edm_proto
}
var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_init()
}
func file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_init() {
	if File_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv4RibEdmProto_KEYS); i {
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
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv4RibEdmProto); i {
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
			RawDescriptor: file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto = out.File
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_rawDesc = nil
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_goTypes = nil
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_bgp_as_information_ipv4_rib_edm_proto_proto_depIdxs = nil
}
