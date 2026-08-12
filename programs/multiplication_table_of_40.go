// multiplication_table_of_40.go
// Tabel perkalian 40.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("40 x %d = %d\n", i, 40*i)
	}
}
