package api

import "net"

type Location struct {
	CountryName    string `json:"country_name"`
	CountryISOCode string `json:"country_iso_code"`
}

type Geo struct {
	Ip       net.IP   `json:"ip"`
	Location Location `json:"location"`
}

type Error struct {
	Message string `json:"message"`
}
