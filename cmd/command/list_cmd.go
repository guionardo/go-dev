package command

import (
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/urfave/cli/v2"
	"strings"
)

var (
	ListCmd = &cli.Command{
		Name:   "list",
		Usage:  "List folders",
		Action: ListAction,
		Before: BeforeActionLoadConfiguration,
	}
)

func ListAction(*cli.Context) error {
	_, err := fmt.Println(strings.Join(configuration.DefaultConfig.Paths.FolderList(), "\n"))

	return err
}
