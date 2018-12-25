package shared

import (
    //"strings"
    "errors"
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
    vals := map[string]ServerCalls {
        "ServerError":-1,
        "ServerPing":0,
        "ServerManifest":1,
        "ServerFileTransfer":2,
        "ServerReturnBackup":3,
        "ServerLs":4,
        "ServerCompile":5,
        "ServerLogin":6}
    
    if _, exists := vals[v]; exists {
        return vals[v], nil
    } else {
        return -1, errors.New(fmt.Sprintf("Unknown Server Call %s", v))
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
