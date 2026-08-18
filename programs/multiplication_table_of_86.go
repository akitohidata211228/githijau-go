// multiplication_table_of_86.go
// Tabel perkalian 86.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("86 x %d = %d\n", i, 86*i)
	}
}
