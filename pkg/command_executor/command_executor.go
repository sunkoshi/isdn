package command_executor

import (
	"bytes"
	"context"
	"os/exec"
)

func ExecuteContext(ctx context.Context, dir string, rawCmd string) (string, string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", rawCmd)
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Dir = dir
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer
	err := cmd.Run()
	if err != nil {
		return "", errorBuffer.String(), err
	}
	return outputBuffer.String(), errorBuffer.String(), nil
}
