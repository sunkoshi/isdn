package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (app *App) handleWs(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
	}
	langs := r.URL.Query().Get("langs")
	app.wm.AddNewNode(1, strings.Split(langs, ","), ws)
}
