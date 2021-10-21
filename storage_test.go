package main

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
