package bencode

import (
    "strings"
    "testing"
)

func TestDecodeInt(t *testing.T) {
    input := strings.NewReader("i32e")

    decoded, _ := Decode(input)

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

    decoded, _ := Decode(input)

    for i, v := range decoded.([]interface{}) {
        if v != expected[i] {
            t.Errorf("Decoded input %v was incorrect, got %v and expected %v", input, v, expected[i])
        }
    }
}

func TestDecodeString(t *testing.T) {
    input := "spam"

    reader := strings.NewReader(input)
    decoded, _ := Decode(reader)

    if decoded != input {
        t.Errorf("Decoded input %v was incorrect, got %v and expected %v", input, decoded, input)
    }
}
