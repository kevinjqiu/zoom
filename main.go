package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const VERSION = "0.1.0"

func actionServe(c *cli.Context) {
	host := c.String("host")
	port := c.Int("port")

	server := ZoomApi{
		Host: host,
		Port: port,
	}

	server.Serve()
}

func actionUpdate(c *cli.Context) {
}

func actionQuery(c *cli.Context) {
}

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
