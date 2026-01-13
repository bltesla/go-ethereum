package rpc

import (
	"encoding/json"
	"testing"
)

type CustomValue struct {
	Val string
}

func (c CustomValue) MarshalJSON() ([]byte, error) {
	return json.Marshal("CUSTOM-VALUE")
}

type CustomPointer struct {
	Val string
}

func (c *CustomPointer) MarshalJSON() ([]byte, error) {
	return json.Marshal("CUSTOM-POINTER")
}

type Container struct {
	Value   CustomValue
	Pointer CustomPointer
}

func TestMarshalJSONRespect(t *testing.T) {
	c := Container{
		Value:   CustomValue{Val: "test"},
		Pointer: CustomPointer{Val: "test"},
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

	if res["value"] != "CUSTOM-VALUE" {
		t.Errorf("Expected CUSTOM-VALUE, got %v", res["value"])
	}

	// This is likely where it fails if it's a pointer receiver and we pass a value
	if res["pointer"] != "CUSTOM-POINTER" {
		t.Errorf("Expected CUSTOM-POINTER, got %v", res["pointer"])
	}
}
