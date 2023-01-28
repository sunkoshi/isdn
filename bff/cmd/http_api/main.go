package main

import (
	"log"
	"net/http"
	"os"

	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/pkg/object_store"
)

type App struct {
	store       *db.Queries
	kv          map[string]string
	objectStore *object_store.ObjectStore
	logger      *log.Logger
	quitters    map[string]chan struct{}
	osSignal    chan os.Signal
	srv         *http.Server
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
	initObjectStore(app)
	initServer(app)
	go initCleaner(app)

	log.Fatal(app.srv.ListenAndServe())
}
