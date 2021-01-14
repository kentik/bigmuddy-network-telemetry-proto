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
// source: cisco_ios_xr_drivers_media_eth_oper/ethernet_interface/berts/bert/eth_showctrl_bert.proto

package cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert

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

// Per port BERT test status information
type EthShowctrlBert_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *EthShowctrlBert_KEYS) Reset() {
	*x = EthShowctrlBert_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthShowctrlBert_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthShowctrlBert_KEYS) ProtoMessage() {}

func (x *EthShowctrlBert_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthShowctrlBert_KEYS.ProtoReflect.Descriptor instead.
func (*EthShowctrlBert_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescGZIP(), []int{0}
}

func (x *EthShowctrlBert_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

type EthShowctrlBert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current test status
	BertStatus *EthernetBertStatus_ `protobuf:"bytes,50,opt,name=bert_status,json=bertStatus,proto3" json:"bert_status,omitempty"`
	// Remaining time for this test in seconds
	TimeLeft uint32 `protobuf:"varint,51,opt,name=time_left,json=timeLeft,proto3" json:"time_left,omitempty"`
	// Port BERT interval
	PortBertInterval uint32 `protobuf:"varint,52,opt,name=port_bert_interval,json=portBertInterval,proto3" json:"port_bert_interval,omitempty"`
}

func (x *EthShowctrlBert) Reset() {
	*x = EthShowctrlBert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthShowctrlBert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthShowctrlBert) ProtoMessage() {}

func (x *EthShowctrlBert) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthShowctrlBert.ProtoReflect.Descriptor instead.
func (*EthShowctrlBert) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescGZIP(), []int{1}
}

func (x *EthShowctrlBert) GetBertStatus() *EthernetBertStatus_ {
	if x != nil {
		return x.BertStatus
	}
	return nil
}

func (x *EthShowctrlBert) GetTimeLeft() uint32 {
	if x != nil {
		return x.TimeLeft
	}
	return 0
}

func (x *EthShowctrlBert) GetPortBertInterval() uint32 {
	if x != nil {
		return x.PortBertInterval
	}
	return 0
}

type EthernetBertStatus_ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State
	BertStateEnabled uint32 `protobuf:"varint,1,opt,name=bert_state_enabled,json=bertStateEnabled,proto3" json:"bert_state_enabled,omitempty"`
	// Flag indicating available data
	DataAvailability uint32 `protobuf:"varint,2,opt,name=data_availability,json=dataAvailability,proto3" json:"data_availability,omitempty"`
	// Receive count (if 0x1 set in flag)
	ReceiveCount uint64 `protobuf:"varint,3,opt,name=receive_count,json=receiveCount,proto3" json:"receive_count,omitempty"`
	// Transmit count (if 0x2 set in flag)
	TransmitCount uint64 `protobuf:"varint,4,opt,name=transmit_count,json=transmitCount,proto3" json:"transmit_count,omitempty"`
	// Received errors (if 0x4 set in flag)
	ReceiveErrors uint64 `protobuf:"varint,5,opt,name=receive_errors,json=receiveErrors,proto3" json:"receive_errors,omitempty"`
	// Bit, block or frame error
	ErrorType string `protobuf:"bytes,6,opt,name=error_type,json=errorType,proto3" json:"error_type,omitempty"`
	// Test pattern
	TestPattern string `protobuf:"bytes,7,opt,name=test_pattern,json=testPattern,proto3" json:"test_pattern,omitempty"`
	// Device being tested
	DeviceUnderTest string `protobuf:"bytes,8,opt,name=device_under_test,json=deviceUnderTest,proto3" json:"device_under_test,omitempty"`
	// Interface being tested
	InterfaceDevice string `protobuf:"bytes,9,opt,name=interface_device,json=interfaceDevice,proto3" json:"interface_device,omitempty"`
}

