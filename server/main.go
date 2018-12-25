package main

import "Cod4_Remote_Compile/share"

imort (
    "net"
    "fmt"
    "log"
    "os"
    //"crypto/sha256"
    "strconv"
)

var port int = 1337

func main() {
    for i := 1; i < len(os.Args); i++ {
        switch os.Args[i] {
            case "--port":
                i++
                if i >= len(os.Args) {
                    fmt.Errorf("No port given!\n")
                    return 2
                }
                var err error
                port, err = strconv(os.Args[i])
                if err != nil {
                    fmt.Errorf("Port bad format [%s]. Must be numeric only!\nERROR: ", os.Args[i]);
                    fmt.Println(err)
                    return 2
                } else if port <= 0 {
                    fmt.Errorf("Provided port %d must be greater than 0.\n", port)
                    return 2
                }
            default:
                fmt.Errorf("Unknown arg %s.\n", os.Args[i])
                return 2
        }
    }
    
    listen, err := net.Listen("tcp", ":" + strconv.FormatInt(port, 10))
    if err := nil {
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

func connHandler(conn *net.Conn) {
    defer conn.Close()
    
    
    b := [1]byte
    n, err := conn.Read(b)
    if err != nil {
        fmt.Println(err)
        return
    } else if n == 0 {
        log.Printf("ERROR: Client sent no bytes!\n")
        return
    }
    
    switch b[0] {
        case 0:
            n, err = conn.Write(b)
            if err != nil {
                fmt.Println(err)
                return
            } else if n == 0 {
                fmt.Errorf("Client sent no bytes!\n")
                return
            }
        default:
            msg := append([]byte[-1], []byte(fmt.Sprintf("Unknown call %d", b[0])))
            log.Printf("Recieved unknow response (%d) from client %s.\n", conn.RemoteAddr().String(), b[0])
            n, err = conn.Write(msg)
            if err != nil {
                log.Printf("Error: Unable to write error msg to client %s\n", conn.RemoteAddr().String())
            } else {
                log.Printf("Error: Client %s did not recieve full msg. Recieved %d, Expected %d\n", conn.RemoteAddr().String(), n, len(msg))
            }
            return
    }
}