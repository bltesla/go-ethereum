package rpc

import (
	"encoding/json"
	"testing"
)

type CustomByteArray [4]byte

func (c *CustomByteArray) MarshalJSON() ([]byte, error) {
	return json.Marshal("CUSTOM-ARRAY")
}

type CustomByteSlice []byte

func (c *CustomByteSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal("CUSTOM-SLICE")
}

type Container2 struct {
	Array CustomByteArray
	Slice CustomByteSlice
}

func TestMarshalJSONNonStruct(t *testing.T) {
	c := Container2{
		Array: CustomByteArray{1, 2, 3, 4},
		Slice: CustomByteSlice{5, 6, 7, 8},
	}

	processed := camelCaseKeys(c)
	b, err := json.Marshal(processed)
	if err != nil {
		t.Fatal(err)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(b, &res); err != nil {
		t.Fatal(err)
	}

	if res["array"] != "CUSTOM-ARRAY" {
		t.Errorf("Expected CUSTOM-ARRAY, got %v", res["array"])
	}

	if res["slice"] != "CUSTOM-SLICE" {
		t.Errorf("Expected CUSTOM-SLICE, got %v", res["slice"])
	}
}
