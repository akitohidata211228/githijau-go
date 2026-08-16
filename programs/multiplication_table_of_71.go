// multiplication_table_of_71.go
// Tabel perkalian 71.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("71 x %d = %d\n", i, 71*i)
	}
}
