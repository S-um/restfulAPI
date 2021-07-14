package main

import (
	"net/http"

	"github.com/sir/restfulAPI/serveHandle"
)

func main() {
	http.ListenAndServe(":3000", serveHandle.NewServeHandler())
}
