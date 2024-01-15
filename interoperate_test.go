package json

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMarshalJSONNumber(t *testing.T) {
	var a any
	decoder := json.NewDecoder(bytes.NewReader([]byte("123")))
	decoder.UseNumber()
	err := decoder.Decode(&a)
	if err != nil {
		t.Fatal(err)
	}

	marshal, err := Marshal(a)
	if err != nil {
		t.Fatal(err)
	}

	if string(marshal) != "123" {
		t.Fatalf("%s != 123", string(marshal))
	}
}

func TestMarshalLocalNumber(t *testing.T) {
	var a any
	decoder := NewDecoder(bytes.NewReader([]byte("123")))
	decoder.UseNumber()
	err := decoder.Decode(&a)
	if err != nil {
		t.Fatal(err)
	}

	marshal, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}

	if string(marshal) != "123" {
		t.Fatalf("%s != 123", string(marshal))
	}
}
