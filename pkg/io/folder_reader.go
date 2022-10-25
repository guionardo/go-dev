package io

type FolderReader func(root string, maxDepth int, acceptFolder func(string) bool, notify func(string)) ([]string, error)
