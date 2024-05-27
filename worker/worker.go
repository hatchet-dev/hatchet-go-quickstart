package worker

import (
	"fmt"
	"hatchet-go-quickstart/internal/types"
	"os"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/hatchet-dev/hatchet/pkg/worker"
)

func Run(check chan string) (func() error, error) {
	token := os.Getenv("HATCHET_CLIENT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("HATCHET_CLIENT_TOKEN is not set")
	}

	_ = os.Setenv("HATCHET_CLIENT_TLS_STRATEGY", "none")

	hatchetClient, err := client.New(
		client.WithToken(token),
	)
	if err != nil {
		return nil, err
	}

	w, err := worker.NewWorker(
		worker.WithClient(hatchetClient),
	)
	if err != nil {
		return nil, err
	}

	err = w.On(worker.Event("test-called"), &worker.WorkflowJob{
		Name:        "event-test",
		Description: "Test workflow.",
		Steps: []*worker.WorkflowStep{
			{
				Function: func(ctx worker.HatchetContext, event *types.TestEvent) (*types.StepResponse, error) {
					fmt.Println("got event: ", event.Name)
					check <- "step1"
					return &types.StepResponse{
						Message: "step1 response",
					}, nil
				},
				Name: "first-event",
			},
			{
				Parents: []string{"first-event"},
				Function: func(ctx worker.HatchetContext, event *types.TestEvent) (*types.StepResponse, error) {
					fmt.Println("got event: ", event.Name)
					check <- "step2"
					return &types.StepResponse{
						Message: "step2 response",
					}, nil
				},
				Name: "second-event",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	cleanup, err := w.Start()
	if err != nil {
		return nil, err
	}

	return cleanup, nil
}
