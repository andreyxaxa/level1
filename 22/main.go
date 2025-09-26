package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1234567890091235869192939102123", 10)
	b.SetString("12312312312331255471", 10)

	fmt.Println(new(big.Int).Mul(a, b)) // *
	fmt.Println(new(big.Int).Div(a, b)) // /
	fmt.Println(new(big.Int).Add(a, b)) // +
	fmt.Println(new(big.Int).Sub(a, b)) // -
}
