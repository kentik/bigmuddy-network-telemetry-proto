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
// source: cisco_ios_xr_clns_isis_oper/isis/instances/instance/nsr_statistics/isis_sh_nsr_stats_global.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics

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

// IS-IS process NSR statistics
type IsisShNsrStatsGlobal_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
}

func (x *IsisShNsrStatsGlobal_KEYS) Reset() {
	*x = IsisShNsrStatsGlobal_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsisShNsrStatsGlobal_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsisShNsrStatsGlobal_KEYS) ProtoMessage() {}

func (x *IsisShNsrStatsGlobal_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsisShNsrStatsGlobal_KEYS.ProtoReflect.Descriptor instead.
func (*IsisShNsrStatsGlobal_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescGZIP(), []int{0}
}

func (x *IsisShNsrStatsGlobal_KEYS) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

type IsisShNsrStatsGlobal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ISIS VM STATE
	IsisVmState uint32 `protobuf:"varint,50,opt,name=isis_vm_state,json=isisVmState,proto3" json:"isis_vm_state,omitempty"`
	// ISIS NSR STATS Data
	IsisNsrStatsData *NsrStatsGblArrType `protobuf:"bytes,51,opt,name=isis_nsr_stats_data,json=isisNsrStatsData,proto3" json:"isis_nsr_stats_data,omitempty"`
}

func (x *IsisShNsrStatsGlobal) Reset() {
	*x = IsisShNsrStatsGlobal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsisShNsrStatsGlobal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsisShNsrStatsGlobal) ProtoMessage() {}

func (x *IsisShNsrStatsGlobal) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsisShNsrStatsGlobal.ProtoReflect.Descriptor instead.
func (*IsisShNsrStatsGlobal) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescGZIP(), []int{1}
}

func (x *IsisShNsrStatsGlobal) GetIsisVmState() uint32 {
	if x != nil {
		return x.IsisVmState
	}
	return 0
}

func (x *IsisShNsrStatsGlobal) GetIsisNsrStatsData() *NsrStatsGblArrType {
	if x != nil {
		return x.IsisNsrStatsData
	}
	return nil
}

type NsrStatsGblType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoOfL1Lsp             uint32 `protobuf:"varint,1,opt,name=no_of_l1_lsp,json=noOfL1Lsp,proto3" json:"no_of_l1_lsp,omitempty"`
	NoOfL2Lsp             uint32 `protobuf:"varint,2,opt,name=no_of_l2_lsp,json=noOfL2Lsp,proto3" json:"no_of_l2_lsp,omitempty"`
	NoOfL1Adj             uint32 `protobuf:"varint,3,opt,name=no_of_l1_adj,json=noOfL1Adj,proto3" json:"no_of_l1_adj,omitempty"`
	NoOfL2Adj             uint32 `protobuf:"varint,4,opt,name=no_of_l2_adj,json=noOfL2Adj,proto3" json:"no_of_l2_adj,omitempty"`
	NoOfLiveInterface     uint32 `protobuf:"varint,5,opt,name=no_of_live_interface,json=noOfLiveInterface,proto3" json:"no_of_live_interface,omitempty"`
	NoOfPtpInterface      uint32 `protobuf:"varint,6,opt,name=no_of_ptp_interface,json=noOfPtpInterface,proto3" json:"no_of_ptp_interface,omitempty"`
	NoOfLanInterface      uint32 `protobuf:"varint,7,opt,name=no_of_lan_interface,json=noOfLanInterface,proto3" json:"no_of_lan_interface,omitempty"`
	NoOfLoopbackInterface uint32 `protobuf:"varint,8,opt,name=no_of_loopback_interface,json=noOfLoopbackInterface,proto3" json:"no_of_loopback_interface,omitempty"`
	NoOfTeTunnels         uint32 `protobuf:"varint,9,opt,name=no_of_te_tunnels,json=noOfTeTunnels,proto3" json:"no_of_te_tunnels,omitempty"`
	NoOfTeLinks           uint32 `protobuf:"varint,10,opt,name=no_of_te_links,json=noOfTeLinks,proto3" json:"no_of_te_links,omitempty"`
	NoOfIpv4Routes        uint32 `protobuf:"varint,11,opt,name=no_of_ipv4_routes,json=noOfIpv4Routes,proto3" json:"no_of_ipv4_routes,omitempty"`
	NoOfIpv6Routes        uint32 `protobuf:"varint,12,opt,name=no_of_ipv6_routes,json=noOfIpv6Routes,proto3" json:"no_of_ipv6_routes,omitempty"`
	Seqnum                uint32 `protobuf:"varint,13,opt,name=seqnum,proto3" json:"seqnum,omitempty"`
}

