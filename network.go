package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

var RouteTable []string

func init() {
	ips, err := net.LookupIP(SEED_DOMAIN)
	if err != nil {
		log.Fatalln(err)
		return
	}
	port := strconv.Itoa(DEFAULT_PORT)
	for _, ip := range ips {
		RouteTable = append(RouteTable, ip.String()+port)
	}
}

func server() {
	port := strconv.Itoa(DEFAULT_PORT)
	ln, err := net.Listen("tcp", ":"+port)
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
	conn := dial(RouteTable[0])
	if conn == nil {
		return
	}
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Println("dial error", err)
	// }
	// fmt.Println(status)
}

func dial(address string) net.Conn {
	conn, err := net.Dial("tcp", ":25000")
	if err != nil {
		fmt.Println("dial err:", err)
		return nil
	}
	return conn
}
