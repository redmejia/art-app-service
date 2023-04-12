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

func QueryType(db *db.DB) *graphql.Object {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"list": ArtistList(db),
			},
		},
	)

	return queryType
}
