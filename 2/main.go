package main

import (
	"fmt"
	"os"
	"sync"
)

func SquareEachElement(arr []int) {
	wg := &sync.WaitGroup{}

	for _, el := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, n*n)
		}(el)
	}
	wg.Wait()
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	SquareEachElement(array)
}
