package http

import (
	"strconv"

	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/util"
)

var httpPort string

func init() {
	httpPort = strconv.Itoa(config.HTTP_PORT)
	injuctEnvVar()
}

func injuctEnvVar() {
	value, ok := util.GetEnvVar("HTTPPort")
	if ok {
		httpPort = value
	}
}
