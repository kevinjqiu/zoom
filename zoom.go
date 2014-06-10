package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/kevinjqiu/zoom/api"
)

const VERSION = "0.1.0"

func actionServe(c *cli.Context) {
	host := c.String("host")
	port := c.Int("port")

	server := api.ZoomApi{
		Host: host,
		Port: port,
	}

	server.Serve()
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
				cli.IntFlag{"port", 5656, "Port"},
				cli.StringFlag{"host", "", "Host"},
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
