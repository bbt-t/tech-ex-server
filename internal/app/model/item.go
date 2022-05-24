package model

import "time"

type Item struct {
	Id       int
	CreateAt time.Time
	JsonData []byte
}
