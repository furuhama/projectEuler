// this is a code for Project Euler No.12

package main

import (
	"fmt"
	"math"
)

const (
	underline = 500
)

func countDivisor(n int) int {
	// escape if n equals 0
	if n == 0 {
		n = 1
	}
	counter := 0
	upperLimit := int(math.Sqrt(float64(n)))
	for i := 1; i < upperLimit+1; i++ {
		if n%i == 0 {
			if i == upperLimit {
				counter++
			} else {
				counter += 2
			}
		}
	}
	return counter
}

func main() {
	triangle := 0
	for i := 1; countDivisor(triangle) <= underline; i++ {
		triangle += i
	}
	fmt.Println("the answer is:", triangle)
}
