package setup

import (
	"os"
	"testing"
)

func Test_detectShell(t *testing.T) {
	tests := []struct {
		name        string
		shell       string
		expectError bool
	}{
		{name: "current", shell: os.Getenv("SHELL"), expectError: false},
		{name: "fail", shell: "not_found", expectError: true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("SHELL", tt.shell)
			rcfile = ""
			detectShell()
			if tt.expectError && len(rcfile) > 0 {
				t.Errorf("Expected error %s", tt.name)
			}

		})
	}
}
