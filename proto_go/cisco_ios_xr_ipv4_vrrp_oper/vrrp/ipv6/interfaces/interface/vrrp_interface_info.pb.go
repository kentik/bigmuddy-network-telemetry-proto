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
// source: cisco_ios_xr_ipv4_vrrp_oper/vrrp/ipv6/interfaces/interface/vrrp_interface_info.proto

package cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface

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

// VRRP Interface statistics
type VrrpInterfaceInfo_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *VrrpInterfaceInfo_KEYS) Reset() {
	*x = VrrpInterfaceInfo_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VrrpInterfaceInfo_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VrrpInterfaceInfo_KEYS) ProtoMessage() {}

func (x *VrrpInterfaceInfo_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VrrpInterfaceInfo_KEYS.ProtoReflect.Descriptor instead.
func (*VrrpInterfaceInfo_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescGZIP(), []int{0}
}

func (x *VrrpInterfaceInfo_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

type VrrpInterfaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IM Interface
	Interface string `protobuf:"bytes,50,opt,name=interface,proto3" json:"interface,omitempty"`
	// Invalid checksum
	InvalidChecksumCount uint32 `protobuf:"varint,51,opt,name=invalid_checksum_count,json=invalidChecksumCount,proto3" json:"invalid_checksum_count,omitempty"`
	// Unknown/unsupported version
	InvalidVersionCount uint32 `protobuf:"varint,52,opt,name=invalid_version_count,json=invalidVersionCount,proto3" json:"invalid_version_count,omitempty"`
	// Invalid vrID
	InvalidVridCount uint32 `protobuf:"varint,53,opt,name=invalid_vrid_count,json=invalidVridCount,proto3" json:"invalid_vrid_count,omitempty"`
	// Bad packet lengths
	InvalidPacketLengthCount uint32 `protobuf:"varint,54,opt,name=invalid_packet_length_count,json=invalidPacketLengthCount,proto3" json:"invalid_packet_length_count,omitempty"`
}

func (x *VrrpInterfaceInfo) Reset() {
	*x = VrrpInterfaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VrrpInterfaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VrrpInterfaceInfo) ProtoMessage() {}

func (x *VrrpInterfaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VrrpInterfaceInfo.ProtoReflect.Descriptor instead.
func (*VrrpInterfaceInfo) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescGZIP(), []int{1}
}

func (x *VrrpInterfaceInfo) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *VrrpInterfaceInfo) GetInvalidChecksumCount() uint32 {
	if x != nil {
		return x.InvalidChecksumCount
	}
	return 0
}

func (x *VrrpInterfaceInfo) GetInvalidVersionCount() uint32 {
	if x != nil {
		return x.InvalidVersionCount
	}
	return 0
}

func (x *VrrpInterfaceInfo) GetInvalidVridCount() uint32 {
	if x != nil {
		return x.InvalidVridCount
	}
	return 0
}

func (x *VrrpInterfaceInfo) GetInvalidPacketLengthCount() uint32 {
	if x != nil {
		return x.InvalidPacketLengthCount
	}
	return 0
}

var File_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDesc = []byte{
	0x0a, 0x54, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x76, 0x72, 0x72, 0x70, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x76, 0x72,
	0x72, 0x70, 0x2f, 0x69, 0x70, 0x76, 0x36, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x72, 0x72,
	0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f,
	0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x76, 0x72, 0x72, 0x70, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x76, 0x72, 0x72, 0x70, 0x2e, 0x69, 0x70, 0x76, 0x36, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x22, 0x41, 0x0a, 0x18, 0x76, 0x72, 0x72, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x02, 0x0a, 0x13, 0x76, 0x72, 0x72, 0x70, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x5f, 0x76, 0x72, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x35, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x56, 0x72, 0x69, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x36, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescData = file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDesc
)

func file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDescData
}

var file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_goTypes = []interface{}{
	(*VrrpInterfaceInfo_KEYS)(nil), // 0: cisco_ios_xr_ipv4_vrrp_oper.vrrp.ipv6.interfaces.interface.vrrp_interface_info_KEYS
	(*VrrpInterfaceInfo)(nil),      // 1: cisco_ios_xr_ipv4_vrrp_oper.vrrp.ipv6.interfaces.interface.vrrp_interface_info
}
var file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_init()
}
func file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_init() {
	if File_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VrrpInterfaceInfo_KEYS); i {
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
		file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VrrpInterfaceInfo); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto = out.File
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_rawDesc = nil
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_goTypes = nil
	file_cisco_ios_xr_ipv4_vrrp_oper_vrrp_ipv6_interfaces_interface_vrrp_interface_info_proto_depIdxs = nil
}
