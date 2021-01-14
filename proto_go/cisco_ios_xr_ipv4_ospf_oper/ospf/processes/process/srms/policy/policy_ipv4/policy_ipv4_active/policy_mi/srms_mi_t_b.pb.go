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
// source: cisco_ios_xr_ipv4_ospf_oper/ospf/processes/process/srms/policy/policy_ipv4/policy_ipv4_active/policy_mi/srms_mi_t_b.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi

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

// SRMS show bag
type SrmsMiTB_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessName string `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	MiId        string `protobuf:"bytes,2,opt,name=mi_id,json=miId,proto3" json:"mi_id,omitempty"`
}

func (x *SrmsMiTB_KEYS) Reset() {
	*x = SrmsMiTB_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SrmsMiTB_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SrmsMiTB_KEYS) ProtoMessage() {}

func (x *SrmsMiTB_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SrmsMiTB_KEYS.ProtoReflect.Descriptor instead.
func (*SrmsMiTB_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescGZIP(), []int{0}
}

func (x *SrmsMiTB_KEYS) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *SrmsMiTB_KEYS) GetMiId() string {
	if x != nil {
		return x.MiId
	}
	return ""
}

type SrmsMiTB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src string `protobuf:"bytes,50,opt,name=src,proto3" json:"src,omitempty"`
	// Router ID
	Router string `protobuf:"bytes,51,opt,name=router,proto3" json:"router,omitempty"`
	// Area (OSPF) or Level (ISIS)
	Area string `protobuf:"bytes,52,opt,name=area,proto3" json:"area,omitempty"`
	Addr *Addr  `protobuf:"bytes,53,opt,name=addr,proto3" json:"addr,omitempty"`
	// Prefix length
	Prefix uint32 `protobuf:"varint,54,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Starting SID
	SidStart uint32 `protobuf:"varint,55,opt,name=sid_start,json=sidStart,proto3" json:"sid_start,omitempty"`
	// SID range
	SidCount uint32 `protobuf:"varint,56,opt,name=sid_count,json=sidCount,proto3" json:"sid_count,omitempty"`
	// Last IP Prefix
	LastPrefix string `protobuf:"bytes,57,opt,name=last_prefix,json=lastPrefix,proto3" json:"last_prefix,omitempty"`
	// Last SID Index
	LastSidIndex uint32 `protobuf:"varint,58,opt,name=last_sid_index,json=lastSidIndex,proto3" json:"last_sid_index,omitempty"`
	// Attached flag
	FlagAttached string `protobuf:"bytes,59,opt,name=flag_attached,json=flagAttached,proto3" json:"flag_attached,omitempty"`
}

func (x *SrmsMiTB) Reset() {
	*x = SrmsMiTB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SrmsMiTB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SrmsMiTB) ProtoMessage() {}

func (x *SrmsMiTB) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SrmsMiTB.ProtoReflect.Descriptor instead.
func (*SrmsMiTB) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescGZIP(), []int{1}
}

func (x *SrmsMiTB) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *SrmsMiTB) GetRouter() string {
	if x != nil {
		return x.Router
	}
	return ""
}

func (x *SrmsMiTB) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *SrmsMiTB) GetAddr() *Addr {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *SrmsMiTB) GetPrefix() uint32 {
	if x != nil {
		return x.Prefix
	}
	return 0
}

func (x *SrmsMiTB) GetSidStart() uint32 {
	if x != nil {
		return x.SidStart
	}
	return 0
}

func (x *SrmsMiTB) GetSidCount() uint32 {
	if x != nil {
		return x.SidCount
	}
	return 0
}

func (x *SrmsMiTB) GetLastPrefix() string {
	if x != nil {
		return x.LastPrefix
	}
	return ""
}

func (x *SrmsMiTB) GetLastSidIndex() uint32 {
	if x != nil {
		return x.LastSidIndex
	}
	return 0
}

func (x *SrmsMiTB) GetFlagAttached() string {
	if x != nil {
		return x.FlagAttached
	}
	return ""
}

type In6AddrTB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *In6AddrTB) Reset() {
	*x = In6AddrTB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *In6AddrTB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*In6AddrTB) ProtoMessage() {}

func (x *In6AddrTB) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use In6AddrTB.ProtoReflect.Descriptor instead.
func (*In6AddrTB) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescGZIP(), []int{2}
}

func (x *In6AddrTB) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Addr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Af string `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	// IPv4
	Ipv4 string `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// IPv6
	Ipv6 *In6AddrTB `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
}

func (x *Addr) Reset() {
	*x = Addr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addr) ProtoMessage() {}

func (x *Addr) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addr.ProtoReflect.Descriptor instead.
func (*Addr) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescGZIP(), []int{3}
}

func (x *Addr) GetAf() string {
	if x != nil {
		return x.Af
	}
	return ""
}

func (x *Addr) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *Addr) GetIpv6() *In6AddrTB {
	if x != nil {
		return x.Ipv6
	}
	return nil
}

var File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDesc = []byte{
	0x0a, 0x79, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x6f, 0x73,
	0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x73, 0x72, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6d, 0x69, 0x2f, 0x73, 0x72, 0x6d, 0x73, 0x5f, 0x6d,
	0x69, 0x5f, 0x74, 0x5f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x67, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f,
	0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x73, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x6d, 0x69, 0x22, 0x4a, 0x0a, 0x10, 0x73, 0x72, 0x6d, 0x73, 0x5f, 0x6d, 0x69, 0x5f,
	0x74, 0x5f, 0x62, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x6d,
	0x69, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x49, 0x64,
	0x22, 0x8d, 0x03, 0x0a, 0x0b, 0x73, 0x72, 0x6d, 0x73, 0x5f, 0x6d, 0x69, 0x5f, 0x74, 0x5f, 0x62,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x72, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x81,
	0x01, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x35, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6d, 0x2e,
	0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76,
	0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x6f, 0x73, 0x70, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x73, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6d, 0x69, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x36, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x69, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x69, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x39, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x69,
	0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x69, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x6c, 0x61, 0x67, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x3b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x6c, 0x61, 0x67, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x22, 0x24, 0x0a, 0x0c, 0x69, 0x6e, 0x36, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x74, 0x5f, 0x62,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x61, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x66, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x70, 0x76, 0x34, 0x12, 0x89, 0x01, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x75, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78,
	0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x73, 0x70, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x6f, 0x73, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6d, 0x69, 0x2e, 0x69, 0x6e, 0x36,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x74, 0x5f, 0x62, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescData = file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDesc
)

func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDescData
}

var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_goTypes = []interface{}{
	(*SrmsMiTB_KEYS)(nil), // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.srms_mi_t_b_KEYS
	(*SrmsMiTB)(nil),      // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.srms_mi_t_b
	(*In6AddrTB)(nil),     // 2: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.in6_addr_t_b
	(*Addr)(nil),          // 3: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.addr
}
var file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_depIdxs = []int32{
	3, // 0: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.srms_mi_t_b.addr:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.addr
	2, // 1: cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.addr.ipv6:type_name -> cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_active.policy_mi.in6_addr_t_b
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_init()
}
func file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_init() {
	if File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SrmsMiTB_KEYS); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SrmsMiTB); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*In6AddrTB); i {
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
		file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addr); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto = out.File
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_rawDesc = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_goTypes = nil
	file_cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_active_policy_mi_srms_mi_t_b_proto_depIdxs = nil
}
