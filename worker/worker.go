package worker

import (
	"fmt"
	"hatchet-go-quickstart/internal/types"
	"os"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/hatchet-dev/hatchet/pkg/worker"
)

func Run(check chan string) error {
	token := os.Getenv("HATCHET_CLIENT_TOKEN")
	if token == "" {
		return fmt.Errorf("HATCHET_CLIENT_TOKEN is not set")
	}

	hatchetClient, err := client.New(
		client.WithToken(token),
	)
	if err != nil {
		return err
	}

	w, err := worker.NewWorker(
		worker.WithClient(hatchetClient),
	)
	if err != nil {
		return err
	}

	err = w.On(worker.Event("test-called"), &worker.WorkflowJob{
		Name:        "event-test",
		Description: "Test workflow.",
		Steps: []*worker.WorkflowStep{
			{
				Function: func(ctx worker.HatchetContext, event *types.TestEvent) error {
					fmt.Println("got event: ", event.Name)
					check <- "step1"
					return nil
				},
				Name: "first-event",
			},
			{
				Parents: []string{"first-event"},
				Function: func(ctx worker.HatchetContext, event *types.TestEvent) error {
					fmt.Println("got event: ", event.Name)
					check <- "step2"
					return nil
				},
				Name: "second-event",
			},
		},
	})
	if err != nil {
		return err
	}

	cleanup, err := w.Start()
	if err != nil {
		return err
	}

	// TODO cleanup
	_ = cleanup

	return nil
}
