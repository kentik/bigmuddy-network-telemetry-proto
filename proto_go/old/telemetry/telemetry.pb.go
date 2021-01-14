// ----------------------------------------------------------------------------
// telemetry.proto - Telemetry protobuf definitions
//
// August 2015, Robert Wills
//
// Copyright (c) 2015 by Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: old/telemetry/telemetry.proto

// Package with obsolete definition of Model Driven Telemetry GPB format message.

package telemetry

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

type TelemetryHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encoding   uint32            `protobuf:"varint,1,opt,name=encoding,proto3" json:"encoding,omitempty"`
	PolicyName string            `protobuf:"bytes,2,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	Version    string            `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Identifier string            `protobuf:"bytes,4,opt,name=identifier,proto3" json:"identifier,omitempty"`
	StartTime  uint64            `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    uint64            `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Tables     []*TelemetryTable `protobuf:"bytes,7,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *TelemetryHeader) Reset() {
	*x = TelemetryHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_telemetry_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryHeader) ProtoMessage() {}

func (x *TelemetryHeader) ProtoReflect() protoreflect.Message {
	mi := &file_old_telemetry_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryHeader.ProtoReflect.Descriptor instead.
func (*TelemetryHeader) Descriptor() ([]byte, []int) {
	return file_old_telemetry_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *TelemetryHeader) GetEncoding() uint32 {
	if x != nil {
		return x.Encoding
	}
	return 0
}

func (x *TelemetryHeader) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *TelemetryHeader) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TelemetryHeader) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *TelemetryHeader) GetStartTime() uint64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TelemetryHeader) GetEndTime() uint64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *TelemetryHeader) GetTables() []*TelemetryTable {
	if x != nil {
		return x.Tables
	}
	return nil
}

type TelemetryTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyPath string   `protobuf:"bytes,1,opt,name=policy_path,json=policyPath,proto3" json:"policy_path,omitempty"`
	Row        [][]byte `protobuf:"bytes,2,rep,name=row,proto3" json:"row,omitempty"`
}

func (x *TelemetryTable) Reset() {
	*x = TelemetryTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_telemetry_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryTable) ProtoMessage() {}

func (x *TelemetryTable) ProtoReflect() protoreflect.Message {
	mi := &file_old_telemetry_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryTable.ProtoReflect.Descriptor instead.
func (*TelemetryTable) Descriptor() ([]byte, []int) {
	return file_old_telemetry_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetryTable) GetPolicyPath() string {
	if x != nil {
		return x.PolicyPath
	}
	return ""
}

func (x *TelemetryTable) GetRow() [][]byte {
	if x != nil {
		return x.Row
	}
	return nil
}

var File_old_telemetry_telemetry_proto protoreflect.FileDescriptor

var file_old_telemetry_telemetry_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6f, 0x6c, 0x64, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0xf5, 0x01, 0x0a, 0x0f, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x31, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x22, 0x43, 0x0a, 0x0e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_old_telemetry_telemetry_proto_rawDescOnce sync.Once
	file_old_telemetry_telemetry_proto_rawDescData = file_old_telemetry_telemetry_proto_rawDesc
)

func file_old_telemetry_telemetry_proto_rawDescGZIP() []byte {
	file_old_telemetry_telemetry_proto_rawDescOnce.Do(func() {
		file_old_telemetry_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_old_telemetry_telemetry_proto_rawDescData)
	})
	return file_old_telemetry_telemetry_proto_rawDescData
}

var file_old_telemetry_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_old_telemetry_telemetry_proto_goTypes = []interface{}{
	(*TelemetryHeader)(nil), // 0: telemetry.TelemetryHeader
	(*TelemetryTable)(nil),  // 1: telemetry.TelemetryTable
}
var file_old_telemetry_telemetry_proto_depIdxs = []int32{
	1, // 0: telemetry.TelemetryHeader.tables:type_name -> telemetry.TelemetryTable
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_old_telemetry_telemetry_proto_init() }
func file_old_telemetry_telemetry_proto_init() {
	if File_old_telemetry_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_old_telemetry_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryHeader); i {
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
		file_old_telemetry_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryTable); i {
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
			RawDescriptor: file_old_telemetry_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_old_telemetry_telemetry_proto_goTypes,
		DependencyIndexes: file_old_telemetry_telemetry_proto_depIdxs,
		MessageInfos:      file_old_telemetry_telemetry_proto_msgTypes,
	}.Build()
	File_old_telemetry_telemetry_proto = out.File
	file_old_telemetry_telemetry_proto_rawDesc = nil
	file_old_telemetry_telemetry_proto_goTypes = nil
	file_old_telemetry_telemetry_proto_depIdxs = nil
}
