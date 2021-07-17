package configuration

import (
	"bufio"
	"encoding/json"
	"os"
)

type ConfigFileType struct {
	DevFolder         string `json:"dev_folder,omitempty"`
	Paths             Paths  `json:"paths,omitempty"`
	ConfigurationFile string `json:"configuration_file,omitempty"`
	MaxSubLevels      int    `json:"max_sub_levels,omitempty"`
}

var DefaultConfig = &ConfigFileType{
	DevFolder:         DefaultDevFolder(),
	Paths:             make(Paths),
	ConfigurationFile: DefaultFolderConfig(),
	MaxSubLevels:      MaximumSubLevel,
}

func (cf *ConfigFileType) Load(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		return err
	}
	var size = stats.Size()
	bytes := make([]byte, size)
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	if err != nil {
		return err
	}

	newPc := make(Paths)
	if err := json.Unmarshal(bytes, &newPc); err != nil {
		return err
	}
	for _, p := range newPc {
		cf.Paths.Set(p)
	}
	cf.ConfigurationFile = fileName
	return nil
}

func (cf *ConfigFileType) Save() error {
	folderJson, _ := json.Marshal(cf)
	file, err := os.OpenFile(cf.ConfigurationFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0655)
	if err != nil {
		return err
	}
	file.Write(folderJson)
	file.Close()
	return nil
}
