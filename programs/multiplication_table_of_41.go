// multiplication_table_of_41.go
// Tabel perkalian 41.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("41 x %d = %d\n", i, 41*i)
	}
}
