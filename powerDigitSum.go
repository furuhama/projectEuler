// this is a code for Project Euler q.16
// 2^1000 は 1000 * log10_2 = 301.02... なので 301,2 桁くらいである

package main

import (
	"fmt"
)

const (
	// try to get 2^1000
	n = 1000
)

func getNextPow(digitSl []int) {
	digitLen := len(digitSl)
	extra := 0
	for i := 0; i < digitLen; i++ {
		digit := digitSl[i]
		digitSl[i] = (digit*2 + extra) % 10
		extra = (digit*2 + extra) / 10
	}
}

func getSumOfDigits(digitSl []int) int {
	digitLen := len(digitSl)
	result := 0
	for i := 0; i < digitLen; i++ {
		result += digitSl[i]
	}
	return result
}

func main() {
	// initialize slice
	digitSlice := make([]int, 400)
	digitSlice[0] = 1

	// iterate the loop for powering by 2
	for i := 0; i < n; i++ {
		getNextPow(digitSlice)
	}

	result := getSumOfDigits(digitSlice)
	fmt.Println("the answer is:", result)
}
