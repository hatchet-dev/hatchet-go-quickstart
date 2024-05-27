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

	if err := worker.Run(); err != nil {
		panic(err)
	}
}
