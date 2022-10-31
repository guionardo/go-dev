package git

import (
	"testing"
)

func Test_isGitUrl(t *testing.T) {

	tests := []struct {
		name       string
		url        string
		want       bool
		wantDomain string
		wantRepo   string
	}{
		{
			name:       "gitlab",
			url:        "git@gitlab.com:wee-ops/wee-api.git",
			want:       true,
			wantDomain: "gitlab.com",
			wantRepo:   "wee-ops/wee-api",
		}, {
			name:       "github",
			url:        "git@github.com:guionardo/go-dev.git",
			want:       true,
			wantDomain: "github.com",
			wantRepo:   "guionardo/go-dev",
		}, {
			name:       "azure",
			url:        "git@ssh.dev.azure.com:v3/AMBEV-SA/AMBEV-BIFROST/ms-beesforce-credit-transformation",
			want:       true,
			wantDomain: "dev.azure.com",
			wantRepo:   "AMBEV-SA/AMBEV-BIFROST/ms-beesforce-credit-transformation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			success, domain, repo := IsGitURL(tt.url)
			if success != tt.want {
				t.Errorf("isGitUrl(%s) = %v, want %v", tt.url, success, tt.want)
			}
			if domain != tt.wantDomain {
				t.Errorf("isGitUrl(%s) = %v, want %v", tt.url, domain, tt.wantDomain)
			}
			if repo != tt.wantRepo {
				t.Errorf("isGitUrl(%s) = %v, want %v", tt.url, repo, tt.wantRepo)
			}
		})
	}
}

// 	var str = `https://gitlab.com/wee-ops/wee-api
// git@gitlab.com:wee-ops/wee-api.git
// https://gitlab.com/wee-ops/wee-api.git

// git@github.com:guionardo/go-dev.git
// https://github.com/guionardo/go-dev.git
// https://github.com/guionardo/go-dev

// git@ssh.dev.azure.com:v3/AMBEV-SA/AMBEV-BIFROST/ms-beesforce-credit-transformation
// https://dev.azure.com/AMBEV-SA/AMBEV-BIFROST/_git/beesforce-metric-api
// https://AMBEV-SA@dev.azure.com/AMBEV-SA/AMBEV-BIFROST/_git/beesforce-metric-api`

func Test_isHttpUrl(t *testing.T) {
	tests := []struct {
		name string
		url  string

		wantSuccess bool
		wantDomain  string
		wantRepo    string
	}{
		{
			name:        "gitlab",
			url:         "https://gitlab.com/wee-ops/wee-api.git",
			wantSuccess: true,
			wantDomain:  "gitlab.com",
			wantRepo:    "wee-ops/wee-api",
		},
		{
			name:        "github",
			url:         "https://github.com/guionardo/go-dev.git",
			wantSuccess: true,
			wantDomain:  "github.com",
			wantRepo:    "guionardo/go-dev",
		},
		{
			name:        "azure",
			url:         "https://AMBEV-SA@dev.azure.com/AMBEV-SA/AMBEV-BIFROST/_git/beesforce-metric-api",
			wantSuccess: true,
			wantDomain:  "dev.azure.com",
			wantRepo:    "AMBEV-SA/AMBEV-BIFROST/beesforce-metric-api",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuccess, gotDomain, gotRepo := IsHttpURL(tt.url)
			if gotSuccess != tt.wantSuccess {
				t.Errorf("isHttpUrl() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("isHttpUrl() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("isHttpUrl() gotRepo = %v, want %v", gotRepo, tt.wantRepo)
			}
		})
	}
}
