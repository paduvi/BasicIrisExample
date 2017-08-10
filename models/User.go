package models

import "time"

type User struct {
	Id        int                `json:"id"`
	Histories map[int]time.Time    `json:"histories"`
}
