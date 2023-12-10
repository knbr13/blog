package main

import "time"

func daysAgo(t time.Time) int {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	hours := int(today.Sub(t).Hours())
	if hours < 0 {
		return 0
	}
	if hours%24 == 0 {
		return hours / 24
	}
	return hours/24 + 1
}

func getOffset() int {
	now := time.Now()
	switch now.Weekday() {
	case time.Sunday:
		return 0
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	}
	panic("unhandled time")
}
