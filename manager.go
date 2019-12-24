package freight

import (
	"context"
	"fmt"

	homedir "github.com/mitchellh/go-homedir"

	// register all tasks
	_ "github.com/ChrisMcKenzie/freight/tasks/script"
)

// Manager manages the resolving projects and tasks in a freight manafest
type Manager struct{}

func (m *Manager) Run(ctx context.Context, cfg *Config) error {
	rd, err := homedir.Expand(cfg.Base.Root)
	if err != nil {
		return err
	}

	cfg.Base.Root = rd

	for _, project := range cfg.Projects {
		fmt.Printf("Syncing %q\n", project.Name)
		if err := project.Resolve(ctx, cfg.Base); err != nil {
			return err
		}
	}

	return nil
}
