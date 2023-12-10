package main

import (
	"log"
)

func main() {
	folders, err := scanGitFolders("/home/fiber/dev")
	if err != nil {
		log.Fatal(err)
	}

	commits, err := processRepos(folders, "abdullahalaadine63@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	printTable(commits)
}
