package worker_manager

import "time"

type FunctionExecutionResult struct {
	RequestID     string        `json:"request_id"`
	ExecutionTime time.Duration `json:"execution_time"`
	Stdout        string        `json:"stdout"`
	Error         string        `json:"error"`
	Output        string        `json:"output"`
}
