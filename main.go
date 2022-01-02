package main

import (
	"log"

	"github.com/smallyunet/tmb/http"
	"github.com/smallyunet/tmb/network"
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
	go network.Server()
	go http.Server()
	<-make(chan int)
}
