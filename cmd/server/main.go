package main

import (
	"hatchet-go-quickstart/server"

	"github.com/hatchet-dev/hatchet/pkg/cmdutils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}

	interruptCh := cmdutils.InterruptChan()

	<-interruptCh
}
