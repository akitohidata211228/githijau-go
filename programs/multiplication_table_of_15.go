// multiplication_table_of_15.go
// Tabel perkalian 15.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("15 x %d = %d\n", i, 15*i)
	}
}
