package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

var ignoreDirectories = []string{
	"node_modules",
	"venv",
	"__pycache__",
	".venv",
	".pytest_cache",
	"target",
	"bin",
	"build",
	".gradle",
	"vendor",
	"gemfile.lock",
	"obj",
	"packages",
	".build",
	"packages",
	".dart_tool",
	"pubspec.lock",
}

func scanGitFolders(root string) ([]string, error) {
	var gitFolders []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %q: %v", path, err)
		}

		if slices.Contains(ignoreDirectories, strings.ToLower(info.Name())) {
			return filepath.SkipDir
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
