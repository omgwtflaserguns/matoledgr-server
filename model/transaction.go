package model

import (
	"time"
)

type Transaction struct {
	Id        int32
	AccountId int32
	ProductId int32
	Price     float32
	Timestamp time.Time
}
