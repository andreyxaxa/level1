package main

import "fmt"

func XORSwap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b
}

func main() {
	a := 5
	b := 200
	fmt.Printf("before swap:\na = %d\nb = %d\n", a, b)

	a, b = XORSwap(a, b)
	fmt.Printf("after swap:\na = %d\nb = %d", a, b)
}
