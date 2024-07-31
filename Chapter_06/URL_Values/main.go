// main.go

// URL_Values demonstrates a map type with methods.

package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // Direct construction
	// Add adds the value to key.
	// It appends to any existing values associated with key.
	m.Add("item", "1")
	m.Add("item", "2")

	// Get returns the first value associated with the given key.
	// It returns "" if there are none.
	fmt.Println(m.Get("lang")) // "en"	"en" is the value associated with "lang" in the map.
	fmt.Println(m.Get("q"))    // ""	There is no "q" key.
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // [1 2] Direct map access.
	fmt.Println(m["item"][0])  // 1
	fmt.Println(m["item"][1])  // 2

	m = nil                    // nil map{}
	fmt.Println(m.Get("item")) // ""
}
