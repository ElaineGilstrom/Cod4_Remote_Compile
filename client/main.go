package main

import "Cod4_Remote_Compile/share"

import (
    "net"
    "fmt"
    "os"
    "bufio"
    //"crypto/sha256"
    //"strconv"
    "strings"
    "time"
    "regexp"
)


var connStr, login string
var callType share.ServerCalls

func main() {
    if len(os.Args) <= 1 {
        if err := handleConfig("config"); err != 0 {
            return
        }
    } else {
        for i := 1; i < (len(os.Args) - 1); i += 2 {
            if err := initHandler(os.Args[i], os.Args[i + 1], "Command line"); err != 0 {
                return
            }
        }
    }
    
    if connStr == "" {
        fmt.Println("No connection was inputed!")
        return
    }
    
    conn, err := net.Dial("tcp", connStr)
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    
    //TODO: Handle connection
    switch callType {
        case share.ServerPing:
            start := time.Now()
            n, err := conn.Write([]byte{0})
            if err != nil {
                fmt.Println(err)
                return
            } else if n == 0 {
                fmt.Println("ERROR: Server read no bytes!\n")
                return
            }
            b := make([]byte, 1)
            n,err = conn.Read(b)
            if err != nil {
                fmt.Println(err)
                return
            } else if n == 0 {
                fmt.Println("ERROR: Server sent no bytes!\n")
                return
            } else if b[0] != 0 {
                fmt.Println("ERROR: Server did not pong! (%d != 0)\n", b[0])
                return
            }
            fmt.Printf("Pong! %v\n", time.Now().Sub(start))
            return
        default:
            fmt.Println("ERROR: %s not implemented!\n", callType.String())
    }
}

func initHandler(arg string, val string, source string) int {
    arg = strings.ToLower(arg)
    var err error//Initializing here for simplicities sake
    
    switch arg {
        case "-c", "--config":
            return handleConfig(val)
        case "-s", "--server":
            m, err := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{1,5}$", val)
            if err != nil {
                panic(err)
            } else if !m {
                fmt.Println("ERROR: Server ip:port [%s] did not match format 255.255.255.255:99999! (From %s)\n", val, source)
                return 2
            } else {
                connStr = val
            }
            return 0
        case "-l", "--login":
            login = val
            return 0
        case "-t", "--type":
            callType, err = share.NewServerCalls(val)
            if err != nil {
                fmt.Println("ERROR: Error from %s: ", source)
                fmt.Println(err)
                return 2
            } else {
                return 0
            }
        default:
            fmt.Println("ERROR: Unknown arg %s. (From %s)\n", val, source)
            return 2
    }
}

func handleConfig(fileName string) int {
    f, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
        return 2
    }
    
    reader := bufio.NewReader(f)
    for l, err := reader.ReadString('\n'); err != nil; l, err = reader.ReadString('\n') {// If there is a better way to do this, please change it or tell me.
        iErr := 0
        if l == "" {
            continue
        } else {
            s := strings.Split(l, "=")
            if len(s) != 2 {
                fmt.Println("ERROR: Config %s not formatted properly!\nA line must include 1 and only 1 equals sign.!\nGot: %s\n", fileName, l)
                return 2
            } else {
                iErr = initHandler("-" + s[0], s[1], "Config " + fileName)
            }
        }
        return iErr
    }
    
    return 0
}