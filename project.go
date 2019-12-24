package freight

import (
	"context"
	"path"

	"github.com/ChrisMcKenzie/freight/tasks"
	getter "github.com/hashicorp/go-getter"
)

type Project struct {
	Name   string
	Remote string
	Path   string

	AfterTasks []tasks.Task
}

func (p *Project) Resolve(ctx context.Context, bc BaseConfig) error {
	dst := path.Join(bc.Root, p.Path, p.Name)

	err := getter.GetAny(dst, p.Remote, getter.WithContext(ctx), getter.WithProgress(defaultProgressBar))
	if err != nil {
		return err
	}

	for _, t := range p.AfterTasks {
		if err := t.Run(ctx, dst); err != nil {
			return err
		}
	}

	return nil
}
