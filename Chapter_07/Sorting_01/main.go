// main.go

// Sorting_01 sorts a slice of integers.

package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 2}

	fmt.Println("Values:", values)
	fmt.Println("Are values sorted?", sort.IntsAreSorted(values))

	sort.Ints(values)
	fmt.Println("\nSorted:", values)
	fmt.Println("Are values sorted?", sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println("\nReversed:", values)
	fmt.Println("Are values sorted?",sort.IntsAreSorted(values))
}
