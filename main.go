package main

import (
	"os"

	"github.com/bradgignac/slingshot/command"
	"github.com/codegangsta/cli"
)

const (
	name    = "slingshot"
	usage   = "Store configuration files to etcd"
	version = "0.0.0"
	author  = "Brad Gignac"
	email   = "bgignac@bradgignac.com"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Author = author
	app.Email = email

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "peer, p",
			Usage: "Etcd peers to connect to",
			Value: &cli.StringSlice{},
		},
	}

	app.Commands = []cli.Command{
		command.Push,
	}

	app.Run(os.Args)
}
