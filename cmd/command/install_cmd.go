package command

import (
	"bitbucket.org/kardianos/osext"
	_ "embed"
	"fmt"
	"github.com/guionardo/go-dev/cmd/configuration"
	"github.com/guionardo/go-dev/cmd/utils"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"os/user"
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
			log.Printf("Configuration file saved @ %s (base folder = %s)\n", newConfig.ConfigurationFile, newConfig.DevFolder)
			goDevBinFile, err := installBinary()
			if err == nil {
				goDevScript, err := installScript(goDevBinFile)
				if err == nil {
					if err = installAlias(goDevScript); err == nil {
						if err = installEnv(newConfig.ConfigurationFile); err == nil {
							log.Println("Alias setup is done")
						}
					}
				}
			}
		}
	}
	if err != nil {
		log.Printf("Install failed: %v\n", err)
	} else{
		installCronUpdate()
	}
	return err
}

func installScript(filename string) (string, error) {
	var scriptFile string

	devSh = strings.ReplaceAll(devSh, "{GO_DEV}", filename)
	scriptFile = path.Join(filepath.Dir(filename), "go-dev.sh")
	err := os.WriteFile(scriptFile, []byte(devSh), 0655)

	return scriptFile, err
}

func useOrCreateFolder(folders []string) (string, error) {
	for _, folder := range folders {
		if utils.PathExists(folder) {
			log.Printf("Using bin folder: %s\n", folder)
			return folder, nil
		}
	}
	err := os.MkdirAll(folders[0], os.ModePerm)
	if err == nil {
		log.Printf("Using created bin folder: %s\n", folders[0])
		return folders[0], nil
	}

	return "", err
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func installBinary() (string, error) {
	userBinFolder, err := useOrCreateFolder([]string{
		path.Join(configuration.HomePath, ".bin"),
		path.Join(configuration.HomePath, "bin"),
	})
	if err != nil {
		return "", err
	}
	currentBinFile, _ := osext.Executable()
	userBinFile := path.Join(userBinFolder, "go-dev")
	nBytes, err := copyFile(currentBinFile, userBinFile)
	if err == nil && nBytes > 0 {
		if err = os.Chmod(userBinFile, 0777); err == nil {
			return userBinFile, nil
		}
	}
	return "", err
}

func installEnv(configurationFile string) error {
	bashRcFile := path.Join(configuration.HomePath, ".bashrc")

	var bashLines []string
	if utils.FileExists(bashRcFile) {
		bashRc, err := os.ReadFile(bashRcFile)
		if err != nil {
			return err
		}
		bashLines = strings.Split(string(bashRc), "\n")

		clearedBashLines := utils.Filter(bashLines,
			func(w string) bool {
				return !strings.Contains(w, "GO_DEV")
			})
		if len(bashLines) != len(clearedBashLines) {
			log.Printf("Removing old configurations from %s\n", bashRcFile)
			bashContent := strings.Join(clearedBashLines, "\n")
			if err = os.WriteFile(bashRcFile, []byte(bashContent), 0644); err != nil {
				log.Printf("Failed to save file %s - %v", bashRcFile, err)
				return err
			}
			bashLines = clearedBashLines
		}
	} else {
		bashLines = make([]string, 0)
	}
	bashLines = append(bashLines, fmt.Sprintf("export GO_DEV_CONFIG=%s", configurationFile))
	bashContent := strings.Join(bashLines, "\n") + "\n"
	err := os.WriteFile(bashRcFile, []byte(bashContent), 0644)
	if err == nil {
		log.Printf("Updated configurations: %s\n", bashRcFile)
	}
	return err
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

func installCronUpdate() {
	user, err := user.Current()
	if err != nil {
		log.Printf("Failed to read user data")
		return
	}
	log.Printf("Update cron tab with this commands:")
	log.Printf("1. Open cron tab configurations for your user:")
	log.Printf("crontab -u ")
	log.Printf("2. Add this line to crontab:")
	log.Printf("0 8 * * * %s/.bin/go-dev update", user.HomeDir)
}

func BeforeInstallAction(context *cli.Context) error {
	configuration.SetupEnvironmentVars(context.String(BaseFolderArg), context.String(ConfigArg))

	if !configuration.DefaultConfig.TryLoad(configuration.ConfigFileName) {
		log.Printf("New configuration file: %s", configuration.ConfigFileName)
	}
	return nil
}
