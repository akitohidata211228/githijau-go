// multiplication_table_of_25.go
// Tabel perkalian 25.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("25 x %d = %d\n", i, 25*i)
	}
}
