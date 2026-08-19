// multiplication_table_of_97.go
// Tabel perkalian 97.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("97 x %d = %d\n", i, 97*i)
	}
}
