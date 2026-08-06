// multiplication_table_of_16.go
// Tabel perkalian 16.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("16 x %d = %d\n", i, 16*i)
	}
}
