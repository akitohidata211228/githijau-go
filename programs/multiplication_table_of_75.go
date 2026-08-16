// multiplication_table_of_75.go
// Tabel perkalian 75.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("75 x %d = %d\n", i, 75*i)
	}
}
