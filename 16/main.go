package main

import "fmt"

// Первый вариант
func quickSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, right []int
	for i, v := range arr {
		// пропускаем опорный элемент
		if i == len(arr)/2 {
			continue
		}
		// если меньше опорного -> кидаем в левый слайс, иначе в правый
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// рекурсивно сортируем. меньшие элементы в left, затем pivot, затем большие элементы в right
	return append(append(quickSort1(left), pivot), quickSort1(right)...)
}

// Второй вариант
func quickSort2(arr []int, left, right int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]

	// два индекса двигаются друг к другу, пока не встретятся, или пока не пересекут друг друга
	for l <= r {
		// если значение по левому индексу меньше опорного элемента -> просто двигаем дальше
		for arr[l] < pivot {
			l++
		}
		// если значение по правому индексу больше опорного элемента -> просто двигаем дальше
		for arr[r] > pivot {
			r--
		}
		// пока они не встретились или не пересеклись, свапаем 2 значения, на которых остановились 2 предыдущих for'a.
		// таким образом, слева окажутся элементы меньше опорного, а справа те, что больше.
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	if left < r {
		quickSort2(arr, left, r)
	}
	if l < right {
		quickSort2(arr, l, right)
	}
}

func main() {
	arr := []int{2, 1, 4, 9, 8, 5, 6, 7, 3, 0}

	// Первый вариант
	sorted := quickSort1(arr)
	fmt.Println(sorted)

	// Второй вариант
	quickSort2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
