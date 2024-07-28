// main.go

/* Findlinks_03 crawls the web breadth-first, starting with the URLs on the command line.
   How to run: $ go run . https://go.dev */

package main

import (
	"fmt"
	"log"
	"os"

	"findlinks-03/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

/* breadthFirst calls f for each item in the worklist. Any items returned by
   f are added to the worklist. f is called at most once for each item. */
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
