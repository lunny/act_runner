package runtime

import (
	"context"

	"gitea.com/gitea/act_runner/client"
	runnerv1 "gitea.com/gitea/proto-go/runner/v1"

	"github.com/sirupsen/logrus"
)

// Defines the Resource Kind and Type.
const (
	Kind = "pipeline"
	Type = "docker"
)

// Runner runs the pipeline.
type Runner struct {
	Machine string
	Environ map[string]string
	Client  client.Client
}

// Run runs the pipeline stage.
func (s *Runner) Run(ctx context.Context, stage *runnerv1.Stage) error {
	l := logrus.
		WithField("runner.ID", stage.Id).
		WithField("runner.BuildID", stage.BuildId)

	l.Info("start running pipeline")
	// TODO: docker runner with stage data
	// task.Run is blocking, so we need to use goroutine to run it in background
	// return task metadata and status to the server
	task := NewTask()

	return task.Run(ctx)
}
