package db

import "net"

type ZoomDataProvider interface {
	GetLocationByGeonameId(string) Location
	GetBlockByIP(net.IP) Block
}
