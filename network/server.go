package network

import (
	"log"
	"net"
)

func Server() {
	ln, err := net.Listen("tcp", ":"+localPort)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// https://stackoverflow.com/questions/24339660/read-whole-data-with-golang-net-conn-read
	buf := make([]byte, 0, tcpBufferSize)
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			log.Fatalln(err)
		}
		buf = append(buf, tmp[:n]...)
		if n < 256 {
			break
		}
	}
	defer conn.Close()
	log.Println(string(buf))
	conn.Write([]byte("Hello, World!"))
}
