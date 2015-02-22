# slingshot

Slingshot is a command-line tool that allows you to easily write configuration values stored in formats like YAML and JSON as well as plain text to etcd.

## Installation

To install Slingshot, you will need to have Go v1.3 or higher installed. Simply run:

```
$ go install github.com/bradgignac/slingshot
```

The slingshot binary is now available in your `$GOPATH`. Once the first version of this tool is officially released, binaries will be published to GitHub.

## Usage

This section describes the commands available in Slingshot. The config files used here are located in the `examples` directory at the root of the repository.

### Writing Configuration

The `push` command provides support for writing config files to etcd.

```
$ slingshot push examples/config.json examples/config.yaml
```

You can push an entire directory or glob of files as well.

```
$ slingshot push examples/**/*
```

#### Changing Data Location

By default, Slingshot uses the path and name of the config file to determine where to store configuration data. In the above examples, the configuration would be written to `/examples/config`. You can also change the location where data is written to in etcd:

```
$ slingshot push --directory custom/location examples/config.json
```

#### Etcd Peers

Slingshot assumes etcd is available at `http://127.0.0.1:4001`. If you need specify an alternate etcd location, use the `--peer` flag to provide one or more URLs for nodes in your etcd cluster.

```
$ slingshot push --peer http://10.10.10.1:4001 examples/config.json
```

For more information about Slingshot commands and their options, run `slingshot --help`.

## License

Slingshot is released under the [MIT License](LICENSE).
