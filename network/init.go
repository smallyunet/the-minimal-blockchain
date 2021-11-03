package network

import (
	"log"
	"net"
	"strconv"
	"the-minimal-blockchain/config"
)

var RouteTable []string
var port string

func init() {
	port = strconv.Itoa(config.DEFAULT_PORT)
	ips, err := net.LookupIP(config.SEED_DOMAIN)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for _, ip := range ips {
		RouteTable = append(RouteTable, ip.String()+":"+port)
	}
	log.Println("RouteTable:", RouteTable)
}
