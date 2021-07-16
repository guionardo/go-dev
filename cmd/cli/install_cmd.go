package command

import "github.com/urfave/cli/v2"

var (
	installRemove bool
	InstallCmd    = &cli.Command{
		Name:  "install",
		Usage: "Install go-dev",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "uninstall",
				Usage:       "Remove installation",
				Destination: &installRemove,
			}},
		Action: InstallAction,
		Before: BeforeInstallAction,
	}
)

func InstallAction(context *cli.Context) error {
	return nil
}

func BeforeInstallAction(context *cli.Context) error {
	return nil
}
