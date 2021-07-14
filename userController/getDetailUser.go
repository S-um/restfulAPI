package userController

import (
	"fmt"
	"net/http"
)

// request with get with user id
func (u UserDB) GetDetailInfo(w http.ResponseWriter, usrId int) {
	fmt.Fprint(w, u[usrId])
}
