package service

import (
	"strconv"

	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/util"
)

// tcp buffer max size
const tcpBufferSize = config.TcpBufferSize

// port of whole network default define
var defaultPort string

// port of local node start
var localPort string

func init() {
	// get local port from config
	localPort = strconv.Itoa(config.LocalPort)

	// get default port from config
	defaultPort = strconv.Itoa(config.DefaultPort)

	injuctEnvVar()
}

func injuctEnvVar() {
	value, ok := util.GetEnvVar("LocalPort")
	if ok {
		localPort = value
	}
}
