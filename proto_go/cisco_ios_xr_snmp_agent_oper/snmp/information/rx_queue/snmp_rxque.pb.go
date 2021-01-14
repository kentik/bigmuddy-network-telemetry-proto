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
// source: cisco_ios_xr_snmp_agent_oper/snmp/information/rx_queue/snmp_rxque.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue

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

// SNMP Incoming queue statistics
type SnmpRxque_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SnmpRxque_KEYS) Reset() {
	*x = SnmpRxque_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpRxque_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpRxque_KEYS) ProtoMessage() {}

func (x *SnmpRxque_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpRxque_KEYS.ProtoReflect.Descriptor instead.
func (*SnmpRxque_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescGZIP(), []int{0}
}

type SnmpRxque struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qlen      uint32 `protobuf:"varint,50,opt,name=qlen,proto3" json:"qlen,omitempty"`
	InMin     uint32 `protobuf:"varint,51,opt,name=in_min,json=inMin,proto3" json:"in_min,omitempty"`
	InAvg     uint32 `protobuf:"varint,52,opt,name=in_avg,json=inAvg,proto3" json:"in_avg,omitempty"`
	InMax     uint32 `protobuf:"varint,53,opt,name=in_max,json=inMax,proto3" json:"in_max,omitempty"`
	PendMin   uint32 `protobuf:"varint,54,opt,name=pend_min,json=pendMin,proto3" json:"pend_min,omitempty"`
	PendAvg   uint32 `protobuf:"varint,55,opt,name=pend_avg,json=pendAvg,proto3" json:"pend_avg,omitempty"`
	PendMax   uint32 `protobuf:"varint,56,opt,name=pend_max,json=pendMax,proto3" json:"pend_max,omitempty"`
	IncomingQ []byte `protobuf:"bytes,57,opt,name=incoming_q,json=incomingQ,proto3" json:"incoming_q,omitempty"`
	PendingQ  []byte `protobuf:"bytes,58,opt,name=pending_q,json=pendingQ,proto3" json:"pending_q,omitempty"`
}

func (x *SnmpRxque) Reset() {
	*x = SnmpRxque{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpRxque) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpRxque) ProtoMessage() {}

func (x *SnmpRxque) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpRxque.ProtoReflect.Descriptor instead.
func (*SnmpRxque) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescGZIP(), []int{1}
}

func (x *SnmpRxque) GetQlen() uint32 {
	if x != nil {
		return x.Qlen
	}
	return 0
}

func (x *SnmpRxque) GetInMin() uint32 {
	if x != nil {
		return x.InMin
	}
	return 0
}

func (x *SnmpRxque) GetInAvg() uint32 {
	if x != nil {
		return x.InAvg
	}
	return 0
}

func (x *SnmpRxque) GetInMax() uint32 {
	if x != nil {
		return x.InMax
	}
	return 0
}

func (x *SnmpRxque) GetPendMin() uint32 {
	if x != nil {
		return x.PendMin
	}
	return 0
}

func (x *SnmpRxque) GetPendAvg() uint32 {
	if x != nil {
		return x.PendAvg
	}
	return 0
}

func (x *SnmpRxque) GetPendMax() uint32 {
	if x != nil {
		return x.PendMax
	}
	return 0
}

func (x *SnmpRxque) GetIncomingQ() []byte {
	if x != nil {
		return x.IncomingQ
	}
	return nil
}

func (x *SnmpRxque) GetPendingQ() []byte {
	if x != nil {
		return x.PendingQ
	}
	return nil
}

var File_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDesc = []byte{
	0x0a, 0x47, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x73,
	0x6e, 0x6d, 0x70, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73,
	0x6e, 0x6d, 0x70, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x72, 0x78,
	0x71, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x63, 0x69, 0x73, 0x63, 0x6f,
	0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x78, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x22, 0x11, 0x0a, 0x0f, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x72, 0x78, 0x71, 0x75, 0x65, 0x5f,
	0x4b, 0x45, 0x59, 0x53, 0x22, 0xf2, 0x01, 0x0a, 0x0a, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x72, 0x78,
	0x71, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x71, 0x6c, 0x65, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x71, 0x6c, 0x65, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x5f, 0x6d, 0x69,
	0x6e, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x69, 0x6e, 0x5f, 0x61, 0x76, 0x67, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x69, 0x6e, 0x41, 0x76, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x18,
	0x35, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x70, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x6e, 0x64, 0x5f,
	0x61, 0x76, 0x67, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x41,
	0x76, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x38,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x18, 0x39, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x51, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescData = file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDesc
)

func file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescData)
	})
	return file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDescData
}

var file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_goTypes = []interface{}{
	(*SnmpRxque_KEYS)(nil), // 0: cisco_ios_xr_snmp_agent_oper.snmp.information.rx_queue.snmp_rxque_KEYS
	(*SnmpRxque)(nil),      // 1: cisco_ios_xr_snmp_agent_oper.snmp.information.rx_queue.snmp_rxque
}
var file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_init() }
func file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_init() {
	if File_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpRxque_KEYS); i {
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
		file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpRxque); i {
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
			RawDescriptor: file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto = out.File
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_rawDesc = nil
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_goTypes = nil
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_rx_queue_snmp_rxque_proto_depIdxs = nil
}
