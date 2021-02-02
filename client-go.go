/*****************************************************************************
 * client-go.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const SEND_BUFFER_SIZE = 2048

/* TODO: client()
 * Open socket and send message from stdin.
 */
func client(server_ip string, server_port string) {
	dest_addr := net.ParseIP(server_ip).String() + ":" + server_port

	tcp_addr, err := net.ResolveTCPAddr("tcp", dest_addr)
	checkErr(err)

	conn, err := net.DialTCP("tcp4", nil, tcp_addr)
	checkErr(err)

	reader := bufio.NewReader(os.Stdin)
	result, err := ioutil.ReadAll(reader)
	checkErr(err)

	for {
		lenResult := len(result)
		i := 0
		n := 0
		for {
			if lenResult < SEND_BUFFER_SIZE {
				n, err = conn.Write(result[i : i+lenResult])
			} else {
				n, err = conn.Write(result[i : i+SEND_BUFFER_SIZE])
			}
			checkErr(err)

			fmt.Println(n)

			i = i + n
			lenResult = lenResult - n
			if lenResult <= 0 {
				return
			}
		}
	}
}

// Main parses command-line arguments and calls client function
func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ./client-go [server IP] [server port] < [message file]")
	}
	server_ip := os.Args[1]
	server_port := os.Args[2]
	client(server_ip, server_port)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
