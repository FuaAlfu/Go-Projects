package main

import(
	"fmt"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"time"
	"bufio"
	"net"
)

//handling TCP connections
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}

			temp := strings.TrimSpace(string(netData))
			if temp == "STOP" {
					break
			}

			result := strconv.Itoa(random()) + "\n"
			c.Write([]byte(string(result)))
	}
	c.Close()
}

func run(){
	arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide a port number!")
                return
        }

        PORT := ":" + arguments[1]
        l, err := net.Listen("tcp4", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()
        rand.Seed(time.Now().Unix())

        for {
                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }
                go handleConnection(c)
        }
}
func main() {
	run()
}