package db

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
	NetworkStartIp              string
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
