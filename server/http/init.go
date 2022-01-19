package http

import (
	"strconv"

	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/util"
)

var httpPort string

func init() {
	httpPort = strconv.Itoa(config.HttpPort)
	injectEnvVar()
}

func injectEnvVar() {
	value, ok := util.GetEnvVar("HTTPPort")
	if ok {
		httpPort = value
	}
}
