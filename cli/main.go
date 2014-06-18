package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const VERSION = "0.1.0"

const DEFAULT_PORT = 5656
const DEFAULT_ADDR = "127.0.0.1"

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
				cli.IntFlag{"port", DEFAULT_PORT, "Port"},
				cli.StringFlag{"host", DEFAULT_ADDR, "Host"},
			},
		},
		{
			Name:   "update",
			Usage:  "Update GeoLite2 database",
			Action: actionUpdate,
		},
		{
			Name:   "query",
			Usage:  "Query the geolocation of an IP",
			Action: actionQuery,
		},
	}

	app.Run(os.Args)
}
