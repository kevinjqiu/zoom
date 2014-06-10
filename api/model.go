package api

type Address struct {
	CountryName    string `json:"country_name"`
	CountryISOCode string `json:"country_iso_code"`
}

type Geo struct {
	Ip      string  `json:"ip"`
	Address Address `json:"address"`
}
