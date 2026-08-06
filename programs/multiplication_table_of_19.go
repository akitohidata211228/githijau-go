// multiplication_table_of_19.go
// Tabel perkalian 19.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("19 x %d = %d\n", i, 19*i)
	}
}
