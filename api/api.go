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
	r.HandleFunc("/geo/{ip}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		println(vars["ip"])
	})
	bindAddr := fmt.Sprintf("%s:%d", api.Host, api.Port)
	log.Printf("Serving Zoom at %s", bindAddr)
	log.Fatalf("%s", http.ListenAndServe(bindAddr, nil))
}
