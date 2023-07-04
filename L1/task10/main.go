package main

import (
	"fmt"
)

func main() {
	temper := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	hashTab := make(map[int][]float32)
	for _, val := range temper {
		key := int(val/10) * 10
		hashTab[key] = append(hashTab[key], val)
	}
	fmt.Println(hashTab)
}
