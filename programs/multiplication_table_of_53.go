// multiplication_table_of_53.go
// Tabel perkalian 53.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("53 x %d = %d\n", i, 53*i)
	}
}
