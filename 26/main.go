package main

import (
	"fmt"
	"strings"
)

func allCharsUnique(str string) bool {
	m := make(map[rune]bool)

	// регистронезависимая проверка
	for _, c := range strings.ToLower(str) {
		if m[c] {
			return false
		}
		m[c] = true
	}

	return true
}

func main() {
	cases := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, s := range cases {
		fmt.Println(allCharsUnique(s))
	}
}
