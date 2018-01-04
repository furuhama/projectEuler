// this is a code for Project Euler q.19

package main

import (
	"fmt"
	"time"
)

// it is not good to use "date" package for this question,
// however it is too hard to implement them all by myself
// (and, many modern programming languages have their own datetime package
// we almost never implement them by ourselves)
func countSundays() int {
	counter := 0
	var weekday time.Weekday
	for i := 1901; i <= 2000; i++ {
		for j := 1; j <= 12; j++ {
			weekday = time.Date(i, time.Month(j), 1, 0, 0, 0, 0, time.Local).Weekday()
			// time.Weekday(0) means Sunday
			if weekday == time.Weekday(0) {
				counter++
			}
		}
	}
	return counter
}

func main() {
	count := countSundays()
	fmt.Println("the answer is:", count)
}
