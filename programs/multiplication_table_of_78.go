// multiplication_table_of_78.go
// Tabel perkalian 78.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("78 x %d = %d\n", i, 78*i)
	}
}
