// this is a code for Project Euler q.20

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorialBig(n int) *big.Int {
	result := factorialInside(big.NewInt(1), n)
	return result
}

func factorialInside(base *big.Int, n int) *big.Int {
	if n != 1 {
		base.Mul(base, big.NewInt(int64(n)))
		factorialInside(base, n-1)
	}
	return base
}

func parseStringGetSum(s string) int {
	resultSum := 0
	for i := 0; i < len(s); i++ {
		num, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			fmt.Println(err)
		}
		resultSum += num
	}
	return resultSum
}

func main() {
	resultString := factorialBig(100).String()
	result := parseStringGetSum(resultString)
	fmt.Println("the answer is:", result)
}
