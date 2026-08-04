// multiplication_table_of_3.go
// Tabel perkalian 3.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("3 x %d = %d\n", i, 3*i)
	}
}
