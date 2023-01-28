package main

import (
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
	r.Post("/function", app.handleCheckAllowance(http.HandlerFunc(app.handleCreateFunction)))
	r.Delete("/function", app.handleCheckAllowance(http.HandlerFunc(app.handleDeleteFunction)))
}
