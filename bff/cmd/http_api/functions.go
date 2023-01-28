package main

import (
	"net/http"
	"strings"

	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/util"
)

func (app *App) handleCreateFunction(rw http.ResponseWriter, r *http.Request) {
	userId := int64(util.ToInt(r.Context().Value("session").(string)))
	var args db.CreateFunctionParams

	// handling function details values
	args.CreatorID = userId
	args.Language = r.FormValue("language")
	args.Name = r.FormValue("name")
	args.Timeout = int64(util.ToInt(r.FormValue("timeout")))

	if r.FormValue("output_type") != "" {
		args.OutputType = r.FormValue("output_type")
	} else {
		args.OutputType = "json"
	}

	// creating function with details
	fn, err := app.store.CreateFunction(r.Context(), args)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// handling file
	file, handler, err := r.FormFile("code")
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid name/ file missing")
		return
	}
	defer file.Close()
	if !(strings.Contains(handler.Header.Get("Content-Type"), "zip")) {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "only .zip allowed")
		return
	}
	fileName := util.ToString(fn.ID) + ".zip"
	_, err = app.objectStore.Put(file, fileName)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
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
