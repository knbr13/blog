package main

import "time"

func main() {
	arr := []int{5, 2, 8, 3, 9, 8, 1, 3, 7, 4}
	quickSort(arr, time.Microsecond)
	for _, v := range arr {
		println(v)
	}
}

func quickSort(arr []int, delay time.Duration) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[len(arr)-1]
	i, j := -1, 0
	for ; j < len(arr); j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	go quickSort(arr[:i+1], delay)
	go quickSort(arr[i+2:], delay)
}
