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
// source: cisco_ios_xr_infra_rmf_oper/redundancy/summary/red_summary.proto

package cisco_ios_xr_infra_rmf_oper_redundancy_summary

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

// rmf show
type RedSummary_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RedSummary_KEYS) Reset() {
	*x = RedSummary_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedSummary_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedSummary_KEYS) ProtoMessage() {}

func (x *RedSummary_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedSummary_KEYS.ProtoReflect.Descriptor instead.
func (*RedSummary_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescGZIP(), []int{0}
}

type RedSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Redundancy Pair
	RedPair []*RedSummaryPair `protobuf:"bytes,50,rep,name=red_pair,json=redPair,proto3" json:"red_pair,omitempty"`
	// Error Log
	ErrLog string `protobuf:"bytes,51,opt,name=err_log,json=errLog,proto3" json:"err_log,omitempty"`
}

func (x *RedSummary) Reset() {
	*x = RedSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedSummary) ProtoMessage() {}

func (x *RedSummary) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedSummary.ProtoReflect.Descriptor instead.
func (*RedSummary) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescGZIP(), []int{1}
}

func (x *RedSummary) GetRedPair() []*RedSummaryPair {
	if x != nil {
		return x.RedPair
	}
	return nil
}

func (x *RedSummary) GetErrLog() string {
	if x != nil {
		return x.ErrLog
	}
	return ""
}

type RedGroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active   string `protobuf:"bytes,1,opt,name=active,proto3" json:"active,omitempty"`
	Standby  string `protobuf:"bytes,2,opt,name=standby,proto3" json:"standby,omitempty"`
	HaState  string `protobuf:"bytes,3,opt,name=ha_state,json=haState,proto3" json:"ha_state,omitempty"`
	NsrState string `protobuf:"bytes,4,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
}

func (x *RedGroupInfo) Reset() {
	*x = RedGroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedGroupInfo) ProtoMessage() {}

func (x *RedGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedGroupInfo.ProtoReflect.Descriptor instead.
func (*RedGroupInfo) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescGZIP(), []int{2}
}

func (x *RedGroupInfo) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

func (x *RedGroupInfo) GetStandby() string {
	if x != nil {
		return x.Standby
	}
	return ""
}

func (x *RedGroupInfo) GetHaState() string {
	if x != nil {
		return x.HaState
	}
	return ""
}

func (x *RedGroupInfo) GetNsrState() string {
	if x != nil {
		return x.NsrState
	}
	return ""
}

// rmf row show
type RedSummaryPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Active node name R/S/I
	Active string `protobuf:"bytes,1,opt,name=active,proto3" json:"active,omitempty"`
	// Standby node name R/S/I
	Standby string `protobuf:"bytes,2,opt,name=standby,proto3" json:"standby,omitempty"`
	// High Availability state Ready/Not Ready
	HaState string `protobuf:"bytes,3,opt,name=ha_state,json=haState,proto3" json:"ha_state,omitempty"`
	// NSR state Configured/Not Configured
	NsrState  string          `protobuf:"bytes,4,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	Groupinfo []*RedGroupInfo `protobuf:"bytes,5,rep,name=groupinfo,proto3" json:"groupinfo,omitempty"`
}

func (x *RedSummaryPair) Reset() {
	*x = RedSummaryPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedSummaryPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedSummaryPair) ProtoMessage() {}

func (x *RedSummaryPair) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedSummaryPair.ProtoReflect.Descriptor instead.
func (*RedSummaryPair) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescGZIP(), []int{3}
}

func (x *RedSummaryPair) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

func (x *RedSummaryPair) GetStandby() string {
	if x != nil {
		return x.Standby
	}
	return ""
}

func (x *RedSummaryPair) GetHaState() string {
	if x != nil {
		return x.HaState
	}
	return ""
}

func (x *RedSummaryPair) GetNsrState() string {
	if x != nil {
		return x.NsrState
	}
	return ""
}

func (x *RedSummaryPair) GetGroupinfo() []*RedGroupInfo {
	if x != nil {
		return x.Groupinfo
	}
	return nil
}

var File_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x72, 0x6d, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x79, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x2f, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x72, 0x6d, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e,
	0x72, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x79, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x5f, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x5b, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x61,
	0x69, 0x72, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f,
	0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x72, 0x6d,
	0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x79, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x52, 0x07, 0x72, 0x65, 0x64, 0x50,
	0x61, 0x69, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4c, 0x6f, 0x67, 0x22, 0x7a, 0x0a, 0x0e,
	0x72, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x73, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x73, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xda, 0x01, 0x0a, 0x10, 0x72, 0x65, 0x64,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x68, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x73,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x73, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x72, 0x6d, 0x66, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61,
	0x6e, 0x63, 0x79, 0x2e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x64, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescData = file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDesc
)

func file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescData)
	})
	return file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDescData
}

var file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_goTypes = []interface{}{
	(*RedSummary_KEYS)(nil), // 0: cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary_KEYS
	(*RedSummary)(nil),      // 1: cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary
	(*RedGroupInfo)(nil),    // 2: cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_group_info
	(*RedSummaryPair)(nil),  // 3: cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary_pair
}
var file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_depIdxs = []int32{
	3, // 0: cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary.red_pair:type_name -> cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary_pair
	2, // 1: cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary_pair.groupinfo:type_name -> cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_group_info
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_init() }
func file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_init() {
	if File_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedSummary_KEYS); i {
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
		file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedSummary); i {
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
		file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedGroupInfo); i {
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
		file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedSummaryPair); i {
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
			RawDescriptor: file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto = out.File
	file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_rawDesc = nil
	file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_goTypes = nil
	file_cisco_ios_xr_infra_rmf_oper_redundancy_summary_red_summary_proto_depIdxs = nil
}
