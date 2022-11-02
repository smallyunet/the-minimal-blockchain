package route

import (
	"github.com/smallyunet/tmb/config"
	"log"
	"net"
	"strconv"
)

// RouteTable {address: timestamp}
var RouteTable = map[string]uint64{
	"127.0.0.1:25010": 0,
	"127.0.0.1:25011": 0,
	"127.0.0.1:25012": 0,
}

func Init() {
	// lookup ip address from seed domain
	ips, err := net.LookupIP(config.SeedDomain)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// append addr to route table
	for _, ip := range ips {
		addr := ip.String() + ":" + strconv.Itoa(config.DefaultPort)
		RouteTable[addr] = 0
	}
}
