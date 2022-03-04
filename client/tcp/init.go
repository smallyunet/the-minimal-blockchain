package tcp

import (
	"log"
	"net"
	"strconv"

	"github.com/smallyunet/tmb/config"
)

// tcp buffer max size
const tcpBufferSize = config.TcpBufferSize

// RouteTable {address: timestamp}
var RouteTable = map[string]uint64{
	"127.0.0.1:25001": 0,
	"127.0.0.1:25002": 0,
}

// port of whole network default define
var defaultPort string

func init() {
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
}
