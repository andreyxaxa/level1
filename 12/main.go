package main

import "fmt"

func getUnique(seq []string) []string {
	m := make(map[string]struct{})
	var res []string

	for _, v := range seq {
		m[v] = struct{}{}
	}

	// для возврата четкого набора уникальных слов. можно и просто мапу 'm' вернуть
	for k := range m {
		res = append(res, k)
	}

	return res
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(getUnique(seq))
}
