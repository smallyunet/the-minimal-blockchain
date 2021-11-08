package main

import (
	"time"

	"github.com/smallyunet/tmb/network"
)

func init() {
}

func main() {
	go network.Server()
	for {
		time.Sleep(time.Second * 2)
		network.Client()
	}
}
