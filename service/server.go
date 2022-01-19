package service

import (
	"encoding/json"
	"log"
	"net"

	"github.com/smallyunet/tmb/storage"
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
	if string(buf) != "" {
		log.Println(string(buf))
		// handle payload data
		var d []storage.KeyValue
		err := json.Unmarshal(buf, &d)
		if err != nil {
			conn.Write([]byte("Error json data format."))
			return
		}
		// TODO should not handle tx rather than block data use consenses
		//block.DataCache = append(block.DataCache, d...)
		//block.DataMsg <- len(d)
		conn.Write([]byte("Handle data success."))
	}
}
