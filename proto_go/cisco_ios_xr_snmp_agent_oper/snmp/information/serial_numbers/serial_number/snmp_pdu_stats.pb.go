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
// source: cisco_ios_xr_snmp_agent_oper/snmp/information/serial_numbers/serial_number/snmp_pdu_stats.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number

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

type SnmpPduStats_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ReqId  uint32 `protobuf:"varint,2,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
	Port   uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *SnmpPduStats_KEYS) Reset() {
	*x = SnmpPduStats_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpPduStats_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpPduStats_KEYS) ProtoMessage() {}

func (x *SnmpPduStats_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpPduStats_KEYS.ProtoReflect.Descriptor instead.
func (*SnmpPduStats_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescGZIP(), []int{0}
}

func (x *SnmpPduStats_KEYS) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *SnmpPduStats_KEYS) GetReqId() uint32 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (x *SnmpPduStats_KEYS) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type SnmpPduStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  NMS address Rx PDU
	Nms string `protobuf:"bytes,50,opt,name=nms,proto3" json:"nms,omitempty"`
	//  SNMP request id per PDU
	RequestId uint32 `protobuf:"varint,51,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// NMS port number
	Port uint32 `protobuf:"varint,52,opt,name=port,proto3" json:"port,omitempty"`
	//  PDU type
	PduType uint32 `protobuf:"varint,53,opt,name=pdu_type,json=pduType,proto3" json:"pdu_type,omitempty"`
	// Is reques dropped due to error
	ErrorStatus uint32 `protobuf:"varint,54,opt,name=error_status,json=errorStatus,proto3" json:"error_status,omitempty"`
	// Serial number per PDU processing
	SerialNum uint32 `protobuf:"varint,55,opt,name=serial_num,json=serialNum,proto3" json:"serial_num,omitempty"`
	// Request inserted into input queue
	InputQ string `protobuf:"bytes,56,opt,name=input_q,json=inputQ,proto3" json:"input_q,omitempty"`
	// De-queue the request from the input queue
	OutputQ uint32 `protobuf:"varint,57,opt,name=output_q,json=outputQ,proto3" json:"output_q,omitempty"`
	// Enqueue the request into pending queue
	PendingQ uint32 `protobuf:"varint,58,opt,name=pending_q,json=pendingQ,proto3" json:"pending_q,omitempty"`
	// Response sent
	ResponseOut uint32 `protobuf:"varint,59,opt,name=response_out,json=responseOut,proto3" json:"response_out,omitempty"`
}

func (x *SnmpPduStats) Reset() {
	*x = SnmpPduStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpPduStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpPduStats) ProtoMessage() {}

func (x *SnmpPduStats) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpPduStats.ProtoReflect.Descriptor instead.
func (*SnmpPduStats) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescGZIP(), []int{1}
}

func (x *SnmpPduStats) GetNms() string {
	if x != nil {
		return x.Nms
	}
	return ""
}

func (x *SnmpPduStats) GetRequestId() uint32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *SnmpPduStats) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *SnmpPduStats) GetPduType() uint32 {
	if x != nil {
		return x.PduType
	}
	return 0
}

func (x *SnmpPduStats) GetErrorStatus() uint32 {
	if x != nil {
		return x.ErrorStatus
	}
	return 0
}

func (x *SnmpPduStats) GetSerialNum() uint32 {
	if x != nil {
		return x.SerialNum
	}
	return 0
}

func (x *SnmpPduStats) GetInputQ() string {
	if x != nil {
		return x.InputQ
	}
	return ""
}

func (x *SnmpPduStats) GetOutputQ() uint32 {
	if x != nil {
		return x.OutputQ
	}
	return 0
}

func (x *SnmpPduStats) GetPendingQ() uint32 {
	if x != nil {
		return x.PendingQ
	}
	return 0
}

func (x *SnmpPduStats) GetResponseOut() uint32 {
	if x != nil {
		return x.ResponseOut
	}
	return 0
}

var File_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDesc = []byte{
	0x0a, 0x5f, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x73,
	0x6e, 0x6d, 0x70, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73,
	0x6e, 0x6d, 0x70, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x6e, 0x6d,
	0x70, 0x5f, 0x70, 0x64, 0x75, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x4a, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e,
	0x73, 0x6e, 0x6d, 0x70, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x58, 0x0a,
	0x13, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x70, 0x64, 0x75, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f,
	0x4b, 0x45, 0x59, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06,
	0x72, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65,
	0x71, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xa6, 0x02, 0x0a, 0x0e, 0x73, 0x6e, 0x6d, 0x70,
	0x5f, 0x70, 0x64, 0x75, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6d,
	0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x64, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x35, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x64, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x37, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x71, 0x18, 0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x51, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x71, 0x18, 0x39, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x51,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x18, 0x3a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x51, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x3b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x75, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescData = file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDesc
)

func file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescData)
	})
	return file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDescData
}

var file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_goTypes = []interface{}{
	(*SnmpPduStats_KEYS)(nil), // 0: cisco_ios_xr_snmp_agent_oper.snmp.information.serial_numbers.serial_number.snmp_pdu_stats_KEYS
	(*SnmpPduStats)(nil),      // 1: cisco_ios_xr_snmp_agent_oper.snmp.information.serial_numbers.serial_number.snmp_pdu_stats
}
var file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_init()
}
func file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_init() {
	if File_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpPduStats_KEYS); i {
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
		file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpPduStats); i {
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
			RawDescriptor: file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto = out.File
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_rawDesc = nil
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_goTypes = nil
	file_cisco_ios_xr_snmp_agent_oper_snmp_information_serial_numbers_serial_number_snmp_pdu_stats_proto_depIdxs = nil
}
