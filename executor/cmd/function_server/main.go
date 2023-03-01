package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/orted-org/isdn/internal/function_executor"
	"github.com/orted-org/isdn/internal/lang_handler"
)

type App struct {
	lh     *lang_handler.LanguageHandler
	conn   *websocket.Conn
	config *ExecutorConfig
}

func main() {

	app := &App{}
	initConfig(app)
	initLanguageHandler(app)
	initConnectionToMaster(app)
	for {
		_, payload, err := app.conn.ReadMessage()
		if err != nil {
			log.Println("Error", err)
			continue
		}
		result, err := app.handleFunctionCall(payload)
		if err != nil {
			log.Println("Error", err)
		} else {
			log.Println("sending result for request id", result.RequestID, result)
			app.conn.WriteMessage(websocket.TextMessage, []byte(JSONStringify(result)))
		}
	}
}

func (app *App) handleFunctionCall(req []byte) (*function_executor.FunctionExecutionResult, error) {

	// parsing the request from req byte
	var args FunctionCallRequest
	err := JSONParse(req, &args)
	if err != nil {
		return &function_executor.FunctionExecutionResult{
			Error: "could not parse request for function call",
		}, err
	}

	// forming new fe
	fe, err := function_executor.New(app.lh, function_executor.FunctionExecutorParams{
		RequestID: fmt.Sprintf("%d", args.RequestID),
		CodeRepo:  fmt.Sprintf("%s/%d.zip", app.config.CodeRepoBaseURL, args.FunctionID),
		Language:  args.Language,
		Input:     JSONStringify(args.Input),
	})
	if err != nil {
		return nil, err
	}

	// running the function
	output := fe.Run(context.Background())
	if err != nil {
		return nil, err
	}
	// fmt.Println("Error:", output.Error)
	// fmt.Println("Stdout:", output.Stdout)
	// fmt.Println("Output:", output.Output)
	// fmt.Println("Execution Time:", output.ExecutionTime)
	return &output, nil
}
