package main

import (
	"fmt"
	"strings"
	"time"
)

var sixEmptySpaces = strings.Repeat(" ", 6)

func buildHeader(start, end time.Time) string {
	s := strings.Builder{}
	for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 1, 0) {
		s.WriteString(fmt.Sprintf("%-16s", current.Month().String()[:3]))
	}
	return s.String()
}

func printTable(commits map[int]int) {
	fmt.Printf("%s%s%s\n", sixEmptySpaces, sixEmptySpaces, buildHeader(sixMonthsAgo, time.Now()))
	s := strings.Builder{}
	for i := 0; i < 7; i++ {
		s.WriteString(fmt.Sprintf("%-12s", getDay(i)))
		for j := daysAgoFromSixMonths+getOffset()-1; j >= 0; j -= 7 {
			s.WriteString(fmt.Sprintf(" %2d ", commits[j-i]))
		}
		fmt.Println(s.String())
		s.Reset()
	}
}

func getDay(i int) string {
	switch i {
	case 0:
		return "Sun"
	case 1:
		return "Mon"
	case 2:
		return "Tue"
	case 3:
		return "Wed"
	case 4:
		return "Thu"
	case 5:
		return "Fri"
	case 6:
		return "Sat"
	}
	return strings.Repeat(" ", 3)
}
