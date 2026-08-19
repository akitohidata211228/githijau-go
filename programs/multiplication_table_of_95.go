// multiplication_table_of_95.go
// Tabel perkalian 95.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("95 x %d = %d\n", i, 95*i)
	}
}
