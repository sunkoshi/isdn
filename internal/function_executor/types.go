package function_executor

import (
	"io"
	"time"

	"github.com/orted-org/isdn/internal/file_manager"
	"github.com/orted-org/isdn/internal/lang_handler"
)

type FunctionExecutorParams struct {
	RequestID string        `json:"request_id"`
	Code      io.Reader     `json:"code"`
	IsZip     bool          `json:"is_zip"`
	Language  string        `json:"lang"`
	Input     FunctionInput `json:"input"`
}

type FunctionExecutor struct {
	params           FunctionExecutorParams
	workingDirectory string
	codeFilePath     string
	fileManger       *file_manager.FileManager
	langHandler      *lang_handler.LanguageHandler
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
