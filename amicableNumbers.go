// this is a code for Project Euler q.21

package main

import (
	"fmt"
	"math"
)

func getSumOfDivisors(n int) int {
	var divisors []int
	divisors = append(divisors, 1)
	maximum := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i < maximum+1; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	sum := 0
	for i := 0; i < len(divisors); i++ {
		sum += divisors[i]
	}
	return sum
}

func searchAmicableNumbers(max int) [][]int {
	var resultSlice [][]int
	var sumOfDivisors int
	for i := 1; i <= max; i++ {
		sumOfDivisors = getSumOfDivisors(i)
		if sumOfDivisors > i && getSumOfDivisors(sumOfDivisors) == i {
			pairOfAmicableNumbers := make([]int, 2)
			pairOfAmicableNumbers[0] = i
			pairOfAmicableNumbers[1] = sumOfDivisors
			resultSlice = append(resultSlice, pairOfAmicableNumbers)
		}
	}
	return resultSlice
}

func sumOfAmicableNumbers(sl [][]int) int {
	result := 0
	for i := 0; i < len(sl); i++ {
		result += sl[i][0]
		result += sl[i][1]
	}
	return result
}

func main() {
	amicableNumberSlice := searchAmicableNumbers(10000)
	fmt.Println(amicableNumberSlice)
	fmt.Println("the answer is:", sumOfAmicableNumbers(amicableNumberSlice))
}
