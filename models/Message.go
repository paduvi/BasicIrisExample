package models

import "time"

type Message struct {
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	CreatedAt	time.Time 	`json:"created_at"`
}
