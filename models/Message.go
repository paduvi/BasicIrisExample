package models

import "time"

type Message struct {
	Id        int            `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
}

type Messages []Message