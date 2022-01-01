package util

import "testing"

func TestGetEnvVar(t *testing.T) {
	value, ok := GetEnvVar("aaaa")
	t.Log(value, ok)
}
