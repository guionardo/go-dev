package shell

import (
	"strings"
	"testing"
)

func TestGetShellInfo(t *testing.T) {
	tests := []struct {
		name          string
		wantShellName string
		wantRcFile    string
		wantErr       bool
	}{
		{
			name:          "bash",
			wantShellName: "bash",
			wantRcFile:    ".bashrc",
			wantErr:       false,
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShellName, gotRcFile, err := GetShellInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShellInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotShellName != tt.wantShellName {
				t.Errorf("GetShellInfo() gotShellName = %v, want %v", gotShellName, tt.wantShellName)
			}
			if !strings.HasSuffix(gotRcFile, tt.wantRcFile) {
				t.Errorf("GetShellInfo() gotRcFile = %v, want %v", gotRcFile, tt.wantRcFile)
			}
		})
	}
}
