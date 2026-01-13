package rpc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEmbeddedFlattening(t *testing.T) {
	val := TestStruct{
		TestBase: TestBase{Height: 100, Type: "exampleType"},
		MyField:  "val1",
	}

	processed := camelCaseKeys(val)
	b, err := json.Marshal(processed)
	if err != nil {
		t.Fatal(err)
	}

	var res map[string]interface{}
	json.Unmarshal(b, &res)

	fmt.Printf("res: %v", res)

	// Should be flattened
	if h, ok := res["height"].(float64); !ok || h != 100 {
		t.Errorf("Expected height: 100, got %v", res["height"])
	}
	if ty, ok := res["type"].(string); !ok || ty != "exampleType" {
		t.Errorf("Expected type: exampleType, got %v", res["type"])
	}

	// Should NOT be nested
	if _, ok := res["testBase"]; ok {
		t.Errorf("Did NOT expect nested testBase, got %v", res)
	}
}
