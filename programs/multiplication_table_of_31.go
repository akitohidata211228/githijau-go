// multiplication_table_of_31.go
// Tabel perkalian 31.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("31 x %d = %d\n", i, 31*i)
	}
}
