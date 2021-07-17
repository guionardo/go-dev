package utils

import "os"

func PathExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}