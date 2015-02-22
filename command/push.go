package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/mailgun/go-etcd/etcd"
)

// Push writes configuration files to etcd.
var Push = cli.Command{
	Name:   "push",
	Usage:  "Write configuration files to etcd",
	Action: push,
}

func push(c *cli.Context) {
	configPaths := c.Args()
	configFiles, err := findConfigFiles(configPaths)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	peers := c.GlobalStringSlice("peer")
	client := etcd.NewClient(peers)
	for _, file := range configFiles {
		fmt.Printf("Uploading %s...\n", file)

		contents, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// TODO: Automatically detect key collisions.

		key := strings.TrimSuffix(file, filepath.Ext(file))
		_, err = client.Set(key, string(contents), 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func findConfigFiles(paths []string) ([]string, error) {
	files := map[string]bool{}

	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() || !isValidConfig(path) {
				return nil
			}

			if _, ok := files[path]; ok {
				return nil
			}

			files[path] = false

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return convertMapToSlice(files), nil
}

func isValidConfig(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".json", ".yaml", ".txt":
		return true
	}

	return false
}

func convertMapToSlice(m map[string]bool) []string {
	i := 0
	s := make([]string, len(m))

	for k := range m {
		s[i] = k
		i++
	}

	return s
}
