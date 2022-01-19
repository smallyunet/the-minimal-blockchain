package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func Send(protocol, address, msg string) string {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, msg)
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	return buf.String()
}

func SendToAll(msg string) {
	for k, _ := range RouteTable {
		// TODO change the default protocol
		conn, err := net.Dial("tcp", k)
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		fmt.Fprint(conn, msg)
		var buf bytes.Buffer
		io.Copy(&buf, conn)
		log.Println(buf.String())
	}
}
