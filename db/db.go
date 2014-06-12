package db

import "net"

type ZoomDataProvider interface {
	GetLocationByGeonameId(string) Location
	GetBlockByIP(net.IP) Block
}

type Location struct {
	GeonameId          string
	ContinentCode      string
	ContinentName      string
	CountryISOCode     string
	CountryName        string
	SubDivisionISOCode string
	SubDivisionName    string
	CityName           string
	MetroCode          string
	TimeZone           string
}

type Block struct {
	networkEndIp                net.IP
	NetworkStartIp              net.IP
	NetworkPrefixLength         int
	GeonameId                   string
	RegisteredCountryGeoNameId  string
	RepresentedCountryGeoNameId string
	PostalCode                  string
	Latitude                    string
	Longitude                   string
	IsAnonymousProxy            bool
	IsSatelliteProvider         bool
}

func (this *Block) NetworkEndIp() net.IP {
	if this.networkEndIp == nil {
		ip := this.NetworkStartIp.To16()
		mask := net.CIDRMask(this.NetworkPrefixLength, 128)

		endIp := net.IP(make([]byte, 16, 16))
		for i, _ := range mask {
			endIp[i] = ip[i] | ^mask[i]
		}
		this.networkEndIp = net.IP(endIp)
	}

	return this.networkEndIp
}
