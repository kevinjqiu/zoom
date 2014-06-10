package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
)

const VERSION = "0.1.0"

func actionServe(c *cli.Context) {
	port := c.String("port")

	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/getcountry", func(rw http.ResponseWriter, r *http.Request) {
		println(r)
	})
	log.Printf("Serving Zoom at :%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Printf("Finished")
}

func actionUpdate(c *cli.Context) {
}

func main() {
	app := cli.NewApp()
	app.Name = "zoom"
	app.Usage = "Command line tool for zoom"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "Start zoom server",
			Action: actionServe,
			Flags: []cli.Flag{
				cli.StringFlag{"port", "5656", "Port"},
			},
		},
		{
			Name:   "update",
			Usage:  "Update GeoLite2 database",
			Action: actionUpdate,
		},
	}

	app.Run(os.Args)
}
