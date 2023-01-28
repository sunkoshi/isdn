package main

import (
	"github.com/go-chi/chi/v5"
)

func initHandler(app *App, r *chi.Mux) {

	// uptime
	// r.Post("/uptime", app.handleCheckAllowance(http.HandlerFunc(app.handleCreateWatchReq)))
	// r.Get("/uptime", app.handleCheckAllowance(http.HandlerFunc(app.handleGetWatchReq)))
	// r.Put("/uptime/{id}", app.handleCheckAllowance(http.HandlerFunc(app.handleUpdateWatchReq)))
	// r.Delete("/uptime/{id}", app.handleCheckAllowance(http.HandlerFunc(app.handleDeleteWatchReq)))

}
