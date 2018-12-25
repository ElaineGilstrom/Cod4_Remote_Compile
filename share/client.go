package share

import (
    "strings"
    //"errors"
    "fmt"
)

//These are the different calls that can be made to the client
type ClientCalls int

const (
    ClientError            ClientCalls = -1
    ClientPong             ClientCalls = 0
    ClientManifest         ClientCalls = 1
    ClientFileTransfer     ClientCalls = 2
    ClientPrintResult      ClientCalls = 3
)

func NewClientCalls(v string) (val ClientCalls, err error) {
    v = strings.ToLower(v)
    
    vals := map[string]ClientCalls {
        "clienterror":-1,
        "clientpong":0,
        "clientmanifest":1,
        "clientfiletransfer":2,
        "clientprintresult":3}
    
    if _, exists := vals[v]; exists {
        return vals[v], nil
    } else {
        v2 := "client" + v
        if _, exists := vals[v2]; exists {
            return vals[v2], nil
        } else {
            return -1, fmt.Errorf("Unknown Client Call %s", v)
        }
    }
}

func (s ClientCalls) String() string {
    calls := [...]string {
        "ClientPong",
        "ClientManifest",
        "ClientFileTransfer",
        "ClientPrintResult"}
    
    if s < ClientPong || s > ClientPrintResult {
        if s == -1 {
            return "ClientError"
        } else {
            return "Unknown"
        }
    } else {
        return calls[s]
    }
}

func ByteToClientCalls(b byte) ClientCalls {
    val := (ClientCalls)(b)
    if (val <= ClientPrintResult && val >= ClientError) {
        return val
    } else {
        return -2
    }
}
