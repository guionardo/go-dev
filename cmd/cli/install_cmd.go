package command

import (
	"bitbucket.org/kardianos/osext"
	_ "embed"
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//go:embed go-dev.sh
var devSh string

var (
	installRemove bool
	devBaseFolder string
	InstallCmd    = &cli.Command{
		Name:  "install",
		Usage: "Install go-dev",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "uninstall",
				Usage:       "Remove installation",
				Destination: &installRemove,
			},
			&cli.PathFlag{
				Name:        BaseFolderArg,
				Usage:       "Development base folder",
				Value:       configuration.DefaultDevFolder,
				Destination: &devBaseFolder,
			},
			&cli.IntFlag{
				Name:        "max-path-level",
				Usage:       "Maximum level of paths",
				Value:       configuration.MaximumSubLevel,
				Destination: &configuration.MaxFolderLevel,
			},
		},
		Action: InstallAction,
		Before: BeforeInstallAction,
	}
)

func InstallAction(*cli.Context) error {
	newConfig := &configuration.ConfigFileType{
		DevFolder:         configuration.DevFolder,
		Paths:             make(configuration.Paths),
		ConfigurationFile: configuration.ConfigFileName,
	}
	var err error
	if err = newConfig.Paths.ReadFolders(configuration.DevFolder, configuration.MaxFolderLevel); err == nil {
		if err = newConfig.Save(); err == nil {
			fmt.Printf("Configuration file saved @ %s (base folder = %s)\n", newConfig.ConfigurationFile, newConfig.DevFolder)
			goDevScript, err := installScript()
			if err == nil {
				if err = installAlias(goDevScript); err == nil {
					if err = installEnv(); err == nil {
						fmt.Println("Alias setup is done")
					}
				}
			}
		}
	}

	return err
}

func installScript() (string, error) {
	filename, err := osext.Executable()
	var scriptFile string
	if err == nil {
		devSh = strings.ReplaceAll(devSh, "{GO_DEV}", filename)
		scriptFile = path.Join(filepath.Dir(filename), "go-dev.sh")
		err = os.WriteFile(scriptFile, []byte(devSh), 0655)
	}
	return scriptFile, err
}

func installEnv() error {
	//TODO: Implementar a gravação de export GO-DEV-CONFIG em .bashrc
	//bashRcFile:=path.Join(configuration.HomePath,".bashrc")
	//var bashRc string
	//if utils.FileExists(bashRcFile){
	//	bashRc,err:=os.ReadFile(bashRc)
	//	if err!=nil{
	//		return err
	//	}
	//	if strings.Contains(bashRcFile,"export GO-DEV-CONFIG="){
	//
	//	}
	//}
	return nil
}

func installAlias(goDevShFile string) error {
	aliasDevLine := fmt.Sprintf("alias dev='source %s'", goDevShFile)
	var err error
	bashRcAliasesFile := path.Join(configuration.HomePath, ".bash_aliases")
	var lines []string
	if utils.FileExists(bashRcAliasesFile) {
		aliases, err := os.ReadFile(bashRcAliasesFile)
		if err == nil {
			lines = strings.Split(string(aliases), "\n")
			aliasIndex := -1
			for index, line := range lines {
				if strings.HasPrefix(line, "alias dev=") {
					aliasIndex = index
					break
				}
			}
			if aliasIndex > -1 {
				lines[aliasIndex] = aliasDevLine
			} else {
				lines = append(lines, aliasDevLine)
			}
		}
	} else {
		lines = append(lines, aliasDevLine)
	}
	content := strings.Join(lines, "\n")
	err = os.WriteFile(bashRcAliasesFile, []byte(content), 0655)

	return err
}

func BeforeInstallAction(context *cli.Context) error {
	configuration.SetupEnvironmentVars(context.String(BaseFolderArg), context.String(ConfigArg))

	if !configuration.DefaultConfig.TryLoad(configuration.ConfigFileName) {
		log.Printf("New configuration file: %s", configuration.ConfigFileName)
	}
	return nil
}
