package util

import (
	"fmt"
	"os"
	"strings"
)

var variables map[string]string

func init() {
	if variables == nil {
		variables = make(map[string]string)
	}
	freshCache()
}

func freshCache() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		variables[pair[0]] = pair[1]
	}
}

func GetEnvVar(name string) (string, bool) {
	value, ok := variables[name]
	if !ok {
		freshCache()
	}
	value, ok = variables[name]
	fmt.Println(variables)
	return value, ok
}
