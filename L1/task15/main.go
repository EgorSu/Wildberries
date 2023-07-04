package main

import (
	"math/rand"
)

/*
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
*/
var justString string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(lenght int) string {
	n := len(letters)
	str := make([]rune, lenght)
	for i := 0; i < lenght; i++ {
		str[i] = letters[rand.Intn(n)]
	}
	return string(str)
}

func someFunc() {

	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
