package api

type Location struct {
	CountryName    string `json:"country_name"`
	CountryISOCode string `json:"country_iso_code"`
}

type Geo struct {
	Ip       string   `json:"ip"`
	Location Location `json:"location"`
}
