package model

type SearchParams struct {
	FirstName string
	LastName  string
	Offset    int
	OrderBy   string // Id, First name, Last name
	OrderType string //ASC OR DESC
}
