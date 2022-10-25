package io

import (
	"os"
	"path"
)

func FolderReaderReadDir(root string, maxDepth int, acceptFolder func(string) bool, notify func(string)) ([]string, error) {
	return readDirs_(root, 1, maxDepth, acceptFolder, notify)
}

func readDirs_(root string, level int, maxLevel int, acceptFolder func(string) bool, notify func(string)) ([]string, error) {
	dirs := make([]string, 0, 1000)
	entries, err := os.ReadDir(root)
	for _, entry := range entries {
		if entry.IsDir() && acceptFolder(entry.Name()) {
			dir := path.Join(root, entry.Name())
			notify(dir)
			dirs = append(dirs, dir)
			if level < maxLevel {
				subDirs, err := readDirs_(dir, level+1, maxLevel, acceptFolder, notify)
				if err == nil && len(subDirs) > 0 {
					dirs = append(dirs, subDirs...)
				}
			}
		}
	}
	return dirs, err
}
