// word.go

// Word provides utilities for word games.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Is 'racecar' a palindrome?", IsPalindrome("racecar"))
	fmt.Println("Is 'kayak' a palindrome?", IsPalindrome("kayak"))
	fmt.Println("Is 'rotator' a palindrome?", IsPalindrome("rotator"))
}

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
