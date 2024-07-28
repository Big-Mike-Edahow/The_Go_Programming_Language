// main.go

/* Issues prints a table of GitHub issues matching the search terms.
   How to use: $ go run . repo:golang/go is:open json decoder */

package main

import (
	"fmt"
	"log"
	"os"

	"go-issues/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}