package router

import (
	"art-app/api/handler"
	"net/http"
)

func Router(a *handler.App) http.Handler {
	mux := http.NewServeMux()
	// ?query={list{id,name,info,price}}
	mux.HandleFunc("/art", a.ArtHandler)

	return mux
}
