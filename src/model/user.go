package model

import "encoding/json"

type User struct {
	// user id
	Id          string   `json:"user_id"`
	Authorities []string `json:"authorities"`
	Scope       []string `json:"scope"`
}

func (u *User) String() string {
	j, _ := json.Marshal(u)
	return string(j)
}
