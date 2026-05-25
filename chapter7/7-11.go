package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	items map[string]dollars
	mu    sync.Mutex
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()
	for item, price := range db.items {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is an empty string")
		return
	}

	_, ok := db.items[item]
	if ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}
	strPrice := req.URL.Query().Get("price")

	price, err := strconv.ParseFloat(strPrice, 32) // float32
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to parse the price %q\n", strPrice)
		return
	}

	db.items[item] = dollars(price)

	fmt.Fprintf(w, "created an item\n")
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is an empty string")
		return
	}

	_, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %q does not exist\n", item)
		return
	}

	delete(db.items, item)
	fmt.Fprintf(w, "deleted an item\n")
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := req.URL.Query().Get("item")
	_, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item does not exist: %q\n", item)
		return
	}
	strPrice := req.URL.Query().Get("price")

	price, err := strconv.ParseFloat(strPrice, 32) // float32
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to parse the price %q\n", strPrice)
		return
	}

	db.items[item] = dollars(price)

	fmt.Fprintf(w, "updated an item\n")
}

func main() {
	db := database{map[string]dollars{"shoes": 50, "socks": 5}, sync.Mutex{}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
