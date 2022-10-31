package git

import (
	"testing"
)

func TestParseGitURL(t *testing.T) {

	tests := []struct {
		name    string
		url     string
		want    GitURL
		wantUrl string
	}{
		{
			name: "gitlab ssh",
			url:  "git@gitlab.com:wee-ops/wee-api.git",
			want: GitURL{
				Success: true,
				Domain:  "gitlab.com",
				Repo:    "wee-ops/wee-api",
			},
			wantUrl: "https://gitlab.com/wee-ops/wee-api",
		},
		{
			name: "github ssh",
			url:  "git@github.com:guionardo/go-dev.git",
			want: GitURL{
				Success: true,
				Domain:  "github.com",
				Repo:    "guionardo/go-dev",
			},
			wantUrl: "https://github.com/guionardo/go-dev",
		},
		{
			name: "azure ssh",
			url:  "git@ssh.dev.azure.com:v3/CUSTOMER-SA/CUSTOMER-NS/ms-credit-api",
			want: GitURL{
				Success: true,
				Domain:  "dev.azure.com/CUSTOMER-SA/CUSTOMER-NS",
				Repo:    "ms-credit-api",
			},
			wantUrl: "https://dev.azure.com/CUSTOMER-SA/CUSTOMER-NS/_git/ms-credit-api",
		},
		{
			name: "gitlab http",
			url:  "https://gitlab.com/wee-ops/wee-api.git",
			want: GitURL{
				Success: true,
				Domain:  "gitlab.com",
				Repo:    "wee-ops/wee-api",
			},
			wantUrl: "https://gitlab.com/wee-ops/wee-api",
		},
		{
			name: "github http",
			url:  "https://github.com/guionardo/go-dev.git",
			want: GitURL{
				Success: true,
				Domain:  "github.com",
				Repo:    "guionardo/go-dev",
			},
			wantUrl: "https://github.com/guionardo/go-dev",
		},
		{
			name: "azure http",
			url:  "https://CUSTOMER-SA@dev.azure.com/CUSTOMER-SA/CUSTOMER-NS/_git/metric-api",
			want: GitURL{
				Success: true,
				Domain:  "dev.azure.com/CUSTOMER-SA/CUSTOMER-NS",
				Repo:    "metric-api",
			},
			wantUrl: "https://dev.azure.com/CUSTOMER-SA/CUSTOMER-NS/_git/metric-api",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseGitURL(tt.url)
			if got.Success != tt.want.Success {
				t.Errorf("ParseGitURL() success expected")
			}
			if got.Domain != tt.want.Domain || got.Repo != tt.want.Repo {
				t.Errorf("ParseGitURL() = %v, want %v", got, tt.want)
			}
			if got.GetURL() != tt.wantUrl {
				t.Errorf("ParseGitURL() = %s, want %v", got.GetURL(), tt.wantUrl)
			}
		})
	}
}
