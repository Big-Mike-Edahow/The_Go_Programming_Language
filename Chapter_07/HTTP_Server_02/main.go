// main.go

// HTTP_Server_02 is an e-commerce server with /list and /price endpoints.

package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func main() {
	db := database{
		"Shoes": 50,
		"Socks": 5,
	}

	log.Println("Serving HTTP on Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "No such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "No such page: %s\n", req.URL)
	}
}
