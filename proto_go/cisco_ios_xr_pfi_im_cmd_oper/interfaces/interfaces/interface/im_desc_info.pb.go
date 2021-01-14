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
// source: cisco_ios_xr_pfi_im_cmd_oper/interfaces/interfaces/interface/im_desc_info.proto

package cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface

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

// Interface description information
type ImDescInfo_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *ImDescInfo_KEYS) Reset() {
	*x = ImDescInfo_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImDescInfo_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImDescInfo_KEYS) ProtoMessage() {}

func (x *ImDescInfo_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImDescInfo_KEYS.ProtoReflect.Descriptor instead.
func (*ImDescInfo_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescGZIP(), []int{0}
}

func (x *ImDescInfo_KEYS) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

type ImDescInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interface
	Interface string `protobuf:"bytes,50,opt,name=interface,proto3" json:"interface,omitempty"`
	// Operational state with no translation of error disable or shutdown
	State string `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	// Line protocol state with no translation of error disable or shutdown
	LineState string `protobuf:"bytes,52,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	// Interface description string
	Description string `protobuf:"bytes,53,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ImDescInfo) Reset() {
	*x = ImDescInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImDescInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImDescInfo) ProtoMessage() {}

func (x *ImDescInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImDescInfo.ProtoReflect.Descriptor instead.
func (*ImDescInfo) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescGZIP(), []int{1}
}

func (x *ImDescInfo) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *ImDescInfo) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ImDescInfo) GetLineState() string {
	if x != nil {
		return x.LineState
	}
	return ""
}

func (x *ImDescInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x70,
	0x66, 0x69, 0x5f, 0x69, 0x6d, 0x5f, 0x63, 0x6d, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x69,
	0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x3c, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f,
	0x70, 0x66, 0x69, 0x5f, 0x69, 0x6d, 0x5f, 0x63, 0x6d, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x22,
	0x3a, 0x0a, 0x11, 0x69, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x4b, 0x45, 0x59, 0x53, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0c,
	0x69, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x34,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x35,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescData = file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDesc
)

func file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescData)
	})
	return file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDescData
}

var file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_goTypes = []interface{}{
	(*ImDescInfo_KEYS)(nil), // 0: cisco_ios_xr_pfi_im_cmd_oper.interfaces.interfaces.interface.im_desc_info_KEYS
	(*ImDescInfo)(nil),      // 1: cisco_ios_xr_pfi_im_cmd_oper.interfaces.interfaces.interface.im_desc_info
}
var file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_init()
}
func file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_init() {
	if File_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImDescInfo_KEYS); i {
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
		file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImDescInfo); i {
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
			RawDescriptor: file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto = out.File
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_rawDesc = nil
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_goTypes = nil
	file_cisco_ios_xr_pfi_im_cmd_oper_interfaces_interfaces_interface_im_desc_info_proto_depIdxs = nil
}
