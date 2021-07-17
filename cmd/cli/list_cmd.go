package command

import (
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
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
	fmt.Println(strings.Join(configuration.Config.FolderList(), "\n"))

	return nil
}

func BeforeListAction(context *cli.Context) error {
	return nil
}