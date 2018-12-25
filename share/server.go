package share

import (
    "strings"
    //"errors"
    "fmt"
)

//These are the different calls that can be made to the server
type ServerCalls int

const (
    ServerError           ServerCalls = -1
    ServerPing            ServerCalls = 0
    ServerManifest        ServerCalls = 1
    ServerFileTransfer    ServerCalls = 2
    ServerReturnBackup    ServerCalls = 3
    ServerLs              ServerCalls = 4
    ServerCompile         ServerCalls = 5
    ServerLogin           ServerCalls = 6
)


func NewServerCalls(v string) (val ServerCalls, err error) {
    v = strings.ToLower(v)
    
    vals := map[string]ServerCalls {
        "servererror":-1,
        "serverping":0,
        "servermanifest":1,
        "serverfiletransfer":2,
        "serverreturnbackup":3,
        "serverls":4,
        "servercompile":5,
        "serverlogin":6}
    
    if _, exists := vals[v]; exists {
        return vals[v], nil
    } else {
        v2 := "server" + v
        if _, exists := vals[v2]; exists {
            return vals[v2], nil
        } else {
            return -1, fmt.Errorf("Unknown Server Call %s", v)
        }
    }
}

func (s ServerCalls) String() string {
    calls := [...]string {
        "ServerPing",
        "ServerManifest",
        "ServerFileTransfer",
        "ServerReturnBackup",
        "ServerLs",
        "ServerCompile",
        "ServerLogin"}
    
    if s < ServerPing || s > ServerLogin {
        if s == -1 {
            return "ServerError"
        } else {
            return "Unknown"
        }
    } else {
        return calls[s]
    }
}

func ByteToServerCalls(b byte) ServerCalls {
    val := (ServerCalls)(b)
    if (val <= ServerLogin && val >= ServerError) {
        return val
    } else {
        return -2
    }
}
