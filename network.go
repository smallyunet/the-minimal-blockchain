package main

import (
	"bufio"
	"fmt"
	"net"
)

func init() {

}

func server() {
	ln, err := net.Listen("tcp", ":25000")
	if err != nil {
		fmt.Println("listen error", err)
	}
	fmt.Println("node is running...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accetp error", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("conn: %v\n", conn)
}

func client() {
	conn, err := net.Dial("tcp", ":25000")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("dial error", err)
	}
	fmt.Println(status)
}
