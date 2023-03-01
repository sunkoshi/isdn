package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initHandler(app *App, r *chi.Mux) {

	// auth
	r.Post("/auth", http.HandlerFunc(app.handleCreateUser))
	r.Get("/auth", app.handleCheckAllowance(http.HandlerFunc(app.handleIfLogin)))
	r.Post("/auth/login", http.HandlerFunc(app.handleLogin))
	r.Delete("/auth", app.handleCheckAllowance(http.HandlerFunc(app.handleLogout)))

	// function
	r.Get("/function", app.handleCheckAllowance(http.HandlerFunc(app.handleGetFunctions)))
	r.Get("/function/run/{id}", http.HandlerFunc(app.handleFunctionCall))
	r.Post("/function", app.handleCheckAllowance(http.HandlerFunc(app.handleCreateFunction)))
	r.Delete("/function", app.handleCheckAllowance(http.HandlerFunc(app.handleDeleteFunction)))

	r.Get("/worker", http.HandlerFunc(app.handleWs))

	r.Get("/code/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		file, err := app.objectStore.Get(idStr)
		if err != nil {
			log.Println(err)
			return
		}
		b, err := io.ReadAll(file)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(b)
	})
}
