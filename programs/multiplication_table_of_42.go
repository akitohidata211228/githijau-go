// multiplication_table_of_42.go
// Tabel perkalian 42.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("42 x %d = %d\n", i, 42*i)
	}
}
