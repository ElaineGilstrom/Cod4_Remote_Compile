package share

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

func ConvIntToBytes(n int, numBytes int) (b []byte, err error) {
    var buf bytes.Buffer
    err = binary.Write(&buf, binary.LittleEndian, n)
    if err != nil {
        return nil, err
    } else {
        b = buf.Bytes()
        if len(b) > numBytes {
            return nil, fmt.Errorf("%d is too large. Cannot be converted to byte array of len < %d.", len(b), numBytes)
        } else if len(b) < numBytes {
            for len(b) < numBytes {
                b = append(b, 0)
            }
        }
        return b, nil
    }
}

func DecodeIntFromBytes(byteArr []byte) (n int) {
    //n64, r := binary.Varint(byteArr)
    switch {
    case len(byteArr) == 1:
        n |= (int)(byteArr[0])
        return n
    case len(byteArr) > 4:
        //TODO: Handle too big int
        panic(fmt.Sprintf("%#v is too big!", byteArr))
    default:
        var m int
        if byteArr[len(byteArr) - 1] & 0x80 != 0 {
            m = -1
            byteArr[len(byteArr) - 1] = byteArr[len(byteArr) - 1] & 0x7F
        } else {
            m = 1
        }
        
        for _, i := range byteArr {
            n <<= 8
            n |= (int)(i)
        }
        
        n *= m
        return n
    }
}
