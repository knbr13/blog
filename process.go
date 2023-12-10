package main

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var sixMonthsAgo time.Time = time.Now().AddDate(0, -6, 0)
var daysAgoFromSixMonths int = daysAgo(sixMonthsAgo)

func processRepos(repos []string, email string) (map[int]int, error) {
	m := map[int]int{}
	var err error
	for _, repo := range repos {
		err = fillCommits(repo, email, m)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func fillCommits(path, email string, commits map[int]int) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	commitIterator, err := repo.Log(&git.LogOptions{Since: &sixMonthsAgo})
	if err != nil {
		return err
	}

	err = commitIterator.ForEach(func(c *object.Commit) error {
		if c.Author.Email != email {
			return nil
		}

		days := daysAgo(c.Author.When)
		commits[days]++
		return nil
	})
	if err != nil {
		return err
	}
	for i := time.Now(); i.After(sixMonthsAgo); i = i.AddDate(0, 0, -1) {
		days := daysAgo(i)
		if _, ok := commits[days]; !ok {
			commits[days] = 0
		}
	}
	for i := daysAgoFromSixMonths; i < daysAgoFromSixMonths+getOffset()-1; i++ {
		commits[i] = 0
	}
	return nil
}
