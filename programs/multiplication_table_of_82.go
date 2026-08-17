// multiplication_table_of_82.go
// Tabel perkalian 82.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("82 x %d = %d\n", i, 82*i)
	}
}
