package main

import (
	"github.com/smallyunet/tmb/server/http"
	"github.com/smallyunet/tmb/server/tcp"
	"log"
	"sync"

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
	var wg sync.WaitGroup
	wg.Add(2)
	go tcp.Server(&wg)
	go http.Server(&wg)
	wg.Wait()
	log.Println("-----------started-------------")
	<-make(chan int)
}
