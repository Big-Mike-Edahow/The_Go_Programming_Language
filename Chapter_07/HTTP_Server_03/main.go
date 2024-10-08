// main.go

// HTTP_Server_03 is an e-commerce server that registers the
// /list and /price endpoints by calling (*http.ServeMux).Handle.

package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()

	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	log.Println("Serving HTTP on Port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%v: %s\n", item, price)
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }
