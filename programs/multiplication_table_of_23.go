// multiplication_table_of_23.go
// Tabel perkalian 23.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("23 x %d = %d\n", i, 23*i)
	}
}
