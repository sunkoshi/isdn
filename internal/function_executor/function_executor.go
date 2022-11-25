package function_executor

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/orted-org/dambda/internal/file_manager"
	"github.com/orted-org/dambda/internal/lang_handler"
	"github.com/orted-org/dambda/pkg/command_executor"
)

func New(langHandler *lang_handler.LanguageHandler, params FunctionExecutorParams) (*FunctionExecutor, error) {
	if !langHandler.IfConfigExists(params.Language) {
		return nil, errors.New("language config does not exists")
	}

	return &FunctionExecutor{
		params:      params,
		fileManger:  file_manager.New("./function"),
		langHandler: langHandler,
	}, nil
}

func (fe *FunctionExecutor) Provision() error {

	// forming filename
	fileName := fmt.Sprintf("code.%s", fe.langHandler.GetExtension(fe.params.Language))

	// saving file in system
	err := fe.fileManger.Put([]string{fe.params.RequestID, fileName}, []byte(fe.params.Code))
	if err != nil {
		return err
	}

	// handling file input
	if fe.params.Input.InputFileName != "" {
		// create file for input
		err = fe.fileManger.Put([]string{fe.params.RequestID, fe.params.Input.InputFileName}, []byte(fe.params.Input.File))
		if err != nil {
			return err
		}
	}

	return nil
}

func (fe *FunctionExecutor) Compile(ctx context.Context) (string, error) {

	compileCmd := fe.langHandler.GetCompileCmd("code", fe.params.Language)

	// if no compile required
	if compileCmd == "" {
		return "", nil
	}
	_, stdErr, err := command_executor.ExecuteContext(ctx, command_executor.SH_SHELL, compileCmd)
	if err != nil {
		return stdErr, err
	}
	return "", nil
}

func (fe *FunctionExecutor) Execute(ctx context.Context) (string, string, error) {
	executionCmd := fe.langHandler.GetExecutionCmd("code", fe.params.Language)
	return command_executor.ExecuteContext(ctx, command_executor.SH_SHELL, executionCmd)
}

func (fe *FunctionExecutor) Run(ctx context.Context) FunctionExecutionResult {
	start := time.Now()
	var result FunctionExecutionResult
	if stdErr, err := fe.Compile(ctx); err != nil {
		result.Error = mergeError(err, stdErr)
	} else {
		stdOut, stdErr, err := fe.Execute(ctx)
		if err != nil {
			result.Error = mergeError(err, stdErr)
		}
		result.Output = stdOut
	}
	result.ExecutionTime = time.Since(start)
	return result
}

func mergeError(err error, stdErr string) string {
	return fmt.Sprintf("ERROR: %s\n%s", err.Error(), stdErr)
}
