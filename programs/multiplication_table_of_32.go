// multiplication_table_of_32.go
// Tabel perkalian 32.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("32 x %d = %d\n", i, 32*i)
	}
}
