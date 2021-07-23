package configuration

import (
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"testing"
)

const (
	folderCount    = 10
	subFolderCount = 4
	letterBytes    = "0123456789ABCDEF"
	permissions    = 0777
)

var (
	testingFolder            string
	testingConfigurationFile string
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func setupFolders() {
	testingFolder, _ = filepath.Abs("./test_dev_folder")
	testingConfigurationFile, _ = filepath.Abs("./test_dev_configuration.json")
	tearDownFolders()
	for nf := 0; nf < folderCount; {
		nf++
		folder := path.Join(testingFolder, RandStringBytes(10))

		for nsf := 0; nsf < subFolderCount; {
			nsf++
			subFolder := path.Join(folder, RandStringBytes(10))
			if err := os.MkdirAll(subFolder, permissions); err != nil {
				log.Fatalf("Failed to create folder %s - %v", subFolder, err)
			} else {
				log.Printf("Folder %s\n", subFolder)
			}
		}
	}
}

func tearDownFolders() {
	if err := os.RemoveAll(testingFolder); err != nil {
		log.Printf("Failed to remove %s - %v", testingFolder, err)
	}
	if err := os.Remove(testingConfigurationFile); err != nil {
		log.Printf("Failed to remove %s - %v", testingConfigurationFile, err)
	}
}

func init() {
	setupFolders()
}

func TestSaveAndLoad(t *testing.T) {
	cf := ConfigFileType{
		DevFolder:         testingFolder,
		Paths:             make(Paths),
		ConfigurationFile: testingConfigurationFile,
		MaxSubLevels:      3,
	}
	defer tearDownFolders()

	if err := cf.Paths.ReadFolders(cf.DevFolder, cf.MaxSubLevels); err != nil {
		t.Errorf("Failed to read folders - %v", err)
	}
	if err := cf.Save(); err != nil {
		t.Errorf("Failed to save configuration - %v", err)
	}

	cf2 := ConfigFileType{}
	if err := cf2.Load(testingConfigurationFile); err != nil {
		t.Errorf("Failed to read configuration - %v", err)
	}

}
