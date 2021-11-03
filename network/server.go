package network

import (
	"log"
	"net"
)

func Server() {
	ln, err := net.Listen("tcp", ":25000")
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
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	log.Println(string(buf[:n]))
	conn.Write([]byte("Hello, World!"))
}
