package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kevinjqiu/zoom/api"
	"github.com/kevinjqiu/zoom/db"
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
			Name: "test",
			Action: func(c *cli.Context) {
				provider := db.NewCsvDataProvider("_data")
				fmt.Println(provider)
				// db.Locations()
				// fmt.Println(db.Blocks())
			},
		},
	}

	app.Run(os.Args)
}
