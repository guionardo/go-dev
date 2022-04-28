package shell

import (
	"log"
	"os"
	"path"
	"strings"
)

var (
	shell    string
	rcfile   string
	homePath string
)

func init() {
	var err error
	homePath, err = os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home dir - %v", err)
	}
}

func GetShellInfo() (shellName string, rcFile string, err error) {
	if len(shell) > 0 {
		return shell, rcfile, nil
	}
	shellName, rcFile, err = detectShell()
	if err == nil {
		shell = shellName
		rcfile = rcFile
	}
	return
}

func detectShell() (shellName string, rcFile string, err error) {
	// Detect shell
	shellName = os.Getenv("SHELL")
	if len(shellName) == 0 {
		log.Printf("WARNING: No SHELL environment detected")
		return
	}
	for _, sh := range []string{
		"bash",
		"zsh",
		"ksh",
	} {
		if strings.HasSuffix(shellName, sh) {
			shellName = sh
			rcFile = path.Join(homePath, "."+sh+"rc")
			break
		}
	}
	if len(rcFile) == 0 {
		log.Printf("WARNING: unexpected shell - %s", shell)
		return
	}
	if _, err = os.Stat(rcFile); err != nil {
		log.Printf("WARNING: error getting shell resource file '%s' - %v", rcFile, err)
	}
	return
}
