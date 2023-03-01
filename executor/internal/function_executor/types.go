package function_executor

import (
	"time"

	"github.com/orted-org/isdn/internal/file_manager"
	"github.com/orted-org/isdn/internal/lang_handler"
)

type FunctionExecutorParams struct {
	RequestID string `json:"request_id"`
	CodeRepo  string `json:"code_repo"`
	Language  string `json:"lang"`
	Input     string `json:"input"`
}

type FunctionExecutor struct {
	params           FunctionExecutorParams
	workingDirectory string
	fileManger       *file_manager.FileManager
	langHandler      *lang_handler.LanguageHandler
}
type FunctionExecutionResult struct {
	RequestID     string        `json:"request_id"`
	ExecutionTime time.Duration `json:"execution_time"`
	Stdout        string        `json:"stdout"`
	Error         string        `json:"error"`
	Output        string        `json:"output"`
}
