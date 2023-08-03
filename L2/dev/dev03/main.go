package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
*/
var ch int

const (
	name = "text.txt"
	path = ""
)

type file struct {
	path  string
	name  string
	lines []line
	col   int
}

func NewFile(path string, name string, col int) (*file, error) {
	f := file{path: path, name: name, lines: make([]*line, 0), col: col}
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
		f.appendLine(line)
	}
	fmt.Println(f)
	return &f, err
}
func (f *file) appendLine(data []byte) {
	str := string(data)
	strArr := strings.Split(str, " ")
	var newLine = &line{key: " ", str: str}
	if len(strArr) > f.col {
		newLine.key = strArr[f.col]
	}
	f.lines = append(f.lines, newLine)
}
func (f *file) String() string {
	var res string
	for _, val := range f.lines {
		res += fmt.Sprintf("%s\n", val.str)
	}
	return res
}

type line struct {
	key string
	str string
}

func sortLines(l []line) {
	if len(l) < 2 {
		return
	}
	j := 0
	p := len(l) - 1
	for i := 0; i < len(l)-1; i++ {
		bigger := l[j].key >= l[p].key
		if l[i].key < l[p].key {
			if i != j && bigger {
				l[i], l[j] = l[j], l[i]
				ch++
			}
			j++
		}
	}
	if j != len(l)-1 {
		l[len(l)-1], l[j] = l[j], l[len(l)-1]
	}
	sortLines(l[:j])
	sortLines(l[j+1:])
}

func main() {

	numCol := flag.Int("k", 0, "number of columns for sort")
	//numberValSort := flag.Bool("n", false, "sort by number value")
	//inverseSort := flag.Bool("r", false, "inverse sort")
	//notShowRepeat := flag.Bool("u", false, "dont show repeat strings")
	flag.Parse()

	newFile, err := NewFile(path, name, *numCol)
	if err != nil {
		log.Fatal(err)
	}
	sortLines(newFile.lines)
	fmt.Print(newFile)
}
