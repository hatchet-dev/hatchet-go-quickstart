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
	if err := worker.Run(ch); err != nil {
		panic(err)
	}
}
