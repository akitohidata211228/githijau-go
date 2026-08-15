// multiplication_table_of_64.go
// Tabel perkalian 64.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("64 x %d = %d\n", i, 64*i)
	}
}
