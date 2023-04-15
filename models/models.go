package models

type Artist struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`       // change to artis_name
	Profession string `json:"profession"` // painter | urban artist (grafitti) | photgrafer
}

type Photos struct {
	ID          int    `json:"id"`
	Description string `json:"discription"` // fix type on db table photo
	Name        string `json:"name"`
	PhotoURL    string `json:"photo_url"`
}
