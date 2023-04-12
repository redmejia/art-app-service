package models

type Artist struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`       // change to artis_name
	Profession string `json:"profession"` // painter | urban artist (grafitti) | photgrafer
}
