// multiplication_table_of_20.go
// Tabel perkalian 20.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("20 x %d = %d\n", i, 20*i)
	}
}
