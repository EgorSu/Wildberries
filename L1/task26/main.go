package main

import (
	"flag"
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	valMap := make(map[rune]struct{})
	str = strings.ToLower(str)
	for _, e := range []rune(str) {
		if _, ok := valMap[e]; ok {
			return false
		}
		valMap[e] = struct{}{}
	}
	return true
}

func main() {

	word := flag.String("word", "abcd", "word for reverse")
	flag.Parse()
	fmt.Printf("word %s is unique: %t\n", *word, isUnique(*word))
}
