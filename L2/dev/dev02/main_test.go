package main

import "testing"

var tests = map[string]string{
	"a4bc2d5e":  "aaaabccddddde",
	"abcd":      "abcd",
	"45":        "",
	"":          "",
	"qwe\\4\\5": "qwe45",
	"qwe\\45":   "qwe44444",
	"qwe\\\\5":  "qwe\\\\\\\\\\",
}

func TestUnpacking(t *testing.T) {
	for key, val := range tests {
		unpack, err := Unpacking(key)
		if val != unpack || err != nil {
			t.Error(
				"For", key,
				"expected", val,
				"got", unpack,
			)
		}
	}
}
