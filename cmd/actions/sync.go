package actions

import (
	"github.com/guionardo/go-dev/cmd/ctx"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/urfave/cli/v2"
)

func SyncAction(c *cli.Context) error {
	c2 := ctx.GetContext(c)

	for _, folder := range c2.Config.DevFolders {
		if c.IsSet(consts.FlagMaxDept) && c.Int(consts.FlagMaxDept) > 0 && folder.MaxDepth != c.Int(consts.FlagMaxDept) {
			folder.MaxDepth = c.Int(consts.FlagMaxDept)

		}
		if err := folder.Sync(); err != nil {
			c2.LastErr = err
			return err
		}
	}

	return nil
}
