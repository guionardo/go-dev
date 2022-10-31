package git

import (
	"regexp"
)

/*
ssh:	git@ssh.dev.azure.com:v3/AMBEV-SA/AMBEV-BIFROST/ms-beesforce-credit-transformation

web: 	https://dev.azure.com/AMBEV-SA/AMBEV-BIFROST/_git/beesforce-metric-api
*/
func isAzureSSH(url string) (success bool, domain string, repo string) {
	azureSSHRE := regexp.MustCompile(`(?m)git@ssh.dev.azure.com:(v[0-9]{1,2})/(.*)/(.*)`)
	matches := azureSSHRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		success, domain, repo = true, "dev.azure.com", matches[2]+"/"+matches[3]
	}
	return
}

/*
ssh: 	git@github.com:guionardo/go-dev.git
https: 	https://github.com/guionardo/go-dev.git
web: 	https://github.com/guionardo/go-dev
*/
func isGitSSH(url string) (success bool, domain string, repo string) {
	gitSSHRE := regexp.MustCompile(`(?m)git@(.*):(.*)/(.*)\.git`)
	matches := gitSSHRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		success, domain, repo = true, matches[1], matches[2]+"/"+matches[3]
	}
	return
}

func IsGitURL(url string) (success bool, domain string, repo string) {
	if success, domain, repo = isAzureSSH(url); success {
		return
	}
	if success, domain, repo = isGitSSH(url); success {
		return
	}
	return
}

/*
https:	https://AMBEV-SA@dev.azure.com/AMBEV-SA/AMBEV-BIFROST/_git/beesforce-metric-api
*/
func isAzureHTTP(url string) (success bool, domain string, repo string) {
	azureHttpRE := regexp.MustCompile(`(?m)https://(.*)@dev.azure.com/(.*)/_git/(.*)`)
	matches := azureHttpRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		success, domain, repo = true, "dev.azure.com", matches[2]+"/"+matches[3]
	}
	return
}

/*
https://gitlab.com/wee-ops/wee-api.git
*/
func isGitHTTP(url string) (success bool, domain string, repo string) {
	gitHttpRE := regexp.MustCompile(`(?m)https://(.*)/(.*)/(.*)\.git`)
	matches := gitHttpRE.FindStringSubmatch(url)
	if len(matches) > 0 {
		success, domain, repo = true, matches[1], matches[2]+"/"+matches[3]
	}
	return
}

func IsHttpURL(url string) (success bool, domain string, repo string) {
	if success, domain, repo = isAzureHTTP(url); success {
		return
	}
	if success, domain, repo = isGitHTTP(url); success {
		return
	}
	return
}

//https://gitlab.com/wee-ops/wee-api
//git@gitlab.com:wee-ops/wee-api.git
//https://gitlab.com/wee-ops/wee-api.git
