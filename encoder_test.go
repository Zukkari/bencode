package bencode

import "testing"

func TestEncodeInt(t *testing.T) {
	encoded, _ := Encode(32)
	expected := "i32e"

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}

func TestEncodeList(t *testing.T) {
	input := make([]int, 3)
	input[0] = 1
	input[1] = 2
	input[2] = 3

	encoded, _ := Encode(input)
	expected := "li1ei2ei3ee"

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}

func TestEncodeString(t *testing.T) {
	input := "test"
	expected := "4:test"

	encoded, _ := Encode(input)

	if encoded != expected {
		t.Errorf("Expected encoded value to be %v but got %v", expected, encoded)
	}
}
