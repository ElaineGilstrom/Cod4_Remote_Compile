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


var connStr, login, manifest, args string
var callType share.ServerCalls

func main() {
    if len(os.Args) <= 1 {
        if err := handleConfig("config"); err != 0 {
            return
        }
    } else if len(os.Args) % 2 == 0 {
        //Display help for now. Non param taking flags may come in the future
        fmt.Printf("%s [-c -s -l -t -m -a]\n")
        //write param defs, mention how they must either be in config or in cmd line
        fmt.Printf("\t%30s Path to a config file. %s --help config for more\n", "-c --config", os.Args[0])
        fmt.Printf("\t%30s Ip and port (255.255.255.255:99999) of server to connect to.\n", "-s --server server")
        fmt.Printf("\t%30s Login/Identity string. Maps will be associated with this string.\n", "-l --login login")
        fmt.Printf("\t%30s The comunication type or objective of this connection. %s --help types for more.\n", "-t --type type", os.Args[0])
        fmt.Printf("\t%30s Manifest of files to upload. %s --help manifest for more.\n", "-m --manifest manifest", os.Args[0])
        fmt.Printf("\t%30s The arguments to pass to the compiler.\n", "-a --args args")
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
        case share.ServerPing://TODO: Add timeout
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
        case "-s", "--server", "server":
            m, err := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{1,5}$", val)
            if err != nil {
                panic(err)
            } else if !m {
                fmt.Println("ERROR: Server ip:port [%s] did not match format 255.255.255.255:99999! (From %s)\n", val, source)
                return 2
            } else {
                connStr = val
                return 0
            }
        case "-l", "--login", "login":
            //TODO: Implement hash. Will either be 64 or 128 bit.
            login = val
            return 0
        case "-t", "--type", "type":
            callType, err = share.NewServerCalls(val)
            if err != nil {
                fmt.Println("ERROR: Error from %s: ", source)
                fmt.Println(err)
                return 2
            } else {
                return 0
            }
        case "-m", "--manifest", "manifest":
            //TODO: Check if manifest name is formatted correctly.
            manifest = val
        case "-a", "--args", "args":
            //TODO: Check if args are formatted correctly.
            args = val
        case "--help":
            printHelpPage(val)
            return 1
        default:
            fmt.Println("ERROR: Unknown arg %s. (From %s)\n", val, source)
            return 2
    }
}

func printHelpPage(page string) {
    switch page {
    case "config":
        fmt.Println("The config can contain any parameter that doesn't start with a dash (-).")
        fmt.Println("It follows the syntax of: <parameter>:<value>")
        fmt.Println("Example: type=Ping")
    case "types":
        fmt.Printf("%16s Checks to see if the server is alive, and checks latency. If server doesn't respond in 2 minutes it times out.\n", "Ping")
        fmt.Printf("%16s Takes the path to a file that contains a list of all files to be included in the map. Then uploads those files.%s --help manifest for more.\n", "Manifest", os.Args[0])
        fmt.Printf("%16s Returns a copy of a map from the backup. %s --help backup for more.\n", "ReturnBackup", os.Args[0])
        fmt.Printf("%16s Takes 'maps' or map name as argument. If maps is given, it lists all maps stored in the current login. If map name is specified, it lists all versions availible for that map.\n", "Ls")
        fmt.Printf("%16s Takes compile mode as argument. Returns compiled files. %s --help compile for more.\n", "Compile", os.Args[0])
        fmt.Printf("%16s Login/Identity string. This sets what maps you have access to. If you want to cooperate with someone else, share this string.\n", "Login")
    case "manifest":
        fmt.Println("This file will include a list of files with their paths relative to the cod4 root direcory. Symlinks not supported!")
        fmt.Println("The manifest's name must follow the syntax: <map name>.manifest. Example: mp_wave.manifest.")
        fmt.Println("The manifest is formatted as a csv. The first column is how this file should be updated. Can be add, update or remove. The second column is the path relative to the cod4 root folder, then third is the file name.")
        fmt.Println("Example line: update,raw/maps/mp/gametypes/,mp_wave.gsc")
    case "backup":
        fmt.Println("When ReturnBackup is run, if a manifest is specified, then all files tagged with add or update will be returned, else, all files for that map will be returned.")
        fmt.Println("The arg should be formatted as <map name>[:date]")
        fmt.Println("Example: mp_wave")
        fmt.Println("Given no manifest was specified, this will return all of the files for the map mp_wave that were in use during the most recent upload.")
        fmt.Println("Example2: mp_wave:20181229_0039")
        fmt.Println("Given no manifest was specified, this will return all files used by the version made on December 29, 2018 at 12:39 AM.")
        fmt.Println("For both these examples, if a manifest was specified, it will return all the files in use and in the manifest tagged with add or update at the time of the specified map version.")
    case "compile":
        fmt.Println("This is still far from being implemented, come back later when I know what this is going to look like.")
    default:
        fmt.Printf("Unknown help page %s!\n", page)
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
                iErr = initHandler(s[0], s[1], "Config " + fileName)
            }
        }
        return iErr
    }
    
    return 0
}