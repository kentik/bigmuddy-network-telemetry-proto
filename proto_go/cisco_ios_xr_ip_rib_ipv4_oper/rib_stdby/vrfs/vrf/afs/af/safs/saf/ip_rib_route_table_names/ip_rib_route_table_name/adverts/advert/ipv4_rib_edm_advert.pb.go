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
// source: cisco_ios_xr_ip_rib_ipv4_oper/rib_stdby/vrfs/vrf/afs/af/safs/saf/ip_rib_route_table_names/ip_rib_route_table_name/adverts/advert/ipv4_rib_edm_advert.proto

package cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert

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

// Route advertisement information
type Ipv4RibEdmAdvert_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName,proto3" json:"route_table_name,omitempty"`
	Address        string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength   uint32 `protobuf:"varint,6,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
}

func (x *Ipv4RibEdmAdvert_KEYS) Reset() {
	*x = Ipv4RibEdmAdvert_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv4RibEdmAdvert_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv4RibEdmAdvert_KEYS) ProtoMessage() {}

func (x *Ipv4RibEdmAdvert_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv4RibEdmAdvert_KEYS.ProtoReflect.Descriptor instead.
func (*Ipv4RibEdmAdvert_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescGZIP(), []int{0}
}

func (x *Ipv4RibEdmAdvert_KEYS) GetVrfName() string {
	if x != nil {
		return x.VrfName
	}
	return ""
}

func (x *Ipv4RibEdmAdvert_KEYS) GetAfName() string {
	if x != nil {
		return x.AfName
	}
	return ""
}

func (x *Ipv4RibEdmAdvert_KEYS) GetSafName() string {
	if x != nil {
		return x.SafName
	}
	return ""
}

func (x *Ipv4RibEdmAdvert_KEYS) GetRouteTableName() string {
	if x != nil {
		return x.RouteTableName
	}
	return ""
}

func (x *Ipv4RibEdmAdvert_KEYS) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Ipv4RibEdmAdvert_KEYS) GetPrefixLength() uint32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

type Ipv4RibEdmAdvert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Next advertising proto
	Ipv4RibEdmAdvert []*Ipv4RibEdmAdvertItem `protobuf:"bytes,50,rep,name=ipv4_rib_edm_advert,json=ipv4RibEdmAdvert,proto3" json:"ipv4_rib_edm_advert,omitempty"`
}

func (x *Ipv4RibEdmAdvert) Reset() {
	*x = Ipv4RibEdmAdvert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv4RibEdmAdvert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv4RibEdmAdvert) ProtoMessage() {}

func (x *Ipv4RibEdmAdvert) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv4RibEdmAdvert.ProtoReflect.Descriptor instead.
func (*Ipv4RibEdmAdvert) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescGZIP(), []int{1}
}

func (x *Ipv4RibEdmAdvert) GetIpv4RibEdmAdvert() []*Ipv4RibEdmAdvertItem {
	if x != nil {
		return x.Ipv4RibEdmAdvert
	}
	return nil
}

type Ipv4RibEdmAdvertItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Protocol advertising the route
	ProtocolId uint32 `protobuf:"varint,1,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
	//   Client advertising the route
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Number of extended communities
	NumberOfExtendedCommunities uint32 `protobuf:"varint,3,opt,name=number_of_extended_communities,json=numberOfExtendedCommunities,proto3" json:"number_of_extended_communities,omitempty"`
	// Extended communities
	ExtendedCommunities []byte `protobuf:"bytes,4,opt,name=extended_communities,json=extendedCommunities,proto3" json:"extended_communities,omitempty"`
	// OSPF area-id flags
	ProtocolOpaqueFlags uint32 `protobuf:"varint,5,opt,name=protocol_opaque_flags,json=protocolOpaqueFlags,proto3" json:"protocol_opaque_flags,omitempty"`
	// OSPF area-id
	ProtocolOpaque uint32 `protobuf:"varint,6,opt,name=protocol_opaque,json=protocolOpaque,proto3" json:"protocol_opaque,omitempty"`
	// Protocol code
	Code int32 `protobuf:"zigzag32,7,opt,name=code,proto3" json:"code,omitempty"`
	// Instance name
	InstanceName []byte `protobuf:"bytes,8,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
}

func (x *Ipv4RibEdmAdvertItem) Reset() {
	*x = Ipv4RibEdmAdvertItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv4RibEdmAdvertItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv4RibEdmAdvertItem) ProtoMessage() {}

func (x *Ipv4RibEdmAdvertItem) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv4RibEdmAdvertItem.ProtoReflect.Descriptor instead.
func (*Ipv4RibEdmAdvertItem) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescGZIP(), []int{2}
}

func (x *Ipv4RibEdmAdvertItem) GetProtocolId() uint32 {
	if x != nil {
		return x.ProtocolId
	}
	return 0
}

func (x *Ipv4RibEdmAdvertItem) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Ipv4RibEdmAdvertItem) GetNumberOfExtendedCommunities() uint32 {
	if x != nil {
		return x.NumberOfExtendedCommunities
	}
	return 0
}

func (x *Ipv4RibEdmAdvertItem) GetExtendedCommunities() []byte {
	if x != nil {
		return x.ExtendedCommunities
	}
	return nil
}

func (x *Ipv4RibEdmAdvertItem) GetProtocolOpaqueFlags() uint32 {
	if x != nil {
		return x.ProtocolOpaqueFlags
	}
	return 0
}

func (x *Ipv4RibEdmAdvertItem) GetProtocolOpaque() uint32 {
	if x != nil {
		return x.ProtocolOpaque
	}
	return 0
}

func (x *Ipv4RibEdmAdvertItem) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Ipv4RibEdmAdvertItem) GetInstanceName() []byte {
	if x != nil {
		return x.InstanceName
	}
	return nil
}

var File_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDesc = []byte{
	0x0a, 0x9a, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2f, 0x72, 0x69, 0x62, 0x5f, 0x73, 0x74, 0x64, 0x62, 0x79, 0x2f, 0x76, 0x72, 0x66, 0x73, 0x2f,
	0x76, 0x72, 0x66, 0x2f, 0x61, 0x66, 0x73, 0x2f, 0x61, 0x66, 0x2f, 0x73, 0x61, 0x66, 0x73, 0x2f,
	0x73, 0x61, 0x66, 0x2f, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x69, 0x70, 0x5f,
	0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x2f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d,
	0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x80, 0x01,
	0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x5f,
	0x72, 0x69, 0x62, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x72, 0x69,
	0x62, 0x5f, 0x73, 0x74, 0x64, 0x62, 0x79, 0x2e, 0x76, 0x72, 0x66, 0x73, 0x2e, 0x76, 0x72, 0x66,
	0x2e, 0x61, 0x66, 0x73, 0x2e, 0x61, 0x66, 0x2e, 0x73, 0x61, 0x66, 0x73, 0x2e, 0x73, 0x61, 0x66,
	0x2e, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x22, 0xd2, 0x01, 0x0a, 0x18, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x65, 0x64,
	0x6d, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x72, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x72, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x66, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x66, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0xe2, 0x01, 0x0a, 0x13, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72,
	0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12, 0xca, 0x01,
	0x0a, 0x13, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x9a, 0x01, 0x2e, 0x63,
	0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x5f, 0x72,
	0x69, 0x62, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x72, 0x69, 0x62,
	0x5f, 0x73, 0x74, 0x64, 0x62, 0x79, 0x2e, 0x76, 0x72, 0x66, 0x73, 0x2e, 0x76, 0x72, 0x66, 0x2e,
	0x61, 0x66, 0x73, 0x2e, 0x61, 0x66, 0x2e, 0x73, 0x61, 0x66, 0x73, 0x2e, 0x73, 0x61, 0x66, 0x2e,
	0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x69, 0x70, 0x5f, 0x72, 0x69, 0x62, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x10, 0x69, 0x70, 0x76, 0x34, 0x52, 0x69,
	0x62, 0x45, 0x64, 0x6d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x22, 0xe6, 0x02, 0x0a, 0x18, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x72, 0x69, 0x62, 0x5f, 0x65, 0x64, 0x6d, 0x5f, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x1e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6f, 0x66, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescData = file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDesc
)

func file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescData)
	})
	return file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDescData
}

var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_goTypes = []interface{}{
	(*Ipv4RibEdmAdvert_KEYS)(nil), // 0: cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv4_rib_edm_advert_KEYS
	(*Ipv4RibEdmAdvert)(nil),      // 1: cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv4_rib_edm_advert
	(*Ipv4RibEdmAdvertItem)(nil),  // 2: cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv4_rib_edm_advert_item
}
var file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv4_rib_edm_advert.ipv4_rib_edm_advert:type_name -> cisco_ios_xr_ip_rib_ipv4_oper.rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.adverts.advert.ipv4_rib_edm_advert_item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_init()
}
func file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_init() {
	if File_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv4RibEdmAdvert_KEYS); i {
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
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv4RibEdmAdvert); i {
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
		file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv4RibEdmAdvertItem); i {
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
			RawDescriptor: file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto = out.File
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_rawDesc = nil
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_goTypes = nil
	file_cisco_ios_xr_ip_rib_ipv4_oper_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_adverts_advert_ipv4_rib_edm_advert_proto_depIdxs = nil
}
