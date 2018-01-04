// this is a code for Project Euler No.5

package main

import "fmt"

func getSmallestMultiple(x, y int) int {
	gcd := getGreatestCommonDivisor(x, y)
	return x * y / gcd
}

func getGreatestCommonDivisor(x, y int) int {
	if x > y {
		x, y = y, x
	}
	for x != 0 {
		x, y = y%x, x
	}
	return y
}

func getSMFromIteration(maxIter int) int {
	result := 1
	for i := 2; i < 21; i++ {
		result = getSmallestMultiple(result, i)
	}
	return result
}

func main() {
	fmt.Println("main function started...")
	result := getSMFromIteration(20)
	fmt.Println(result)
}
