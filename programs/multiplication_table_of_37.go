// multiplication_table_of_37.go
// Tabel perkalian 37.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("37 x %d = %d\n", i, 37*i)
	}
}
