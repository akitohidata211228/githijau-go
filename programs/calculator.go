// calculator.go
// Kalkulator dua angka.

package main

import "fmt"

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("pembagian nol")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("operasi tidak dikenal")
}

func main() {
	res, err := calculate(10, 4, "+")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 + 4 =", res)
	}
}
