package graphql

import (
	db "art-app/models/DB"

	"github.com/graphql-go/graphql"
)

var artistType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Artist",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"profession": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var photosType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Photos",
		Fields: graphql.Fields{
			"artist": &graphql.Field{
				Type: artistType,
			},
			"art_id": &graphql.Field{
				Type: graphql.String,
			},
			"art_description": &graphql.Field{
				Type: graphql.String,
			},
			"photo_url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func QueryType(db *db.DB) *graphql.Object {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"list":        ArtistList(db),
				"art_list":    ArtList(db),
				"artInput":    Art(db),
				"firtsNthArt": FirstNthPieces(db),
			},
		},
	)

	return queryType
}
