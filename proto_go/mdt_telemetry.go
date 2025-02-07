package telemetry
import ("reflect"
	cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters "github.com/kentik/bigmuddy-network-telemetry-proto/proto_go/cisco_ios_xr_infra_statsd_oper/infra_statistics/interfaces/interface/latest/generic_counters")
//
// ProtoMsgs are an ordered collection of reflect.Type objects. These
// objects can be used to unmarshal a binary protobuf stream into
// a go structure. The index into the array is the ProtoMsgType.
type ProtoMsgs struct {
	msgset []reflect.Type
}

type ProtoMsgType int
const (
	// PROTO_CONTENT_MSG corresponds to content field serialised
	// in TelemetryRowGPB message.
	PROTO_CONTENT_MSG = 0
	// PROTO_KEYS_MSG corresponds to keys field serialised
	// in TelemetryRowGPB message.
	PROTO_KEYS_MSG    = 1
	//  PROTO_LEGACY_COMBINED_KEYS_AND_CONTENT_MSG, not populated
	//  anymore, but used to support legacy cases where keys
	//  and content used to be combined.
	PROTO_LEGACY_COMBINED_KEYS_AND_CONTENT_MSG = 2
)

//
// ProtoKey identifies which proto messages we are looking for. The
// encoding_path field in telemetry containing the yang path where the
// content in the message is rooted at is used to set
// EncodingPath. The Version field can be extracted from the Telemetry
// header and reflects the yang model version. If no Version is
// provided, the default for the package is used. As of IOS-XR 6.1.1,
// IOS-XR telemetry does not produce a version field yet in the
// Telemetry header.
type ProtoKey struct {
	EncodingPath string
	Version string
}

// EncodingPathsSupported provides a list of gather path ProtoKeys
// supported by package (for reporting purposes).
func EncodingPathsSupported() []*ProtoKey {
	paths := make([]*ProtoKey, len(schema2message))
	i := 0
	for protoKey := range schema2message {
		path := protoKey
		paths[i] = &path
		i++
	}
	return paths
}

//
// EncodingPathToMessageReflectionSet returns an opaque structure
// tracking relect.Type for Messages pertaining to the ProtoKey p.
func EncodingPathToMessageReflectionSet(p *ProtoKey) *ProtoMsgs {

	//
	// Given the underlying tree nature of the paths, it might be
	// more appropriate to move to a more efficient data structure
	// as the number of paths increases.
	set := &ProtoMsgs{
		msgset: schema2message[*p],
	}

	if set.msgset == nil {
		return nil
	}

	return set
}

//
// MessageReflection returns reflection type of the object pertaining
// to the encoding path, and qualified by ProtoMsgType t
func (p *ProtoMsgs) MessageReflection(t ProtoMsgType) reflect.Type {

	l := p.msgset

	if l == nil {
		return nil
	}

	if len(l) > int(t) {
		return l[t]
	}

	return nil
}

// It would be very useful if the registry of messages where
// exported. Something like:
//
// func RegisteredMessages() map[string]reflect.Type {
//	return protoTypes
// }
//
// We could then simply use the map. Until then...
//

var schema2message = map[ProtoKey][]reflect.Type{ProtoKey{EncodingPath:"Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters",Version:""}:[]reflect.Type{reflect.TypeOf((*cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters.IfstatsbagGeneric)(nil)),reflect.TypeOf((*cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters.IfstatsbagGeneric_KEYS)(nil))}}