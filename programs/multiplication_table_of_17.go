// multiplication_table_of_17.go
// Tabel perkalian 17.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("17 x %d = %d\n", i, 17*i)
	}
}
