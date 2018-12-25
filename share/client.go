package shared

import (
    //"strings"
    "errors"
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
    vals := map[string]ClientCalls {
        "error":-1,
        "pong":0,
        "manifest":1,
        "fileTransfer":2,
        "printResult":3}
    
    if _, exists := vals[v]; exists {
        return vals[v], nil
    } else {
        return -1, errors.New(fmt.Sprintf("Unknown Client Call %s", v))
    }
}

func (s ClientCalls) String() string {
    calls := [...]string {
        "pong",
        "manifest",
        "fileTransfer",
        "printResult"}
    
    if s < pong || s > printResult {
        if s == -1 {
            return "error"
        } else {
            return "Unknown"
        }
    } else {
        return calls[s]
    }
}
