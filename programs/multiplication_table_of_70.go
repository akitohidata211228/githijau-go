// multiplication_table_of_70.go
// Tabel perkalian 70.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("70 x %d = %d\n", i, 70*i)
	}
}
