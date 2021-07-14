package serveHandle

import (
	"fmt"
	"net/http"
)

type index struct{}

// on default
func (i *index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index html")
}

// NewServeHandler make a new restfulAPI serve handler
func NewServeHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", &index{})

	return mux
}
