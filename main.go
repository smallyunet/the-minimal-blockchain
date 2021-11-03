package main

import (
	"the-minimal-blockchain/network"
	"time"
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
