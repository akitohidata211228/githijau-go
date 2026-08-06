// multiplication_table_of_14.go
// Tabel perkalian 14.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("14 x %d = %d\n", i, 14*i)
	}
}
