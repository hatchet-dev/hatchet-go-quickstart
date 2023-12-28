package main

import (
	"context"
	"fmt"
	"hatchet-go-quickstart/internal/events"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/hatchet-dev/hatchet/pkg/cmdutils"
	"github.com/hatchet-dev/hatchet/pkg/worker"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	ctx, cancel := cmdutils.NewInterruptContext()

	defer cancel()

	hatchetClient, err := client.New(
		client.WithHostPort("127.0.0.1", 7077),
	)

	if err != nil {
		panic(err)
	}

	w, err := worker.NewWorker(
		worker.WithClient(hatchetClient),
	)

	if err != nil {
		panic(err)
	}

	err = w.On(worker.Event("test-called"), &worker.WorkflowJob{
		Name:        "event-test",
		Description: "Test workflow.",
		Timeout:     "60s",
		Steps: []worker.WorkflowStep{
			{
				Function: func(ctx context.Context, event *events.TestEvent) error {
					fmt.Println("got event: ", event.Name)
					return nil
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	w.Start(ctx)
}
