package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {

	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	res := make([]string, 0)
	for _, e := range arr {
		if !slices.Contains(res, e) {
			res = append(res, e)
		}
	}
	fmt.Println(res)
}
