package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		// Get n as a string.
		var nString string
		fmt.Printf("N: ")
		fmt.Scanln(&nString)

		// If the n string is blank, break out of the loop.
		if len(nString) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(nString, 10, 64)
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}

// fibonacci returns the nth Fibonacci number and uses recursion.
// F(0) = 0 and F(1) = 1 and F(n) = F(n-1) + F(n-2) for n > 1.
func fibonacci(n int64) int64 {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	}

	// Return the sum of the previous two Fibonacci numbers.
	return fibonacci(n-1) + fibonacci(n-2)
}
