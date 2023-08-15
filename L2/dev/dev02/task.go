package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func unpackSimbol(s string) string {
	simbol := []rune(s)[len([]rune(s))-2]
	num := []rune(s)[len([]rune(s))-1]
	val, err := strconv.Atoi(string(num))
	if err != nil {
		return ""
	}
	res := strings.Repeat(string(simbol), val)
	return res
}
func unpackEscape(s string) string {
	return string([]rune(s)[1])
}
func Unpacking(str string) (string, error) {
	correctReg, err := regexp.Compile(`[a-zA-Z]`)
	if err != nil {
		return "", err
	}
	isCorrect := correctReg.MatchString(str)
	if !isCorrect {
		return "", errors.New("string is not correct")
	}

	unpackReg, err := regexp.Compile(`((\\\\)|[0-9a-zA-Z])\d`)
	if err != nil {
		return "", err
	}

	unpackEs, err := regexp.Compile(`\\\d`)
	if err != nil {
		return "", err
	}

	res := unpackReg.ReplaceAllStringFunc(str, unpackSimbol)
	res = unpackEs.ReplaceAllStringFunc(res, unpackEscape)
	return res, nil
}
