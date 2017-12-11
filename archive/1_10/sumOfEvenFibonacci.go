package main

import "fmt"

// updateFibonacci gets slice
// it appends a next fibonacci number to the last
// it returns slice
func updateFibonacci(i int, m []int) []int {
	if i == 0 {
		m = append(m, 1)
		return m
	} else if i == 1 {
		m = append(m, 2)
		return m
	} else {
		a := m[i-2]
		b := m[i-1]

		m = append(m, a+b)
		return m
	}
}

func main() {
	// set slice
	var memory []int
	sum := 0
	// for iteration
	i := 0
	for {
		memory = updateFibonacci(i, memory)
		// if memory[i] is larger than 4 mil, break loop
		// else, calculate sum, and increment i
		if memory[i] >= 4000000 {
			break
		} else {
			// if memory[i] is even number, add it to sum
			if memory[i]%2 == 0 {
				sum += memory[i]
			}
			i++
		}
	}
	fmt.Printf("memory: %d\nThe sum is %d", memory, sum)
}
