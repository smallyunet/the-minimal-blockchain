package network

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func Client() {
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "test")
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	log.Println(buf.String())
}
