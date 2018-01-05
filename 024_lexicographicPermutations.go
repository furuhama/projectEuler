package main

import (
	"fmt"
	"strconv"
)

func removeElementFromSlice(sl []int, i int) []int {
	if i > len(sl)-1 {
		fmt.Println("an error occured! second argument is too large")
		return sl
	}
	newSl := make([]int, 0, len(sl)-1)
	for j := 0; j < len(sl); j++ {
		if j != i {
			newSl = append(newSl, sl[j])
		}
	}
	return newSl
}

func getDivAndModFromNaturalNum(num int, natural int) (int, int) {
	naturalNumber := 1
	for i := 1; i <= natural; i++ {
		naturalNumber *= i
	}
	div := num / naturalNumber
	mod := num % naturalNumber
	if num%naturalNumber == 0 {
		div--
		mod += naturalNumber
	}
	return div, mod
}

func getLexPermFromIndex(index int) string {
	if index > 3628800 { // 3628800 is 10!
		fmt.Println("an error occured! index is too large!")
		return "0"
	} else if index <= 0 {
		fmt.Println("an error occured! index is lower than 1!")
	}

	// initialize slices
	digitList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	targetNum := make([]int, 0, 10)
	var div int
	var mod int
	for i := 1; i < 10; i++ {
		div, mod = getDivAndModFromNaturalNum(index, 10-i)
		targetNum = append(targetNum, digitList[div])
		digitList = removeElementFromSlice(digitList, div)
		index = mod
	}
	targetNum = append(targetNum, digitList[0])
	// convert slice to []rune
	result := make([]rune, 0, 10)
	for i := 0; i < len(targetNum); i++ {
		result = append(result, []rune(strconv.Itoa(targetNum[i]))...)
	}
	return string(result)
}

func main() {
	ans := getLexPermFromIndex(1000000)
	fmt.Println("the answer is:", ans)
}
