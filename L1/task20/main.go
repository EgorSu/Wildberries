package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func wordReverse(srt string) (string, error) {
	if srt == "" {
		return "", errors.New("string is empty")
	}
	res := strings.Split(srt, " ")
	i := 0
	j := len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return strings.Join(res, " "), nil
}

func main() {

	var str string
	for _, arg := range os.Args[1:] {
		str += fmt.Sprintf(" %s", arg)
	}
	result, err := wordReverse(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
