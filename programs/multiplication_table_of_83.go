// multiplication_table_of_83.go
// Tabel perkalian 83.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("83 x %d = %d\n", i, 83*i)
	}
}
