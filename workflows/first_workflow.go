package workflows

import (
	"fmt"
	"strings"

	v1 "github.com/hatchet-dev/hatchet/pkg/v1"
	"github.com/hatchet-dev/hatchet/pkg/v1/task"
	"github.com/hatchet-dev/hatchet/pkg/v1/workflow"
	"github.com/hatchet-dev/hatchet/pkg/worker"
)

type SimpleInput struct {
	Message string
}

type LowerOutput struct {
	TransformedMessage string
}

type SimpleResult struct {
	ToLower LowerOutput
}

func FirstWorkflow(hatchet *v1.HatchetClient) workflow.WorkflowDeclaration[SimpleInput, SimpleResult] {

	simple := v1.WorkflowFactory[SimpleInput, SimpleResult](
		workflow.CreateOpts[SimpleInput]{
			Name: "first-workflow",
		},
		hatchet,
	)

	simple.Task(
		task.CreateOpts[SimpleInput]{
			Name: "ToLower",
			Fn: func(input SimpleInput, ctx worker.HatchetContext) (*LowerOutput, error) {
				fmt.Println("ToLower task called")
				return &LowerOutput{
					TransformedMessage: strings.ToLower(input.Message),
				}, nil
			},
		},
	)

	return simple
}
