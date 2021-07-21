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
	ConfigFile              ConfigFileType
	MaxFolderLevel          int
	DevFolder               string
	ConfigFileName          string
	DefaultDevFolder        string
	DefaultFolderConfigFile string
	DefaultOutputFile       string
)

func CurrentDir() string {
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}
	return "."
}

func SetupBaseEnvironment() error {
	var err error
	HomePath, err = os.UserHomeDir()
	if err == nil {
		DefaultFolderConfigFile = path.Join(HomePath, ".dev_folders_go.json")
		DefaultDevFolder = path.Join(HomePath, "dev")
		DefaultOutputFile = path.Join(HomePath, ".dev_folders_go.out")
	}
	utils.SetOutput(DefaultOutputFile)

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
	ConfigFileName = configurationFile
}

func SetupEnvironment() {
	ConfigFile = ConfigFileType{
		DevFolder:         path.Join(HomePath, "dev"),
		Paths:             make(Paths),
		ConfigurationFile: ConfigFileName,
		MaxSubLevels:      MaximumSubLevel,
	}
	var err error
	if NeedUpdateConfigFile(ConfigFileName, false) {
		if !utils.FileExists(ConfigFileName) {
			log.Printf("Configuration file will be created: %s\n", ConfigFileName)
		} else {
			err = ConfigFile.Load(ConfigFileName)
			if err != nil {
				log.Printf("Error reading configuration file. Will be recreated: %v\n", err)
			}
		}
		if err = ConfigFile.Paths.ReadFolders(DevFolder, MaxFolderLevel); err == nil {
			if err = ConfigFile.Save(); err == nil {
				log.Printf("Updated configuration file: %s\n", ConfigFileName)
			} else {
				log.Fatalf("Failed to save configuration file: %s - %v\n", ConfigFileName, err)
			}
		} else {
			log.Fatalf("Failed to read folders: %v\n", err)
		}
	} else {
		err = ConfigFile.Load(ConfigFileName)
		if err != nil {
			if err = os.Remove(ConfigFileName); err == nil {
				log.Printf("Error reading configuration file. Will be recreated on next run: %v\n", err)
			} else {
				log.Fatalf("Failed to remove invalid configuration file %s - %v\n", ConfigFileName, err)
			}
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
	if time.Now().Unix()-info.ModTime().Unix() > MaximumAge {
		log.Printf("Reloading paths after %v", time.Duration(MaximumAge))
		return true
	}
	return false
}
