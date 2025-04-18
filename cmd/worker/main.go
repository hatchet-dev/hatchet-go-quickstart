package main

import (
	workflows "hatchet-go-quickstart/workflows"

	"github.com/hatchet-dev/hatchet/pkg/cmdutils"
	v1 "github.com/hatchet-dev/hatchet/pkg/v1"
	"github.com/hatchet-dev/hatchet/pkg/v1/worker"
	"github.com/hatchet-dev/hatchet/pkg/v1/workflow"
)

func main() {

	hatchet, err := v1.NewHatchetClient()

	if err != nil {
		panic(err)
	}

	worker, err := hatchet.Worker(
		worker.WorkerOpts{
			Name: "first-workflow-worker",
			Workflows: []workflow.WorkflowBase{
				workflows.FirstWorkflow(hatchet),
			},
		},
	)

	if err != nil {
		panic(err)
	}

	// we construct an interrupt context to handle Ctrl+C
	// you can pass in your own context.Context here to the worker
	interruptCtx, cancel := cmdutils.NewInterruptContext()

	defer cancel()

	err = worker.StartBlocking(interruptCtx)

	if err != nil {
		panic(err)
	}
}
