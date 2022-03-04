package model

import (
	"encoding/base64"
	"fmt"
)

type User struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Created string `json:"created"`
	Balance string `json:"balance"`
}

func (u User) GetIndex() string {
	return u.First[0:1]
}

func (u User) GetUniqueId() string {
	allUserFields := fmt.Sprintf("%s_%s_%s_%s_%s_%s", u.First, u.Last, u.Email, u.Address, u.Created, u.Balance)
	encodedText := base64.StdEncoding.EncodeToString([]byte(allUserFields))

	return encodedText
}
