// multiplication_table_of_84.go
// Tabel perkalian 84.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("84 x %d = %d\n", i, 84*i)
	}
}
