package telemetry
import ("reflect"
	cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters "github.com/kentik/bigmuddy-network-telemetry-proto/proto_go/cisco_ios_xr_infra_statsd_oper/infra_statistics/interfaces/interface/latest/generic_counters")
import (
        "encoding/json"
	"fmt"
	"testing"
	"time"
	"math/rand"
	"io/ioutil"
)

func getProtoKeysTest() []*ProtoKey {
	keys := make([]*ProtoKey, 0, len(schema2message))
	for k, _ := range schema2message {
		keys = append(keys, &k)
	}
	return keys
}

func BenchmarkMessageSetLookup(b *testing.B) {
	//
	// Build random sequence of keys (don't just rely on map
	// random iteration)
	rand.Seed(time.Now().UnixNano())
	max := b.N
	keys := getProtoKeysTest()
	indexMap := rand.Perm(len(keys))
	testPaths := make([]*ProtoKey, 0, max)

	k := 0
	for i := 0; i < max; i++ {
		if k >= len(keys) {
			k = 0
		}
		testPaths = append(testPaths, keys[indexMap[k]])
		k++
	}

	b.ResetTimer()
	for i := 0; i < max; i++ {
		s := EncodingPathToMessageReflectionSet(testPaths[i])
		if s == nil {
			b.Fatalf("Expected msg set for %v",
				&testPaths[i])
		}
	}
}

func ExampleEncodingPathToMessageReflectionSet() {
	var t reflect.Type
	//
	// Extract encoding_path field from Telemetry message,
	// and use pass it in
	s := EncodingPathToMessageReflectionSet(
		&ProtoKey{
			EncodingPath: "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters",
			Version: ""})

	// Use returned object to extract the KEYs message to
	// unmarshal keys field from TelemetryRowGPB
	t = s.MessageReflection(PROTO_KEYS_MSG)
	fmt.Print(t)
	//Output: *cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters.IfstatsbagGeneric_KEYS
}

func ExampleProtoMsgs_MessageReflection() {
	var t reflect.Type
	//
	// Extract encoding_path field from Telemetry message,
	// and use pass it in
	s := EncodingPathToMessageReflectionSet(
		&ProtoKey{
			EncodingPath: "Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters",
			Version: ""})

	// Use returned object to extract the KEYs message to
	// unmarshal keys field from TelemetryRowGPB
	t = s.MessageReflection(PROTO_KEYS_MSG)
	fmt.Print(t)
	//Output: *cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters.IfstatsbagGeneric_KEYS
}

type Yang2ProtoTestTableEntry struct {
	EncodingPath string
	Version      string
	Present      bool
	Keys         reflect.Type
	Content      reflect.Type
}

func yang2ProtoTestHelper(tb *testing.T, table *[]Yang2ProtoTestTableEntry) {
	for _, subt := range *table {
		tb.Run(
			fmt.Sprintf("%s%s", subt.EncodingPath, subt.Version),
			func(t *testing.T) {
				s := EncodingPathToMessageReflectionSet(
					&ProtoKey{
						EncodingPath: subt.EncodingPath,
						Version:      subt.Version})
				if subt.Present {
					if s == nil {
						t.Fatalf(
							"Expected msg set for %s%s, didn't get one",
							subt.EncodingPath, subt.Version)
					}

					msg := s.MessageReflection(PROTO_KEYS_MSG)
					if msg != subt.Keys {
						t.Fatalf(
							"Expected keys %s, and got %+v",
							subt.Keys, msg)
					}
					msg = s.MessageReflection(PROTO_CONTENT_MSG)
					if msg != subt.Content {
						t.Fatalf(
							"Expected content %s, and got %+v",
							subt.Content, msg)
					}
					// Negative test
					msg = s.MessageReflection(PROTO_CONTENT_MSG + 1000)
					if msg != nil {
						t.Fatalf(
							"Expected no msg at index, and got %+v", msg)
					}

				} else {
					if s != nil {
						t.Fatalf(
							"Didn't expect msg set for %s%s, got %+v",
							subt.EncodingPath, subt.Version, s)
					}
				}
			})
	}
}

func TestYang2Proto(tb *testing.T) {
	yang2ProtoTestHelper(tb, &Yang2ProtoTestTable)
}


func TestBasePathXlationMap(t *testing.T) {

    bpxJSON, err := ioutil.ReadFile("basepathxlation.json")
    if err != nil {
        t.Fatal("Failed to open base path xlation map basepathxlation.json:", err)
    }
    basePathXlation := map[string]string{}
    err = json.Unmarshal(bpxJSON, &basePathXlation)
    if err != nil {
        t.Fatal("Failed to unmarshal to expected structure:", err)
    }
}

var Yang2ProtoTestTable = []Yang2ProtoTestTableEntry{
	{"", "", false, nil, nil},
	{"Cisco-IOS-XR-infra-statsd-oper:infra-statistics", "", false, nil, nil},
	{"Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-NONEXIST", "", false, nil, nil},
	{"Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters/TOOLONG", "", false, nil, nil},
{"Cisco-IOS-XR-infra-statsd-oper:infra-statistics/interfaces/interface/latest/generic-counters","",true,reflect.TypeOf((*cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters.IfstatsbagGeneric_KEYS)(nil)),reflect.TypeOf((*cisco_ios_xr_infra_statsd_oper_infra_statistics_interfaces_interface_latest_generic_counters.IfstatsbagGeneric)(nil))}}