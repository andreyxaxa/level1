package main

import "fmt"

func binarySearch(arr []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}

	// ищем середину
	mi := lo + (hi-lo)/2

	// если сразу встали на искомое значение -> вернем его индекс
	if arr[mi] == target {
		return mi
	}
	// если мы меньше, чем искомое значение, значит оно где-то справа, там где всё больше, чем текущее.
	if arr[mi] < target {
		return binarySearch(arr, target, mi+1, hi)
	}
	return binarySearch(arr, target, lo, mi-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(binarySearch(arr, 3, 0, len(arr)-1))  //  2
	fmt.Println(binarySearch(arr, 10, 0, len(arr)-1)) // -1
}
