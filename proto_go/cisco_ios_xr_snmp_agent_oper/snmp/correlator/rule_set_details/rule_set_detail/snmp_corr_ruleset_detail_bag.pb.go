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
// source: cisco_ios_xr_snmp_agent_oper/snmp/correlator/rule_set_details/rule_set_detail/snmp_corr_ruleset_detail_bag.proto

package cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail

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

// Correlation Ruleset detail information
type SnmpCorrRulesetDetailBag_KEYS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RulesetName string `protobuf:"bytes,1,opt,name=ruleset_name,json=rulesetName,proto3" json:"ruleset_name,omitempty"`
}

func (x *SnmpCorrRulesetDetailBag_KEYS) Reset() {
	*x = SnmpCorrRulesetDetailBag_KEYS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpCorrRulesetDetailBag_KEYS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpCorrRulesetDetailBag_KEYS) ProtoMessage() {}

func (x *SnmpCorrRulesetDetailBag_KEYS) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpCorrRulesetDetailBag_KEYS.ProtoReflect.Descriptor instead.
func (*SnmpCorrRulesetDetailBag_KEYS) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescGZIP(), []int{0}
}

func (x *SnmpCorrRulesetDetailBag_KEYS) GetRulesetName() string {
	if x != nil {
		return x.RulesetName
	}
	return ""
}

type SnmpCorrRulesetDetailBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ruleset Name
	RuleSetName string `protobuf:"bytes,50,opt,name=rule_set_name,json=ruleSetName,proto3" json:"rule_set_name,omitempty"`
	// Rules contained in a ruleset
	Rules []*SnmpCorrRuleSummaryBag `protobuf:"bytes,51,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *SnmpCorrRulesetDetailBag) Reset() {
	*x = SnmpCorrRulesetDetailBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpCorrRulesetDetailBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpCorrRulesetDetailBag) ProtoMessage() {}

func (x *SnmpCorrRulesetDetailBag) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpCorrRulesetDetailBag.ProtoReflect.Descriptor instead.
func (*SnmpCorrRulesetDetailBag) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescGZIP(), []int{1}
}

func (x *SnmpCorrRulesetDetailBag) GetRuleSetName() string {
	if x != nil {
		return x.RuleSetName
	}
	return ""
}

func (x *SnmpCorrRulesetDetailBag) GetRules() []*SnmpCorrRuleSummaryBag {
	if x != nil {
		return x.Rules
	}
	return nil
}

// Correlation Rule summary information
type SnmpCorrRuleSummaryBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Correlation Rule Name
	RuleName string `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	// Applied state of the rule It could be not applied, applied or applied to all
	RuleState string `protobuf:"bytes,2,opt,name=rule_state,json=ruleState,proto3" json:"rule_state,omitempty"`
	// Number of buffered traps correlated to this rule
	BufferedTrapsCount uint32 `protobuf:"varint,3,opt,name=buffered_traps_count,json=bufferedTrapsCount,proto3" json:"buffered_traps_count,omitempty"`
}

func (x *SnmpCorrRuleSummaryBag) Reset() {
	*x = SnmpCorrRuleSummaryBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnmpCorrRuleSummaryBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnmpCorrRuleSummaryBag) ProtoMessage() {}

func (x *SnmpCorrRuleSummaryBag) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnmpCorrRuleSummaryBag.ProtoReflect.Descriptor instead.
func (*SnmpCorrRuleSummaryBag) Descriptor() ([]byte, []int) {
	return file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescGZIP(), []int{2}
}

func (x *SnmpCorrRuleSummaryBag) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *SnmpCorrRuleSummaryBag) GetRuleState() string {
	if x != nil {
		return x.RuleState
	}
	return ""
}

func (x *SnmpCorrRuleSummaryBag) GetBufferedTrapsCount() uint32 {
	if x != nil {
		return x.BufferedTrapsCount
	}
	return 0
}

var File_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto protoreflect.FileDescriptor

var file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDesc = []byte{
	0x0a, 0x70, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x73,
	0x6e, 0x6d, 0x70, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73,
	0x6e, 0x6d, 0x70, 0x2f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f,
	0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x4d, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72,
	0x5f, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72,
	0x2e, 0x73, 0x6e, 0x6d, 0x70, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x22, 0x46, 0x0a, 0x21, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x61,
	0x67, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x1c, 0x73, 0x6e,
	0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x62, 0x61, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7f,
	0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x33, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x69, 0x2e,
	0x63, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x69, 0x6f, 0x73, 0x5f, 0x78, 0x72, 0x5f, 0x73, 0x6e, 0x6d,
	0x70, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x6e, 0x6d,
	0x70, 0x2e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x73, 0x6e,
	0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x67, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22,
	0x8a, 0x01, 0x0a, 0x1a, 0x73, 0x6e, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x72, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x67, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x70, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescOnce sync.Once
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescData = file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDesc
)

func file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescGZIP() []byte {
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescOnce.Do(func() {
		file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescData)
	})
	return file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDescData
}

var file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_goTypes = []interface{}{
	(*SnmpCorrRulesetDetailBag_KEYS)(nil), // 0: cisco_ios_xr_snmp_agent_oper.snmp.correlator.rule_set_details.rule_set_detail.snmp_corr_ruleset_detail_bag_KEYS
	(*SnmpCorrRulesetDetailBag)(nil),      // 1: cisco_ios_xr_snmp_agent_oper.snmp.correlator.rule_set_details.rule_set_detail.snmp_corr_ruleset_detail_bag
	(*SnmpCorrRuleSummaryBag)(nil),        // 2: cisco_ios_xr_snmp_agent_oper.snmp.correlator.rule_set_details.rule_set_detail.snmp_corr_rule_summary_bag
}
var file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_depIdxs = []int32{
	2, // 0: cisco_ios_xr_snmp_agent_oper.snmp.correlator.rule_set_details.rule_set_detail.snmp_corr_ruleset_detail_bag.rules:type_name -> cisco_ios_xr_snmp_agent_oper.snmp.correlator.rule_set_details.rule_set_detail.snmp_corr_rule_summary_bag
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_init()
}
func file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_init() {
	if File_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpCorrRulesetDetailBag_KEYS); i {
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
		file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpCorrRulesetDetailBag); i {
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
		file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnmpCorrRuleSummaryBag); i {
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
			RawDescriptor: file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_goTypes,
		DependencyIndexes: file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_depIdxs,
		MessageInfos:      file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_msgTypes,
	}.Build()
	File_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto = out.File
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_rawDesc = nil
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_goTypes = nil
	file_cisco_ios_xr_snmp_agent_oper_snmp_correlator_rule_set_details_rule_set_detail_snmp_corr_ruleset_detail_bag_proto_depIdxs = nil
}
