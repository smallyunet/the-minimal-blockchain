package tcp

import (
	"encoding/json"
	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/consensus"
	"log"
	"net"
	"sync"
)

func Server(wg *sync.WaitGroup) {
	ln, err := net.Listen("tcp", ":"+localPort)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("TCP server running on port", localPort)
	wg.Done()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
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
		// handle block data
		var b block.Block
		err := json.Unmarshal(buf, &b)
		if err != nil {
			conn.Write([]byte(localAddress + ": Receive error json data format."))
			return
		}
		consensus.HandleBlock(&b)
		conn.Write([]byte(localAddress + ": Receive data success."))
	}
}
