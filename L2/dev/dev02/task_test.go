package main

import (
	"testing"
)

type returnVal struct {
	str    string
	errStr string
}

var tests = map[string]returnVal{
	"a4bc2d5e":  {"aaaabccddddde", ""},
	"abcd":      {"abcd", ""},
	"45":        {"", "string is not correct"},
	"":          {"", "string is not correct"},
	"qwe\\4\\5": {"qwe45", ""},
	"qwe\\45":   {"qwe44444", ""},
	"qwe\\\\5":  {"qwe\\\\\\\\\\", ""},
}

func TestUnpacking(t *testing.T) {
	for key, val := range tests {
		unpack, err := Unpacking(key)
		if err != nil && val.errStr != err.Error() {
			t.Errorf("For key %s expected error \"%s\", got error \"%s\"", key, val.errStr, err.Error())
		}
		if val.str != unpack {
			t.Errorf("For key %s expected val %s, got %s", key, val.str, unpack)
		}
	}
}
