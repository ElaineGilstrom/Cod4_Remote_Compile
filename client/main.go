package main

imort (
	"net"
	"fmt"
	"os"
	//"crypto/sha256"
	"strconv"
)

var port int = 1337

func main() {
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
			case "-c":
				i++
				if i >= len(os.Args) {
					fmt.Errorf("No config given!\n")
					return 2
				}
			default:
				fmt.Errorf("Unknown arg %s.\n", os.Args[i])
				return 2
		}
	}
	
	conn, err := net.Dial("tcp", connStr)
	if err != nil {
		panic(err)
	}
	
	//TODO: Handle connection
}
