package function_executor

import (
	"time"

	"github.com/orted-org/dambda/internal/file_manager"
	"github.com/orted-org/dambda/internal/lang_handler"
)

type FunctionExecutorParams struct {
	RequestID string        `json:"request_id"`
	Code      string        `json:"code"`
	Language  string        `json:"lang"`
	Input     FunctionInput `json:"input"`
}

type FunctionExecutor struct {
	params      FunctionExecutorParams
	fileManger  *file_manager.FileManager
	langHandler *lang_handler.LanguageHandler
}

type FunctionInput struct {
	File          string `json:"file"`
	InputFileName string `json:"file_name"`
	Stdin         string `json:"stdin"`
	CLA           string `json:"cla"`
}

type FunctionExecutionResult struct {
	ExecutionTime time.Duration `json:"execution_time"`
	Output        string        `json:"output"`
	Error         string        `json:"error"`
}
