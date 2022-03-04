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
var localIP string
var localAddress string

func init() {
	// get local port from config
	localPort = strconv.Itoa(config.LocalPort)
	localIP = config.LocalIP
	injectEnvVar()
	localAddress = localIP + ":" + localPort
}

func injectEnvVar() {
	value, ok := util.GetEnvVar("LocalPort")
	if ok {
		localPort = value
	}
	value, ok = util.GetEnvVar("LocalIP")
	if ok {
		localIP = value
	}
}
