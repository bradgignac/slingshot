package command

import (
	"fmt"
	"os"

	"github.com/bradgignac/slingshot/config"
	"github.com/bradgignac/slingshot/store"
	"github.com/codegangsta/cli"
)

// Push writes configuration files to etcd.
var Push = cli.Command{
	Name:   "push",
	Usage:  "Write configuration files to etcd",
	Action: push,
}

func push(c *cli.Context) {
	paths := c.Args()
	files, err := config.FindFiles(paths)
	if err != nil {
		fmt.Println(err)
	}

	prefix := c.GlobalString("key")
	peers := c.GlobalStringSlice("peer")
	store := store.NewEtcdStore(prefix, peers)

	for _, file := range files.ToSlice() {
		fmt.Printf("Uploading %v...\n", file)

		err := store.Upload(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
