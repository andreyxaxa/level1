package main

import "fmt"

func findIntersection(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	var res []int

	// записываем все значения из 'arr1' как ключи в мапу, проставляем 'true'
	for _, v := range arr1 {
		m[v] = true
	}

	// проходимся по значениям 'arr2', если в мапе по этим ключам уже стоит 'true' - пересечение
	for _, v := range arr2 {
		// если там уже 'true'
		if m[v] {
			res = append(res, v)
			m[v] = false
		}
	}

	return res
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}
	fmt.Println(findIntersection(arr1, arr2))

	arr3 := []int{1, 2, 3, 10, 20, 30, 40, 90, 100}
	arr4 := []int{2, 3, 4, 90, 100, 9, 11, 22, 30}
	fmt.Println(findIntersection(arr3, arr4))
}
