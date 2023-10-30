package main

import "fmt"

func main() {
	var n int64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, factorial(n))
	}
	fmt.Println()
}

func factorial(n int64) int64 {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}
