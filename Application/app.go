package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func GenerateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Application Console"
	app.Usage = "Search ips and server names on the internet"

	flags := []cli.Flag{
		cli.StringSliceFlag{
			Name:  "host",
			Value: &cli.StringSlice{"google.com.br"},
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "findIp",
			Usage:  "Search IP addresses on the internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "findServer",
			Usage:  "Search server names on the internet",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIps(c *cli.Context) {

	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {

	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
