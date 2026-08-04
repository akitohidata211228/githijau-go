// multiplication_table_of_2.go
// Tabel perkalian 2.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %d = %d\n", i, 2*i)
	}
}
