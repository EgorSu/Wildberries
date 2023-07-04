package main

import (
	"errors"
	"flag"
	"fmt"
)

func letterReverse(str string) (string, error) {
	res := []rune(str) //разбиваем строку на срез символов
	if len(res) == 0 {
		return "", errors.New("string is empty")
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res), nil
}
func main() {
	word := flag.String("word", "главрыба", "word for reverse")
	flag.Parse()
	reverseWord, err := letterReverse(*word)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reverseWord)
}
