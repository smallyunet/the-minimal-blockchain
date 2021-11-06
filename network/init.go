package network

import (
	"log"
	"net"
	"strconv"
	"the-minimal-blockchain/config"
)

var RouteTable = map[string]uint64{
	"127.0.0.1:25000": 0,
	"127.0.0.1:25001": 0,
}
var defaultPort string
var localPort string

func init() {
	// get local port from config
	localPort = strconv.Itoa(config.LOCAL_PORT)

	// get default port from config
	defaultPort = strconv.Itoa(config.DEFAULT_PORT)
	// lookup ip address from seed domain
	ips, err := net.LookupIP(config.SEED_DOMAIN)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// append addr to route table
	for _, ip := range ips {
		addr := ip.String() + ":" + defaultPort
		RouteTable[addr] = 0
	}
	log.Println("RouteTable:", RouteTable)
}
