package actions

import (
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/urfave/cli/v2"
)

func SyncAction(c *cli.Context) error {
	c2 := ctx.GetContext(c)
	for _, folder := range c2.Config.DevFolders {
		if err := folder.Sync(); err != nil {
			c2.LastErr = err
			return err
		}
	}
	return nil
}
