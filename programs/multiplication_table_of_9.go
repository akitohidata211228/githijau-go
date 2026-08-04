// multiplication_table_of_9.go
// Tabel perkalian 9.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("9 x %d = %d\n", i, 9*i)
	}
}
