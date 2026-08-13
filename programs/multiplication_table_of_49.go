// multiplication_table_of_49.go
// Tabel perkalian 49.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("49 x %d = %d\n", i, 49*i)
	}
}
