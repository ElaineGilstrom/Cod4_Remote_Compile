package main

imort (
	"net"
	"fmt"
	"os"
	"bufio"
	//"crypto/sha256"
	"strconv"
	"strings"
)


var conn, login string

func main() {
	for i := 1; i < (len(os.Args) + 1); i += 2 {
		if err := initHandler(os.Args[i], os.Args[i + 1], "Command line"); err != 0 {
			return err
		}
	}
	
	conn, err := net.Dial("tcp", connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	//TODO: Handle connection
}

func initHandler(arg string, val string, source string) {
	switch arg {
		case "-c", "--config":
			f, err := os.Open(val)
			if err != nil {
				fmt.Println(err)
				return 2
			}
			
			reader := bufio.NewReader(f)
			for l, err := reader.ReadString('\n'); err != nil; l, err := reader.ReadString('\n') {// If there is a better way to do this, please change it or tell me.
				iErr := 0
				if l == "" {
					continue
				} else {
					s := strings.Split(l, "=")
					if len(s) != 2 {
						fmt.Errorf("Config not formatted properly!\nA line must include 1 and only 1 equals sign.!\nGot: %s\n", l)
						return 2
					} else {
						iErr = initHandler(s[0], s[1], "Config " + val)
					}
				}
				return iErr
			}
		case "-s", "--server":
			m, err := regexp.MatchString("[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{1,5}", val)
			if err != nil {
				panic(err)
			} else if !m {
				fmt.Errorf("Server ip:port [%s] did not match format 255.255.255.255:99999! (From %s)\n", val, source)
				return 2
			} else {
				conn = val
			}
			return 0
		case "-l", "--login":
			login = val
			return 0
		default:
			fmt.Errorf("Unknown arg %s. (From %s)\n", val, source)
			return 2
	}
}