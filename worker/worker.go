package worker

import (
	"context"
	"fmt"
	"hatchet-go-quickstart/internal/events"
	"os"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/hatchet-dev/hatchet/pkg/worker"
)

func Run() error {
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
				Function: func(ctx context.Context, event *events.TestEvent) error {
					fmt.Println("got event: ", event.Name)
					return nil
				},
				Name: "first-event",
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
