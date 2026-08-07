// multiplication_table_of_27.go
// Tabel perkalian 27.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("27 x %d = %d\n", i, 27*i)
	}
}
