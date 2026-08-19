// multiplication_table_of_91.go
// Tabel perkalian 91.

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("91 x %d = %d\n", i, 91*i)
	}
}
