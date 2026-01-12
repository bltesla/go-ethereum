package rpc

import (
	"encoding/json"
	"testing"
)

type TestBase struct {
	Height int
	Type   string
}
type TestStruct struct {
	TestBase
	MyField       string
	Tagged        string `json:"tagged"`
	Nested        NestedStruct
	NestedTagged  NestedStruct `json:"nestedTagged"`
	HTTPServerUrl string
	REF           string
}

type NestedStruct struct {
	InnerField string
}

func TestDefaultMarshaling(t *testing.T) {
	val := TestStruct{
		TestBase:      TestBase{Height: 100, Type: "exampleType"},
		MyField:       "val1",
		Tagged:        "val2",
		Nested:        NestedStruct{InnerField: "inner1"},
		NestedTagged:  NestedStruct{InnerField: "inner2"},
		HTTPServerUrl: "http://example.com",
		REF:           "refvalue",
	}

	// We need to call the internal helper camelCaseKeys directly to verify its behavior
	// because json.Marshal(struct) bypasses the RPC logic.
	processed := camelCaseKeys(val)
	b, err := json.Marshal(processed)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Default JSON: %s", string(b))

	// Check standard behavior
	var res map[string]interface{}
	json.Unmarshal(b, &res)

	if _, ok := res["MyField"]; ok {
		t.Errorf("Did NOT expect MyField (PascalCase), got %v", res)
	}
	if _, ok := res["myField"]; !ok {
		t.Errorf("Expected myField (camelCase), got %v", res)
	}

	// Check tagged field (should remain "tagged")
	if _, ok := res["tagged"]; !ok {
		t.Errorf("Expected 'tagged' (json tag), got %v", res)
	}

	// Check nested untagged field
	nested, ok := res["nested"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected nested is map, got %T", res["nested"])
	}
	if _, ok := nested["innerField"]; !ok {
		t.Errorf("Expected nested.innerField (camelCase), got %v", nested)
	}
}
