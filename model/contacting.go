package model

import (
// "database/sql"

// "github.com/lib/pq"
)

type Contact struct {
	ID          int    `json:"-" `
	FirstName   string `json:"" `
	LastName    string `json:""`
	Text        string `json:""`
	PhoneNumber string `json:""`
	TypeService string `json:""`

	Date string `json:""`
}
