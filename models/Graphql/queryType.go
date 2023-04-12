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
