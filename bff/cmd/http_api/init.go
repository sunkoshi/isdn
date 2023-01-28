package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
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

		// waiting to gracefully shutdown all the processes
		app.logger.Println("quitting application in 3s")
		time.Sleep(time.Second * 3)

		// finally shutting down the server
		app.logger.Println("shutting down the http server")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		app.srv.Shutdown(ctx)
	}()
}

func initDB() (db.Store, error) {
	var err error

	// TODO: pass db info from config
	tDB, err := sql.Open("sqlite3", "../../db.db")
	if err != nil {
		return nil, err
	}
	q, err := db.Prepare(context.Background(), tDB)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func initServer(app *App) {
	r := chi.NewRouter()
	initHandler(app, r)

	// TODO: send address from config file
	srv := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}
	app.srv = &srv
}
