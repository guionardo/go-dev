package configuration

import (
	"github.com/guionardo/go-dev/cmd/utils"
	"log"
	"os"
	"path"
	"time"
)

const MaximumAge = 86400
const MaximumSubLevel = 4

var (
	HomePath                string
	ConfigurationFile       ConfigFileType
	MaxFolderLevel          int
	DevFolder               string
	ConfigurationFileName   string
	DefaultDevFolder        string
	DefaultFolderConfigFile string
)

func SetupBaseEnvironment() error {
	var err error
	HomePath, err = os.UserHomeDir()
	if err == nil {
		DefaultFolderConfigFile = path.Join(HomePath, ".dev_folders_go.json")
		DefaultDevFolder = path.Join(HomePath, "dev")
	}

	return err
}

func SetupEnvironmentVars(devFolder string, configurationFile string) {
	if len(devFolder) == 0 {
		devFolder = DefaultDevFolder
	}
	DevFolder = devFolder
	if len(configurationFile) == 0 {
		configurationFile = DefaultFolderConfigFile
	}
	ConfigurationFileName = configurationFile
}

func SetupEnvironment() {
	ConfigurationFile = ConfigFileType{
		DevFolder:         path.Join(HomePath, "dev"),
		Paths:             make(Paths),
		ConfigurationFile: ConfigurationFileName,
		MaxSubLevels:      MaximumSubLevel,
	}
	var err error
	if NeedUpdateConfigFile(ConfigurationFileName, false) {
		if !utils.FileExists(ConfigurationFileName) {
			log.Printf("Configuration file will be created: %s\n", ConfigurationFileName)
		} else {
			err = ConfigurationFile.Load(ConfigurationFileName)
			if err != nil {
				log.Printf("Error reading configuration file. Will be recreated: %v\n", err)
			}
		}
		if err = ConfigurationFile.Paths.ReadFolders(DevFolder, MaxFolderLevel); err == nil {
			if err = ConfigurationFile.Save(); err == nil {
				log.Printf("Updated configuration file: %s\n", ConfigurationFileName)
			} else {
				log.Fatalf("Failed to save configuration file: %s - %v\n", ConfigurationFileName, err)
			}
		} else {
			log.Fatalf("Failed to read folders: %v\n", err)
		}
	} else {
		err = ConfigurationFile.Load(ConfigurationFileName)
		if err != nil {
			os.Remove(ConfigurationFileName)
			log.Printf("Error reading configuration file. Will be recreated on next run: %v\n", err)
		}
	}
}

func NeedUpdateConfigFile(filename string, force bool) bool {
	if force {
		return true
	}
	info, err := os.Stat(filename)
	if err != nil {
		return true
	}
	return time.Now().Unix()-info.ModTime().Unix() > MaximumAge
}
