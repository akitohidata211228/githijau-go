// multiplication_table_of_45.go
// Tabel perkalian 45.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("45 x %d = %d\n", i, 45*i)
	}
}
