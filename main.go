package main

import (
	"fmt"
	"time"
)

var visualization = [][][]int{}

func main() {
	arr := []int{5, 2, 8, 3, 9, 8, 1, 3, 7, 4}
	visualization = append(visualization, append([][]int{}, arr))
	quickSort(&arr, time.Microsecond)
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println(visualization)
}

func quickSort(arr *[]int, delay time.Duration) {
	m := 1
	n := 1
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
	if len(visualization) <= m {
		visualization = append(visualization, append([][]int{}, x))
	}
	visualization[m] = append(visualization[m], x, append([]int{}, pivot))
	m++
	quickSort(&x, delay)
	if len(visualization) <= n {
		visualization = append(visualization, append([][]int{}, y))
	}
	visualization[n] = append(visualization[n], y, append([]int{}, pivot))
	n++
	quickSort(&y, delay)
}
