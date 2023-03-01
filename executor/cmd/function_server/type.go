package main

import (
	"time"
)

type FunctionCallRequest struct {
	Input      map[string]interface{} `json:"input"`
	Language   string                 `json:"language"`
	RequestID  int                    `json:"request_id"`
	FunctionID int                    `json:"function_id"`
}
type FunctionExecutionResult struct {
	ExecutionTime time.Duration `json:"execution_time"`
	Stdout        string        `json:"stdout"`
	Error         string        `json:"error"`
	Output        string        `json:"output"`
}
type ExecutorConfig struct {
	MasterURL       string `json:"master_url"`
	CodeRepoBaseURL string `json:"code_repo_base_url"`
}
