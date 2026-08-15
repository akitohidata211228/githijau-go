// multiplication_table_of_62.go
// Tabel perkalian 62.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("62 x %d = %d\n", i, 62*i)
	}
}
