// multiplication_table_of_8.go
// Tabel perkalian 8.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("8 x %d = %d\n", i, 8*i)
	}
}
