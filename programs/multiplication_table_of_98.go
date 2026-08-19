// multiplication_table_of_98.go
// Tabel perkalian 98.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("98 x %d = %d\n", i, 98*i)
	}
}
