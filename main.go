package main

import (
	"fmt"
	"time"
)

var visualization = [][]interface{}{}

/*
	[
		[[1, 4, 2 , 7, 3],2,[3, 32, 5 ,2 ,7]],
        [[1, 4, 2 , 7, 3],2,[3, 32], 5, [2 ,7],
	]
*/

func main() {
	arr := []int{5, 2, 8, 3}
	visualization = append(visualization, append([]interface{}{}, arr))
	quickSort(&arr)
	fmt.Println("final results: ", visualization)
}

func quickSort(arr *[]int) {
	m := 1
	n := 1
	if len(*arr) <= 1 {
		return
	}
	pivot := (*arr)[len((*arr))-1]
	fmt.Println("pivot: ", pivot)
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
	fmt.Println("y: ", y)
	if len(visualization) <= m && len(x) > 0 {
		visualization = append(visualization, []interface{}{})
		visualization[m] = append(visualization[m], x, pivot)
		m++
	}
	printAndDelay(visualization)
	quickSort(&x)
	if len(visualization) <= n && len(y) > 0 {
		visualization = append(visualization, []interface{}{})
		visualization[n] = append(visualization[n], y)
		n++
	}
	printAndDelay(visualization)
	quickSort(&y)
}

func printAndDelay(x interface{}) {
	fmt.Println(x)
	time.Sleep(time.Second)
}
