package bencode

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Decoder struct {
	r io.Reader

	buff []byte
}

type NoBytesReadError struct {
}

func (err NoBytesReadError) Error() string {
	return "could not read any bytes"
}

func (d Decoder) decode() (interface{}, error) {
	n, err := d.r.Read(d.buff)
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, NoBytesReadError{}
	}

	switch b := d.buff[0]; b {
	case 'i':
		return d.decodeInt()
	case 'l':
		return d.decodeList()
	case 'd':
		return d.decodeDict()
	default:
		return d.decodeString(b)
	}
}

func (d Decoder) decodeList() (interface{}, error) {
	return nil, nil
}

func (d Decoder) decodeInt() (int64, error) {
	intBuff := make([]byte, 64)

	offset := 0
	for {
		n, err := d.r.Read(d.buff)

		for i := 0; i < n; i++ {
			intBuff[offset] = d.buff[i]
			offset++
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
	}

	byteReader := bytes.NewBuffer(intBuff)
	return binary.ReadVarint(byteReader)
}

func (d Decoder) decodeDict() (interface{}, error) {
	return nil, nil
}

func (d Decoder) decodeString(initial byte) (string, error) {
	return "nil", nil
}

func Decode(reader io.Reader) (interface{}, error) {
	decoder := Decoder{
		reader,
		make([]byte, 1),
	}

	return decoder.decode()
}
