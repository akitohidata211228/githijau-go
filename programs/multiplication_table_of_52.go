// multiplication_table_of_52.go
// Tabel perkalian 52.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("52 x %d = %d\n", i, 52*i)
	}
}
