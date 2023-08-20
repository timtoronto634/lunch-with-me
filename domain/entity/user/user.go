package user

import "github.com/google/uuid"

type UserID uuid.UUID

type User struct {
	ID    UserID
	Name  string
	Email string
}
