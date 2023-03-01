package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/orted-org/isdn/internal/lang_handler"
)

func initLanguageHandler(app *App) {
	lh, err := lang_handler.New()
	if err != nil {
		panic(err)
	}
	app.lh = lh
}

func initConfig(app *App) {
	file, err := os.ReadFile(path.Join("config", "executor_config.json"))
	if err != nil {
		panic(err)
	}
	var config ExecutorConfig
	err = JSONParse(file, &config)
	if err != nil {
		panic(err)
	}
	app.config = &config
}

func initConnectionToMaster(app *App) {
	log.Printf("connecting to master at %s", app.config.MasterURL)
	suppotedLangs := strings.Join(app.lh.GetSupportedLanguages(), ",")
	urlWithLangs := fmt.Sprintf("%s?langs=%s", app.config.MasterURL, suppotedLangs)
	log.Printf("formed url for connection with master %s", urlWithLangs)
	conn, _, err := websocket.DefaultDialer.Dial(urlWithLangs, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	log.Println("connected to master")
	app.conn = conn
}
