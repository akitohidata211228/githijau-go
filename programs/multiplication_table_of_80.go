// multiplication_table_of_80.go
// Tabel perkalian 80.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("80 x %d = %d\n", i, 80*i)
	}
}
