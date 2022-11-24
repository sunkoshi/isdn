package command_executor

import (
	"context"
	"os/exec"
)

const SH_SHELL = "sh"
const BASH_SHELL = "bash"
const ZSH_SHELL = "zsh"

func ExecuteContext(ctx context.Context, shell string, cmd string) (string, error) {
	r := exec.CommandContext(ctx, "bash", "-c", cmd)
	output, err := r.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func Execute(ctx context.Context, shell string, cmd string) (string, error) {
	r := exec.CommandContext(ctx, cmd)
	output, err := r.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
