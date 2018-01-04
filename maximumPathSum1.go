// this is a code for Project Euler q.18

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	pSize   = 15
	pyramid = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
)

func makeArray(text string) [pSize][pSize]int {
	// set array
	var pathArray [pSize][pSize]int

	// parse the pyramid(str -> int)
	stringArray := strings.Split(text, "\n")
	for i := 0; i < len(stringArray); i++ {
		numStrArray := strings.Split(stringArray[i], " ")
		for j := 0; j < len(numStrArray); j++ {
			num, err := strconv.Atoi(numStrArray[j])
			if err != nil {
				fmt.Println(err)
			}
			pathArray[i][j] = num
		}
	}
	return pathArray
}

func returnMaximumPathSum(pathArray [pSize][pSize]int) int {
	// init the slice with the lowest layer of pathArray
	pathSumSlice := make([]int, pSize)
	for i := 0; i < pSize; i++ {
		pathSumSlice[i] = pathArray[pSize-1][i]
	}
	var bigger int

	// the main algorithm from here
	for i := 0; i < pSize-1; i++ {
		// think of pathArray[pSize-2-i][XXX]
		// this layer has pSize-1-i elements
		for j := 0; j < pSize-1-i; j++ {
			bigger = int(math.Max(float64(pathSumSlice[j]), float64(pathSumSlice[j+1])))
			pathSumSlice[j] = pathArray[pSize-2-i][j] + bigger
		}
	}
	return pathSumSlice[0]
}

func main() {
	pyramidArray := makeArray(pyramid)
	ans := returnMaximumPathSum(pyramidArray)
	fmt.Println("the answer is:", ans)
}
