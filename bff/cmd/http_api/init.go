package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/internal/worker_manager"
	"github.com/orted-org-isdn-bff/pkg/object_store"
)

// function to cleanup the open resources
func initCleaner(app *App) {
	app.osSignal = make(chan os.Signal, 1)
	signal.Notify(app.osSignal, os.Interrupt)
	go func() {
		<-app.osSignal
		app.logger.Println("starting cleaning up")

		app.logger.Println("removing all the go routines running services")
		for _, v := range app.quitters {
			v <- struct{}{}
		}

		app.logger.Println("closing db connection")
		app.store.Close()

		// finally shutting down the server
		app.logger.Println("shutting down the http server")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		app.srv.Shutdown(ctx)
	}()
}

func initDB(app *App) {
	var err error

	tDB, err := sql.Open("sqlite3", "./db/dev.db")
	if err != nil {
		panic(err)
	}
	q, err := db.Prepare(context.Background(), tDB)
	if err != nil {
		panic(err)
	}
	app.store = q
}
func initObjectStore(app *App) {
	os, err := object_store.NewObjectStore("./code_repo")
	if err != nil {
		panic(err)
	}
	app.objectStore = os
}

func initWorkerManager(app *App) {
	app.wm = worker_manager.New(app.store)
}

func initServer(app *App) {
	r := chi.NewRouter()
	initHandler(app, r)

	srv := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}
	app.srv = &srv
}
