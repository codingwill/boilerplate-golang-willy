package models

import (
	"database/sql"
)

type Person struct {
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Email     string `db:"email" json:"email"`
	Id        int64  `db:"id" json:"id"`
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func (person Person) ServiceGetById(id int) (result Person) {

	return result
}
