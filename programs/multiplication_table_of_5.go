// multiplication_table_of_5.go
// Tabel perkalian 5.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("5 x %d = %d\n", i, 5*i)
	}
}
