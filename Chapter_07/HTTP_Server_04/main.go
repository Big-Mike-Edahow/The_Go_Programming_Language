// main.go

// HTTP_Server_04 is an e-commerce server that registers the
// /list and /price endpoints by calling (*http.ServeMux).HandleFunc.

package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]int

func main() {
	db := database{"shoes": 50,"socks": 5}
	mux := http.NewServeMux()

	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	
	log.Println("Serving HTTP on Port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%v: $%d\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
