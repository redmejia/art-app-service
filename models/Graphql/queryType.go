package graphql

import (
	db "art-app/models/DB"

	"github.com/graphql-go/graphql"
)

func ArtistList(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(artistType),
		Description: "Get all artists",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			artistList := db.GetAll()
			return artistList, nil
		},
	}
}

func ArtList(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(photosType),
		Description: "Get all art photos",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			photos := db.GetAllPhotos()
			return photos, nil
		},
	}
}

func Art(db *db.DB) *graphql.Field {
	return &graphql.Field{
		Type:        photosType,
		Description: "Get single art by art_id",
		Args: graphql.FieldConfigArgument{
			"art_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			artId, ok := p.Args["art_id"].(int)
			if ok {
				photo := db.GetArtById(artId)
				return photo, nil
			}
			return nil, nil
		},
	}
}
