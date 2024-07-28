// main.go

// Rev reverses a slice.

package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Original slice:\n%v\n\n",a)

	// Reverse the order of the slice.
	reverse(a[:])
	fmt.Printf("Reversed slice:\n%v\n\n",a) // "[5 4 3 2 1 0]"

	// Rotate s left by two positions.
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Printf("Rotated left by two Positions:\n%v\n",s) // "[2 3 4 5 0 1]"
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
