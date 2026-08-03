// palindrome.go
// Cek apakah sebuah kata palindrom.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(text string) bool {
	var cleaned []rune
	for _, r := range strings.ToLower(text) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}
	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}
	return true
}

func main() {
	for _, s := range []string{"Racecar", "Hello", "Kasur ini rusak"} {
		fmt.Printf("%q -> %v\n", s, isPalindrome(s))
	}
}
