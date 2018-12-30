package share

import (
	"os"
	"io/ioutil"
	"net"
    "fmt"
    //"bytes"
)

type NetFile struct {
    local bool
    nblen []byte//Length of name as byte array
    name string
    fblen []byte//File length as byte array
    contense []byte
}

func NewNetFile(fileName string) (nf *NetFile, err error) {
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        return nil, fmt.Errorf("File %s does not exist!", fileName)
    }
    nf = new(NetFile)
    
    nf.local = true
    nf.name = fileName
    
    file, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }
    
    l, err := ConvIntToBytes(len(fileName), 1)
    if err != nil {
        return nil, err
    }
    nf.nblen = l
    
    contense, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }
    nf.contense = contense
    
    l, err = ConvIntToBytes(len(contense), 4)
    if err != nil {
        return nil, err
    }
    nf.fblen = l
    
    return nf, nil
}

func NewNetFileFromConn(conn net.Conn) (nf *NetFile, err error) {
    nf = new(NetFile)
    
    nf.local = false
    
    nf.nblen = make([]byte, 4)
    n, err := conn.Read(nf.nblen)
    if err != nil {
        return nil, err
    } else if n != 4 {
        return nil, fmt.Errorf("Recieved wrong number of bytes! Got: %d, Expected: 4.",n)
    }
    
    nl := DecodeIntFromBytes(nf.nblen)
    tmp := make([]byte, nl)
    n, err = conn.Read(tmp)
    if err != nil {
        return nil, err
    } else if n != nl {
        return nil, fmt.Errorf("Recieved wrong number of bytes! Got: %d, Expected: %s.", n, nl)
    }
    nf.name = string(tmp)
    
    nf.fblen = make([]byte, 4)
    n, err = conn.Read(nf.nblen)
    if err != nil {
        return nil, err
    } else if n != 4 {
        return nil, fmt.Errorf("Recieved wrong number of bytes! Got: %d, Expected: 4.",n)
    }
    
    fl := DecodeIntFromBytes(nf.fblen)
    nf.contense = make([]byte, fl)
    n, err = conn.Read(nf.contense)
    if err != nil {
        return nil, err
    } else if n != fl {
        return nil, fmt.Errorf("Recieved wrong number of bytes! Got: %d, Expected: %d.", n, fl)
    }
    
    return nf, nil
}

func (nf *NetFile) Send(conn net.Conn, header byte) {
    msg := [][]byte{[]byte{header}}
    msg = append(msg, nf.nblen)
    msg = append(msg, []byte(nf.name))
    msg = append(msg, nf.fblen)
    msg = append(msg, nf.contense)
    if !WriteMsg(conn, msg ...) {
        panic("Write Failed!")
    }
}

func (nf *NetFile) WriteToDir(dir string) (err error) {
    return ioutil.WriteFile(dir + nf.name, nf.contense, 0644)
}
