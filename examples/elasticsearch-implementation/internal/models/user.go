package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
}
