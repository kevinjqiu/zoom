package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ZoomApi struct {
	Host string
	Port int
}

func (api *ZoomApi) Serve() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/getcountry", func(rw http.ResponseWriter, r *http.Request) {
		println(r)
	})
	bindAddr := fmt.Sprintf("%s:%d", api.Host, api.Port)
	log.Printf("Serving Zoom at %s", bindAddr)
	http.ListenAndServe(bindAddr, nil)
	log.Printf("Finished")
}
