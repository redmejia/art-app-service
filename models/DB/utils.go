package db

import (
	"art-app/models"
	"context"
	"log"
	"time"
)

func (db *DB) GetAll() []models.Artist {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var artists []models.Artist

	row, err := db.Db.QueryContext(ctx, `select a_id, artist_name, profession from artists`)
	if err != nil {
		log.Println("error rows ", err)
	}

	for row.Next() {
		var artist models.Artist
		row.Scan(&artist.ID, &artist.Name, &artist.Profession)
		artists = append(artists, artist)
	}

	return artists

}
