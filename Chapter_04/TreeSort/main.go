// main.go

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"go-sort-test/treesort"
)

func main() {
	data := TestSort(&testing.T{})
	fmt.Println("Tree Sort of Test Data:\n", data)
}

func TestSort(t *testing.T)[]int {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
	return data
}
