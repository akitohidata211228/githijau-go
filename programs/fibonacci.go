// fibonacci.go
// Deret Fibonacci: 15 suku pertama.

package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
