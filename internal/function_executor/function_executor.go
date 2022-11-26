package function_executor

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/orted-org/isdn/internal/file_manager"
	"github.com/orted-org/isdn/internal/lang_handler"
	"github.com/orted-org/isdn/pkg/command_executor"
	"github.com/orted-org/isdn/util"
)

const BASE_DIR = "/home/hs/Documents/Projects/isdn/functions"
const TEMPLATES_DIR = "/home/hs/Documents/Projects/isdn/templates"

func New(langHandler *lang_handler.LanguageHandler, params FunctionExecutorParams) (*FunctionExecutor, error) {
	if !langHandler.IfConfigExists(params.Language) {
		return nil, errors.New("language config does not exists")
	}

	return &FunctionExecutor{
		params:      params,
		fileManger:  file_manager.New(""),
		langHandler: langHandler,
	}, nil
}

func (fe *FunctionExecutor) Compile(ctx context.Context) (string, error) {
	compileCmd := fe.langHandler.GetCompileCmd(fe.workingDirectory, "code", fe.params.Language)

	// if no compile required
	if compileCmd == "" {
		return "", nil
	}

	log.Println("compiling with", compileCmd)

	// only taking stdErr and err
	_, stdErr, err := command_executor.ExecuteContext(ctx, command_executor.SH_SHELL, compileCmd)
	if err != nil {
		return stdErr, err
	}
	return "", nil
}

func (fe *FunctionExecutor) Execute(ctx context.Context) (string, string, error) {
	executionCmd := fe.langHandler.GetExecutionCmd(fe.workingDirectory, "code", fe.params.Language)
	log.Println("executing with", executionCmd)
	return command_executor.ExecuteContext(ctx, command_executor.SH_SHELL, executionCmd)
}

func (fe *FunctionExecutor) Run(ctx context.Context) FunctionExecutionResult {
	var result FunctionExecutionResult
	start := time.Now()
	defer func() {
		log.Println("cleaning resources")
		err := fe.Clean()
		if err != nil {
			result.Error = fmt.Sprintf("ERROR: could not clean provisioned resources\n%s", err.Error())
		}
		result.ExecutionTime = time.Since(start)
	}()

	err := fe.Provision()
	if err != nil {
		result.Error = fmt.Sprintf("ERROR: could not provision resources\n%s", err.Error())
		return result
	}

	if stdErr, err := fe.Compile(ctx); err != nil {
		// merging stdErr and err
		result.Error = mergeError(err, stdErr)
		return result
	}

	stdOut, stdErr, err := fe.Execute(ctx)
	if err != nil {
		// merging stdErr and err
		result.Error = mergeError(err, stdErr)
		return result
	}
	result.Stdout = stdOut
	functionOutput, err := os.ReadFile(path.Join(fe.workingDirectory, "output.out"))
	if err != nil {
		result.Error = util.MergeErrors(errors.New("could not extract output"), err).Error()
		return result
	}
	result.Output = string(functionOutput)

	return result
}

func mergeError(err error, stdErr string) string {
	return fmt.Sprintf("ERROR: %s\n%s", err.Error(), stdErr)
}
