// multiplication_table_of_35.go
// Tabel perkalian 35.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("35 x %d = %d\n", i, 35*i)
	}
}
