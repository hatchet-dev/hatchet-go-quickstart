package main

import (
	"hatchet-go-quickstart/worker"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ch := make(chan string, 5)
	cleanup, err := worker.Run(ch)
	if err != nil {
		panic(err)
	}
	if err := cleanup(); err != nil {
		panic(err)
	}
}
