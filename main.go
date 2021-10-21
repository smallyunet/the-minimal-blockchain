package main

func main() {
	// go server()
	// go client()
	// <-make(chan interface{})
	height, _ := GetHeight()
	Set(height)
}
