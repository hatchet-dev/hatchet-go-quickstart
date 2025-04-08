package workflows

import (
	"fmt"
	"strings"

	"github.com/hatchet-dev/hatchet/pkg/client/create"
	v1 "github.com/hatchet-dev/hatchet/pkg/v1"
	"github.com/hatchet-dev/hatchet/pkg/v1/factory"
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

func FirstWorkflow(hatchet v1.HatchetClient) workflow.WorkflowDeclaration[SimpleInput, SimpleResult] {
	simple := factory.NewWorkflow[SimpleInput, SimpleResult](
		create.WorkflowCreateOpts[SimpleInput]{
			Name: "first-workflow",
		},
		hatchet,
	)

	simple.Task(
		create.WorkflowTask[SimpleInput, SimpleResult]{
			Name: "ToLower",
		},
		func(ctx worker.HatchetContext, input SimpleInput) (any, error) {
			fmt.Println("ToLower task called")
			return &LowerOutput{
				TransformedMessage: strings.ToLower(input.Message),
			}, nil
		},
	)

	return simple
}
