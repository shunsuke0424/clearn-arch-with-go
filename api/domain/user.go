package domain

import "time"

type Users []User

type User struct {
	ID        uint      `json:"id" param:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
