package main

import (
	"art-app/api/handler"
	"art-app/api/router"
	db "art-app/models/DB"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	dbCon, err := db.DSNConn(os.Getenv("DSN"))
	if err != nil {
		errorLog.Fatal("CONNECTION ERROR : ", err)
	}

	app := handler.App{
		DB:       db.DB{Db: dbCon},
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	srv := &http.Server{
		Addr:         ":80",
		Handler:      router.Router(&app),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Println("ARTAPP-SERVICE")
	infoLog.Println("http://localhost:8080/")
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal("Service Faild : ", err)
	}

}
