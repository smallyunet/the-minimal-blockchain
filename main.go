package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("miss args")
		return
	}
	mode := args[1]
	if mode == "server" {
		server()
	}
	if mode == "client" {
		client()
	}
}
