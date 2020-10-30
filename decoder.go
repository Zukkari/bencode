package bencode

import (
    "io"
    "strconv"
)

type Decoder struct {
    r    io.Reader
    buff []byte
}

func (d *Decoder) next() (byte, error) {
    _, err := d.r.Read(d.buff)

    if err != nil {
        return 0, err
    }

    return d.buff[0], nil
}

func (d *Decoder) decode() (interface{}, error) {
    next, err := d.next()
    if err != nil {
        return nil, err
    }

    return d.decodeByte(next)
}

func (d *Decoder) decodeByte(b byte) (interface{}, error) {
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

func (d *Decoder) decodeList() ([]interface{}, error) {
    buff := make([]interface{}, 0)

    for {
        next, err := d.next()

        if next != 'e' {
            decoded, err := d.decodeByte(next)

            if err != nil {
                return nil, err
            }

            buff = append(buff, decoded)
        } else {
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

func (d *Decoder) decodeInt() (int, error) {
    intBuff := make([]byte, 64)
    offset := 0

    for {
        next, err := d.next()

        if next != 'e' {
            intBuff[offset] = next
            offset++
        } else {
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

func (d *Decoder) decodeString(b byte) (string, error) {
    number := make([]byte, 1)
    number[0] = b

    for {
        next, err := d.next()

        if next == ':' {
            break
        } else {
            number = append(number, next)
        }

        if err == io.EOF {
            break
        } else if err != nil {
            return "", err
        }
    }

    size, err := strconv.Atoi(string(number))
    if err != nil {
        return "", err
    }

    buff := make([]byte, size)
    for i := range buff {
        next, err := d.next()

        buff[i] = next

        if err != nil {
            return "", err
        }
    }

    return string(buff), nil
}

func (d *Decoder) decodeDict() (interface{}, error) {
    return nil, nil
}

func Decode(reader io.Reader) (interface{}, error) {
    decoder := &Decoder{
        reader,
        make([]byte, 1),
    }

    return decoder.decode()
}
