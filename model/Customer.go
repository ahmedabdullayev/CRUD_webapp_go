package model

import "time"

type Customer struct {
	Id        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Gender    string
	Email     string
	Address   string
}
