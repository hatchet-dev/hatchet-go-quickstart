package main

import (
	"hatchet-go-quickstart/server"
	"hatchet-go-quickstart/worker"
	"net/http"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	go func() {
		if err := server.Run(); err != nil {
			t.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)

	worker.Run()

	time.Sleep(5 * time.Second)

	if _, err := http.Get("http://localhost:1323/test"); err != nil {
		t.Fatal(err)
	}
}
