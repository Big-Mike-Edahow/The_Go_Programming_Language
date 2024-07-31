// main.go

// HTTP_Server_01 is a rudimentary e-commerce server.

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
		"<p><b>Shoes</b></p>": 50,
		"<p><b>Socks</b></p>": 5,
	}

	log.Println("Starting HTTP Server on Port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
