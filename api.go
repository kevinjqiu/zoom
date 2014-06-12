package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinjqiu/zoom/db"
)

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

type ZoomApi struct {
	Host string
	Port int
}

func jsonResponder(handler func(map[string]string) (int, interface{})) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		statusCode, result := handler(vars)

		jsonResponse, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(statusCode)
		rw.Write(jsonResponse)
	}
}

func (api *ZoomApi) Serve() {
	provider := db.NewCsvDataProvider("_data")

	r := mux.NewRouter()
	r.HandleFunc("/geo/{ip}", jsonResponder(func(vars map[string]string) (int, interface{}) {
		log.Printf("geoip requested: %s", vars["ip"])
		ip := net.ParseIP(vars["ip"])
		if ip == nil {
			return 400, Error{fmt.Sprintf("%q is not a valid IP address", vars["ip"])}
		}

		block := provider.GetBlockByIP(ip)
		location := provider.GetLocationByGeonameId(block.GeonameId)

		geo := Geo{
			Ip: ip,
			Location: Location{
				CountryName:    location.CountryName,
				CountryISOCode: location.CountryISOCode,
			},
		}

		return 200, geo
	}))

	http.Handle("/", r)

	bindAddr := fmt.Sprintf("%s:%d", api.Host, api.Port)
	log.Printf("Serving Zoom at %s", bindAddr)
	log.Fatalf("%s", http.ListenAndServe(bindAddr, nil))
}