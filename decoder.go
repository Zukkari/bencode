package bencode

import (
    "io"
    "strconv"
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

    return d.decodeFrom(d.buff[0])
}

func (d Decoder) decodeFrom(b byte) (interface{}, error) {
    switch b {
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
    buff := make([]interface{}, 0)

    processor := func(n int) (bool, error) {
        for i := 0; i < n; i++ {
            if read := d.buff[i]; read != 'e' {
                decoded, err := d.decodeFrom(read)

                if err != nil {
                    return false, err
                }

                buff = append(buff, decoded)
            } else {
                return false, nil
            }
        }

        return true, nil
    }

    for {
        n, err := d.r.Read(d.buff)

        cont, err := processor(n)
        if !cont {
            break
        }

        if err == io.EOF {
            break
        } else if err != nil {
            return nil, err
        }
    }

    return buff, nil
}

func (d Decoder) decodeInt() (int, error) {
    intBuff := make([]byte, 64)
    offset := 0

    processor := func(n int) bool {
        for i := 0; i < n; i++ {
            if read := d.buff[i]; read != 'e' {
                intBuff[offset] = read
                offset++
            } else {
                return false
            }
        }

        return true
    }

    for {
        n, err := d.r.Read(d.buff)

        cont := processor(n)
        if !cont {
            break
        }

        if err == io.EOF {
            break
        } else if err != nil {
            return 0, err
        }
    }

    return strconv.Atoi(string(intBuff[:offset]))
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
