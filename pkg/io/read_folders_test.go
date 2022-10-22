package io

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestPaths_ReadFolders(t *testing.T) {
	root := t.TempDir()
	folders := make([]string, 27)
	for i := 0; i < 27; i++ {
		folders[i] = path.Join(root, fmt.Sprintf("folder%d/sub%d/sub%d", i%3, i%9, i))
	}
	for _, folder := range folders {
		os.MkdirAll(folder, 0777)
	}

	t.Run("default", func(t *testing.T) {
		got, err := ReadFolders(root, 3)
		if err != nil {
			t.Errorf("ReadFolders() error = %v", err)
			return
		}
		if len(got) != 40 {
			t.Errorf("ReadFolders() = %v, want %v", len(got), 40)
		}
		return
	})

}
