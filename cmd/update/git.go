package update

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v45/github"
	"github.com/guionardo/go-dev/pkg/consts"
)

// TODO: Implementar update via github
func getGithubVersion() (string, error) {
	client := github.NewClient(http.DefaultClient)
	release, response, err := client.Repositories.GetLatestRelease(context.Background(), "guionardo", consts.AppName)
	if err != nil {
		log.Printf("Failed to get latest release: %v", err)
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		log.Printf("Failed to get latest release: %v", response.StatusCode)
		return "", err
	}
	return release.GetTagName(), nil
}
