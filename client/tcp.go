package client

import (
	"bytes"
	"fmt"
	"github.com/smallyunet/tmb/route"
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
	for k, _ := range route.RouteTable {
		conn, err := net.Dial("tcp", k)
		if err != nil {
			log.Println("Broadcast block data:", err)
			return
		}
		defer conn.Close()
		fmt.Fprint(conn, msg)
		var buf bytes.Buffer
		io.Copy(&buf, conn)
	}
}
