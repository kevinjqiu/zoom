package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kevinjqiu/zoom/api"
	"github.com/kevinjqiu/zoom/db"
)

func actionServe(c *cli.Context) {
	host := c.String("host")
	port := c.Int("port")

	server := api.ZoomApi{
		Host:         host,
		Port:         port,
		DataProvider: db.NewCsvDataProvider("_data"),
	}

	server.Serve()
}

func actionUpdate(c *cli.Context) {
}

func actionQuery(c *cli.Context) {
	provider := db.NewCsvDataProvider("_data")
	for _, ipStr := range c.Args() {
		targetIP := net.ParseIP(ipStr)
		if targetIP == nil {
			fmt.Fprintf(os.Stderr, "%s is not a valid ip address\n", ipStr)
			continue
		}
		block := provider.GetBlockByIP(targetIP)
		loc := provider.GetLocationByGeonameId(block.GeonameId)
		fmt.Printf("%s: %s\n", ipStr, loc.CountryName)
	}
}
