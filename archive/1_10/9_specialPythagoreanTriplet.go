// this is a code for Project Euler No.9

package main

/*
a^2 + b^2 = c^2
a + b + c = 1000

pattern 1:
a, b, c: even

pattern 2:
a(or b), c: odd
b: even

a, b < c
so, c is 334 at least
*/

import (
	"fmt"
	"math"
)

func searchPythagoreanNum() (int, int, int) {
	var a, b, c int
	// c is 334 at least
	for c = 334; c < 1000; c++ {
		// b is Sqrt(c * c / 2) at least
		halfSqrt := int(math.Sqrt(float64(c * c / 2)))
		for b = halfSqrt; b < c; b++ {
			a = 1000 - (b + c)
			if a*a == c*c-b*b {
				return a, b, c
			}
		}
		// hoge
	}
	return 0, 0, 0
}

func main() {
	a, b, c := 0, 0, 0
	a, b, c = searchPythagoreanNum()
	result := a * b * c
	fmt.Println("a, b, c: ", a, b, c, "\nthe answer is ", result)
}
