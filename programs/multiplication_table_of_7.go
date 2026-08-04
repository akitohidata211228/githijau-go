// multiplication_table_of_7.go
// Tabel perkalian 7.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("7 x %d = %d\n", i, 7*i)
	}
}
