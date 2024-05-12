package model

import "time"

type Event struct {
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Info   string    `json:"info"`
	Date   time.Time `json:"date"`
}

type Storage struct {
	Title string    `json:"title"`
	Info  string    `json:"info"`
	Date  time.Time `json:"date"`
}
