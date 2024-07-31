// main.go

// HTTP_Server_05 is an e-commerce server that registers the
// /list and /price endpoint by calling http.HandleFunc.

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

	http.HandleFunc("/", db.indexHandler)
	http.HandleFunc("/list", db.listHandler)
	http.HandleFunc("/price", db.priceHandler)

	log.Println("Starting HTTP Server on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", logRequest(http.DefaultServeMux)))
}

func (db database) indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to Shoe Emporium!")
}

func (db database) listHandler(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) priceHandler(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%v: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
