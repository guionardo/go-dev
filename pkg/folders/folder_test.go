package folders

import "testing"

func TestFolder_Match(t *testing.T) {
	tests := []struct {
		name string
		path string
		args []string
		want bool
	}{
		{name: "empty", path: "", args: []string{}, want: false},
		{name: "github", path: "/home/username/dev/github.com", args: []string{"github"}, want: true},
		{name: "github provider", path: "/home/username/dev/github.com/provider", args: []string{"github", "provider"}, want: true},
		{name: "github tester", path: "/home/username/dev/github.com/tester", args: []string{"github", "provider"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Folder{Path: tt.path}
			if got := f.Match(tt.args); got != tt.want {
				t.Errorf("Folder.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
