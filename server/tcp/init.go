package tcp

import (
	"strconv"

	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/util"
)

// tcp buffer max size
const tcpBufferSize = config.TcpBufferSize


// port of local node start
var localPort string

func init() {
	// get local port from config
	localPort = strconv.Itoa(config.LocalPort)

	injectEnvVar()
}

func injectEnvVar() {
	value, ok := util.GetEnvVar("LocalPort")
	if ok {
		localPort = value
	}
}
