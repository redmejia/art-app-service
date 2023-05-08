package handler

import (
	"encoding/json"
	"io"
	"net/http"

	artql "art-app/models/Graphql"

	"github.com/graphql-go/graphql"
)

func (a *App) ArtHandler(w http.ResponseWriter, r *http.Request) {

	a.InfoLog.Println("Go!")

	query, err := io.ReadAll(r.Body)
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	// a.InfoLog.Println(">>>>>>>>>>>> -------- ", string(query))

	var target map[string]interface{}

	err = json.Unmarshal(query, &target)
	if err != nil {
		a.ErrorLog.Fatal("Umarshal error : ", err)
	}

	// a.InfoLog.Println("QUERY: ", target["query"])

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: artql.QueryType(&a.DB),
		},
	)

	// result := artql.ExecuteQuery(r.URL.Query().Get("query"), schema)
	result := artql.ExecuteQuery(target["query"].(string), schema)

	bytes, err := json.Marshal(result)
	if err != nil {
		a.ErrorLog.Println("ERROR : ", err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(bytes)
}
