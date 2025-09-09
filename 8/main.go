package main

import "fmt"

func SetBit(number *int64, pos uint) {
	*number |= 1 << pos
}

func UnsetBit(number *int64, pos uint) {
	*number &^= 1 << pos
}

func main() {
	var x int64 = 5 // 00000101

	// Установим 1-ый бит в '1'
	fmt.Printf("%d == %b -> ", x, x)
	SetBit(&x, 1)
	fmt.Printf("%d == %b\n", x, x)

	// Установим 0-ой бит в '0'
	fmt.Printf("%d == %b -> ", x, x)
	UnsetBit(&x, 0)
	fmt.Printf("%d == %b\n", x, x)
}
