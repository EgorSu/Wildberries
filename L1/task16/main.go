package main

import (
	"fmt"
	"math/rand"
)

func quikSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	j := 0
	p := arr[len(arr)-1]
	for i := 0; i < len(arr)-1; i++ {
		bigger := arr[j] >= p
		if arr[i] < p {
			if i != j && bigger {
				arr[i], arr[j] = arr[j], arr[i]
			}
			j++
		}
	}
	if j != len(arr)-1 {
		arr[len(arr)-1], arr[j] = arr[j], arr[len(arr)-1]
	}
	quikSort(arr[:j])
	quikSort(arr[j+1:])

}
func main() {
	n := rand.Intn(100)
	array := make([]int, n)
	for i := range array {
		array[i] = rand.Intn(100)
	}
	fmt.Println(array)
	quikSort(array)
	fmt.Println(array)
	//проверка, выведет wrong если массим не отсортирован
	max := 0
	for _, e := range array {
		if max > e {
			fmt.Println("wrong", max, e)
			break
		}
		max = e
	}
}
