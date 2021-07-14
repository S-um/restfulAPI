package userController

import (
	"fmt"
	"net/http"
)

// request with get with nothing
func (u UserDB) GetList(w http.ResponseWriter) {
	fmt.Fprintf(w, "User Data\n\n")
	for _, val := range u {
		fmt.Fprintln(w, val.Id)
	}
}
