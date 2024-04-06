package httphandler

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollar

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d)}

func (d *database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range *d {
		fmt.Fprintf(w, "The item: %v is discounted at %v\n", item, price)
	}
}
func main() {
	db := &database{"drinks": 100, "chair": 90}
	log.Fatal(http.ListenAndServe(":8080", db))
}