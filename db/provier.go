package db

type ZoomDataProvider interface {
	GetLocationByGeonameId(string) Location
	GetBlockByIP(string) Block
}
