package main

import (
	"github.com/smallyunet/tmb/pool"
	"github.com/smallyunet/tmb/route"
	"github.com/smallyunet/tmb/server"
	"github.com/smallyunet/tmb/storage"
)

func init() {
}

func main() {
	storage.Init()
	route.Init()
	server.InitHttp()
	server.InitTcp()

	go server.HttpServer()
	go server.TcpServer()
	go pool.AcceptTx()
	go pool.AcceptBlock()

	<-make(chan int)
}
