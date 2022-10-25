# go-dev

[![.github/workflows/release.yaml](https://github.com/guionardo/go-dev/actions/workflows/release.yaml/badge.svg)](https://github.com/guionardo/go-dev/actions/workflows/release.yaml)
[![Go](https://github.com/guionardo/go-dev/actions/workflows/go.yml/badge.svg)](https://github.com/guionardo/go-dev/actions/workflows/go.yml)
[![CodeQL](https://github.com/guionardo/go-dev/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/guionardo/go-dev/actions/workflows/codeql-analysis.yml)

CLI for open and initiate development folder projects.

Just go in terminal and "dev my_project", and the tool will find the project into folders and open it.

## Installing

### Using go install

```bash
go install github.com/guionardo/go-dev
``` 

### Download from releases

Only for linux-amd64

Go to [release page](https://github.com/guionardo/go-dev/releases) and download the latest go-dev-v*-linux-amd64.tar.gx

Extract the go-dev file into a folder. Assure that this folder is in PATH variable.

## Check the installation

```bash
❯ go-dev --version
go-dev version 0.0.0
```



## Environment

* GO_DEV_CONFIG = Configuration file (default = ~/.config/go-dev/go-dev.yaml)

## Commands

```shell
❯ go-dev --help
NAME:
   go-dev - Go to your projects

USAGE:
   go-dev [global options] command [command options] [arguments...]

VERSION:
   1.3.0

DESCRIPTION:
   Builder Info: guionardo@ambevtech-guionardo - 2022-10-24 19:09:36

AUTHOR:
   Guionardo Furlan <guionardo@gmail.com>

COMMANDS:
   go       Execute go command
   setup    Setup go-dev
   sync     Sync go-dev
   init     Initialize go-dev for bash
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config-file value  (default: "/home/guionardo/.config/go-dev/go-dev.yaml") [$GO_DEV_CONFIG]
   --debug              (default: false) [$GO_DEV_DEBUG]
   --output value       (default: "/home/guionardo/.config/go-dev/go-dev.sh")
   --help, -h           show help (default: false)
   --version, -v        print the version (default: false)
```

