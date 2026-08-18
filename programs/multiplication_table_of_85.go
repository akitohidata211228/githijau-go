// multiplication_table_of_85.go
// Tabel perkalian 85.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("85 x %d = %d\n", i, 85*i)
	}
}
