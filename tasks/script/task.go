package script

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/ChrisMcKenzie/freight/tasks"
)

func init() {
	tasks.RegisterTask("script", func() tasks.Task { return &Task{} })
}

// Task define a Task that can be used in config that run a given set of commands
type Task struct {
	Command string `mapstructure:"command" hcl:"command"`
}

func (t *Task) Run(ctx context.Context, cwd string) error {
	cmd := exec.CommandContext(ctx, "bash", "-x", "-c", t.Command)
	cmd.Dir = cwd

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", stdoutStderr)

	return nil
}
