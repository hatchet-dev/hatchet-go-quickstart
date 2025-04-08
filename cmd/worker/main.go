package main

import (
	workflows "hatchet-go-quickstart/workflows"

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

	err = worker.StartBlocking()

	if err != nil {
		panic(err)
	}
}
