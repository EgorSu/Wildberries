package main

import (
	"bufio"
	"flag"
	"fmt"
	"l1/go/pkg/mod/golang.org/x/exp@v0.0.0-20221031165847-c99f073a8326/slices"
	"l1/go/pkg/mod/golang.org/x/exp@v0.0.0-20230626212559-97b1e661b5df/slices"
	"log"
	"os"
	"strings"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной
утилитой (man grep — смотрим описание и основные параметры).
Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/
const (
	name = "text.txt"
	path = ""
)

type file struct {
	path string
	name string
	str  []string
	ids  []int
}

func NewFile(path, name string) (*file, error) {
	f := file{path: path, name: name, str: make([]string, 0)}
	data, err := os.ReadFile(path + name)
	if err != nil {
		return nil, err
	}
	for {
		adv, line, err := bufio.ScanLines(data, true)
		if err != nil {
			break
		}
		if adv == 0 {
			break
		}
		if adv <= len(data) {
			data = data[adv:]
		}
		f.str = append(f.str, string(line))
	}
	return &f, nil
}
func (f *file) String() string {
	result := ""
	for _, id := range f.ids {
		result += f.str[id] + "\n"
	}
}
func (f *file) findPatternId(pattern string, ignore bool) []int {
	idContain := make([]int, 0)
	for id, str := range f.str {
		if ignore {
			str = strings.ToLower(str)
			pattern = strings.ToLower(pattern)
		}
		if strings.Contains(str, pattern) {
			idContain = append(idContain, id)
		}
	}
	return idContain
}
func (f *file) Find(pattern string) {
	f.ids = f.findPatternId(pattern, false)
}
func (f *file) FindNear(pattern string, numLines int) {
	f.ids = f.findPatternId(pattern, false)
	if numLines == 0 {
		fmt.Println(f)
		return
	}
	for _, val := range f.ids {
		if numLines > 0 {
			if val+numLines > len(f.str) {
				val = len(f.str)
			}
		} else {
			if val < -numLines {
				val = 0
			}
		}
	}
}
func (f *file) FindContext(pattern string, numLines int) {
	ids := f.findPatternId(pattern, false)
	result := make([]string, 0)
	if numLines == 0 {
		f.Find(pattern)
		return
	}
	for _, val := range ids {
		valMax := val + numLines
		valMin := val - numLines
		if valMax > len(f.str) {
			valMax = len(f.str) - 1
		}
		if valMin < numLines {
			valMin = 0
		}
		for i := valMin; i <= valMax; i++ {
			if !slices.Contains(result, f.str[i]) {
				result = append(result, f.str[i])
			}
		}
	}
	for r := range result {
		fmt.Println(r)
	}
}
func (f *file) FindCount(pattern string) {
	f.ids = f.findPatternId(pattern, false)
	fmt.Println(len(f.ids))
}
func (f *file) FindIgnore(pattern string) {
	ids := f.findPatternId(pattern, true)
	result := make([]int, 0)
	for i := 0; i < len(f.str); i++ {
		if !slices.Contains(ids, i) {
			result = append(result, i)
		}
	}
	f.ids = result
}
func (f *file) FindReverse(pattern string) {
	ids := f.findPatternId(pattern, false)
	for id, val := range f.str {
		if !slices.Contains(ids, id) {
			fmt.Println(val)
		}
	}
}
func (f *file) FindFixed(pattern string) {
	idContain := make([]int, 0)
	for id, str := range f.str {
		if str == pattern {
			idContain = append(idContain, id)
		}
	}
	f.ids = idContain
}
func (f *file) FindLineNum(pattern string) {
	ids := f.findPatternId(pattern, false)
	for i := range ids {
		fmt.Printf("%v ", i)
	}
}
func main() {
	var pattern string
	after := flag.Int("A", 0, "after")
	before := flag.Int("B", 0, "before")
	context := flag.Int("C", 0, "context")
	count := flag.Bool("c", false, "count")
	ignore := flag.Bool("i", false, "ignore-case")
	invert := flag.Bool("v", false, "invert")
	fixed := flag.Bool("f", false, "fixed")
	printLineNum := flag.Bool("n", false, "line num")
	flag.Parse()
	pattern = flag.Args[1]
	myFile, err := NewFile(path, name)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case *after != 0:
		myFile.FindNear(pattern, *after)
	case *before != 0:
		myFile.FindNear(pattern, *before)
	case *context != 0:
		myFile.FindContext(pattern, *context)
	case *count:
		myFile.FindCount(pattern)
	case *ignore:
		myFile.FindIgnore(pattern)
	case *invert:
		myFile.FindReverse(pattern)
	case *fixed:
		myFile.FindFixed(pattern)
	case *printLineNum:
		myFile.FindLineNum(pattern)
	default:
		myFile.Find(pattern)
	}

}
