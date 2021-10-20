package models

import "time"

type Password struct {
	FirstCode   string    `json:"first_code"`
	SecondCode  string    `json:"second_code"`
	SpecialCode string    `json:"special_code"`
	RandomCode  string    `json:"random_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