func (x *NsrStatsGblType) Reset() {
	*x = NsrStatsGblType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsrStatsGblType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsrStatsGblType) ProtoMessage() {}

func (x *NsrStatsGblType) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsrStatsGblType.ProtoReflect.Descriptor instead.
func (*NsrStatsGblType) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescGZIP(), []int{2}
}

func (x *NsrStatsGblType) GetNoOfL1Lsp() uint32 {
	if x != nil {
		return x.NoOfL1Lsp
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfL2Lsp() uint32 {
	if x != nil {
		return x.NoOfL2Lsp
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfL1Adj() uint32 {
	if x != nil {
		return x.NoOfL1Adj
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfL2Adj() uint32 {
	if x != nil {
		return x.NoOfL2Adj
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfLiveInterface() uint32 {
	if x != nil {
		return x.NoOfLiveInterface
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfPtpInterface() uint32 {
	if x != nil {
		return x.NoOfPtpInterface
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfLanInterface() uint32 {
	if x != nil {
		return x.NoOfLanInterface
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfLoopbackInterface() uint32 {
	if x != nil {
		return x.NoOfLoopbackInterface
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfTeTunnels() uint32 {
	if x != nil {
		return x.NoOfTeTunnels
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfTeLinks() uint32 {
	if x != nil {
		return x.NoOfTeLinks
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfIpv4Routes() uint32 {
	if x != nil {
		return x.NoOfIpv4Routes
	}
	return 0
}

func (x *NsrStatsGblType) GetNoOfIpv6Routes() uint32 {
	if x != nil {
		return x.NoOfIpv6Routes
	}
	return 0
}

func (x *NsrStatsGblType) GetSeqnum() uint32 {
	if x != nil {
		return x.Seqnum
	}
	return 0
}

type NsrStatsGblArrType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Self *NsrStatsGblType   `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Peer []*NsrStatsGblType `protobuf:"bytes,2,rep,name=peer,proto3" json:"peer,omitempty"`
}

func (x *NsrStatsGblArrType) Reset() {
	*x = NsrStatsGblArrType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsrStatsGblArrType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsrStatsGblArrType) ProtoMessage() {}

func (x *NsrStatsGblArrType) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsrStatsGblArrType.ProtoReflect.Descriptor instead.
func (*NsrStatsGblArrType) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescGZIP(), []int{3}
}

func (x *NsrStatsGblArrType) GetSelf() *NsrStatsGblType {
	if x != nil {
		return x.Self
	}
	return nil
}

func (x *NsrStatsGblArrType) GetPeer() []*NsrStatsGblType {
	if x != nil {
		return x.Peer
	}
	return nil
}

var File_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDesc = []byte{
	0x0a, 0x61, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x63,
	0x6c, 0x6e, 0x73, 0x5f, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x73,
	0x69, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x5f, 0x6e, 0x73, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x42, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78,
	0x72, 0x5f, 0x63, 0x6c, 0x6e, 0x73, 0x5f, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x69, 0x73, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x44, 0x0a, 0x1d, 0x69, 0x73, 0x69, 0x73, 0x5f,
	0x73, 0x68, 0x5f, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xca, 0x01,
	0x0a, 0x18, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x5f, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73,
	0x69, 0x73, 0x5f, 0x76, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x69, 0x73, 0x69, 0x73, 0x56, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x89,
	0x01, 0x0a, 0x13, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x63,
	0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x63, 0x6c, 0x6e, 0x73,
	0x5f, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x69, 0x73, 0x69, 0x73, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x62, 0x6c, 0x5f,
	0x61, 0x72, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x10, 0x69, 0x73, 0x69, 0x73, 0x4e, 0x73,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x9c, 0x04, 0x0a, 0x12, 0x6e,
	0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x62, 0x6c, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x31, 0x5f, 0x6c, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x6f, 0x4f, 0x66, 0x4c, 0x31, 0x4c,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x32, 0x5f, 0x6c,
	0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x6f, 0x4f, 0x66, 0x4c, 0x32,
	0x4c, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x31, 0x5f,
	0x61, 0x64, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x6f, 0x4f, 0x66, 0x4c,
	0x31, 0x41, 0x64, 0x6a, 0x12, 0x1f, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x32,
	0x5f, 0x61, 0x64, 0x6a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x6f, 0x4f, 0x66,
	0x4c, 0x32, 0x41, 0x64, 0x6a, 0x12, 0x2f, 0x0a, 0x14, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x6e, 0x6f, 0x4f, 0x66, 0x4c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f,
	0x70, 0x74, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x6e, 0x6f, 0x4f, 0x66, 0x50, 0x74, 0x70, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c,
	0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x6e, 0x6f, 0x4f, 0x66, 0x4c, 0x61, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x6f,
	0x6f, 0x70, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x6e, 0x6f, 0x4f, 0x66, 0x4c, 0x6f, 0x6f, 0x70,
	0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x10, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x6f, 0x4f, 0x66, 0x54, 0x65, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x0e, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f,
	0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x6e, 0x6f, 0x4f, 0x66, 0x54, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x6e,
	0x6f, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e, 0x6f, 0x4f, 0x66, 0x49, 0x70, 0x76, 0x34,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x6e, 0x6f, 0x5f, 0x6f, 0x66, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x6e, 0x6f, 0x4f, 0x66, 0x49, 0x70, 0x76, 0x36, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x65, 0x71, 0x6e, 0x75, 0x6d, 0x22, 0xf0, 0x01, 0x0a, 0x16, 0x6e, 0x73,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x62, 0x6c, 0x5f, 0x61, 0x72, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x6a, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x56, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78,
	0x72, 0x5f, 0x63, 0x6c, 0x6e, 0x73, 0x5f, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x69, 0x73, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x67, 0x62, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x73, 0x65, 0x6c, 0x66,
	0x12, 0x6a, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x56,
	0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x63, 0x6c,
	0x6e, 0x73, 0x5f, 0x69, 0x73, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x69, 0x73, 0x69,
	0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x6e, 0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x62,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescData = file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDesc
)

func file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescData)
	})
	return file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDescData
}

var file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_goTypes = []interface{}{
	(*IsisShNsrStatsGlobal_KEYS)(nil), // 0: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.isis_sh_nsr_stats_global_KEYS
	(*IsisShNsrStatsGlobal)(nil),      // 1: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.isis_sh_nsr_stats_global
	(*NsrStatsGblType)(nil),           // 2: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_type
	(*NsrStatsGblArrType)(nil),        // 3: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_arr_type
}
var file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_depIdxs = []int32{
	3, // 0: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.isis_sh_nsr_stats_global.isis_nsr_stats_data:type_name -> cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_arr_type
	2, // 1: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_arr_type.self:type_name -> cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_type
	2, // 2: cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_arr_type.peer:type_name -> cisco_ios_xr_clns_isis_oper.isis.instances.instance.nsr_statistics.nsr_stats_gbl_type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_init()
}
func file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_init() {
	if File_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsisShNsrStatsGlobal_KEYS); i {
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
		file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsisShNsrStatsGlobal); i {
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
		file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsrStatsGblType); i {
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
		file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsrStatsGblArrType); i {
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
			RawDescriptor: file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto = out.File
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_rawDesc = nil
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_goTypes = nil
	file_cisco_ios_xr_clns_isis_oper_isis_instances_instance_nsr_statistics_isis_sh_nsr_stats_global_proto_depIdxs = nil
}
