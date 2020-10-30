package bencode

import (
	"strings"
	"testing"
)

func TestDecodeInt(t *testing.T) {
	input := strings.NewReader("i32e")

	decoded, err := Decode(input)

	if err != nil {
		t.Errorf("Got error when decoding: %v", err)
	}

	if decoded != 32 {
		t.Errorf("Decoded input %v was incorrect, got %v and expected: %v", input, decoded, 32)
	}
}

func TestDecodeListInt(t *testing.T) {
	input := strings.NewReader("li32ei42ei57ee")

	expected := make([]int, 3)
	expected[0] = 32
	expected[1] = 42
	expected[2] = 57

	decoded, err := Decode(input)

	if err != nil {
		t.Errorf("Got error when decoding: %v", err)
	}

	for i, v := range decoded.([]interface{}) {
		if v != expected[i] {
			t.Errorf("Decoded input %v was incorrect, got %v and expected %v", input, v, expected[i])
		}
	}
}

func TestDecodeString(t *testing.T) {
	input := "4:spam"

	reader := strings.NewReader(input)
	decoded, err := Decode(reader)

	if err != nil {
		t.Errorf("Got error when decoding: %v", err)
	}

	if decoded != "spam" {
		t.Errorf("Decoded input %v was incorrect, got %v and expected %v", input, decoded, "spam")
	}
}

func TestDecodeDict(t *testing.T) {
	input := "d3:bar4:spam3:fooi42ee"

	reader := strings.NewReader(input)
	decoded, err := Decode(reader)

	if err != nil {
		t.Errorf("Got error when decoding: %v", err)
	}

	decodedMap := decoded.(map[string]interface{})

	if decodedMap["bar"] != "spam" {
		t.Errorf("Decoded input %v was incorrect, got %v and expected %v", "bar", decoded, "spam")
	}

	if decodedMap["foo"] != 42 {
		t.Errorf("Decoded input %v was incorrect, got %v and expected %v", "foo", decoded, 42)
	}
}
