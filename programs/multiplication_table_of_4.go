// multiplication_table_of_4.go
// Tabel perkalian 4.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("4 x %d = %d\n", i, 4*i)
	}
}
