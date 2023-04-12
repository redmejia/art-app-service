package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	artql "art-app/models/Graphql"

	"github.com/graphql-go/graphql"
)

func (a *App) ArtHandler(w http.ResponseWriter, r *http.Request) {

	a.InfoLog.Println("Go!")

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: artql.QueryType(&a.DB),
		},
	)

	result := artql.ExecuteQuery(r.URL.Query().Get("query"), schema)

	// artists := a.DB.GetAll()

	fmt.Println("this ", result)

	bytes, err := json.Marshal(result)
	if err != nil {
		a.ErrorLog.Println("ERROR : ", err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(bytes)
}
