package tasks

import "context"

type Task interface {
	Run(ctx context.Context, cwd string) error
}

type TaskFunc func() Task

var registry map[string]TaskFunc

func init() {
	registry = make(map[string]TaskFunc)
}

func RegisterTask(name string, t TaskFunc) {
	registry[name] = t
}

func GetTask(name string) Task {
	return registry[name]()
}
