// multiplication_table_of_99.go
// Tabel perkalian 99.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("99 x %d = %d\n", i, 99*i)
	}
}
