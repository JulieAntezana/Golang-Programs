// The Challenge:
// The population of rabbits at a lab doubles every year.

// The lab was started initially with 7 rabbits. Your program needs to calculate the number of rabbits you will have after the number of years provided as input.

// Sample Input:
// 4

// Sample Output:
// 112

// Use a loop to calculate the number of rabbits after each year.
// Explanation: Let's break down the population growth by year:
// Year 0: 7
// Year 1: 14
// Year 2: 28
// Year 3: 56
// Year 4: 112

package main

import "fmt"

func main() {
	var years int
	fmt.Scanln(&years)

	s := 7

	for i := 0; i < years; i++ {
		s = s * 2
	}
	fmt.Println(s)
}
