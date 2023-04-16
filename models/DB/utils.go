package db

import (
	"art-app/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (db *DB) GetAll() []models.Artist {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var artists []models.Artist

	row, err := db.Db.QueryContext(ctx, `select id, artist_name, profession from accounts`)
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
	photosByte, err := json.MarshalIndent(&photos, "", " ")
	if err != nil {
		log.Println("ERORR INDENT: ", err)
	}
	log.Println(string(photosByte))

	return photos

}

func (db *DB) GetArtById(artID int) models.Photos {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var photo models.Photos // single art photo

	row := db.Db.QueryRowContext(ctx,
		`SELECT 
			acc.id,
			acc.artist_name,
			acc.profession,
			ap.art_id,
			ap.art_description,
			ap.photo_url
		FROM accounts acc
		JOIN art_photos ap ON acc.id = ap.id
		WHERE ap.art_id = $1`, artID,
	)

	err := row.Scan(
		&photo.Artist.ID,
		&photo.Artist.Name,
		&photo.Artist.Profession,
		&photo.ArtId,
		&photo.ArtDescription,
		&photo.PhotoURL,
	)

	if err != nil {
		log.Println("ERORR ", err)
	}

	photoByte, err := json.MarshalIndent(&photo, "", " ")
	if err != nil {
		log.Println("ERORR INDENT: ", err)
	}
	log.Println(string(photoByte))

	return photo
}
