package userController

import (
	"encoding/json"
	"time"
)

// User Struct
type UserInfo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (u UserInfo) String() string {
	data, _ := json.Marshal(u)
	return string(data)
}
