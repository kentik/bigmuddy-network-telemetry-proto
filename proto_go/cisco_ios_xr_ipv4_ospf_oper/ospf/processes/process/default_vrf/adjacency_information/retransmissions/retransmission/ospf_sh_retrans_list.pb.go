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
// source: cisco_ios_xr_ipv4_ospf_oper/ospf/processes/process/default_vrf/adjacency_information/retransmissions/retransmission/ospf_sh_retrans_list.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission

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

// OSPF Retransmission List
type OspfShRetransList_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName     string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName   string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
}

func (x *OspfShRetransList_KEYS) Reset() {
	*x = OspfShRetransList_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShRetransList_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShRetransList_KEYS) ProtoMessage() {}

func (x *OspfShRetransList_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShRetransList_KEYS.ProtoReflect.Descriptor instead.
func (*OspfShRetransList_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescGZIP(), []int{0}
}

func (x *OspfShRetransList_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *OspfShRetransList_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *OspfShRetransList_KEYS) GetNeighborAddress() string {
	if x != nil {
		return x.NeighborAddress
	}
	return ""
}

type OspfShRetransList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Neighbor ID
	RetransmissionNeighborId string `protobuf:"bytes,50,opt,name=retransmission_neighbor_id,json=retransmissionNeighborId,proto3" json:"retransmission_neighbor_id,omitempty"`
	// Neighbor IP Address
	RetransmissionNeighborIpAddress string `protobuf:"bytes,51,opt,name=retransmission_neighbor_ip_address,json=retransmissionNeighborIpAddress,proto3" json:"retransmission_neighbor_ip_address,omitempty"`
	// Retransmission list interface
	RetransmissionInterfaceName string `protobuf:"bytes,52,opt,name=retransmission_interface_name,json=retransmissionInterfaceName,proto3" json:"retransmission_interface_name,omitempty"`
	// Amount of time remaining on retransmission timer (ms)
	RetransmissionTimer uint32 `protobuf:"varint,53,opt,name=retransmission_timer,json=retransmissionTimer,proto3" json:"retransmission_timer,omitempty"`
	// Retransmission queue length
	RetransmissionCount uint32 `protobuf:"varint,54,opt,name=retransmission_count,json=retransmissionCount,proto3" json:"retransmission_count,omitempty"`
	// List of Area scope entries
	RetransmissionAreaDbList []*OspfShLsaSum `protobuf:"bytes,55,rep,name=retransmission_area_db_list,json=retransmissionAreaDbList,proto3" json:"retransmission_area_db_list,omitempty"`
	// List of AS Scope entries
	RetransmissionAsdbList []*OspfShLsaSum `protobuf:"bytes,56,rep,name=retransmission_asdb_list,json=retransmissionAsdbList,proto3" json:"retransmission_asdb_list,omitempty"`
}

func (x *OspfShRetransList) Reset() {
	*x = OspfShRetransList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShRetransList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShRetransList) ProtoMessage() {}

func (x *OspfShRetransList) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShRetransList.ProtoReflect.Descriptor instead.
func (*OspfShRetransList) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescGZIP(), []int{1}
}

func (x *OspfShRetransList) GetRetransmissionNeighborId() string {
	if x != nil {
		return x.RetransmissionNeighborId
	}
	return ""
}

func (x *OspfShRetransList) GetRetransmissionNeighborIpAddress() string {
	if x != nil {
		return x.RetransmissionNeighborIpAddress
	}
	return ""
}

func (x *OspfShRetransList) GetRetransmissionInterfaceName() string {
	if x != nil {
		return x.RetransmissionInterfaceName
	}
	return ""
}

func (x *OspfShRetransList) GetRetransmissionTimer() uint32 {
	if x != nil {
		return x.RetransmissionTimer
	}
	return 0
}

func (x *OspfShRetransList) GetRetransmissionCount() uint32 {
	if x != nil {
		return x.RetransmissionCount
	}
	return 0
}

func (x *OspfShRetransList) GetRetransmissionAreaDbList() []*OspfShLsaSum {
	if x != nil {
		return x.RetransmissionAreaDbList
	}
	return nil
}

