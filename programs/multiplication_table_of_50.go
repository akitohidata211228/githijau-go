// multiplication_table_of_50.go
// Tabel perkalian 50.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("50 x %d = %d\n", i, 50*i)
	}
}
