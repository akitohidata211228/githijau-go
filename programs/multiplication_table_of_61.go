// multiplication_table_of_61.go
// Tabel perkalian 61.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("61 x %d = %d\n", i, 61*i)
	}
}
