package network

import (
	"log"
	"net"
	"strconv"

	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/util"
)

// tcp buffer max size
const tcpBufferSize = config.TcpBufferSize

// {address: timestamp}
var RouteTable = map[string]uint64{
	"127.0.0.1:25000": 0,
	"127.0.0.1:25001": 0,
}

// port of whole network default define
var defaultPort string

// port of local node start
var localPort string

func init() {
	// get local port from config
	localPort = strconv.Itoa(config.LocalPort)

	// get default port from config
	defaultPort = strconv.Itoa(config.DefaultPort)
	// lookup ip address from seed domain
	ips, err := net.LookupIP(config.SeedDomain)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// append addr to route table
	for _, ip := range ips {
		addr := ip.String() + ":" + defaultPort
		RouteTable[addr] = 0
	}
	log.Println("RouteTable initialized size:", len(RouteTable))

	injuctEnvVar()
}

func injuctEnvVar() {
	value, ok := util.GetEnvVar("LocalPort")
	if ok {
		localPort = value
	}
}
