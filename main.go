// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	arr := []int{5, 2, 8, 3, 9, 8, 1, 3, 7, 4}
// 	quickSort(&arr, time.Microsecond)
// 	for _, v := range arr {
// 		fmt.Println(v)
// 	}
// }

// func quickSort(arr *[]int, delay time.Duration) {
// 	if len(*arr) <= 1 {
// 		return
// 	}
// 	pivot := (*arr)[len((*arr))-1]
// 	i, j := -1, 0
// 	for ; j < len((*arr)); j++ {
// 		if (*arr)[j] < pivot {
// 			i++
// 			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
// 		}
// 	}
// 	(*arr)[i+1], (*arr)[len((*arr))-1] = (*arr)[len((*arr))-1], (*arr)[i+1]
// 	x := (*arr)[:i+1]
// 	y := (*arr)[i+2:]
// 	quickSort(&x, delay)
// 	quickSort(&y, delay)
// }

package main

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

func main() {
	arr := []int{5, 2, 8, 3, 9, 8, 1, 3, 7, 4}
	quickSort(&arr, time.Millisecond*500) // Increase or decrease delay as needed
	fmt.Println("Sorted Array:")
	printArrayWithColors(arr)
}

func quickSort(arr *[]int, delay time.Duration) {
	if len(*arr) <= 1 {
		return
	}
	pivot := (*arr)[len((*arr))-1]
	i, j := -1, 0
	for ; j < len((*arr)); j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[len((*arr))-1] = (*arr)[len((*arr))-1], (*arr)[i+1]
	x := (*arr)[:i+1]
	y := (*arr)[i+2:]

	// Print the current state of the array with colors
	printArrayWithColors(*arr)

	time.Sleep(delay)

	quickSort(&x, delay)
	quickSort(&y, delay)
}

func printArrayWithColors(arr []int) {
	// Create color styles
	cyan := color.FgCyan.Render
	yellow := color.FgYellow.Render
	// magenta := color.FgMagenta.Render

	// Print the array with colors
	for i, v := range arr {
		if i != 0 {
			fmt.Print(" ")
		}
		if i == 0 || i == len(arr)-1 {
			fmt.Print(cyan("[", v, "]"))
		} else {
			fmt.Print(yellow("[", v, "]"))
		}
	}
	fmt.Println()
}
