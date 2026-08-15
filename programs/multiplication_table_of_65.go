// multiplication_table_of_65.go
// Tabel perkalian 65.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("65 x %d = %d\n", i, 65*i)
	}
}
