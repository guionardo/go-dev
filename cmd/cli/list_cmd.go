package command

import (
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

var (
	ListCmd   = &cli.Command{
		Name:      "list",
		Usage:     "List folders",
		Action:    ListAction,
		Before:    BeforeListAction,
	}
)

func ListAction(context *cli.Context) error {
	fmt.Println(strings.Join(configuration.DefaultConfig.Paths.FolderList(), "\n"))

	return nil
}

func BeforeListAction(context *cli.Context) error {
	configuration.SetupEnvironmentVars(context.String("basefolder"), context.String("config"))

	if !configuration.DefaultConfig.TryLoad(configuration.ConfigurationFileName) {
		log.Fatalf("Failed to read configuration file %s",configuration.ConfigurationFileName)
	}
	return nil
}