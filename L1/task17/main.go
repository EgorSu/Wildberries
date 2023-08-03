package main

import (
	"fmt"
)

func search(arr []int, slice, val int) int {
	if len(arr) == 0 {
		return -1
	}
	avVal := arr[len(arr)/2]
	switch {
	case avVal > val:
		return search(arr[:len(arr)/2], slice, val)
	case avVal < val:
		return search(arr[len(arr)/2+1:], len(arr)/2+1+slice, val)
	}
	return slice
}

func main() {
	a := []int{1, 3, 6, 9, 12, 43, 90} 
	fmt.Println(search(a, 0, 1))
	fmt.Println(search(a, 0, 91))
	fmt.Println(search(a, 0, 7))
	fmt.Println(search(a, 0, 12))
	fmt.Println(search(a, 0, 0))
	fmt.Println(search([]int{2}, 0, 1))
}
