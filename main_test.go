package main

import (
	"hatchet-go-quickstart/server"
	"hatchet-go-quickstart/worker"
	"net/http"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestExample(t *testing.T) {
	_ = godotenv.Load()

	if err := server.Run(); err != nil {
		t.Fatal(err)
	}

	ch := make(chan string, 5)
	cleanup, err := worker.Run(ch)
	defer func() {
		if err := cleanup(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := http.Get("http://localhost:1323/test"); err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	if <-ch != "step1" {
		t.Fatal("step1 not received")
	}
	if <-ch != "step2" {
		t.Fatal("step2 not received")
	}
}
