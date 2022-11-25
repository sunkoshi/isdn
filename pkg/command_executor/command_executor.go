package command_executor

import (
	"bytes"
	"context"
	"os/exec"
)

const SH_SHELL = "sh"
const BASH_SHELL = "bash"
const ZSH_SHELL = "zsh"

func ExecuteContext(ctx context.Context, shell string, rawCmd string) (string, string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", rawCmd)
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer
	err := cmd.Run()
	if err != nil {
		return "", errorBuffer.String(), err
	}
	return outputBuffer.String(), errorBuffer.String(), nil
}
