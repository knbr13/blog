// package main

// import (
// 	"fmt"
// 	"time"
// )

// var visualization = [][]interface{}{}

// /*
// 	[
// 		[[1, 4, 2 , 7, 3],2,[3, 32, 5 ,2 ,7]],
//         [[1, 4, 2 , 7, 3],2,[3, 32], 5, [2 ,7],
// 	]
// */

// func main() {
// 	arr := []int{5, 2, 8, 3}
// 	visualization = append(visualization, append([]interface{}{}, arr))
// 	quickSort(&arr, time.Microsecond)
// 	for _, v := range arr {
// 		fmt.Println(v)
// 	}
// 	fmt.Println(visualization)
// }

// func quickSort(arr *[]int, delay time.Duration) {
// 	m := 1
// 	n := 1
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
// 	if len(visualization) <= m {
// 		visualization = append(visualization, []interface{}{})
// 	}
// 	visualization[m] = append(visualization[m], x, pivot)
// 	m++
// 	quickSort(&x, delay)
// 	if len(visualization) <= n {
// 		visualization = append(visualization, []interface{}{})
// 	}
// 	visualization[n] = append(visualization[n], y)
// 	n++
// 	quickSort(&y, delay)
// }

// func printAndDelay(x interface{}) {
// 	fmt.Println(x)
// 	time.Sleep(time.Second)
// }

package main

import (
	"fmt"
	"time"
)

type Element struct {
	Value int
	Color string
}

var visualization = [][]Element{}

func main() {
	arr := []int{5, 2, 8, 3}
	visualization = append(visualization, makeVisualizationArray(arr, ""))
	quickSort(&arr, time.Millisecond*500)
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
		visualization = append(visualization, []Element{})
	}
	visualization[m] = append(visualization[m], makeVisualizationArray(x, fmt.Sprintf("pivot(%d)", pivot))...)
	m++
	quickSort(&x, delay)
	if len(visualization) <= n {
		visualization = append(visualization, []Element{})
	}
	visualization[n] = append(visualization[n], makeVisualizationArray(y, "")...)
	n++
	quickSort(&y, delay)
}

func makeVisualizationArray(arr []int, color string) []Element {
	elements := make([]Element, len(arr))
	for i, v := range arr {
		elements[i] = Element{Value: v, Color: color}
	}
	return elements
}
