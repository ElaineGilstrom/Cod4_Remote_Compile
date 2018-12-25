package cod4_remote_compile_shared

import (
    "strings"
    "errors"
    "fmt"
)

//These are the different calls that can be made to the server
type ServerCalls int

const (
    ping            ServerCalls = 0
    manifest        ServerCalls = 1
    fileTransfer    ServerCalls = 2
    returnBackup    ServerCalls = 3
    ls              ServerCalls = 4
    compile         ServerCalls = 5
    login           ServerCalls = 6
)


func NewServerCalls(v string) (val ServerCalls, err error) {
    vals := map[string]ServerCalls {
        "ping":0,
        "manifest":1,
        "fileTransfer":2,
        "returnBackup":3,
        "ls":4,
        "compile":5,
        "login":6
    }
    
    if _, exists := vals[v]; exists {
        return vals[v], nil
    } else {
        return -1, errors.New(fmt.Sprintf("Unknown Server Call %s", v))
    }
}

func (s ServerCalls) String() string {
    calls := [...]string {
        "ping",
        "manifest",
        "fileTransfer",
        "returnBackup",
        "ls",
        "compile",
        "login"
    }
    
    if s < ping || s > login {
        return "Unknown"
    } else {
        return calls[s]
    }
}
