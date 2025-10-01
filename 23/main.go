package main

import "fmt"

func removeFromSlice(s []*string, idx int) []*string {
	// Сдвинем элементы
	copy(s[idx:], s[idx+1:])

	// GC увидит, что ссылок нет
	s[len(s)-1] = nil

	// Уменьшаем длину на 1
	return s[:len(s)-1]
}

func main() {
	s1, s2, s3, s4, s5, s6, s7 := "one", "two", "three", "four", "five", "six", "seven"
	sl := []*string{&s1, &s2, &s3, &s4, &s5, &s6, &s7}

	sl = removeFromSlice(sl, 2)

	for _, v := range sl {
		fmt.Printf("%s ", *v)
	}
}
