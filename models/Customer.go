package models

import "time"

type Customer struct {
	Id        int
	FirstName string    `json:"firstName" db:"first_name" validate:"required,max=100"`
	LastName  string    `json:"lastName" db:"last_name" validate:"required,max=100"`
	BirthDate time.Time `json:"birthDate" db:"birth_date"` //check in controller
	Gender    string    `json:"gender" db:"gender" validate:"required,oneof=Female Male"`
	Email     string    `json:"email" db:"email" validate:"required,email,max=55"`
	Address   string    `json:"address" db:"address" validate:"max=200"`
}
