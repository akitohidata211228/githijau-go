// multiplication_table_of_73.go
// Tabel perkalian 73.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("73 x %d = %d\n", i, 73*i)
	}
}
