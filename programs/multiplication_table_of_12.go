// multiplication_table_of_12.go
// Tabel perkalian 12.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("12 x %d = %d\n", i, 12*i)
	}
}
