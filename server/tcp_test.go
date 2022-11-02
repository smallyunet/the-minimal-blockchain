package server

import (
	"sync"
	"testing"
)

func TestServer(t *testing.T) {
	wg := &sync.WaitGroup{}
	Server(wg)
}
