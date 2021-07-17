package configuration

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var (
	DevFolder       string
	DevFolderConfig string
	Config          = make(Paths)
)

const MaximumAge = 86400
const MaximumSubLevel = 4

func DefaultFolderConfig() string {
	var home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, ".dev_folders_go.json")
}

func DefaultDevFolder() string {
	var home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, "dev")
}
func Setup() {
	var home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	DevFolder = path.Join(home, "dev")
	DevFolderConfig = path.Join(home, ".dev_folders_go.json")
	if NeedUpdateConfigFile(DevFolderConfig, false) {
		Config.ReadFolders(DevFolder, MaximumSubLevel)
		Config.Save(DevFolderConfig)
	} else {
		err := Config.Load(DevFolderConfig)
		if err != nil {
			os.Remove(DevFolderConfig)
			log.Fatalf("Erro na leitura do arquivo de configuração (será reconstruído) %v", err)
		}
	}
}

func ReadFolders() {
	Config.ReadFolders(DevFolder, MaximumSubLevel)
	Config.Save(DevFolderConfig)
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

func FolderChoice(pastas []string) string {
	switch len(pastas) {
	case 0:
		log.Println("Não foi encontrada pasta")
		return ""
	case 1:
		return pastas[0]
	}

	var choice = -1
	for choice < 0 {
		for i, s := range pastas {
			fmt.Printf("%d - %s\n", i, s)
		}
		fmt.Printf("Escolha uma pasta: ")
		reader := bufio.NewReader(os.Stdin)
		n, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		choice, err = strconv.Atoi(strings.Split(n, "\n")[0])
		if err != nil {
			log.Fatal("Opção inválida")
		}
		if choice >= 0 && choice < len(pastas) {
			return pastas[choice]
		}

		choice = -1
	}

	return ""
}

func GetFolders() string {
	return strings.Join(Config.FolderList(), "\n")
}
