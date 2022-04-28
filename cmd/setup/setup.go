package setup

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/guionardo/go-dev/cmd/utils"
)

type Setup struct {
	SetupDir    string
	SetupFile   string
	AliasFile   string
	UserHomeDir string
	shell       string
	rcfile      string
}

var (
	//go:embed shell_dev.sh
	shellDev string
)

func CreateSetup(setupDir string) (*Setup, error) {
	if err := utils.CreatePath(setupDir); err != nil {
		return nil, err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("ERROR: Failed to get user home dir - %v", err)
	}

	setup := &Setup{
		SetupDir:    setupDir,
		SetupFile:   path.Join(setupDir, "go_dev.json"),
		UserHomeDir: home,
	}
	setup.detectShell()

	return setup, nil
}

func (setup *Setup) GetAliasFile() (string, error) {
	aliasFile := path.Join(setup.SetupDir, "go_dev.sh")
	if !utils.FileExists(aliasFile) {
		err := createAliasFile(aliasFile)
		if err != nil {
			return "", err
		}
	}
	return aliasFile, nil
}

func createAliasFile(aliasFile string) error {
	var err error
	if file, err := os.Create(aliasFile); err == nil {
		_, err = file.WriteString(shellDev)
	}
	return err
}

func (setup *Setup) detectShell() {
	// Detect shell
	shell = os.Getenv("SHELL")

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
			setup.shell = sh
			setup.rcfile = path.Join(setup.UserHomeDir, "."+sh+"rc")
			break
		}
	}
	if len(setup.rcfile) == 0 {
		log.Printf("WARNING: unexpected shell - %s", shell)
		return
	}
	if _, err := os.Stat(rcfile); err != nil {
		log.Printf("WARNING: error getting shell resource file '%s' - %v", setup.rcfile, err)
		return
	}
}

func (setup *Setup) EnableAlias() error {
	aliasFile, err := setup.GetAliasFile()
	if err != nil {
		return err
	}

	body, err := os.ReadFile(setup.rcfile)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("source %s", aliasFile)
	if strings.Contains(string(body), cmd) {
		return nil
	}

	file, err := os.OpenFile(setup.rcfile, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("\n%s\n", cmd))
	return err
}

func (setup *Setup) DisableAlias() error {
	aliasFile, err := setup.GetAliasFile()
	if err != nil {
		return err
	}

	body, err := os.ReadFile(setup.rcfile)
	if err != nil {
		return err
	}

	cmd := fmt.Sprintf("source %s", aliasFile)
	strBody := string(body)
	if !strings.Contains(strBody, cmd) {
		return nil
	}
	strBody = strings.ReplaceAll(strBody, cmd, "")

	for strings.Contains("\n\n", strBody) {
		strBody = strings.ReplaceAll(strBody, "\n\n", "\n")
	}
	body = []byte(strBody)
	return os.WriteFile(setup.rcfile, body, 0666)
}
