package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/internal/worker_manager"
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
	log.Println(handler.Header.Get("Content-Type"))
	if !(strings.Contains(handler.Header.Get("Content-Type"), "zip") || strings.Contains(handler.Header.Get("Content-Type"), "octet-stream")) {
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

func (app *App) handleFunctionCall(rw http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "function id missing")
		return
	}
	var arg map[string]interface{}
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "could not parse body")
		return
	}

	fn, err := app.store.GetFunctionById(r.Context(), int64(util.ToInt(idStr)))
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(rw, http.StatusBadRequest, nil, "no such function exists")
			return
		}
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	fnCall, err := app.store.CreateFunctionCall(r.Context(), fn.ID)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	var dataChan chan worker_manager.FunctionExecutionResult = make(chan worker_manager.FunctionExecutionResult)
	err = app.wm.AddNewFunctionCallJob(fn.ID, fn.Language, fnCall.ID, arg, dataChan)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	res := <-dataChan

	app.store.UpdateFunctionCall(r.Context(), db.UpdateFunctionCallParams{
		Output: res.Output,
		Stdout: res.Stdout,
		Error:  res.Stdout,
		Cost:   int64(res.ExecutionTime),
	})

	if fn.OutputType == "json" {
		var finalRes map[string]interface{}
		util.JSONParse([]byte(res.Output), &finalRes)
		sendResponse(rw, http.StatusOK, finalRes, "")
		return
	}

	sendHTMLResponse(rw, http.StatusOK, res.Output)
}
