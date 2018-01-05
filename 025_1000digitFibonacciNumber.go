package main

import (
	"fmt"
	"math/big"
)

func getEnoughDigitFibonacci(n int) int {
	// set a & b, as big.Int type
	a := big.NewInt(1)
	b := big.NewInt(0)
	counter := 1
	minBigInt := new(big.Int)
	minBigInt.Exp(big.NewInt(10), big.NewInt(int64(n-1)), nil)
	// iterate loop while a is less or equal than minBigInt
	for a.Cmp(minBigInt) <= 0 {
		// this is good fibonacci algorithm
		b.Add(a, b)
		a, b = b, a
		counter++
	}
	return counter
}

func main() {
	result := getEnoughDigitFibonacci(1000)
	fmt.Println("the answer is:", result)
}
