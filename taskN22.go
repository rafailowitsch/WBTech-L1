package main

import (
	"fmt"
	"math/big"
)

// функция taskN22 демонстрирует операции с большими числами
func taskN22() {
	a_base := big.NewInt(2)
	a_exponent := uint(20)
	a := new(big.Int).Exp(a_base, big.NewInt(int64(a_exponent)), nil)
	a.Add(a, big.NewInt(1)) // a = 2^20 + 1

	b_base := big.NewInt(2)
	b_exponent := uint(20)
	b := new(big.Int).Exp(b_base, big.NewInt(int64(b_exponent)), nil)
	b.Add(b, big.NewInt(2)) // b = 2^20 + 2

	result := new(big.Int)

	fmt.Println("a =", a.String())
	fmt.Println("b =", b.String())
	fmt.Println("a + b =", result.Add(a, b).String())
	fmt.Println("a - b =", result.Sub(a, b).String())
	fmt.Println("a * b =", result.Mul(a, b).String())
	fmt.Println("a / b =", result.Div(a, b).String())
}
