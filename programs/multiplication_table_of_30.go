// multiplication_table_of_30.go
// Tabel perkalian 30.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("30 x %d = %d\n", i, 30*i)
	}
}
