package server

import (
	"encoding/json"
	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/pool"
	"github.com/smallyunet/tmb/util"
	"log"
	"net"
	"strconv"
)

// port of local node start
var localPort string
var localIP string
var localAddress string

func InitTcp() {
	localPort = strconv.Itoa(config.LocalPort)
	localIP = config.LocalIP
	localAddress = localIP + ":" + localPort

	value, ok := util.GetEnvVar(config.LocalPortFlag)
	if ok {
		localPort = value
	}
	value, ok = util.GetEnvVar(config.LocalIPFlag)
	if ok {
		localIP = value
	}
}

func TcpServer() {
	ln, err := net.Listen("tcp", ":"+localPort)
	if err != nil {
		log.Fatalln(err)
	}
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
	buf := make([]byte, 0, config.TcpBufferSize)
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
		pool.PushBlockToPool(&b)
		conn.Write([]byte(localAddress + ": Receive data success."))
	}
}
