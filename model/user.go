package model

import (
	pb "github.com/omgwtflaserguns/matomat-server/generated"
	"time"
)

type User struct {
	Id       int32
	Username string
	Hash     []byte
}

func GetProtoUserFromUser(user User) *pb.User {
	return &pb.User{Username: user.Username}
}

type Login struct {
	User    User
	Cookie  string
	Created time.Time
}
