package userController

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserDB map[int]UserInfo

func (u UserDB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		if vars["id"] != "" {
			if usrId, err := strconv.Atoi(vars["id"]); err != nil {
				fmt.Fprintln(w, "Wrong User Id")
			} else {
				u.GetDetailInfo(w, usrId)
			}
		} else {
			u.GetList(w)
		}
	case "POST":
		u.CreateUser(w, r)
	}
}
