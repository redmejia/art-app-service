package models

type Artist struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`       // change to artis_name
	Profession string `json:"profession"` // painter | urban artist (grafitti) | photgrafer
}

type Photos struct {
	Artist         `json:"artist"`
	ArtId          int    `json:"art_id"`
	ArtDescription string `json:"art_description"`
	PhotoURL       string `json:"photo_url"`
}
