package storage

import (
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
	block := &Block{
		Payload: "Hello World",
	}
	err := Set(1, block)
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
	err := Add("Hello World")
	if err != nil {
		t.Fatal(err)
	}
}
