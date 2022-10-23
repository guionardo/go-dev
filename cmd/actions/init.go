package actions

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/guionardo/go-dev/pkg/config"
	"github.com/urfave/cli/v2"
)

//go:embed init.sh
var initScript string

func InitAction(c *cli.Context) error {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	initScript = strings.ReplaceAll(initScript, "{GO_DEV}", executable)
	initScript = strings.ReplaceAll(initScript, "{GO_OUTPUT}", config.DefaultOutputFile)

	cmds:=make([]string,len(c.App.Commands))
	for i, cmd := range c.App.Commands {
		cmds[i] = cmd.Name
	}
	initScript = strings.ReplaceAll(initScript, "{GO_CMDS}", strings.Join(cmds, " | "))
	
	fmt.Print(initScript)
	return nil
}
