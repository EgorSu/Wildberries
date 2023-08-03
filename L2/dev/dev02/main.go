package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func unpackSimbol(s string) string {
	simbol := []rune(s)[0]
	num := []rune(s)[1]
	val, err := strconv.Atoi(string(num))
	if err != nil {
		return ""
	}
	res := strings.Repeat(string(simbol), val)
	return res
}

func Unpacking(str string) (string, error) {
	correctReg, err := regexp.Compile(`[a-zA-Z]`)
	if err != nil {
		return "", err
	}
	unpackReg, err := regexp.Compile(`([a-zA-Z]\d)`)
	if err != nil {
		return "", err
	}
	isCorrect := correctReg.MatchString(str)
	if !isCorrect {
		return "", errors.New("string is not correct")
	}
	res := unpackReg.ReplaceAllStringFunc(str, unpackSimbol)
	//fmt.Println(string(res))
	return res, nil
}

func main() {
	str1 := "a4bc2d5e"
	Unpacking(str1)
}
