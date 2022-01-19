package tcp

import "testing"

func TestSend(t *testing.T) {
	res := Send("tcp", "127.0.0.1:25000", "Hello, Server!")
	t.Log(res)
}
