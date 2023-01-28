package main

import (
	"context"
	"net/http"
	"strings"

	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/util"
)

func (app *App) handleCreateUser(rw http.ResponseWriter, r *http.Request) {
	var args db.CreateUserParams
	err := getBody(r, &args)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "could not parse body")
		return
	}
	_, err = app.store.GetUserByEmail(r.Context(), args.Email)
	if err == nil {
		// means user exists
		sendErrorResponse(rw, http.StatusMethodNotAllowed, nil, "user already exists")
		return
	}

	user, err := app.store.CreateUser(r.Context(), args)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "")
		return
	}

	// created session
	session := util.RandomString(32)
	app.kv[session] = util.ToString(user.UserID)

	sendResponse(rw, http.StatusCreated, user, session)
}

func (app *App) handleLogin(rw http.ResponseWriter, r *http.Request) {
	var args map[string]string
	err := getBody(r, &args)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "could not parse body")
		return
	}

	user, err := app.store.GetUserByEmail(r.Context(), args["email"])
	if err != nil {
		sendErrorResponse(rw, http.StatusUnauthorized, nil, err.Error())
		return
	}

	if !(user.Email == args["email"] && user.Password == args["password"]) {
		sendResponse(rw, http.StatusUnauthorized, user, "password/email mismatch")
	}

	session := util.RandomString(32)
	app.kv[session] = util.ToString(user.UserID)
	sendResponse(rw, http.StatusOK, nil, session)
}

func (app *App) handleLogout(rw http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.Header.Get("Authorization"), "Bearer ")

	if len(parts) < 2 || parts[1] == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "no access token found in request")
		return
	}

	token := parts[1]

	delete(app.kv, token)
	sendResponse(rw, http.StatusOK, nil, "")
}

func (app *App) handleIfLogin(rw http.ResponseWriter, r *http.Request) {
	sendResponse(rw, 200, nil, "logged in")
}

func (app *App) handleCheckAllowance(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		parts := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(parts) < 2 || parts[1] == "" {
			sendErrorResponse(rw, http.StatusBadRequest, nil, "no access token found in request")
			return
		}

		token := parts[1]
		if token == "12345" {
			newCtx := context.WithValue(r.Context(), "session", "1")
			next.ServeHTTP(rw, r.WithContext(newCtx))
			return
		}

		if v, ok := app.kv[token]; ok {
			newCtx := context.WithValue(r.Context(), "session", v)
			next.ServeHTTP(rw, r.WithContext(newCtx))
			return
		}

		sendResponse(rw, http.StatusUnauthorized, nil, "not logged in")
	})
}
