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
// source: cisco_ios_xr_infra_statsd_oper/infra_statistics/interfaces/interface/protocols/protocol/ifstatsbag_proto.proto

package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol

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

// Interface counters per protocol
type IfstatsbagProto_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	ProtocolName  string `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
}

func (x *IfstatsbagProto_KEYS) Reset() {
	*x = IfstatsbagProto_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IfstatsbagProto_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IfstatsbagProto_KEYS) ProtoMessage() {}

func (x *IfstatsbagProto_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IfstatsbagProto_KEYS.ProtoReflect.Descriptor instead.
func (*IfstatsbagProto_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescGZIP(), []int{0}
}

func (x *IfstatsbagProto_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *IfstatsbagProto_KEYS) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

type IfstatsbagProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bytes received
	BytesReceived uint64 `protobuf:"varint,50,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	// Packets received
	PacketsReceived uint64 `protobuf:"varint,51,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	// Bytes sent
	BytesSent uint64 `protobuf:"varint,52,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	// Packets sent
	PacketsSent uint64 `protobuf:"varint,53,opt,name=packets_sent,json=packetsSent,proto3" json:"packets_sent,omitempty"`
	// Protocol number
	Protocol uint32 `protobuf:"varint,54,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Time when counters were last written (in seconds)
	LastDataTime uint32 `protobuf:"varint,55,opt,name=last_data_time,json=lastDataTime,proto3" json:"last_data_time,omitempty"`
	// Input data rate in 1000's of bps
	InputDataRate uint64 `protobuf:"varint,56,opt,name=input_data_rate,json=inputDataRate,proto3" json:"input_data_rate,omitempty"`
	// Input packets per second
	InputPacketRate uint64 `protobuf:"varint,57,opt,name=input_packet_rate,json=inputPacketRate,proto3" json:"input_packet_rate,omitempty"`
	// Output data rate in 1000's of bps
	OutputDataRate uint64 `protobuf:"varint,58,opt,name=output_data_rate,json=outputDataRate,proto3" json:"output_data_rate,omitempty"`
	// Output packets per second
	OutputPacketRate uint64 `protobuf:"varint,59,opt,name=output_packet_rate,json=outputPacketRate,proto3" json:"output_packet_rate,omitempty"`
}

func (x *IfstatsbagProto) Reset() {
	*x = IfstatsbagProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IfstatsbagProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IfstatsbagProto) ProtoMessage() {}

func (x *IfstatsbagProto) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IfstatsbagProto.ProtoReflect.Descriptor instead.
func (*IfstatsbagProto) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescGZIP(), []int{1}
}

func (x *IfstatsbagProto) GetBytesReceived() uint64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

func (x *IfstatsbagProto) GetPacketsReceived() uint64 {
	if x != nil {
		return x.PacketsReceived
	}
	return 0
}

func (x *IfstatsbagProto) GetBytesSent() uint64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *IfstatsbagProto) GetPacketsSent() uint64 {
	if x != nil {
		return x.PacketsSent
	}
	return 0
}

func (x *IfstatsbagProto) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *IfstatsbagProto) GetLastDataTime() uint32 {
	if x != nil {
		return x.LastDataTime
	}
	return 0
}

func (x *IfstatsbagProto) GetInputDataRate() uint64 {
	if x != nil {
		return x.InputDataRate
	}
	return 0
}

func (x *IfstatsbagProto) GetInputPacketRate() uint64 {
	if x != nil {
		return x.InputPacketRate
	}
	return 0
}

func (x *IfstatsbagProto) GetOutputDataRate() uint64 {
	if x != nil {
		return x.OutputDataRate
	}
	return 0
}

func (x *IfstatsbagProto) GetOutputPacketRate() uint64 {
	if x != nil {
		return x.OutputPacketRate
	}
	return 0
}

var File_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDesc = []byte{
	0x0a, 0x6e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x69, 0x66, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x62, 0x61, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x57, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x63, 0x0a, 0x15, 0x69, 0x66, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x62, 0x61, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x4b, 0x45,
	0x59, 0x53, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x94,
	0x03, 0x0a, 0x10, 0x69, 0x66, 0x73, 0x74, 0x61, 0x74, 0x73, 0x62, 0x61, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x6e, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x53, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f,
	0x73, 0x65, 0x6e, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x38, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x39, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x3b, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescData = file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDesc
)

func file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescData)
	})
	return file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDescData
}

var file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_goTypes = []interface{}{
	(*IfstatsbagProto_KEYS)(nil), // 0: cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.protocols.protocol.ifstatsbag_proto_KEYS
	(*IfstatsbagProto)(nil),      // 1: cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.protocols.protocol.ifstatsbag_proto
}
var file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_init()
}
func file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_init() {
	if File_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IfstatsbagProto_KEYS); i {
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
		file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IfstatsbagProto); i {
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
			RawDescriptor: file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto = out.File
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_rawDesc = nil
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_goTypes = nil
	file_cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_protocols_protocol_ifstatsbag_proto_proto_depIdxs = nil
}
