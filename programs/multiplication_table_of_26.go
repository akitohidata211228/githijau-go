// multiplication_table_of_26.go
// Tabel perkalian 26.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("26 x %d = %d\n", i, 26*i)
	}
}
