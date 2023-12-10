package main

import "time"

func daysAgo(t time.Time) int {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	hours := int(t.Sub(today).Hours())
	if hours < 0 {
		return 0
	}
	if hours%24 == 0 {
		return hours / 24
	}
	return hours/24 + 1
}
