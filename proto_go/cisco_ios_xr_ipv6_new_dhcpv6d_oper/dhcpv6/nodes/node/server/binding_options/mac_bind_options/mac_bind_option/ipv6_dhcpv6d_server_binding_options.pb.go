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
// source: cisco_ios_xr_ipv6_new_dhcpv6d_oper/dhcpv6/nodes/node/server/binding_options/mac_bind_options/mac_bind_option/ipv6_dhcpv6d_server_binding_options.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option

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

// DHCPv6 server binding inserted option values
type Ipv6Dhcpv6DServerBindingOptions_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName   string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	MacAddress string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
}

func (x *Ipv6Dhcpv6DServerBindingOptions_KEYS) Reset() {
	*x = Ipv6Dhcpv6DServerBindingOptions_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv6Dhcpv6DServerBindingOptions_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv6Dhcpv6DServerBindingOptions_KEYS) ProtoMessage() {}

func (x *Ipv6Dhcpv6DServerBindingOptions_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv6Dhcpv6DServerBindingOptions_KEYS.ProtoReflect.Descriptor instead.
func (*Ipv6Dhcpv6DServerBindingOptions_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescGZIP(), []int{0}
}

func (x *Ipv6Dhcpv6DServerBindingOptions_KEYS) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Ipv6Dhcpv6DServerBindingOptions_KEYS) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

type Ipv6Dhcpv6DServerBindingOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Client MAC address
	MacAddress string `protobuf:"bytes,50,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// Client DUID
	Duid string `protobuf:"bytes,51,opt,name=duid,proto3" json:"duid,omitempty"`
	// DNS address count
	DnsCount uint32 `protobuf:"varint,52,opt,name=dns_count,json=dnsCount,proto3" json:"dns_count,omitempty"`
	// DNS addresses
	DnsAddress []*IPV6AddressType `protobuf:"bytes,53,rep,name=dns_address,json=dnsAddress,proto3" json:"dns_address,omitempty"`
	// Client Option 17 value
	Opt17 string `protobuf:"bytes,54,opt,name=opt17,proto3" json:"opt17,omitempty"`
}

func (x *Ipv6Dhcpv6DServerBindingOptions) Reset() {
	*x = Ipv6Dhcpv6DServerBindingOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv6Dhcpv6DServerBindingOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv6Dhcpv6DServerBindingOptions) ProtoMessage() {}

func (x *Ipv6Dhcpv6DServerBindingOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv6Dhcpv6DServerBindingOptions.ProtoReflect.Descriptor instead.
func (*Ipv6Dhcpv6DServerBindingOptions) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescGZIP(), []int{1}
}

func (x *Ipv6Dhcpv6DServerBindingOptions) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *Ipv6Dhcpv6DServerBindingOptions) GetDuid() string {
	if x != nil {
		return x.Duid
	}
	return ""
}

func (x *Ipv6Dhcpv6DServerBindingOptions) GetDnsCount() uint32 {
	if x != nil {
		return x.DnsCount
	}
	return 0
}

func (x *Ipv6Dhcpv6DServerBindingOptions) GetDnsAddress() []*IPV6AddressType {
	if x != nil {
		return x.DnsAddress
	}
	return nil
}

func (x *Ipv6Dhcpv6DServerBindingOptions) GetOpt17() string {
	if x != nil {
		return x.Opt17
	}
	return ""
}

// IPV6 Address type
type IPV6AddressType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IPV6AddressType) Reset() {
	*x = IPV6AddressType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPV6AddressType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPV6AddressType) ProtoMessage() {}

func (x *IPV6AddressType) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPV6AddressType.ProtoReflect.Descriptor instead.
func (*IPV6AddressType) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescGZIP(), []int{2}
}

func (x *IPV6AddressType) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDesc = []byte{
	0x0a, 0x96, 0x01, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d,
	0x61, 0x63, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6d, 0x61, 0x63, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x69, 0x70, 0x76, 0x36, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x6c, 0x63, 0x69, 0x73, 0x63, 0x6f,
	0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6e, 0x65, 0x77,
	0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x64, 0x68,
	0x63, 0x70, 0x76, 0x36, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x63, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x63, 0x5f, 0x62, 0x69, 0x6e, 0x64,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x28, 0x69, 0x70, 0x76, 0x36, 0x5f,
	0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x4b,
	0x45, 0x59, 0x53, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0xae, 0x02, 0x0a, 0x23, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76,
	0x36, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x75,
	0x69, 0x64, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x75, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x34, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x64, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x0b,
	0x64, 0x6e, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x35, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x7d, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72,
	0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36,
	0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x6d, 0x61, 0x63, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x6d, 0x61, 0x63, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x50, 0x56, 0x36, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x64, 0x6e, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x70, 0x74, 0x31, 0x37, 0x18, 0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x70, 0x74,
	0x31, 0x37, 0x22, 0x27, 0x0a, 0x0f, 0x49, 0x50, 0x56, 0x36, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescData = file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDesc
)

func file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDescData
}

var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_goTypes = []interface{}{
	(*Ipv6Dhcpv6DServerBindingOptions_KEYS)(nil), // 0: cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options_KEYS
	(*Ipv6Dhcpv6DServerBindingOptions)(nil),      // 1: cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options
	(*IPV6AddressType)(nil),                      // 2: cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.IPV6AddressType
}
var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options.dns_address:type_name -> cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.IPV6AddressType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_init()
}
func file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_init() {
	if File_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv6Dhcpv6DServerBindingOptions_KEYS); i {
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
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv6Dhcpv6DServerBindingOptions); i {
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
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPV6AddressType); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto = out.File
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_rawDesc = nil
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_goTypes = nil
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option_ipv6_dhcpv6d_server_binding_options_proto_depIdxs = nil
}
