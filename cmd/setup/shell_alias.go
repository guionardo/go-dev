package setup

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/guionardo/go-dev/cmd/configuration"
)

var (
	shell  string
	rcfile string

)

const (
	EnvVar   string = "GO_DEV_EXIT_COMMAND"
	DevStart string = "GO_DEV START- DO NOT MODIFY"
	DevEnd   string = "GO_DEV END - DO NOT MODIFY"
)

func detectShell() {
	// Detect shell
	shell = os.Getenv("SHELL")
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("ERROR: Failed to get user home dir - %v", err)
	}
	if len(shell) == 0 {
		log.Printf("WARNING: No SHELL environment detected")
		return
	}
	for _, sh := range []string{
		"bash",
		"zsh",
		"ksh",
	} {
		if strings.HasSuffix(shell, sh) {
			shell = sh
			rcfile = path.Join(home, "."+sh+"rc")
			break
		}
	}
	if len(rcfile) == 0 {
		log.Printf("WARNING: unexpected shell - %s", shell)
		return
	}
	if _, err = os.Stat(rcfile); err != nil {
		log.Printf("WARNING: error getting shell resource file '%s' - %v", rcfile, err)
		return
	}
}

// Returns alias file name.
// If folder is empty, will use config dir for current user
// Folder will be created if not exists
func AliasFileName(folder string) (string, error) {
	var err error
	if len(folder) == 0 {
		folder, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		folder = path.Join(folder, "go_dev")
	}
	if _, err = os.Stat(folder); os.IsNotExist(err) {
		err = os.Mkdir(folder, 0666)
	}
	if err != nil {
		return "", err
	}
	if info, err := os.Stat(folder); err == nil && info.IsDir() {
		return path.Join(folder, "go_dev.sh"), nil
	}
	return "", err
}

func CreateAliasFile(folder string) error {
	aliasFile,err:=AliasFileName(folder)
	if err!=nil{
		return err
	}

	
}
func AddAliasFunctionToRC(resourceFile string) error {
	home, _ := os.UserHomeDir()
	description := fmt.Sprintf("Automatically built - %s", configuration.MetaData.ToString())
	text := shellDev
	for _, replace := range []struct {
		from string
		to   string
	}{
		{from: "HOME", to: home},
		{from: "ENVVAR", to: EnvVar},
		{from: "DESCRIPTION", to: description},
		{from: "DEV_START", to: DevStart},
		{from: "DEV_END", to: DevEnd},
	} {
		text = strings.ReplaceAll(text, "#"+replace.from+"#", replace.to)
	}

	file, err := os.OpenFile(resourceFile, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(text)

	return nil
}

func RemoveAliasFunctionFromRC(resourceFile string) error {
	if _, err := os.Stat(resourceFile); err != nil {
		return err
	}
	source, err := os.Open(resourceFile)
	if err != nil {
		return err
	}
	defer source.Close()
	destiny, err := os.CreateTemp("", "go_dev_*")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(source)
	started := false
	removed := false
	for scanner.Scan() {
		line := scanner.Text()
		if !started && strings.Trim(line, " \n\r") == DevStart {
			started = true
			removed = true
			continue
		}
		if started && strings.Trim(line, " \n\r") == DevEnd {
			started = false
			continue
		}
		if !started {
			destiny.WriteString(line)
		}
	}
	if !removed {
		log.Printf("No need to update %s", resourceFile)
		return nil
	}

}

func init() {
	detectShell()
}
