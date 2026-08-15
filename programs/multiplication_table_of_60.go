// multiplication_table_of_60.go
// Tabel perkalian 60.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("60 x %d = %d\n", i, 60*i)
	}
}
