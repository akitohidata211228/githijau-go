// multiplication_table_of_21.go
// Tabel perkalian 21.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("21 x %d = %d\n", i, 21*i)
	}
}
