package handler

import (
	db "art-app/models/DB"
	"log"
)

type App struct {
	DB       db.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
