package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
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
	max := getMaxValue(commits)
	for i := 0; i < 7; i++ {
		s.WriteString(fmt.Sprintf("%-12s", getDay(i)))
		for j := daysAgoFromSixMonths + getOffset() - 1; j >= 0; j -= 7 {
			s.WriteString(printCell(commits[j-i], max))
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

func printCell(n, maxVal int) string {
	var printingColor color.Style
	if n == 0 {
		printingColor = color.New(color.FgBlack, color.BgDarkGray)
		return printingColor.Sprint("  0 ")
	}
	if n <= maxVal/4 {
		printingColor = color.New(color.FgBlack, color.BgLightGreen)
		return printingColor.Sprintf(" %2d ", n)
	} else if n <= maxVal/2 {
		printingColor = color.New(color.FgBlack, color.BgGreen)
		return printingColor.Sprintf(" %2d ", n)
	}
	printingColor = color.New(color.FgLightGreen, color.BgBlack)
	return printingColor.Sprintf(" %2d ", n)
}
