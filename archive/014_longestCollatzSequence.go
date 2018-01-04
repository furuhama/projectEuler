// this is a code for Project Euler q.14
// this program takes really a lot of time ...(need to refactor)
package main

import (
	"fmt"
)

const (
	// N is an array size
	N = 1000000
)

var stepList [N]int

func applyRuleOnce(n int) int {
	var result int
	if n%2 == 0 {
		result = n / 2
	} else {
		result = 3*n + 1
	}
	return result
}

func updateStepList(sq []int, stepList *[N]int) {
	// length of sq
	sqLen := len(sq)
	// set the base step
	lowestStep := stepList[sq[sqLen-1]-1]
	for i := 1; i < sqLen; i++ {
		if sq[sqLen-i-1] < N+1 {
			stepList[sq[sqLen-i-1]-1] = lowestStep + i
		}
	}
}

func makeSequence(n int, stepList *[N]int) []int {
	// apply rule until the number becomes 1
	// and make them slice
	resultList := make([]int, 0, 10*n)
	resultList = append(resultList, n)
	for n != 1 {
		n = applyRuleOnce(n)
		resultList = append(resultList, n)
		if n < N {
			if checkList(n, stepList) {
				return resultList
			}
		}
	}
	resultList = append(resultList, 1)
	return resultList
}

func checkList(n int, stepList *[N]int) bool {
	resultBool := false
	if stepList[n-1] != 0 {
		resultBool = true
	}
	return resultBool
}

func returnHighestIndex(maxIndex int, stepList *[N]int) int {
	chainLen := 1
	resultIndex := 1
	for i := 0; i < maxIndex; i++ {
		if stepList[i] > chainLen {
			chainLen = stepList[i]
			resultIndex = i + 1
		}
	}
	return resultIndex
}

func main() {
	// initialize array
	stepList[0] = 1
	stepList[1] = 2

	// iterate under 1,000,001
	for i := 2; i < N; i++ {
		if stepList[i-1] == 0 {
			sl := makeSequence(i, &stepList)
			updateStepList(sl, &stepList)
		}
	}
	resultIndex := returnHighestIndex(N, &stepList)
	// for i := 0; i < 30; i++ {
	// 	fmt.Println(i+1, stepList[i])
	// }
	fmt.Println(resultIndex)
}
