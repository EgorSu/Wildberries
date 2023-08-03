package main

import (
	"sort"
	"strings"
)

type setWords struct {
	runesWords map[string][]string
	isOne      map[string]struct{}
}

func (sw *setWords) add(word string) {
	word = strings.ToLower(word)
	runeWord := []rune(word)
	sort.Slice(runeWord, func(i, j int) bool {
		if int(runeWord[i]) != int(runeWord[j]) {
			return int(runeWord[i]) < int(runeWord[j])
		}
		return true
	})
	strSortWord := string(runeWord)
	if _, ok := sw.runesWords[strSortWord]; !ok {
		sw.isOne[strSortWord] = struct{}{}
	} else {
		delete(sw.isOne, strSortWord)
	}
	sw.runesWords[strSortWord] = append(sw.runesWords[strSortWord], word)
}
func (sw *setWords) get() *map[string][]string {
	result := make(map[string][]string)
	for key, val := range sw.runesWords {
		if _, ok := sw.isOne[key]; !ok {
			result[val[0]] = val
		}
	}
	return &result
}

func main() {

}
