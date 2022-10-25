package update

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/google/go-github/v45/github"
	"github.com/guionardo/go-dev/pkg/consts"
	"github.com/guionardo/go-dev/pkg/logger"
	"github.com/guionardo/go-dev/pkg/tools"
	"github.com/schollz/progressbar/v3"
)

func RunGitUpdate() error {
	version, url, err := getGithubVersion()
	if err != nil {
		logger.Error("Error getting version from github: %s", err)
		return nil
	}
	logger.Debug("GitHub update - local version: %s - remote version: %s", consts.Version, version)
	vLocal, err := tools.VersionParse(consts.Version)
	if err != nil {
		vLocal = tools.NewVersion(0, 0, 0)
	}
	vRemote, err := tools.VersionParse(version)
	if err != nil {
		logger.Error("Error parsing version from github: %s", err)
		return nil
	}
	if !vRemote.IsNewerThan(&vLocal) {
		logger.Debug("No update available")
		return nil
	}
	logger.Debug("Update available: %s", url)

	filename, err := dowloadFile(url)
	if err != nil {
		logger.Error("Error downloading file: %s", err)
		return nil
	}
	logger.Debug("Downloaded file: %s", filename)
	defer os.Remove(filename)

	destiny, err := os.MkdirTemp("", consts.AppName)
	if err != nil {
		logger.Error("Error creating temp folder: %s", err)
		return nil
	}
	defer os.RemoveAll(destiny)
	files, err := ExtractTarGz(filename, destiny)
	if err != nil {
		logger.Error("Error extracting file: %s", err)
		return nil
	}
	logger.Debug("Extracted files: %v", files)

	binary_file := ""
	for _, file := range files {
		if strings.HasSuffix(file, consts.AppName) {
			logger.Debug("Found executable: %s", file)
			binary_file = file
			break
		}
	}
	if binary_file == "" {
		logger.Error("Error finding executable file")
		return nil
	}
	var downloadedVersion tools.Version
	if downloadedVersion, err = testVersion(binary_file); err != nil {
		logger.Error("Error testing version: %s", err)
		return nil
	}
	if !downloadedVersion.IsNewerThan(&vLocal) {
		logger.Error("Downloaded version %s is not newer than local version %s", downloadedVersion, vLocal)
		return nil
	}

	
	return nil
}

func testVersion(filename string) (version tools.Version, err error) {
	logger.Debug("Testing version: %s", filename)
	err = os.Chmod(filename, 0755)
	if err != nil {
		return
	}
	cmd := exec.Command(filename, "--version")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return version, err
	}
	if err := cmd.Start(); err != nil {
		return version, err
	}
	rawVersion, err := ioutil.ReadAll(stdout)
	if err != nil {
		return version, err
	}
	version, err = tools.VersionParse(string(rawVersion))
	if err != nil {
		return
	}

	logger.Debug("Version: %s", version)
	return
}

// TODO: Implementar update via github
func getGithubVersion() (version string, releaseUrl string, err error) {
	client := github.NewClient(http.DefaultClient)
	release, response, err := client.Repositories.GetLatestRelease(context.Background(), "guionardo", consts.AppName)
	if err != nil {
		return "", "", fmt.Errorf("Failed to get latest release: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("Failed to get latest release: %v", response.StatusCode)
	}
	suffix := runtime.GOOS + "-" + runtime.GOARCH // "-linux-amd64"
	for _, asset := range release.Assets {
		if strings.HasSuffix(*asset.Name, suffix+".tar.gz") {
			return *release.TagName, *asset.BrowserDownloadURL, nil
		}
	}

	return release.GetTagName(), "", nil
}

func dowloadFile(url string) (filename string, err error) {
	tmpfile, err := os.CreateTemp("", "go-dev")
	if err != nil {
		return "", err
	}
	filename = tmpfile.Name()
	defer tmpfile.Close()
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading update",
	)
	written, err := io.Copy(io.MultiWriter(tmpfile, bar), resp.Body)
	if err != nil {
		os.Remove(tmpfile.Name())
		return "", err
	}
	if written != resp.ContentLength {
		os.Remove(tmpfile.Name())
		return "", fmt.Errorf("Failed to download file: size written differs from expected")
	}
	return filename, nil
}
