package bencode

import "fmt"

type Encoder struct {
	value interface{}
}

type UnsupportedType string

func (e UnsupportedType) Error() string {
	return fmt.Sprintf("Unsupported type: %v", e)
}

func (e *Encoder) encode() (string, error) {
	switch e.value.(type) {
	case int, int8, int16, int32, int64:
		return e.encodeInt(), nil
	case []interface{}:
		return e.encodeList()
	case string:
		return e.encodeString(), nil
	case map[string]interface{}:
		return e.encodeDict()
	}

	return "", UnsupportedType(fmt.Sprintf("%T", e.value))
}

func (e *Encoder) encodeInt() string {
	return fmt.Sprintf("i%de", e.value)
}

func (e *Encoder) encodeList() (string, error) {
	buff := make([]byte, 1)
	buff[0] = 'l'

	for _, v := range e.value.([]interface{}) {
		encoded, err := Encode(v)
		if err != nil {
			return "", err
		}

		buff = append(buff, []byte(encoded)...)
	}

	buff = append(buff, 'e')
	return string(buff), nil
}

func (e *Encoder) encodeString() string {
	return fmt.Sprintf("%v:%v", len(e.value.(string)), e.value)
}

func (e *Encoder) encodeDict() (string, error) {
	return "nil", nil
}

func Encode(i interface{}) (string, error) {
	encoder := &Encoder{value: i}
	return encoder.encode()
}
