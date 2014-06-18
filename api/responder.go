package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

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
