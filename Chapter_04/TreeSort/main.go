// main.go

// TreeSort provides insertion sort using an unbalanced binary tree.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"go-sort-test/treesort"
)

func main() {
	TestSort(&testing.T{})
}

func TestSort(t *testing.T){
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Printf("Unsorted data:\n%v\n", data)
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
	fmt.Println("\nTree Sort of Test Data:\n", data)
}
