package main

import (
	"testing"
)

var tests = map[string][]string{
	"сон":   {"сон", "нос"},
	"сонм":  {},
	"колун": {"колун", "кулон"},
	"кулон": {},
	"нос":   {},
	"ром":   {"ром", "мор"},
	"мор":   {},
	"лиса":  {"лиса", "сила"},
	"сила":  {},
	"перс":  {"перс", "репс", "серп"},
	"репс":  {},
	"серп":  {},
	"аскет": {"аскет", "секта", "сетка", "тесак", "теска"},
	"секта": {},
	"сетка": {},
	"тесак": {},
	"теска": {},
}
var words = []string{"сон", "нос", "сонм", "колун", "кулон", "ром", "мор", "ЛИСА", "сила", "ПЕРс", "РеПС", "серп", "аскет", "секта", "сЕтКА", "тесак", "теска"}

func TestM(t *testing.T) {
	dic := setWords{make(map[string][]string), make(map[string]struct{})}
	for _, word := range words {
		dic.add(word)
	}
	result := dic.get()
	for key, val := range tests {
		resultVal := (*result)[key]
		if len(val) != len(resultVal) {
			t.Error("arrays have different len for key:", key, "\n",
				"test values:", val, "\n",
				"result valeus", resultVal, "\n")
			continue
		}
		for i := 0; i < len(val); i++ {
			if val[i] != resultVal[i] {
				t.Error("arrays are not equal \n",
					"key:", key, "\n",
					"test value:", val[i], "\n",
					"result value:", resultVal[i], "\n")
			}
		}

	}
}
