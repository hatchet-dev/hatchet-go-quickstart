package main

import (
	"context"
	"fmt"

	workflows "hatchet-go-quickstart/workflows"

	v1 "github.com/hatchet-dev/hatchet/pkg/v1"
)

func main() {

	hatchet, err := v1.NewHatchetClient()

	if err != nil {
		panic(err)
	}

	simple := workflows.FirstWorkflow(hatchet)

	result, err := simple.Run(context.Background(), workflows.SimpleInput{
		Message: "Hello, World!",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(result.ToLower.TransformedMessage)
}
