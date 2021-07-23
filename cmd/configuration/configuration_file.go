package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/guionardo/go-dev/cmd/utils"
	"log"
	"os"
	"time"
)

type ConfigFileType struct {
	DevFolder         string `json:"dev_folder,omitempty"`
	Paths             Paths  `json:"paths,omitempty"`
	ConfigurationFile string `json:"configuration_file,omitempty"`
	MaxSubLevels      int    `json:"max_sub_levels,omitempty"`
}

var DefaultConfig = &ConfigFileType{
	DevFolder:         DefaultDevFolder,
	Paths:             make(Paths),
	ConfigurationFile: DefaultFolderConfigFile,
	MaxSubLevels:      MaximumSubLevel,
}

func (cf *ConfigFileType) TryLoad(fileName string) bool {
	if !utils.FileExists(fileName) {
		log.Printf("Failed to load configuration: File not found %s", fileName)
		return false
	}
	err := cf.Load(fileName)
	if err == nil {
		if NeedUpdateConfigFile(fileName, false) {
			err = cf.Paths.ReadFolders(cf.DevFolder, cf.MaxSubLevels)
			if err != nil {
				log.Printf("Failed to load configuration: Read folders failed %v", err)
			}
		}
		return true
	}
	log.Printf("Failed to read configuration file %s: %v\n", fileName, err)
	if utils.FileExists(fileName) {
		var newFile = fmt.Sprintf("%s.%s.error", fileName, time.Now().Format("20060102150405"))
		err = os.Rename(fileName, newFile)
		if err == nil {
			log.Printf("Invalid file %s moved to %s\n", fileName, newFile)
		} else {
			err = os.Remove(fileName)
			if err == nil {
				log.Printf("Invalid file %s was removed\n", fileName)
			} else {
				log.Fatalf("Failed to remove invalid file %s: %v", fileName, err)
			}
		}
	}
	return false
}

func (cf *ConfigFileType) Load(fileName string) error {
	fileContent, err := os.ReadFile(fileName)
	if err == nil {
		newCf := &ConfigFileType{
			DevFolder:         "",
			Paths:             make(Paths),
			ConfigurationFile: "",
			MaxSubLevels:      0,
		}
		if err = json.Unmarshal(fileContent, &newCf); err == nil {
			if cf.Paths == nil {
				cf.Paths = make(Paths)
			}
			for _, p := range newCf.Paths {
				if err = cf.Paths.Set(p); err != nil {
					log.Printf("Failed to add folder %s - %v", p.Path, err)
				}
			}
			cf.ConfigurationFile = fileName
		}
	}
	return err

	//file, err := os.Open(fileName)
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//stats, statsErr := file.Stat()
	//if statsErr != nil {
	//	return err
	//}
	//var size = stats.Size()
	//bytes := make([]byte, size)
	//buffer := bufio.NewReader(file)
	//_, err = buffer.Read(bytes)
	//if err != nil {
	//	return err
	//}
	//newCf := &ConfigFileType{
	//	DevFolder:         "",
	//	Paths:             make(Paths),
	//	ConfigurationFile: "",
	//	MaxSubLevels:      0,
	//}
	//
	//if err := json.Unmarshal(bytes, &newCf); err != nil {
	//	return err
	//}
	//for _, p := range newCf.Paths {
	//	cf.Paths.Set(p)
	//}
	//cf.ConfigurationFile = fileName
	//return nil
}

func (cf *ConfigFileType) Save() error {
	folderJson, err := json.Marshal(cf)
	if err == nil {
		err = os.WriteFile(cf.ConfigurationFile, folderJson, 0655)
	}
	return err
}
