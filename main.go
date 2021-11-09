package main

import (
	"github.com/smallyunet/tmb/network"
)

func init() {
}

func main() {
	go network.Server()
	<-make(chan int)
}
