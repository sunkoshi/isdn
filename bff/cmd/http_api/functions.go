package main

import (
	"net/http"

	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/util"
)

func (app *App) handleCreateFunction(rw http.ResponseWriter, r *http.Request) {
	userId := int64(util.ToInt(r.Context().Value("session").(string)))
	var args db.CreateFunctionParams
	err := getBody(r, &args)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "could not parse body")
		return
	}

	args.CreatorID = userId
	fn, err := app.store.CreateFunction(r.Context(), args)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "")
		return
	}
	sendResponse(rw, http.StatusCreated, fn, "")
}

func (app *App) handleGetFunctions(rw http.ResponseWriter, r *http.Request) {
	userId := int64(util.ToInt(r.Context().Value("session").(string)))
	fns, err := app.store.GetFunctionsByCreatorId(r.Context(), userId)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "")
		return
	}
	sendResponse(rw, http.StatusOK, fns, "")
}

func (app *App) handleDeleteFunction(rw http.ResponseWriter, r *http.Request) {
	userId := int64(util.ToInt(r.Context().Value("session").(string)))
	fnId := int64(util.ToInt(r.URL.Query().Get("id")))
	err := app.store.DeleteFunctionsByIdAndCreatorId(r.Context(), db.DeleteFunctionsByIdAndCreatorIdParams{
		ID:        fnId,
		CreatorID: userId,
	})
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "")
		return
	}
	sendResponse(rw, http.StatusOK, nil, "")
}