func (x *OspfShRetransList) GetRetransmissionAsdbList() []*OspfShLsaSum {
	if x != nil {
		return x.RetransmissionAsdbList
	}
	return nil
}

// LSA Summary Entry
type OspfShLsaSum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LSA Type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	// Age of the LSA (s)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsId string `protobuf:"bytes,3,opt,name=header_ls_id,json=headerLsId,proto3" json:"header_ls_id,omitempty"`
	// Router ID of the Advertising Router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber uint32 `protobuf:"varint,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	// Checksum of the LSA
	HeaderLsaChecksum uint32 `protobuf:"varint,6,opt,name=header_lsa_checksum,json=headerLsaChecksum,proto3" json:"header_lsa_checksum,omitempty"`
}

func (x *OspfShLsaSum) Reset() {
	*x = OspfShLsaSum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OspfShLsaSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OspfShLsaSum) ProtoMessage() {}

func (x *OspfShLsaSum) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OspfShLsaSum.ProtoReflect.Descriptor instead.
func (*OspfShLsaSum) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescGZIP(), []int{2}
}

func (x *OspfShLsaSum) GetHeaderLsaType() string {
	if x != nil {
		return x.HeaderLsaType
	}
	return ""
}

func (x *OspfShLsaSum) GetHeaderLsaAge() uint32 {
	if x != nil {
		return x.HeaderLsaAge
	}
	return 0
}

func (x *OspfShLsaSum) GetHeaderLsId() string {
	if x != nil {
		return x.HeaderLsId
	}
	return ""
}

func (x *OspfShLsaSum) GetHeaderAdvertisingRouter() string {
	if x != nil {
		return x.HeaderAdvertisingRouter
	}
	return ""
}

func (x *OspfShLsaSum) GetHeaderSequenceNumber() uint32 {
	if x != nil {
		return x.HeaderSequenceNumber
	}
	return 0
}

func (x *OspfShLsaSum) GetHeaderLsaChecksum() uint32 {
	if x != nil {
		return x.HeaderLsaChecksum
	}
	return 0
}

var File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDesc = []byte{
	0x0a, 0x8e, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x6f,
	0x73, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72,
	0x66, 0x2f, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x72,
	0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x73, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f,
	0x73, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72,
	0x66, 0x2e, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x19, 0x6f, 0x73, 0x70, 0x66, 0x5f,
	0x73, 0x68, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62,
	0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xd4, 0x05, 0x0a, 0x14, 0x6f, 0x73,
	0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x1a, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x4b, 0x0a, 0x22, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x69, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x72, 0x65,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x69, 0x67,
	0x68, 0x62, 0x6f, 0x72, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x42, 0x0a,
	0x1d, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x34,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x13, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x36, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xc4, 0x01, 0x0a, 0x1b, 0x72, 0x65, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f,
	0x64, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x37, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x84, 0x01,
	0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70,
	0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x72, 0x66, 0x2e,
	0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x6c, 0x73, 0x61,
	0x5f, 0x73, 0x75, 0x6d, 0x52, 0x18, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x44, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0xbf,
	0x01, 0x0a, 0x18, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x73, 0x64, 0x62, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x38, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x84, 0x01, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78,
	0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x76, 0x72, 0x66, 0x2e, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68,
	0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x73, 0x75, 0x6d, 0x52, 0x16, 0x72, 0x65, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x64, 0x62, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xa3, 0x02, 0x0a, 0x0f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x73, 0x68, 0x5f, 0x6c, 0x73, 0x61,
	0x5f, 0x73, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c,
	0x73, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x41,
	0x67, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x4c, 0x73, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x16, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x14, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x6c, 0x73, 0x61, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x73, 0x61, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescData = file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDesc
)

func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDescData
}

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_goTypes = []interface{}{
	(*OspfShRetransList_KEYS)(nil), // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list_KEYS
	(*OspfShRetransList)(nil),      // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list
	(*OspfShLsaSum)(nil),           // 2: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum
}
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list.retransmission_area_db_list:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum
	2, // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list.retransmission_asdb_list:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_init()
}
func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_init() {
	if File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShRetransList_KEYS); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShRetransList); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OspfShLsaSum); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto = out.File
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_rawDesc = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_goTypes = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission_ospf_sh_retrans_list_proto_depIdxs = nil
}
