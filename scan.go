package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func scanGitFolders(root, email string) ([]string, error) {
	var gitFolders []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %q: %v", path, err)
		}

		if info.IsDir() && info.Name() == ".git" {
			gitFolder := filepath.Dir(path)
			gitFolders = append(gitFolders, gitFolder)
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return gitFolders, nil
}
