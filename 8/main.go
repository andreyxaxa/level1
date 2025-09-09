package main

import "fmt"

func SetBit(number int64, pos uint) int64 {
	return number | (1 << pos)
}

func UnsetBit(number int64, pos uint) int64 {
	return number &^ (1 << pos)
}

func main() {
	var x int64 = 5 // 00000101

	// Установим 1-ый бит в '1'
	y := SetBit(x, 1)
	fmt.Printf("%d == %b -> %d == %b\n", x, x, y, y)

	// Установим 0-ой бит в '0'
	z := UnsetBit(x, 0)
	fmt.Printf("%d == %b -> %d == %b\n", x, x, z, z)
}
