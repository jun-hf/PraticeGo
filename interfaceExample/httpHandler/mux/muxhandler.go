package mux

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type database map[string]dollar

var mu sync.Mutex

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
	fmt.Fprintf(w, "%v is priced at %v", itemName, v)
}

func (d database) testPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", r.Method)
}

/*
database = 

map of string and dollar
cola = $10
nugget = $29
eggs = $2

I want to create a CRUD for the above db
GET /items -> list out all the items
GET /items/id=? -> get the specfic item
POST /items -> create a new item give back the item name and price -> post in JSON
PUT /items/id=?itemName=?&price=? -> change the price of the item
DELETE /items/id=?

-> because it is concurrent access of the db instance we need to put a lock into it
-> create id for each item

due to amount of effort I have result to 
lock and simple CRUD
GET/items?itemName=<> -> get the specific item 
PUT /items?itemName=<>?newPrice=<> -> change the price
DELETE /items?itemName=<> -> delete the item form db
POST /items?itemName=<>?price=<> -> create the new item
*/


func (d database) getItem(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("itemName")
	mu.Lock()
	p, ok := d[itemName]
	mu.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Item: %v not found", itemName)
	}
	fmt.Fprintf(w, "The item: %v is priced at %v", itemName, p)
}

func main() {
	db := database{"drinks": 100, "chair": 90}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("GET /items/{itemName}", db.getItem)
	// every time we get a req it is passed to a new go routine handler, so remember to lock any resources that is outside the handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}