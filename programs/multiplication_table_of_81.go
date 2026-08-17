// multiplication_table_of_81.go
// Tabel perkalian 81.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("81 x %d = %d\n", i, 81*i)
	}
}
