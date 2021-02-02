/*****************************************************************************
 * server-go.go
 * Name:
 * NetId:
 *****************************************************************************/

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const RECV_BUFFER_SIZE = 2048

/* TODO: server()
 * Open socket and wait for client to connect
 * Print received message to stdout
 */

var sem = make(chan int, 5)

func server(server_port string) {
	service := ":" + server_port
	tcp_addr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)

	ln, err := net.ListenTCP("tcp", tcp_addr)
	checkErr(err)

	i := 0

	for {
		i = i + 1
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		sem <- 1
		handleClient(conn)
		conn.Close()
		<-sem
	}
}

// Main parses command-line arguments and calls server function
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./server-go [server port]")
	}
	server_port := os.Args[1]
	server(server_port)
}

func handleClient(conn net.Conn) {
	emptStr := ""

	var buf [RECV_BUFFER_SIZE]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s", emptStr)
			return
		}
		emptStr = emptStr + string(buf[0:n])
	}
	fmt.Fprintf(os.Stdout, "%s", emptStr)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
