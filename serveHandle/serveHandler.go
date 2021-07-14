package serveHandle

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sir/restfulAPI/userController"
)

type Index struct{}

// on default page
func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index html")
}

// NewServeHandler make a new restfulAPI serve handler
func NewServeHandler() http.Handler {
	mux := mux.NewRouter()
	userController := make(userController.UserDB)

	mux.Handle("/", &Index{})
	mux.Handle("/users", userController)
	mux.Handle("/users/{id:[0-9]+}", userController)

	return mux
}
