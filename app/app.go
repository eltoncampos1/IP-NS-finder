package app

import (
	"github.com/urfave/cli"
	"log"
	"net"
	"fmt"
)

func Generate() *cli.App {
  app := cli.NewApp()
  app.Name = "Cli APP"
  app.Usage = "Find IPs and server names from internet"
 
  flags := []cli.Flag{
  	    cli.StringFlag{
			Name: "host",
			Value: "google.com",
	    },
	}


  app.Commands = []cli.Command{
	{
	  Name: "ip",
	  Usage: "Find Ips from internet addresses",
	  Flags: flags, 
	  Action: findIps,
	},
	{
	  Name: "servers",
	  Usage: "Find the server name from internet",
	  Flags: flags,
	  Action: findServers,
	},
  }

  return app
}

func findIps(c *cli.Context){
  host := c.String("host")
  ips, err := net.LookupIP(host)
  if err != nil {
     log.Fatal(err)
  }

  for _, ip := range ips {
	fmt.Println(ip)
  }
}

func findServers(c *cli.Context){
  host := c.String("host")
  servers, err := net.LookupNS(host)
  if err != nil {
     log.Fatal(err)
  }

  for _, server := range servers {
	fmt.Println(server)
  }
}
