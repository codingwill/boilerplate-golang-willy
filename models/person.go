package models

import (
	"database/sql"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func (person Person) ServiceGetById(id int) (result Person) {

	return result
}
