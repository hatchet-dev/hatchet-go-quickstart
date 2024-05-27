package main

import (
	"hatchet-go-quickstart/server"
	"hatchet-go-quickstart/worker"
	"net/http"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	if err := server.Run(); err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	ch := make(chan string, 5)
	if err := worker.Run(ch); err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	if _, err := http.Get("http://localhost:1323/test"); err != nil {
		t.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	if <-ch != "step1" {
		t.Fatal("step1 not received")
	}
	if <-ch != "step2" {
		t.Fatal("step2 not received")
	}
}
