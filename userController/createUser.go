package userController

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

func (u UserDB) getRandID() int {
	numStr, err := rand.Int(rand.Reader, big.NewInt(10000))
	num := int(numStr.Int64())
	_, exist := u[num]
	if err != nil || exist {
		return -1
	}
	return num
}

func (u UserDB) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(UserInfo)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}

	user.CreatedAt = time.Now()
	var num int
	for num = u.getRandID(); num < 0; num = u.getRandID() {
	}
	user.Id = num
	u[user.Id] = *user
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}
