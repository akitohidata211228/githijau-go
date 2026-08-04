// multiplication_table_of_10.go
// Tabel perkalian 10.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("10 x %d = %d\n", i, 10*i)
	}
}
