package main

import (
	"github.com/smallyunet/tmb/server/http"
	"github.com/smallyunet/tmb/server/tcp"
	"log"

	"github.com/smallyunet/tmb/storage"
)

func init() {
}

func main() {
	h, err := storage.GetHeight()
	if err != nil {
		panic(err)
	}
	log.Println("Current block height:", h)
	go tcp.Server()
	go http.Server()
	<-make(chan int)
}
