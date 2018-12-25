package main

import "Cod4_Remote_Compile/share"

import (
    "net"
    "fmt"
    "log"
    "os"
    //"crypto/sha256"
    "strconv"
)

var port int64 = 1337

func main() {
    for i := 1; i < len(os.Args); i++ {
        switch os.Args[i] {
            case "--port":
                i++
                if i >= len(os.Args) {
                    fmt.Println("ERROR: No port given!\n")
                    return
                }
                var err error
                port, err = strconv.ParseInt(os.Args[i], 10, 64)
                if err != nil {
                    fmt.Println("ERROR: Port bad format [%s]. Must be numeric only!\nERROR: ", os.Args[i]);
                    fmt.Println(err)
                    return
                } else if port <= 0 {
                    fmt.Println("ERROR: Provided port %d must be greater than 0.\n", port)
                    return
                }
            default:
                fmt.Println("ERROR: Unknown arg %s.\n", os.Args[i])
                return
        }
    }
    
    listen, err := net.Listen("tcp", ":" + strconv.FormatInt(port, 10))
    if err != nil {
        panic(err)
    }
    
    for {
        conn, err := listen.Accept()
        if err != nil {
            panic(err)
        }
        
        go connHandler(conn)
    }
}

func connHandler(conn net.Conn) {
    defer conn.Close()
    
    
    b := make([]byte, 1)
    n, err := conn.Read(b)
    if err != nil {
        fmt.Println(err)
        return
    } else if n == 0 {
        log.Printf("ERROR: Client sent no bytes!\n")
        return
    }
    
    switch share.ByteToServerCalls(b[0]) {
        case share.ServerPing:
            n, err = conn.Write(b)
            if err != nil {
                log.Println(err)
                return
            } else if n == 0 {
                log.Println("ERROR: Client sent no bytes!\n")
                return
            }
        default:
            log.Printf("Recieved unknow response (%d) from client %s.\n", conn.RemoteAddr().String(), b[0])
            errStr := []byte(fmt.Sprintf("Unknown call %d", b[0]))
            errLen, err := share.ConvIntToByteArr(len(errStr), 4)
            if err != nil {
                fmt.Printf("Error: Unable to format int %d.\n", len(errStr))
                fmt.Println(err)
                return
            }
            _ = share.WriteMsg(conn, []byte{0xFF}, errLen, errStr)
            return
    }
}
