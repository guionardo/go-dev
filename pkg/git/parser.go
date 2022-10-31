package git

import (
	"fmt"
	"regexp"
)

const (
	AzureFormat = "https://%s/_git/%s"
	GitFormat   = "https://%s/%s"
)

type GitURL struct {
	Success bool
	Domain  string
	Repo    string
	Format  string
}

func (g GitURL) GetURL() string {
	return fmt.Sprintf(g.Format, g.Domain, g.Repo)
}

func isAzureSSH(url string) GitURL {
	azureSSHRE := regexp.MustCompile(`(?m)git@ssh.dev.azure.com:(v[0-9]{1,2})/(.*)/(.*)`)
	matches := azureSSHRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		return GitURL{
			Success: true,
			Domain:  "dev.azure.com" + "/" + matches[2],
			Repo:    matches[3],
			Format:  AzureFormat,
		}
	}
	return GitURL{}
}

/*
ssh: 	git@github.com:guionardo/go-dev.git
https: 	https://github.com/guionardo/go-dev.git
web: 	https://github.com/guionardo/go-dev
*/
func isGitSSH(url string) GitURL {
	gitSSHRE := regexp.MustCompile(`(?m)git@(.*):(.*)/(.*)\.git`)
	matches := gitSSHRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		return GitURL{
			Success: true,
			Domain:  matches[1],
			Repo:    matches[2] + "/" + matches[3],
			Format:  GitFormat,
		}

	}
	return GitURL{}
}

func isAzureHTTP(url string) (gu GitURL) {
	azureHttpRE := regexp.MustCompile(`(?m)https://(.*)@dev.azure.com/(.*)/_git/(.*)`)
	matches := azureHttpRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		gu = GitURL{
			Success: true,
			Domain:  "dev.azure.com" + "/" + matches[2],
			Repo:    matches[3],
			Format:  AzureFormat,
		}
	}
	return
}

/*
https://gitlab.com/wee-ops/wee-api.git
*/
func isGitHTTP(url string) GitURL {
	gitHttpRE := regexp.MustCompile(`(?m)https://(.*)/(.*)/(.*)\.git`)
	matches := gitHttpRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		return GitURL{
			Success: true,
			Domain:  matches[1],
			Repo:    matches[2] + "/" + matches[3],
			Format:  GitFormat,
		}
	}
	return GitURL{}
}

//https://gitlab.com/wee-ops/wee-api
//git@gitlab.com:wee-ops/wee-api.git
//https://gitlab.com/wee-ops/wee-api.git

func ParseGitURL(url string) GitURL {
	for _, parseFunc := range []func(string) GitURL{isAzureHTTP, isAzureSSH, isGitHTTP, isGitSSH} {
		if gitURL := parseFunc(url); gitURL.Success {
			return gitURL
		}
	}
	return GitURL{}
}
