package model

import "time"

type User struct {
	Id       int32
	Username string
	Hash     []byte
}

type Login struct {
	User    User
	Cookie  string
	Created time.Time
}
