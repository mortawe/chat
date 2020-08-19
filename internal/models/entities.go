package models

import "time"

type ID int

type User struct {
	ID        ID        `json:"id"`
	Username  string    `json:"username" db:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type Chat struct {
	ID        ID        `json:"id"`
	Name      string    `json:"name"`
	Users     []ID      `json:"users"`
	CreatedAt time.Time `json:"created_at"`
}

type Message struct {
	ID        ID        `json:"id"`
	ChatID    ID        `json:"chat"`
	AuthorID  ID        `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
