package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinjqiu/zoom/db"
)

type ZoomApi struct {
	Host string
	Port int
}

func (api *ZoomApi) Serve() {
	provider := db.NewCsvDataProvider("_data")

	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/geo/{ip}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ip := vars["ip"]
		block := provider.GetBlockByIP(ip)
		location := provider.GetLocationByGeonameId(block.GeonameId)
		geo := Geo{}
		geo.Ip = ip
		geo.Location = Location{
			CountryName:    location.CountryName,
			CountryISOCode: location.CountryISOCode,
		}

		result, err := json.MarshalIndent(geo, "", "    ")
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		rw.Write(result)
	})
	bindAddr := fmt.Sprintf("%s:%d", api.Host, api.Port)
	log.Printf("Serving Zoom at %s", bindAddr)
	log.Fatalf("%s", http.ListenAndServe(bindAddr, nil))
}