func (x *EthernetBertStatus_) Reset() {
	*x = EthernetBertStatus_{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthernetBertStatus_) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthernetBertStatus_) ProtoMessage() {}

func (x *EthernetBertStatus_) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthernetBertStatus_.ProtoReflect.Descriptor instead.
func (*EthernetBertStatus_) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescGZIP(), []int{2}
}

func (x *EthernetBertStatus_) GetBertStateEnabled() uint32 {
	if x != nil {
		return x.BertStateEnabled
	}
	return 0
}

func (x *EthernetBertStatus_) GetDataAvailability() uint32 {
	if x != nil {
		return x.DataAvailability
	}
	return 0
}

func (x *EthernetBertStatus_) GetReceiveCount() uint64 {
	if x != nil {
		return x.ReceiveCount
	}
	return 0
}

func (x *EthernetBertStatus_) GetTransmitCount() uint64 {
	if x != nil {
		return x.TransmitCount
	}
	return 0
}

func (x *EthernetBertStatus_) GetReceiveErrors() uint64 {
	if x != nil {
		return x.ReceiveErrors
	}
	return 0
}

func (x *EthernetBertStatus_) GetErrorType() string {
	if x != nil {
		return x.ErrorType
	}
	return ""
}

func (x *EthernetBertStatus_) GetTestPattern() string {
	if x != nil {
		return x.TestPattern
	}
	return ""
}

func (x *EthernetBertStatus_) GetDeviceUnderTest() string {
	if x != nil {
		return x.DeviceUnderTest
	}
	return ""
}

func (x *EthernetBertStatus_) GetInterfaceDevice() string {
	if x != nil {
		return x.InterfaceDevice
	}
	return ""
}

var File_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDesc = []byte{
	0x0a, 0x59, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x65, 0x74, 0x68,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x62, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x62,
	0x65, 0x72, 0x74, 0x2f, 0x65, 0x74, 0x68, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x74, 0x72, 0x6c,
	0x5f, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x41, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x65, 0x74, 0x68, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x22, 0x3f,
	0x0a, 0x16, 0x65, 0x74, 0x68, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x74, 0x72, 0x6c, 0x5f, 0x62,
	0x65, 0x72, 0x74, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xd9, 0x01, 0x0a, 0x11, 0x65, 0x74, 0x68, 0x5f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x74, 0x72, 0x6c,
	0x5f, 0x62, 0x65, 0x72, 0x74, 0x12, 0x79, 0x0a, 0x0b, 0x62, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x58, 0x2e, 0x63, 0x69, 0x73,
	0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x65, 0x74, 0x68, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x62, 0x65, 0x72, 0x74, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x62, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x52, 0x0a, 0x62, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x33, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x65, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0xfe, 0x02, 0x0a, 0x15,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x62, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x62, 0x65, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10,
	0x64, 0x61, 0x74, 0x61, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x75, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescData = file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDesc
)

func file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescData)
	})
	return file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDescData
}

var file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_goTypes = []interface{}{
	(*EthShowctrlBert_KEYS)(nil), // 0: cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.berts.bert.eth_showctrl_bert_KEYS
	(*EthShowctrlBert)(nil),      // 1: cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.berts.bert.eth_showctrl_bert
	(*EthernetBertStatus_)(nil),  // 2: cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.berts.bert.ethernet_bert_status_
}
var file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.berts.bert.eth_showctrl_bert.bert_status:type_name -> cisco_ios_xr_drivers_media_eth_oper.ethernet_interface.berts.bert.ethernet_bert_status_
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_init()
}
func file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_init() {
	if File_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthShowctrlBert_KEYS); i {
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
		file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthShowctrlBert); i {
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
		file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthernetBertStatus_); i {
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
			RawDescriptor: file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto = out.File
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_rawDesc = nil
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_goTypes = nil
	file_cisco_ios_xr_drivers_media_eth_oper_ethernet_interface_berts_bert_eth_showctrl_bert_proto_depIdxs = nil
}
