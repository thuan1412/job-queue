package worker

import (
	"job-queue/pkg/shared"
	"time"
)

type Task struct {
	ID       int32  `json:"id"`
	UUID     string `json:"uuid"`
	TaskType string `json:"task_type"`
	Args     any
	ArgsStr  string `json:"args"`
	Status   string `json:"status"`
	Handler  shared.Handler
	Retry    time.Duration
}

func (t *Task) Handle() error {
	return t.Handler(t.Args)
}
