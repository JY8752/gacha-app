package model

import "time"

type User struct {
	Id        string
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}
