package main

imort (
	"net"
	"fmt"
	"os"
	"bufio"
	//"crypto/sha256"
	"strconv"
)



func main() {
	var conn string
	var login string
	
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
			case "-c", "--config":
				i++
				if i >= len(os.Args) {
					fmt.Errorf("No config given!\n")
					return 2
				}
			case "-s", "--server":
				i++
				if i >= len(os.Args) {
					fmt.Errorf("No server given!\n")
					return 2
				}
				
				m, err := regexp.MatchString("[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}:[0-9]{1,5}", os.Args[i])
				if err != nil {
					panic(err)
				} else if !m {
					fmt.Errorf("Server ip:port [%s] did not match format 255.255.255.255:99999!\n", os.Args[i])
					return 2
				} else {
					conn = os.Args[i]
				}
			case "-l", "--login":
				i++
				if i >= len(os.Args) {
					fmt.Errorf("No login given!\n")
					return 2
				}
				
				login = os.Args[i]
			default:
				fmt.Errorf("Unknown arg %s.\n", os.Args[i])
				return 2
		}
	}
	
	conn, err := net.Dial("tcp", connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	
	//TODO: Handle connection
}
