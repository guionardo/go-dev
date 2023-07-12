package actions

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/guionardo/go-dev/pkg/git"
	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/urfave/cli/v2"
)

func UrlAction(c *cli.Context) error {
	justPrint := c.Bool("just-print")

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	// Check folder .git
	gitFolder := path.Join(cwd, ".git")
	if _, err := os.Stat(gitFolder); os.IsNotExist(err) {
		return fmt.Errorf("folder %s is not a git repository", cwd)
	}

	// Run git config --get remote.origin.url
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return fmt.Errorf("current repository has no remote origin - %v", err)
	}
	output := strings.ReplaceAll(strings.SplitN(string(out), "\n", 1)[0], "\n", "")
	url, err := getHttpUrl(output)
	if err != nil {
		return err
	}

	if justPrint {
		_, err = fmt.Println(url)
	} else {
		logger.Info("Opening %s", url)
		err = openInBrowser(url)

	}
	return err
}

func getHttpUrl(url string) (string, error) {
	gu, err := git.Parse(url)
	if !gu.Success || err != nil {
		return "", fmt.Errorf("invalid git url: %s", url)
	}
	return gu.GetURL(), nil
}

func openInBrowser(url string) (err error) {
	switch runtime.GOOS {
	case "linux":
		command := "xdg-open"
		if len(os.Getenv("WSL_DISTRO_NAME")) > 0 {
			command = "sensible-browser"
		}
		err = exec.Command(command, url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return
}
