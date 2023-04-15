package router

import (
	"art-app/api/handler"
	"net/http"
)

func Router(a *handler.App) http.Handler {

	mux := http.NewServeMux()
	// ?query={list{id,name,info,price}}
	mux.HandleFunc("/art", a.ArtHandler)

	var fs = http.FileServer(http.Dir("/app/img"))
	mux.Handle("/art/img/", http.StripPrefix("/art/img/", fs))

	return mux
}
