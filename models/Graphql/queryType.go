package graphql

import (
	db "art-app/models/DB"
	"strconv"

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

// first new art pieces
func FirstNthPieces(db *db.DB) *graphql.Field {
	var firstNth = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "FirstNth",
			Fields: graphql.InputObjectConfigFieldMap{
				"nth": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		},
	)

	return &graphql.Field{
		Type:        graphql.NewList(photosType),
		Description: "Get Nth new piece of art",
		Args: graphql.FieldConfigArgument{
			"nth_num": &graphql.ArgumentConfig{
				Type: firstNth,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			nthNumber, ok := p.Args["nth_num"].(map[string]interface{})

			if ok {
				firstPhotos := db.GetAllPhotos()
				return firstPhotos[0:nthNumber["nth"].(int)], nil
			}
			return nil, nil
		},
	}
}

func Art(db *db.DB) *graphql.Field {
	// input object
	var input = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "Input",
			Fields: graphql.InputObjectConfigFieldMap{
				"art_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	)

	return &graphql.Field{
		Type:        photosType,
		Description: "Get single art by art_id",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: input,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			inputD, ok := p.Args["input"].(map[string]interface{})
			id, _ := strconv.Atoi(inputD["art_id"].(string))

			if ok {
				photo := db.GetArtById(id)
				return photo, nil
			}
			return nil, nil
		},
	}

}
