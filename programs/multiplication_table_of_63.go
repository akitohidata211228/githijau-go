// multiplication_table_of_63.go
// Tabel perkalian 63.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("63 x %d = %d\n", i, 63*i)
	}
}
