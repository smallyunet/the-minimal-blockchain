package storage

import (
	"github.com/smallyunet/tmb/block"
	"testing"
)

func TestGetHeight(t *testing.T) {
	height, err := GetHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(height)
}

func TestSet(t *testing.T) {
	b := &block.Block{
		Payload: []block.KeyValue{{Key: "a", Value: "b"}},
	}
	err := Set(1, b)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	b, err := Get(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*b)
}

func TestAdd(t *testing.T) {
	b, err := Add([]block.KeyValue{{Key: "a", Value: "b"}})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*b)
}
