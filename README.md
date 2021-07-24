# go-dev

[![.github/workflows/release.yaml](https://github.com/guionardo/go-dev/actions/workflows/release.yaml/badge.svg)](https://github.com/guionardo/go-dev/actions/workflows/release.yaml)
[![Go](https://github.com/guionardo/go-dev/actions/workflows/go.yml/badge.svg)](https://github.com/guionardo/go-dev/actions/workflows/go.yml)
[![CodeQL](https://github.com/guionardo/go-dev/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/guionardo/go-dev/actions/workflows/codeql-analysis.yml)

CLI for open and initiate development folder projects.

Just go in terminal and "dev my_project", and the tool will find the project into folders and open it.

## Installing

Docker must be available.

```shell
 bash <(curl -s https://raw.githubusercontent.com/guionardo/go-dev/develop/install.sh)
```

## Environment

* GO_DEV_CONFIG = Configuration file (default = ~/.dev_folders_go.json)

## Commands

```shell
└─ $ ▶ bin/go-dev help
NAME:
   go-dev - Go to your projects

USAGE:
   go-dev [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Guionardo Furlan <guionardo@gmail.com>

COMMANDS:
   go       Go to folder
   setup    Setup configuration for folder
   list     List folders
   update   Update folders
   install  Install go-dev
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value  Configuration file (default: "/home/guionardo/.dev_folders_go.json") [$GO_DEV_CONFIG]
   --output value  Output file for command execution (default: "/home/guionardo/.dev_folders_go.out") [$GO_DEV_OUTPUT]
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)

```

## Links

* https://github.com/marketplace/actions/go-release-binaries
* https://pkg.go.dev/github.com/urfave/cli/v2#App
