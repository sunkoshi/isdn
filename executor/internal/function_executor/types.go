package function_executor

import (
	"io"
	"time"

	"github.com/orted-org/isdn/internal/file_manager"
	"github.com/orted-org/isdn/internal/lang_handler"
)

type FunctionExecutorParams struct {
	RequestID string    `json:"request_id"`
	Code      io.Reader `json:"code"`
	IsZip     bool      `json:"is_zip"`
	Language  string    `json:"lang"`
	Input     string    `json:"input"`
}

type FunctionExecutor struct {
	params           FunctionExecutorParams
	workingDirectory string
	codeFilePath     string
	fileManger       *file_manager.FileManager
	langHandler      *lang_handler.LanguageHandler
}
type FunctionExecutionResult struct {
	ExecutionTime time.Duration `json:"execution_time"`
	Stdout        string        `json:"stdout"`
	Error         string        `json:"error"`
	Output        string        `json:"output"`
}
