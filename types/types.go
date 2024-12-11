package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(ID int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"FirstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"passward" validate:"required,min=8,max=16"`
}
