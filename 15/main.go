package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100]) // копируем
}

func main() {
	someFunc()
}
