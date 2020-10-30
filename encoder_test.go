package bencode

import "testing"

func TestEncodeInt(t *testing.T) {
	encoded, err := Encode(32)
	expected := "i32e"

	if err != nil {
		t.Errorf("Got error when encoding: %v", err)
	}

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}

func TestEncodeList(t *testing.T) {
	input := make([]interface{}, 3)
	input[0] = 1
	input[1] = 2
	input[2] = 3

	encoded, err := Encode(input)
	expected := "li1ei2ei3ee"

	if err != nil {
		t.Errorf("Got error when encoding: %v", err)
	}

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}

func TestEncodeString(t *testing.T) {
	input := "test"
	expected := "4:test"

	encoded, err := Encode(input)

	if err != nil {
		t.Errorf("Got error when encoding: %v", err)
	}

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}

func TestEncodeDict(t *testing.T) {
	input := make(map[string]interface{})
	input["a"] = 42
	input["b"] = "spam"

	arr := make([]interface{}, 2)
	arr[0] = 42
	arr[1] = "eggs"

	input["c"] = arr

	expected := "d1:ai42e1:b4:spam1:cli42e4:eggsee"

	encoded, err := Encode(input)

	if err != nil {
		t.Errorf("Got error when encoding: %v", err)
	}

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}
