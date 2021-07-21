// Code generated by protoc-gen-go.
// source: ifstatsbag_generic.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters is a generated protocol buffer package.

It is generated from these files:
	ifstatsbag_generic.proto

It has these top-level messages:
	IfstatsbagGeneric_KEYS
	IfstatsbagGeneric
*/
package cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Generic set of interface counters
type IfstatsbagGeneric_KEYS struct {
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *IfstatsbagGeneric_KEYS) Reset()                    { *m = IfstatsbagGeneric_KEYS{} }
func (m *IfstatsbagGeneric_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IfstatsbagGeneric_KEYS) ProtoMessage()               {}
func (*IfstatsbagGeneric_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IfstatsbagGeneric_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IfstatsbagGeneric struct {
	// Packets received
	PacketsReceived uint64 `protobuf:"varint,50,opt,name=packets_received,json=packetsReceived" json:"packets_received,omitempty"`
	// Bytes received
	BytesReceived uint64 `protobuf:"varint,51,opt,name=bytes_received,json=bytesReceived" json:"bytes_received,omitempty"`
	// Packets sent
	PacketsSent uint64 `protobuf:"varint,52,opt,name=packets_sent,json=packetsSent" json:"packets_sent,omitempty"`
	// Bytes sent
	BytesSent uint64 `protobuf:"varint,53,opt,name=bytes_sent,json=bytesSent" json:"bytes_sent,omitempty"`
	// Multicast packets received
	MulticastPacketsReceived uint64 `protobuf:"varint,54,opt,name=multicast_packets_received,json=multicastPacketsReceived" json:"multicast_packets_received,omitempty"`
	// Broadcast packets received
	BroadcastPacketsReceived uint64 `protobuf:"varint,55,opt,name=broadcast_packets_received,json=broadcastPacketsReceived" json:"broadcast_packets_received,omitempty"`
	// Multicast packets sent
	MulticastPacketsSent uint64 `protobuf:"varint,56,opt,name=multicast_packets_sent,json=multicastPacketsSent" json:"multicast_packets_sent,omitempty"`
	// Broadcast packets sent
	BroadcastPacketsSent uint64 `protobuf:"varint,57,opt,name=broadcast_packets_sent,json=broadcastPacketsSent" json:"broadcast_packets_sent,omitempty"`
	// Total output drops
	OutputDrops uint32 `protobuf:"varint,58,opt,name=output_drops,json=outputDrops" json:"output_drops,omitempty"`
	// Output queue drops
	OutputQueueDrops uint32 `protobuf:"varint,59,opt,name=output_queue_drops,json=outputQueueDrops" json:"output_queue_drops,omitempty"`
	// Total input drops
	InputDrops uint32 `protobuf:"varint,60,opt,name=input_drops,json=inputDrops" json:"input_drops,omitempty"`
	// Input queue drops
	InputQueueDrops uint32 `protobuf:"varint,61,opt,name=input_queue_drops,json=inputQueueDrops" json:"input_queue_drops,omitempty"`
	// Received runt packets
	RuntPacketsReceived uint32 `protobuf:"varint,62,opt,name=runt_packets_received,json=runtPacketsReceived" json:"runt_packets_received,omitempty"`
	// Received giant packets
	GiantPacketsReceived uint32 `protobuf:"varint,63,opt,name=giant_packets_received,json=giantPacketsReceived" json:"giant_packets_received,omitempty"`
	// Received throttled packets
	ThrottledPacketsReceived uint32 `protobuf:"varint,64,opt,name=throttled_packets_received,json=throttledPacketsReceived" json:"throttled_packets_received,omitempty"`
	// Received parity packets
	ParityPacketsReceived uint32 `protobuf:"varint,65,opt,name=parity_packets_received,json=parityPacketsReceived" json:"parity_packets_received,omitempty"`
	// Unknown protocol packets received
	UnknownProtocolPacketsReceived uint32 `protobuf:"varint,66,opt,name=unknown_protocol_packets_received,json=unknownProtocolPacketsReceived" json:"unknown_protocol_packets_received,omitempty"`
	// Total input errors
	InputErrors uint32 `protobuf:"varint,67,opt,name=input_errors,json=inputErrors" json:"input_errors,omitempty"`
	// Input CRC errors
	CrcErrors uint32 `protobuf:"varint,68,opt,name=crc_errors,json=crcErrors" json:"crc_errors,omitempty"`
	// Input overruns
	InputOverruns uint32 `protobuf:"varint,69,opt,name=input_overruns,json=inputOverruns" json:"input_overruns,omitempty"`
	// Framing-errors received
	FramingErrorsReceived uint32 `protobuf:"varint,70,opt,name=framing_errors_received,json=framingErrorsReceived" json:"framing_errors_received,omitempty"`
	// Input ignored packets
	InputIgnoredPackets uint32 `protobuf:"varint,71,opt,name=input_ignored_packets,json=inputIgnoredPackets" json:"input_ignored_packets,omitempty"`
	// Input aborts
	InputAborts uint32 `protobuf:"varint,72,opt,name=input_aborts,json=inputAborts" json:"input_aborts,omitempty"`
	// Total output errors
	OutputErrors uint32 `protobuf:"varint,73,opt,name=output_errors,json=outputErrors" json:"output_errors,omitempty"`
	// Output underruns
	OutputUnderruns uint32 `protobuf:"varint,74,opt,name=output_underruns,json=outputUnderruns" json:"output_underruns,omitempty"`
	// Output buffer failures
	OutputBufferFailures uint32 `protobuf:"varint,75,opt,name=output_buffer_failures,json=outputBufferFailures" json:"output_buffer_failures,omitempty"`
	// Output buffers swapped out
	OutputBuffersSwappedOut uint32 `protobuf:"varint,76,opt,name=output_buffers_swapped_out,json=outputBuffersSwappedOut" json:"output_buffers_swapped_out,omitempty"`
	// Applique
	Applique uint32 `protobuf:"varint,77,opt,name=applique" json:"applique,omitempty"`
	// Number of board resets
	Resets uint32 `protobuf:"varint,78,opt,name=resets" json:"resets,omitempty"`
	// Carrier transitions
	CarrierTransitions uint32 `protobuf:"varint,79,opt,name=carrier_transitions,json=carrierTransitions" json:"carrier_transitions,omitempty"`
	// Availability bit mask
	AvailabilityFlag uint32 `protobuf:"varint,80,opt,name=availability_flag,json=availabilityFlag" json:"availability_flag,omitempty"`
	// Time when counters were last written (in seconds)
	LastDataTime uint32 `protobuf:"varint,81,opt,name=last_data_time,json=lastDataTime" json:"last_data_time,omitempty"`
	// Number of seconds since last clear counters
	SecondsSinceLastClearCounters uint32 `protobuf:"varint,82,opt,name=seconds_since_last_clear_counters,json=secondsSinceLastClearCounters" json:"seconds_since_last_clear_counters,omitempty"`
	// SysUpTime when counters were last reset (in seconds)
	LastDiscontinuityTime uint32 `protobuf:"varint,83,opt,name=last_discontinuity_time,json=lastDiscontinuityTime" json:"last_discontinuity_time,omitempty"`
	// Seconds since packet received
	SecondsSincePacketReceived uint32 `protobuf:"varint,84,opt,name=seconds_since_packet_received,json=secondsSincePacketReceived" json:"seconds_since_packet_received,omitempty"`
	// Seconds since packet sent
	SecondsSincePacketSent uint32 `protobuf:"varint,85,opt,name=seconds_since_packet_sent,json=secondsSincePacketSent" json:"seconds_since_packet_sent,omitempty"`
}

func (m *IfstatsbagGeneric) Reset()                    { *m = IfstatsbagGeneric{} }
func (m *IfstatsbagGeneric) String() string            { return proto.CompactTextString(m) }
func (*IfstatsbagGeneric) ProtoMessage()               {}
func (*IfstatsbagGeneric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IfstatsbagGeneric) GetPacketsReceived() uint64 {
	if m != nil {
		return m.PacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetPacketsSent() uint64 {
	if m != nil {
		return m.PacketsSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetMulticastPacketsReceived() uint64 {
	if m != nil {
		return m.MulticastPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBroadcastPacketsReceived() uint64 {
	if m != nil {
		return m.BroadcastPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetMulticastPacketsSent() uint64 {
	if m != nil {
		return m.MulticastPacketsSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetBroadcastPacketsSent() uint64 {
	if m != nil {
		return m.BroadcastPacketsSent
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputDrops() uint32 {
	if m != nil {
		return m.OutputDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputQueueDrops() uint32 {
	if m != nil {
		return m.OutputQueueDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputDrops() uint32 {
	if m != nil {
		return m.InputDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputQueueDrops() uint32 {
	if m != nil {
		return m.InputQueueDrops
	}
	return 0
}

func (m *IfstatsbagGeneric) GetRuntPacketsReceived() uint32 {
	if m != nil {
		return m.RuntPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetGiantPacketsReceived() uint32 {
	if m != nil {
		return m.GiantPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetThrottledPacketsReceived() uint32 {
	if m != nil {
		return m.ThrottledPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetParityPacketsReceived() uint32 {
	if m != nil {
		return m.ParityPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetUnknownProtocolPacketsReceived() uint32 {
	if m != nil {
		return m.UnknownProtocolPacketsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputErrors() uint32 {
	if m != nil {
		return m.InputErrors
	}
	return 0
}

func (m *IfstatsbagGeneric) GetCrcErrors() uint32 {
	if m != nil {
		return m.CrcErrors
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputOverruns() uint32 {
	if m != nil {
		return m.InputOverruns
	}
	return 0
}

func (m *IfstatsbagGeneric) GetFramingErrorsReceived() uint32 {
	if m != nil {
		return m.FramingErrorsReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputIgnoredPackets() uint32 {
	if m != nil {
		return m.InputIgnoredPackets
	}
	return 0
}

func (m *IfstatsbagGeneric) GetInputAborts() uint32 {
	if m != nil {
		return m.InputAborts
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputErrors() uint32 {
	if m != nil {
		return m.OutputErrors
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputUnderruns() uint32 {
	if m != nil {
		return m.OutputUnderruns
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputBufferFailures() uint32 {
	if m != nil {
		return m.OutputBufferFailures
	}
	return 0
}

func (m *IfstatsbagGeneric) GetOutputBuffersSwappedOut() uint32 {
	if m != nil {
		return m.OutputBuffersSwappedOut
	}
	return 0
}

func (m *IfstatsbagGeneric) GetApplique() uint32 {
	if m != nil {
		return m.Applique
	}
	return 0
}

func (m *IfstatsbagGeneric) GetResets() uint32 {
	if m != nil {
		return m.Resets
	}
	return 0
}

func (m *IfstatsbagGeneric) GetCarrierTransitions() uint32 {
	if m != nil {
		return m.CarrierTransitions
	}
	return 0
}

func (m *IfstatsbagGeneric) GetAvailabilityFlag() uint32 {
	if m != nil {
		return m.AvailabilityFlag
	}
	return 0
}

func (m *IfstatsbagGeneric) GetLastDataTime() uint32 {
	if m != nil {
		return m.LastDataTime
	}
	return 0
}

func (m *IfstatsbagGeneric) GetSecondsSinceLastClearCounters() uint32 {
	if m != nil {
		return m.SecondsSinceLastClearCounters
	}
	return 0
}

func (m *IfstatsbagGeneric) GetLastDiscontinuityTime() uint32 {
	if m != nil {
		return m.LastDiscontinuityTime
	}
	return 0
}

func (m *IfstatsbagGeneric) GetSecondsSincePacketReceived() uint32 {
	if m != nil {
		return m.SecondsSincePacketReceived
	}
	return 0
}

func (m *IfstatsbagGeneric) GetSecondsSincePacketSent() uint32 {
	if m != nil {
		return m.SecondsSincePacketSent
	}
	return 0
}

func init() {
	proto.RegisterType((*IfstatsbagGeneric_KEYS)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.latest.generic_counters.ifstatsbag_generic_KEYS")
	proto.RegisterType((*IfstatsbagGeneric)(nil), "cisco_ios_xr_infra_statsd_oper.infra_statistics.interfaces.interface.latest.generic_counters.ifstatsbag_generic")
}

func init() { proto.RegisterFile("ifstatsbag_generic.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x95, 0x7d, 0x6f, 0x1c, 0x35,
	0x10, 0xc6, 0x15, 0x09, 0x55, 0x64, 0x92, 0xeb, 0x8b, 0xdb, 0x24, 0xe6, 0xa4, 0x40, 0xae, 0x80,
	0x94, 0x02, 0x3a, 0xa4, 0x34, 0x14, 0x4a, 0x03, 0x34, 0xcd, 0x0b, 0x0d, 0x2d, 0x4d, 0x7a, 0x97,
	0xfe, 0x81, 0x84, 0x64, 0xf9, 0xbc, 0xde, 0xc3, 0xea, 0x9e, 0xbd, 0xf5, 0x4b, 0x4a, 0xbf, 0x23,
	0x1f, 0x0a, 0x79, 0xec, 0xec, 0x6d, 0x6f, 0xef, 0xbf, 0xe4, 0x79, 0x7e, 0xcf, 0xcc, 0xd8, 0x3b,
	0xbb, 0x07, 0x54, 0x95, 0xce, 0x73, 0xef, 0x26, 0x7c, 0xca, 0xa6, 0x52, 0x4b, 0xab, 0xc4, 0xb0,
	0xb6, 0xc6, 0x1b, 0xf2, 0xb7, 0x50, 0x4e, 0x18, 0xa6, 0x8c, 0x63, 0xff, 0x5a, 0xa6, 0x74, 0x69,
	0x39, 0x43, 0xb4, 0x60, 0xa6, 0x96, 0x76, 0x38, 0x57, 0x94, 0xf3, 0x4a, 0xb8, 0xa1, 0xd2, 0x5e,
	0xda, 0x92, 0x0b, 0xd9, 0xfa, 0x73, 0x58, 0x71, 0x2f, 0x9d, 0x1f, 0xe6, 0xd2, 0x4c, 0x98, 0x10,
	0x2d, 0x77, 0xff, 0x29, 0x6c, 0x75, 0x3b, 0xb3, 0x17, 0x27, 0x7f, 0x8d, 0xc9, 0xd7, 0x70, 0xb3,
	0xc9, 0x33, 0xcd, 0x67, 0x92, 0xae, 0xec, 0xac, 0xec, 0xae, 0x8e, 0x7a, 0x8d, 0xfa, 0x8a, 0xcf,
	0xe4, 0xfd, 0xff, 0x7a, 0x40, 0xba, 0x25, 0xc8, 0x03, 0xb8, 0x5d, 0x73, 0xf1, 0x56, 0x7a, 0xc7,
	0xac, 0x14, 0x52, 0x5d, 0xc9, 0x82, 0xee, 0xed, 0xac, 0xec, 0x7e, 0x32, 0xba, 0x95, 0xf5, 0x51,
	0x96, 0x63, 0xa3, 0xc9, 0x07, 0x2f, 0x5b, 0xe0, 0x43, 0x04, 0x7b, 0xa8, 0x36, 0xd8, 0x00, 0xd6,
	0xaf, 0x2b, 0x3a, 0xa9, 0x3d, 0xdd, 0x47, 0x68, 0x2d, 0x6b, 0x63, 0xa9, 0x3d, 0xd9, 0x06, 0x48,
	0x95, 0x10, 0xf8, 0x01, 0x81, 0x55, 0x54, 0xd0, 0x3e, 0x80, 0xfe, 0x2c, 0x54, 0x5e, 0x09, 0xee,
	0x3c, 0xeb, 0x4c, 0xf7, 0x08, 0x71, 0xda, 0x10, 0x17, 0x0b, 0x63, 0x1e, 0x40, 0x7f, 0x62, 0x0d,
	0x2f, 0x96, 0xa7, 0x7f, 0x4c, 0xe9, 0x86, 0x58, 0x4c, 0xef, 0xc3, 0x66, 0xb7, 0x37, 0x8e, 0xf9,
	0x13, 0x26, 0xef, 0x2d, 0xf6, 0xc5, 0x89, 0xf7, 0x61, 0xb3, 0xdb, 0x13, 0x53, 0x8f, 0x53, 0x6a,
	0xb1, 0x1f, 0xa6, 0x06, 0xb0, 0x6e, 0x82, 0xaf, 0x83, 0x67, 0x85, 0x35, 0xb5, 0xa3, 0x3f, 0xef,
	0xac, 0xec, 0xf6, 0x46, 0x6b, 0x49, 0x3b, 0x8e, 0x12, 0xf9, 0x0e, 0x48, 0x46, 0xde, 0x05, 0x19,
	0x64, 0x06, 0x9f, 0x20, 0x78, 0x3b, 0x39, 0xaf, 0xa3, 0x91, 0xe8, 0x2f, 0x60, 0x4d, 0xe9, 0x79,
	0xbd, 0x03, 0xc4, 0x00, 0xa5, 0x04, 0x7c, 0x03, 0x77, 0x12, 0xd0, 0xae, 0xf6, 0x0b, 0x62, 0xb7,
	0xd0, 0x68, 0x15, 0xdb, 0x83, 0x0d, 0x1b, 0xf4, 0x92, 0x2b, 0xfc, 0x15, 0xf9, 0xbb, 0xd1, 0x5c,
	0x72, 0x7b, 0x53, 0xc5, 0x97, 0x85, 0x7e, 0xc3, 0xd0, 0x3d, 0x74, 0x97, 0x3c, 0x31, 0xff, 0x8f,
	0x35, 0xde, 0x57, 0xb2, 0xe8, 0x26, 0x9f, 0x62, 0x92, 0x36, 0xc4, 0x62, 0xfa, 0x11, 0x6c, 0xd5,
	0xdc, 0x2a, 0xff, 0xa1, 0x1b, 0x3d, 0xc4, 0xe8, 0x46, 0xb2, 0x17, 0x73, 0x67, 0x30, 0x08, 0xfa,
	0xad, 0x36, 0xef, 0x35, 0xc3, 0x37, 0x58, 0x98, 0xaa, 0x5b, 0xe1, 0x19, 0x56, 0xf8, 0x3c, 0x83,
	0x17, 0x99, 0x5b, 0x2c, 0x35, 0x80, 0xf5, 0x74, 0xad, 0xd2, 0x5a, 0x63, 0x1d, 0x3d, 0x4a, 0x0f,
	0x12, 0xb5, 0x13, 0x94, 0xe2, 0xca, 0x0b, 0x2b, 0xae, 0x81, 0x63, 0x04, 0x56, 0x85, 0x15, 0xd9,
	0xc6, 0x97, 0x38, 0x56, 0x30, 0x57, 0xd2, 0xda, 0xa0, 0x1d, 0x3d, 0x41, 0xa4, 0x87, 0xea, 0x79,
	0x16, 0xe3, 0x59, 0x4b, 0xcb, 0x67, 0x4a, 0x4f, 0x73, 0xa5, 0xf9, 0xa4, 0xa7, 0xe9, 0xac, 0xd9,
	0x4e, 0x65, 0x9b, 0x01, 0xf7, 0x60, 0x23, 0x95, 0x57, 0x53, 0x6d, 0xec, 0xfc, 0x96, 0xe9, 0xef,
	0xe9, 0x59, 0xa2, 0x79, 0x96, 0xbc, 0x7c, 0xb8, 0xf9, 0xa1, 0xf8, 0xc4, 0x58, 0xef, 0xe8, 0xf3,
	0xd6, 0xa1, 0x0e, 0x51, 0x22, 0x5f, 0x42, 0x2f, 0x6f, 0x67, 0x3e, 0xd7, 0x19, 0x32, 0x79, 0xab,
	0xf3, 0xd1, 0x1e, 0x40, 0x5e, 0x54, 0x16, 0x74, 0x91, 0x0f, 0xf7, 0x47, 0x5a, 0xb9, 0xa4, 0xbf,
	0xb9, 0x96, 0xe3, 0xfa, 0x64, 0x74, 0x12, 0xca, 0x52, 0x5a, 0x56, 0x72, 0x55, 0x05, 0x2b, 0x1d,
	0x7d, 0x91, 0xd6, 0x27, 0xb9, 0xcf, 0xd0, 0x3c, 0xcd, 0x1e, 0x79, 0x02, 0xfd, 0x8f, 0x52, 0x8e,
	0xb9, 0xf7, 0xbc, 0xae, 0x65, 0xc1, 0x4c, 0xf0, 0xf4, 0x25, 0x26, 0xb7, 0xda, 0x49, 0x37, 0x4e,
	0xfe, 0x79, 0xf0, 0xa4, 0x0f, 0x9f, 0xf2, 0xba, 0xae, 0xd4, 0xbb, 0x20, 0xe9, 0x9f, 0x88, 0x36,
	0xff, 0x93, 0x4d, 0xb8, 0x61, 0xa5, 0x8b, 0xd7, 0xf4, 0x0a, 0x9d, 0xfc, 0x1f, 0xf9, 0x1e, 0xee,
	0x0a, 0x6e, 0xad, 0x92, 0x96, 0x79, 0xcb, 0xb5, 0x53, 0x5e, 0x19, 0xed, 0xe8, 0x39, 0x42, 0x24,
	0x5b, 0x97, 0x73, 0x87, 0x7c, 0x0b, 0x77, 0xf8, 0x15, 0x57, 0x15, 0x9f, 0xa8, 0x2a, 0x2e, 0x6a,
	0x59, 0xf1, 0x29, 0xbd, 0x48, 0x2f, 0x71, 0xdb, 0x38, 0xad, 0xf8, 0x94, 0x7c, 0x05, 0x37, 0xab,
	0xf8, 0x19, 0x29, 0xb8, 0xe7, 0xcc, 0xab, 0x99, 0xa4, 0xaf, 0xd3, 0xad, 0x46, 0xf5, 0x98, 0x7b,
	0x7e, 0xa9, 0x66, 0x92, 0x3c, 0x87, 0x81, 0x93, 0xc2, 0xe8, 0xc2, 0x31, 0xa7, 0xb4, 0x90, 0x0c,
	0x33, 0xa2, 0x92, 0xdc, 0x36, 0xbf, 0x1a, 0x74, 0x84, 0xc1, 0xed, 0x0c, 0x8e, 0x23, 0xf7, 0x92,
	0x3b, 0x7f, 0x14, 0xa9, 0xa3, 0x0c, 0xc5, 0x9d, 0x4a, 0xfd, 0xe2, 0xef, 0x97, 0xf6, 0x4a, 0x87,
	0x38, 0x22, 0x36, 0x1e, 0xa7, 0x9d, 0xc2, 0xc6, 0x6d, 0x17, 0x27, 0x38, 0x84, 0xed, 0x8f, 0x27,
	0x48, 0x3b, 0x35, 0xdf, 0xc8, 0x4b, 0x4c, 0xf7, 0xdb, 0xdd, 0xd3, 0x6e, 0x35, 0x6b, 0xf9, 0x18,
	0x3e, 0x5b, 0x5a, 0x02, 0xbf, 0x9c, 0x6f, 0x30, 0xbe, 0xd9, 0x8d, 0xc7, 0x6f, 0xe7, 0xe4, 0x06,
	0xbe, 0xb3, 0x0f, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x08, 0x28, 0xde, 0xab, 0x91, 0x07, 0x00,
	0x00,
}
