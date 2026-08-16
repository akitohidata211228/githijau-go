// multiplication_table_of_72.go
// Tabel perkalian 72.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("72 x %d = %d\n", i, 72*i)
	}
}
