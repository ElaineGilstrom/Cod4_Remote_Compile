package share

import (
	"net"
	"log"
)

func WriteMsg(msg ... []byte) bool {
    for _, d := range msg {
        if !Write(conn, d) {
            return false
        }
    }
    return true
}

//Writes bytes b to connection con. Returns if send was successfull
func Write(conn net.Conn, b []byte) bool {
    n, err := conn.Write(b)
    if err != nil {
        log.Printf("Error: Unable to write error msg to %s\n", conn.RemoteAddr().String())
        return false
    } else if n != len(b) {
        log.Printf("Error: Client %s did not recieve full msg. Recieved %d, Expected %d\n", conn.RemoteAddr().String(), n, len(msg))
        return false
    } else {
        return true
    }
}
