package db

import "net"

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
	ip := this.NetworkStartIp.To16()
	mask := net.CIDRMask(this.NetworkPrefixLength, 128)

	endIp := net.IP(make([]byte, 16, 16))
	for i, _ := range mask {
		endIp[i] = ip[i] | ^mask[i]
	}

	return net.IP(endIp)
}
