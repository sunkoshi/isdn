package main

import (
	"log"
	"net/http"
	"os"

	db "github.com/orted-org-isdn-bff/db/sqlc"
)

type App struct {
	// db store
	store *db.Queries
	kv    map[string]string

	//logger
	logger *log.Logger

	// service quitter signal channel map
	quitters map[string]chan struct{}

	// channel for os signals
	osSignal chan os.Signal

	// http server
	srv *http.Server
}

var (
	lo = log.New(os.Stdout, "",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {

	app := &App{
		quitters: make(map[string]chan struct{}),
		logger:   lo,
		kv:       map[string]string{},
	}

	initDB(app)
	initServer(app)
	go initCleaner(app)

	log.Fatal(app.srv.ListenAndServe())
}
