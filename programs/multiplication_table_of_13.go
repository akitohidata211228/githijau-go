// multiplication_table_of_13.go
// Tabel perkalian 13.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("13 x %d = %d\n", i, 13*i)
	}
}
