// multiplication_table_of_89.go
// Tabel perkalian 89.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("89 x %d = %d\n", i, 89*i)
	}
}
