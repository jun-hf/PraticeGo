package mux

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollar

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d)}

func (d database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "The item %v is priced at %v", item, price)
	}
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("item")
	v, ok:= d[itemName]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v not found", itemName)
	}
	fmt.Fprintf(w, "%v is priced at", itemName, v)
}

func main() {
	db := &database{"drinks": 100, "chair": 90}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe(":8080", mux))
}