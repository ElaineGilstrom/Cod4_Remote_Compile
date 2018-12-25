package main

imort (
    "net"
    "fmt"
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
}