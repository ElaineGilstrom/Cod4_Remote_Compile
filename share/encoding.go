package share

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

func ConvIntToByteArr(n int, numBytes int) (b []byte, err error) {
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
