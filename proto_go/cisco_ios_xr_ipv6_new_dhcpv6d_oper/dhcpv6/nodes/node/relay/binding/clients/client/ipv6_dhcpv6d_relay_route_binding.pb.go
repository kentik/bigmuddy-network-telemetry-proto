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
// source: cisco_ios_xr_ipv6_new_dhcpv6d_oper/dhcpv6/nodes/node/relay/binding/clients/client/ipv6_dhcpv6d_relay_route_binding.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client

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

// DHCPv6 relay route entry
type Ipv6Dhcpv6DRelayRouteBinding_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *Ipv6Dhcpv6DRelayRouteBinding_KEYS) Reset() {
	*x = Ipv6Dhcpv6DRelayRouteBinding_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv6Dhcpv6DRelayRouteBinding_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv6Dhcpv6DRelayRouteBinding_KEYS) ProtoMessage() {}

func (x *Ipv6Dhcpv6DRelayRouteBinding_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv6Dhcpv6DRelayRouteBinding_KEYS.ProtoReflect.Descriptor instead.
func (*Ipv6Dhcpv6DRelayRouteBinding_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescGZIP(), []int{0}
}

func (x *Ipv6Dhcpv6DRelayRouteBinding_KEYS) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Ipv6Dhcpv6DRelayRouteBinding_KEYS) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type Ipv6Dhcpv6DRelayRouteBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Client DUID
	Duid string `protobuf:"bytes,50,opt,name=duid,proto3" json:"duid,omitempty"`
	// Client unique identifier
	ClientId uint32 `protobuf:"varint,51,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// length of prefix
	PrefixLength uint32 `protobuf:"varint,52,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	// DHCPV6 IPv6 Prefix allotted to client
	Prefix string `protobuf:"bytes,53,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// DHCPv6 client/subscriber Vrf name
	VrfName string `protobuf:"bytes,54,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	// Client route lifetime
	Lifetime uint32 `protobuf:"varint,55,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// Client route remaining lifetime
	RemLifeTime uint32 `protobuf:"varint,56,opt,name=rem_life_time,json=remLifeTime,proto3" json:"rem_life_time,omitempty"`
	// Interface name
	InterfaceName string `protobuf:"bytes,57,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// Next hop is our address
	NextHopAddr string `protobuf:"bytes,58,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	// IA_ID of this IA
	IaId uint32 `protobuf:"varint,59,opt,name=ia_id,json=iaId,proto3" json:"ia_id,omitempty"`
	// Relay Profile name
	RelayProfileName string `protobuf:"bytes,60,opt,name=relay_profile_name,json=relayProfileName,proto3" json:"relay_profile_name,omitempty"`
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) Reset() {
	*x = Ipv6Dhcpv6DRelayRouteBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipv6Dhcpv6DRelayRouteBinding) ProtoMessage() {}

func (x *Ipv6Dhcpv6DRelayRouteBinding) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipv6Dhcpv6DRelayRouteBinding.ProtoReflect.Descriptor instead.
func (*Ipv6Dhcpv6DRelayRouteBinding) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescGZIP(), []int{1}
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetDuid() string {
	if x != nil {
		return x.Duid
	}
	return ""
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetClientId() uint32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetPrefixLength() uint32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetVrfName() string {
	if x != nil {
		return x.VrfName
	}
	return ""
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetLifetime() uint32 {
	if x != nil {
		return x.Lifetime
	}
	return 0
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetRemLifeTime() uint32 {
	if x != nil {
		return x.RemLifeTime
	}
	return 0
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetNextHopAddr() string {
	if x != nil {
		return x.NextHopAddr
	}
	return ""
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetIaId() uint32 {
	if x != nil {
		return x.IaId
	}
	return 0
}

func (x *Ipv6Dhcpv6DRelayRouteBinding) GetRelayProfileName() string {
	if x != nil {
		return x.RelayProfileName
	}
	return ""
}

var File_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDesc = []byte{
	0x0a, 0x78, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69,
	0x70, 0x76, 0x36, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x2f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64,
	0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x51, 0x63, 0x69, 0x73, 0x63,
	0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6e, 0x65,
	0x77, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x64,
	0x68, 0x63, 0x70, 0x76, 0x36, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x61, 0x0a,
	0x25, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36, 0x64, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0xf9, 0x02, 0x0a, 0x20, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x64, 0x68, 0x63, 0x70, 0x76, 0x36,
	0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x75, 0x69, 0x64, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x33, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x72, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x72, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65,
	0x6d, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x38, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x4c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x39, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65,
	0x78, 0x74, 0x48, 0x6f, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x69, 0x61, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescData = file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDesc
)

func file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescData)
	})
	return file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDescData
}

var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_goTypes = []interface{}{
	(*Ipv6Dhcpv6DRelayRouteBinding_KEYS)(nil), // 0: cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.binding.clients.client.ipv6_dhcpv6d_relay_route_binding_KEYS
	(*Ipv6Dhcpv6DRelayRouteBinding)(nil),      // 1: cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.binding.clients.client.ipv6_dhcpv6d_relay_route_binding
}
var file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_init()
}
func file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_init() {
	if File_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv6Dhcpv6DRelayRouteBinding_KEYS); i {
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
		file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipv6Dhcpv6DRelayRouteBinding); i {
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
			RawDescriptor: file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto = out.File
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_rawDesc = nil
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_goTypes = nil
	file_cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client_ipv6_dhcpv6d_relay_route_binding_proto_depIdxs = nil
}
