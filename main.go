// package main

// import (
// 	"fmt"
// 	"time"
// )

// var visualization = [][][]int{}

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
)

var visualization = [][][]int{}

func main() {
	arr := []int{5, 2, 8, 3, 9, 8, 1, 3, 7, 4}
	fmt.Println("Original Array:")
	printArray(arr)
	fmt.Println("----------------------------")
	quickSort(&arr, time.Millisecond*500)
	fmt.Println("----------------------------")
	fmt.Println("Sorted Array:")
	printArray(arr)

	fmt.Println("----------------------------")
	fmt.Println("Visualization:")
	printVisualization()
}

func quickSort(arr *[]int, delay time.Duration) {
	if len(*arr) <= 1 {
		return
	}
	pivot := (*arr)[len(*arr)-1]
	i, j := -1, 0
	for ; j < len(*arr); j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[len(*arr)-1] = (*arr)[len(*arr)-1], (*arr)[i+1]
	x := (*arr)[:i+1]
	y := (*arr)[i+2:]

	// Capture the current state of the array
	visualization = append(visualization, [][]int{x, {pivot}, y})

	quickSort(&x, delay)
	quickSort(&y, delay)
}

func printArray(arr []int) {
	for i, v := range arr {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func printVisualization() {
	for _, step := range visualization {
		for _, line := range step {
			printArray(line)
		}
		fmt.Println("----------------------------")
		time.Sleep(time.Second) // Add a delay between steps
	}
}
