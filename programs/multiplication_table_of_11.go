// multiplication_table_of_11.go
// Tabel perkalian 11.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("11 x %d = %d\n", i, 11*i)
	}
}
