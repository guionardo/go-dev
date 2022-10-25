package update

import (
	"strings"
	"testing"
)

func Test_getGithubVersion(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "getGithubVersion",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := getGithubVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("getGithubVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.HasPrefix(got, "v") {
				t.Errorf("getGithubVersion() = %v, want start with 'v'", got)
			}
		})
	}
}
