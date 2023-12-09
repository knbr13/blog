package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func scanGitRepositories(root string) ([]string, error) {
	var result []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking %q: %v", path, err)
		}

		if info.IsDir() && info.Name() == ".git" {
			parent := filepath.Dir(path)
			result = append(result, parent)
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking %q: %v", root, err)
	}
	return result, nil
}
