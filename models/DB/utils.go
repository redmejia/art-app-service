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

	row, err := db.Db.QueryContext(ctx, `select a_id, artist_name, profession from accounts`)
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

func (db *DB) GetAllPhotos() []models.Photos {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var photos []models.Photos

	row, err := db.Db.QueryContext(ctx,
		`SELECT 
			acc.id,
			acc.artist_name,
			acc.profession,
			ap.art_id,
			ap.art_description,
			ap.photo_url
		FROM accounts acc
		JOIN art_photos ap ON acc.id = ap.id`,
	)

	if err != nil {
		log.Println("error rows ", err)
	}

	for row.Next() {
		var photo models.Photos
		row.Scan(
			&photo.Artist.ID,
			&photo.Artist.Name,
			&photo.Artist.Profession,
			&photo.ArtId,
			&photo.ArtDescription,
			&photo.PhotoURL,
		)
		photos = append(photos, photo)
	}
	log.Println("DATA : ", photos)

	return photos

}
