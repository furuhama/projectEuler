// this is a code for Project Euler q.15

package main

import (
	"fmt"
	"math/big"
)

func factrial(n int64) *big.Int {
	if n == 1 || n == 0 {
		return big.NewInt(1)
	} else {
		return big.NewInt(1).Mul(big.NewInt(n), factrial(n-1))
	}
}

func main() {
	result := new(big.Int).Div(factrial(40), new(big.Int).Mul(factrial(20), factrial(20)))
	fmt.Println("the answer is:", result)
}
